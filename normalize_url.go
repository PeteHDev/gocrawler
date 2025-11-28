package main

import (
	"net/url"
	"strings"
)

func normalizeURL(urlstr string) (string, error) {
	structURL, err := url.Parse(urlstr)
	if err != nil {
		return "", err
	}

	noTrailPath := strings.TrimSuffix(structURL.Path, "/")
	return structURL.Host + noTrailPath, nil
}
