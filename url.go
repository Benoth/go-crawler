package main

import (
    "strings"
    "time"
    "net/http"
    "github.com/parnurzeal/gorequest"
    "github.com/PuerkitoBio/goquery"
)

type Url struct {
    uri string
    visited bool
    response gorequest.Response
    body string
    duration time.Duration
}

func (url *Url) parseUrl() ([]error) {
    response, body, duration, requestError := getUrl(url.uri)

    url.response = response
    url.body = body
    url.duration = duration
    url.visited = true

    if len(requestError) > 0 {
        return requestError
    }

    responseCode := response.StatusCode
    location, locationExist := response.Header["Location"]

    if responseCode == 301 && locationExist {
        location := location[0]

        visitedURL[location] = Url{uri: location}

        return nil
    }

    // Parse
    doc, parserErr := goquery.NewDocumentFromResponse(response)
    if parserErr != nil {
        errors := []error{parserErr}
        return errors
    }

    // Find the review items
    // log.Println("Title for :", url.uri, "-", doc.Find("title").Text())
    doc.Find("a").Each(func(i int, element *goquery.Selection) {
        href, hrefExist := element.Attr("href")
        _, exists := visitedURL[href]
        // log.Println("Found", href)
        if hrefExist && ! exists && href != "#" {
            if strings.HasPrefix(href, baseURL) {
                visitedURL[href] = Url{uri: href}
            } else if strings.HasPrefix(href, "/") {
                visitedURL[href] = Url{uri: baseURL + strings.TrimLeft(href, "/")}
            } else {
                // log.Println("Wrong base url for", href)
            }
        }
    })

    return nil
}

func getUrl(url string) (gorequest.Response, string, time.Duration, []error) {
    timeStart := time.Now()

    response, body, err := request.
        Get(url).
        RedirectPolicy(func(req gorequest.Request, via []gorequest.Request) error {
            return http.ErrUseLastResponse
        }).End()

    return response, body, time.Since(timeStart), err
}