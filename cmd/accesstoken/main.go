package main

import (
	"bufio"
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/recentralized/instagram/integration"
)

func main() {
	ctx := context.Background()

	oauth := integration.NewOAuth()

	state := "mystate"

	authURL := oauth.GetAuthorizeURL(state)

	fmt.Printf("\nVisit this url:\n%s\n", authURL)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter the resulting url:\n")
	redir, _ := reader.ReadString('\n')
	redir = strings.Trim(redir, "\n")

	redirURL, err := url.Parse(redir)
	if err != nil {
		fmt.Printf("Could not parse url: %v\n", err)
		os.Exit(1)
	}

	code, err := oauth.GetCodeFromRedirect(redirURL, state)
	if err != nil {
		fmt.Printf("Could not get code from url: %v\n", err)
		os.Exit(1)
	}

	token, err := oauth.GetAccessToken(ctx, code)
	if err != nil {
		fmt.Printf("Could not get access token: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n%s\n", token)
}
