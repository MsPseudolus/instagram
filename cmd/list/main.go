package main

import (
	"context"
	"encoding/json"
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
		out   string
	)

	flag.IntVar(&count, "count", 3, "number of records to return")
	flag.StringVar(&maxID, "maxid", "", "max_id for request")
	flag.StringVar(&out, "out", "json", "output: json,go,raw")
	flag.Parse()

	ctx := context.Background()

	api := integration.NewAPI()
	api.KeepRawBody = out == "raw"

	params := url.Values{}
	params.Set("count", strconv.Itoa(count))
	params.Set("max_id", maxID)

	resp, err := api.GetRecentMedia(ctx, params)
	if err != nil {
		log.Fatalf("GetRecentMedia: %s", err)
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
