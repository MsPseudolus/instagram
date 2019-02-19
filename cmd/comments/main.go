package main

import (
	"context"
	"encoding/json"
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
		out string
	)

	flag.StringVar(&id, "id", "", "id of media to get comments")
	flag.StringVar(&out, "out", "json", "output: json,go,raw")
	flag.Parse()

	ctx := context.Background()

	api := integration.NewAPI()
	api.KeepRawBody = out == "raw"

	resp, err := api.GetMediaRecentComments(ctx, id)
	if err != nil {
		log.Fatalf("GetMediaRecentComments: %s", err)
	}

	switch out {
	case "raw":
		_, err = io.Copy(os.Stdout, api.RawBody)
	case "go":
		pretty.Fprintf(os.Stdout, "%# v", resp)
	case "json":
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		err = enc.Encode(resp)
	}
	if err != nil {
		log.Fatalf("Writing output: %v", err)
	}
	os.Exit(0)
}
