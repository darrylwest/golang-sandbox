package main

/**
 * keypairgen creates a list of box pub/priv keys and encrypts then writes to the secondary database; at the same time, the
 * list of public keys are written to a json file (boxkeys.json) to make available to clients that need a peer pub/priv key that
 * the keyservice will recognize.  communications to the server should always begin by encrypting the request message with a
 * selected public key; the service will then look up the key/pair from the secondard database and decrypt the client message.
 */

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/darrylwest/naledi-key-service/keyservice"
	"golang.org/x/crypto/nacl/box"
	"golang.org/x/crypto/nacl/secretbox"
	"gopkg.in/redis.v3"
	"io"
	"os"
	"path"
)

var (
	config     *keyservice.Config
	ErrEncrypt = errors.New("encryption error")
	ErrDecrypt = errors.New("decryption error")
)

func Encrypt(key *[keyservice.KeySize]byte, message []byte) ([]byte, error) {
	nonce, err := keyservice.GenerateNonce()
	if err != nil {
		return nil, ErrEncrypt
	}

	out := make([]byte, len(nonce))
	copy(out, nonce[:])
	out = secretbox.Seal(out, message, nonce, key)

	return out, nil
}

func Decrypt(key *[keyservice.KeySize]byte, message []byte) ([]byte, error) {
	if len(message) < (keyservice.NonceSize + secretbox.Overhead) {
		return nil, ErrDecrypt
	}

	var nonce [keyservice.NonceSize]byte
	copy(nonce[:], message[:keyservice.NonceSize])
	out, ok := secretbox.Open(nil, message[keyservice.NonceSize:], &nonce, key)
	if !ok {
		return nil, ErrDecrypt
	}

	return out, nil
}

type BoxPair struct {
	pub  [keyservice.KeySize]byte
	priv [keyservice.KeySize]byte
}

func createPairs(n int) *[]BoxPair {
	pairs := make([]BoxPair, n, n)

	for i := 0; i < n; i++ {
		pub, priv, _ := box.GenerateKey(rand.Reader)

		bp := BoxPair{}
		copy(bp.pub[:], pub[:])
		copy(bp.priv[:], priv[:])

		pairs[i] = bp
	}

	return &pairs
}

func writePubKeys(pairs *[]BoxPair, filename string) error {
	fmt.Println("create and write box keys to", filename)
	f, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer f.Close()

	n, _ := io.WriteString(f, "{\n  \"boxkeys\":[\n")
	count := n
	for i, pair := range *pairs {
		pubkey := hex.EncodeToString(pair.pub[:])

		var str string
		if i == 0 {
			str = fmt.Sprintf("    \"%s\"", pubkey)
		} else {
			str = fmt.Sprintf(",\n    \"%s\"", pubkey)
		}

		n, _ = io.WriteString(f, str)
		count += n
	}
	n, _ = io.WriteString(f, "\n  ]\n}\n")
	count += n

	fmt.Printf("bytes written: %d to %s...\n", count, filename)

	return nil
}

func saveKeys(keys *[]BoxPair) error {
	secretkey := config.GetPrivateLocalKey()

	client := redis.NewClient(config.GetSecondaryRedisOptions())

	for _, pair := range *keys {
		key := createDbKey(pair.pub)

		fmt.Println("key:", key)

		enc, _ := Encrypt(secretkey, pair.priv[:])
		value := hex.EncodeToString(enc)
		fmt.Println("val:", value)

		// convert to strings and save to database set(key, value)...
		err := client.Set(key, value, 0).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func createDbKey(pub [keyservice.KeySize]byte) string {
	hash := sha512.New()
	hash.Write(pub[:])
	key := "BoxKey:" + hex.EncodeToString(hash.Sum(nil))

	return key
}

// return the pub/priv keys
func readKey(spub string) *BoxPair {
	key, err := hex.DecodeString(spub)
	if err != nil {
		panic(err)
	}

	box := new(BoxPair)
	copy(box.pub[:], key)

	client := redis.NewClient(config.GetSecondaryRedisOptions())

	dbkey := createDbKey(box.pub)
	val, err := client.Get(dbkey).Result()

	if err != nil {
		panic(err)
	}

	dec, err := hex.DecodeString(val)

	value, err := Decrypt(config.GetPrivateLocalKey(), dec)

	if err != nil {
		panic(err)
	}

	copy(box.priv[:], value)

	return box
}

func readConfig() *keyservice.Config {
	filename := path.Join(os.Getenv("HOME"), ".keyservice", "config.json")
	config, err := keyservice.ReadConfig(filename)

	if err != nil {
		panic(err)
	}

	return config
}

func main() {

	config = readConfig()

	if false {
		pairs := createPairs(100)
		writePubKeys(pairs, "boxkeys.json")
		saveKeys(pairs)
	} else {

		skey := "587b2d753c8409bbf876e7f9dc682b01a411cdd2ce6f0c66046d69c6343c1a1d"
		pk, _ := hex.DecodeString(skey)

		var key *[keyservice.KeySize]byte = new([keyservice.KeySize]byte)
		copy(key[:], pk)
		fmt.Println("db key:", createDbKey(*key))

		box := readKey(skey)
		fmt.Printf("%x %x\n", box.pub, box.priv)
	}
}
