# Go Akamai-open/edgegrid client

[![GoDoc](https://godoc.org/github.com/chonthu/go-edgegrid?status.svg)](https://godoc.org/github.com/chonthu/go-edgegrid)

## Installation

    go get github.com/chonthu/go-edgegrid
    
## Usage

    package main

    import (
        "log"

        "github.com/chonthu/go-edgegrid"
    )

    func main() {
        api := edgegrid.NewFromIni(".edgerc")
        resp, err := api.Send("POST", "/ccu/v3/invalidate/url/production", "{ \"objects\" : [ \"http://example.com\" ]}")
        if err != nil {
            log.Println(err)
        }

        log.Panicln(resp)
    }

## Author

Nithin Meppurathu - Twitter: @meppurathu
