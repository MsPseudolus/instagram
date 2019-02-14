package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/kr/pretty"
	"github.com/recentralized/instagram/integration"
)

func main() {

	var (
		count int
		maxID string
		raw   bool
	)

	flag.IntVar(&count, "count", 3, "number of records to return")
	flag.StringVar(&maxID, "maxid", "", "max_id for request")
	flag.BoolVar(&raw, "raw", false, "write raw output or parsed output")
	flag.Parse()

	ctx := context.Background()

	api := integration.NewAPI()
	api.KeepRawBody = raw

	params := url.Values{}
	params.Set("count", strconv.Itoa(count))
	params.Set("max_id", maxID)

	resp, err := api.GetRecentMedia(ctx, params)
	if err != nil {
		log.Fatalf("GetRecentMedia: %s", err)
	}

	if raw {
		io.Copy(os.Stdout, api.RawBody)
		os.Exit(0)
	}

	pretty.Fprintf(os.Stdout, "%# v", resp)
	os.Exit(0)
}
