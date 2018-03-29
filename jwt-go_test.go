package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"testing"
)

func TestJWT(t *testing.T) {

	privateKeyData, err := ioutil.ReadFile(".\\resources\\rsa_private_key_512.pem")
	if err != nil {
		t.Error(err)
		panic(err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)

	if err != nil {
		t.Error(err)
		panic(err)
	}

	publiceKeyData, err := ioutil.ReadFile(".\\resources\\rsa_public_key_512.pem")

	if err != nil {
		t.Error(err)
		panic(err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publiceKeyData)

	if err != nil {
		t.Error(err)
		panic(err)
	}

	inputClaims := TestToken{
		Test: "nhTest",
		StandardClaims: jwt.StandardClaims{
			Issuer:    "test",
			ExpiresAt: 0,
		},
	}

	//privateKey = privateKey
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, inputClaims)

	tokenStr, err := token.SignedString(privateKey)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	fmt.Println(tokenStr)

	recieveToken, err := jwt.ParseWithClaims(tokenStr, &TestToken{}, func(recieveToken *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		t.Error(err)
		panic(err)
	}
	fmt.Println(recieveToken.Valid)
	if claims, ok := recieveToken.Claims.(*TestToken); ok && recieveToken.Valid {
		fmt.Println(claims.Test)
	}
}

type TestToken struct {
	jwt.StandardClaims
	Test string `json:"test"`
}
