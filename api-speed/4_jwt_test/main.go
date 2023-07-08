package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cbot918/liby/jwty"
	"github.com/golang-jwt/jwt/v4"
)

var lg = fmt.Println

const secret = "AllYourBase"

func main() {
	jwty.New()
	token := createJWT()
	verifyJWT(token)
}

func createJWT() string {
	mySigningKey := []byte(secret)

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(10 * time.Second).Unix(),
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Fatal(err)
	}
	return ss
}

func verifyJWT(tokenString string) {
	// sample token string taken from the New example
	// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("verified")
	lg(token)

	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	fmt.Println(claims["foo"], claims["nbf"])
	// } else {
	// 	lg("1")
	// 	fmt.Println(err)
	// }
}
