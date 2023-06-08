package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	ghCredHelper "github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v52/github"
)

func main() {
	id := int64(220213)
	key, err := ioutil.ReadFile("../pansachin-test.2023-06-07.private-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	transport, err := ghCredHelper.NewAppsTransport(http.DefaultTransport, id, []byte(key))
	if err != nil {
		log.Fatal(err)
	}

	c := github.NewClient(
		&http.Client{
			Transport: transport,
		},
	)

	i, r, err := c.Apps.GetInstallation(
		context.Background(),
		// 38348259,
		38348246,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)

	fmt.Printf("App ID: %v\n", *i.AppID)
	fmt.Printf("App NodeID: %v\n", *i.GetAccount().NodeID)
	fmt.Printf("Login: %v\n", *i.GetAccount().Login)
	fmt.Printf("AvatarURL: %v\n", *i.GetAccount().AvatarURL)
	fmt.Printf("HTMLURL: %v\n", *i.GetAccount().HTMLURL)
	// fmt.Printf("Name: %v\n", *i.GetAccount().Name)
}
