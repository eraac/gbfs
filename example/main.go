package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Eraac/gbfs"
)

func main() {
	c, err := gbfs.NewHTTPClient(
		gbfs.HTTPOptionClient(http.Client{Timeout: 10 * time.Second}),
		gbfs.HTTPOptionBaseURL("https://gbfs.fordgobike.com/gbfs"),
		gbfs.HTTPOptionLanguage("en"),
		gbfs.HTTPOptionForceURL(gbfs.FeedKeyAutoDiscovery, "https://gbfs.fordgobike.com/gbfs/gbfs.json"),
	)

	if err != nil {
		panic(err)
	}

	si, err := c.StationsStatus()

	if err != nil {
		panic(err)
	}

	for _, s := range si.Stations {
		fmt.Printf("Is renting: %t\n", s.IsRenting)
	}

	fmt.Printf("Last updated: %d\n", si.JSON.LastUpdated)
}
