package main

import (
	"context"
	"flag"
	"io"
	"log"
	"os"

	"github.com/kr/pretty"
	"github.com/recentralized/instagram/integration"
)

func main() {

	var (
		id  string
		raw bool
	)

	flag.StringVar(&id, "id", "", "id of media to get comments")
	flag.BoolVar(&raw, "raw", false, "write raw output or parsed output")
	flag.Parse()

	ctx := context.Background()

	api := integration.NewAPI()
	api.KeepRawBody = raw

	resp, err := api.GetMediaRecentComments(ctx, id)
	if err != nil {
		log.Fatalf("GetMediaRecentComments: %s", err)
	}

	if raw {
		io.Copy(os.Stdout, api.RawBody)
		os.Exit(0)
	}

	pretty.Fprintf(os.Stdout, "%# v", resp)
	os.Exit(0)
}
