package main

import (
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T){
	tests := map[string]struct{
		url	string
		expectedURL	string
		errorContains string
	}{
		"empty url": {url: "", expectedURL: "", errorContains: "URL is Empty, please provide a URL"},
		"remove scheme": {url: "https://blog.boot.dev/path", expectedURL: "blog.boot.dev/path" },
		"remove last slash":{url: "http://blog.boot.dev/path/", expectedURL: "blog.boot.dev/path" },
		"remove both": {url: "https://blog.boot.dev/path/", expectedURL: "blog.boot.dev/path" },
		"invalid URL": {url: ":\\invalidURL", expectedURL: "", errorContains: "please provide a correct URL"},
		"lowercase capital letters": {url: "https://BLOG.boot.dev/PATH", expectedURL: "blog.boot.dev/path"},
	}

	for name, tc := range tests{
		t.Run(name, func(t *testing.T) {
			actual, err:= normalizeURL(tc.url)
			if err != nil  && !strings.Contains(err.Error(), tc.errorContains){
				t.Errorf("Test %s FAIL: unexpected error: %v", name, err)
			}
			if actual != tc.expectedURL{
				t.Errorf("Test %s FAIL: expected URL: %s, actual: %s", name, tc.expectedURL, actual)
			}
		})
	}
}

