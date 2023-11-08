package main

import (
	_ "embed"
	"fmt"

	"github.com/GianOrtiz/apple-auth-go"
)

const (
	TEAM_ID   = "QVCMB47W2M"
	CLIENT_ID = "com.honganhdev.signintestservice"
	KEY_ID    = "W25JK5YU26"
)

//go:embed AuthKey_W25JK5YU26.p8
var key []byte

func main() {
	fmt.Println(key)
	appleAuth, err := apple.New(CLIENT_ID, TEAM_ID, KEY_ID, "/path/to/apple-sign-in-key.p8")
	if err != nil {
		panic(err)
	}

	// Validate authorization code from a mobile app.
	tokenResponse, err := appleAuth.ValidateCode("<AUTHORIZATION-CODE>")
	if err != nil {
		panic(err)
	}

	user, err := apple.GetUserInfoFromIDToken(tokenResponse.IDToken)
	if err != nil {
		panic(err)
	}

	// User Apple unique identification.
	fmt.Println(user.UID)
	// User email if the user provided it.
	fmt.Println(user.Email)
}
