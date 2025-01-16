package main

import (
	"errors"
	"testing"
)

func TestNormalizeURL(t *testing.T){
	tests := map[string]struct{
		url	string
		expectedURL	string
		expectedError error
	}{
		"empty url": {url: "", expectedURL: "", expectedError: ErrorEmptyURL},
		"remove scheme": {url: "https://blog.boot.dev/path", expectedURL: "blog.boot.dev/path", expectedError: nil},
		"remove last slash":{url: "http://blog.boot.dev/path/", expectedURL: "blog.boot.dev/path", expectedError: nil},
		"remove both": {url: "https://blog.boot.dev/path/", expectedURL: "blog.boot.dev/path", expectedError: nil},
		"Invalid URL": {url: "https:path", expectedURL: "", expectedError: nil},
	}

	for name, tc := range tests{
		t.Run(name, func(t *testing.T) {
			actual, err:= normalizeURL(tc.url)
			if err != nil  && !errors.Is(err, tc.expectedError){
				t.Errorf("Test %s FAIL: unexpected error: %v", name, err)
			}
			if actual != tc.expectedURL{
				t.Errorf("Test %s FAIL: expected URL: %s, actual: %s", name, tc.expectedURL, actual)
			}
		})
	}
}