package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	username := "darryl@bluelasso.com"
	password := "secret!"

	fmt.Println("authenticate a user/password from the command line")
	// host := "https://data.bluelasso.com"
	host := "http://localhost:16001"

	ctx := new(Context)
	ctx.apikey = "5a66072c-0f9d-11e5-826f-d77d588b4b69"

	key := createKey(username, password)

	ctx.url = host + "/VEAccess/user/authenticate/" + key

	// fmt.Println( ctx )

	authenticate(ctx)
}

type Context struct {
	url    string
	apikey string
}

func createKey(username, password string) string {
	skey := username + "!" + password

	hasher := sha256.New()
	hasher.Write([]byte(skey))

	key := hex.EncodeToString(hasher.Sum(nil))

	return key
}

func authenticate(context *Context) (status int, err error) {
	url := context.url

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Add("X-API-Key", context.apikey)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return resp.StatusCode, err
	}

	defer resp.Body.Close()

	fmt.Printf("status code: %d from URL: %s\n", resp.StatusCode, resp.Request.URL)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return resp.StatusCode, err
	}

	fmt.Printf("%s\n", body)

	return resp.StatusCode, nil
}
