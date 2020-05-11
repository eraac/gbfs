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

	var ss gbfsspec.FeedStationInformation

	if err := c.Get(gbfsspec.FeedKeyStationInformation, &ss); err != nil {
		panic(err)
	}

	for _, s := range ss.Data.Stations {
		fmt.Printf("Station name: %s\n", s.Name)
	}

	fmt.Printf("Last updated: %s\n", ss.LastUpdated.ToTime().String())
}
