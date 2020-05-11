# GBFS Client

A Golang client implementing the [GBFS](https://github.com/NABSA/gbfs) (**G**eneral **B**ikeshare **F**eed **S**pecification)

## Usage

```shell script
go get github.com/Eraac/gbfs
```

Basic example
```go
package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/Eraac/gbfs"
    "github.com/Eraac/gbfs/spec/v2.0"
)

c, err := gbfs.NewHTTPClient(
    gbfs.HTTPOptionBaseURL("https://gbfs.fordgobike.com/gbfs"), // required
    gbfs.HTTPOptionLanguage("en"), // optional if the provider doesn't don't specify the language in the URL
    gbfs.HTTPOptionClient(http.Client{Timeout: 10 * time.Second}), // optional, set a custom http client
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

fmt.Printf("Last updated: %s\n", ss.LastUpdated.ToTime().String())
```

You can use your own structure by implementing the `Feed` interface. Useful when the provider doesn't respect the specification.
```go
type Feed interface {
    FeedKey() string
    IsExpired() bool
}
``` 
