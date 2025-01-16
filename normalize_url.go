package main

import (
	"errors"
	"net/url"
	"strings"
)

var ErrorEmptyURL = errors.New("URL is Empty, please provide a URL")
var ErrorEmptyHost = errors.New("please provide a correct URL")

func normalizeURL(URL string) (string, error){
	if URL == ""{
		return "", ErrorEmptyURL
	}
	parsedURL, err := url.Parse(URL)
	if err != nil {
		return "", err
	}
	trimmedPath := strings.TrimSuffix(parsedURL.Path, "/")
	normalizedURL := parsedURL.Host + trimmedPath

	return normalizedURL, nil
}