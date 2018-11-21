package main

import (
    "log"
    // "fmt"
    "time"
    // "strings"
    // "sync"
    // "github.com/jawher/mow.cli"
    "github.com/parnurzeal/gorequest"
    // "github.com/PuerkitoBio/goquery"
    "gopkg.in/go-playground/pool.v3"
)

var visitedURL map[string]Url = make(map[string]Url)
var request = gorequest.New().Timeout(1000 * time.Millisecond)

// var baseURL string = "http://www.exemple.com/"
var baseURL string = "https://getpsalm.org/"

func main() {
    log.Println("Starting...")

    p := pool.NewLimited(4)
    batch := p.Batch()
    defer p.Close()


    url := Url{uri: baseURL}
    log.Println("Parsing " + url.uri)
    url.parseUrl()
    visitedURL[baseURL] = url

    go func() {
        // Copy the map to avoid overwriting
        work := make(map[string]Url)
        for k,v := range visitedURL {
            work[k] = v
        }
        for _, newUrl := range work {
            if newUrl.visited == false {
                batch.Queue(handleUrl(newUrl))
            }
        }

        // DO NOT FORGET THIS OR GOROUTINES WILL DEADLOCK
        // if calling Cancel() it calles QueueComplete() internally
        batch.QueueComplete()
    }()


    for crawl := range batch.Results() {
        if err := crawl.Error(); err != nil {
            // handle error
        }

        // use return value (url object)
        // log.Println(crawl.Value())
    }

    nop := 0
    yep := 0
    for _, newUrl := range visitedURL {
        if newUrl.visited == false {
            nop += 1
        } else {
            yep += 1
        }
    }

    // log.Println(visitedURL[baseURL])
    log.Println("Not done", nop)
    log.Println("Done ", yep)
    log.Println("Visited URLs", len(visitedURL))
}

func handleUrl(url Url) pool.WorkFunc  {
    return func(wu pool.WorkUnit) (interface{}, error) {
        if wu.IsCancelled() {
            // return values not used
            return nil, nil
        }

        log.Println("Parsing " + url.uri)

        url.parseUrl()

        visitedURL[url.uri] = url

        return url, nil // everything ok, send nil as 2nd parameter if no error
    }
}
