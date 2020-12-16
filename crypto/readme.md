# linux command line encrypt/decrypt

```bash

openssl chacha20 -k <key> -iter 20 -e -in plain_text_file -out encrypted_file
openssl chacha20 -k <key> -iter 20 -d -in rencrypted_file -out plain_text_file

```

## References

* https://godoc.org/crypto/cipher for basic cipher tools
* https://godoc.org/golang.org/x/crypto/scrypt to refactor to use a good crypto password hash
* https://golang.org/pkg/crypto/rsa/

## AWS Key Management Store

```bash

aws kms encrypt --key-id ab3682a5-85d7-43e6-9f7a-38af0f751275 --plaintext 'this is a test' --output text --query CiphertextBlob | base64 --decode > blob.enc
aws kms decrypt --ciphertext-blob fileb://blob.enc --output text --query Plaintext | base64 --decode

```

###### darryl.west | 2020.03.02

