package main

import (
	"net/url"
	"strings"
)

func normalizeURL(urlStr string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	normU := u.Host + u.Path
	normU = strings.ToLower(normU)
	normU = strings.TrimSuffix(normU, "/")
	return normU, nil
}