package main

import (
	"net/url"
)

func normalizeURL(urlstr string) (string, error) {
	structURL, err := url.Parse(urlstr)
	if err != nil {
		return "", err
	}

	var noTrailPath string
	if structURL.Path[len(structURL.Path)-1] == '/' {
		noTrailPath = structURL.Path[:len(structURL.Path)-1]
	} else {
		noTrailPath = structURL.Path
	}
	return structURL.Host + noTrailPath, nil
}
