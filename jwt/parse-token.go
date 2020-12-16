package main

// @see https://jwt.io/ for list of clients and a validation ui
// @see https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac

import (
    "fmt"
    "github.com/dgrijalva/jwt-go"
)

var (
    // generated with uuid --bytes
    hmacSampleSecret = []byte("d7e99acc1184c1e3e0eed341faba9988c558c342ee0a56c1")
)

func Validate(tkn *jwt.Token) (interface{}, error) {
    if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("Unexpected %v", tkn.Header["alg"])
    }

    return hmacSampleSecret, nil
}

func parseToken(token string) {
    tk, err := jwt.Parse(token, Validate)
    if err != nil {
        fmt.Printf("error: %s\n", err)
        return
    }

    fmt.Printf("token: %v\n", tk)
    fmt.Printf("is token valid: %v\n", tk.Valid)

    if claims, ok := tk.Claims.(jwt.MapClaims); ok {
        for k,v := range claims {
            fmt.Printf("%v=%v\n", k, v)
        }
    } else {
        fmt.Println(err)
    }
}

func main() {

    parseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJiczE1NWVzaGd3cHNkNDY2IiwibmFtZSI6ImRhcnJ5bC53ZXN0QHJhaW5jaXR5c29mdHdhcmUuY29tIiwiaWF0IjoxNTE2MjM5MDIyfQ.tfkN-tONvPe_SLCBVESosT9lze4j-vXgzqFc9I5-H-I")

}
