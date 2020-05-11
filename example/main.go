package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Eraac/gbfs"
	"github.com/Eraac/gbfs/spec/v2.0"
)

func main() {
	c, err := gbfs.NewHTTPClient(
		gbfs.HTTPOptionClient(http.Client{Timeout: 10 * time.Second}),
		gbfs.HTTPOptionBaseURL("https://gbfs.fordgobike.com/gbfs"),
		gbfs.HTTPOptionLanguage("en"),
	)
	if err != nil {
		panic(err)
	}

	var si gbfsspec.FeedStationInformation

	if err := c.Get(gbfsspec.FeedKeyStationInformation, &si); err != nil {
		panic(err)
	}

	for _, s := range si.Data.Stations {
		fmt.Printf("Station name: %s\n", s.Name)
	}

	fmt.Printf("Last updated: %s\n", si.LastUpdated.ToTime().String())
}
