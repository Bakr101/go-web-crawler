package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T){
	tests := map[string]struct{
		inputURL		string
		inputBody		string
		expectedURLs	[]string
		errorContains 	string	
	}{
		"absolute and relative URLs" : {
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<html>
					<body>
						<a href="/path/one">
							<span>Boot.dev</span>
						</a>
						<a href="https://other.com/path/one">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
				`,
			expectedURLs: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		"relative URL": {
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<html>
					<body>
						<a href="/path/one">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
				`,
			expectedURLs: []string{"https://blog.boot.dev/path/one"},
		},
		"absolute URL" :{
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<html>
					<body>
						<a href="https://blog.boot.dev">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
				`,
			expectedURLs: []string{"https://blog.boot.dev"},
		},
		"no href":{
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<html>
					<body>
						<a>
							<span>Boot.dev></span>
						</a>
					</body>
				</html>
				`,
			expectedURLs: nil,
		},
		"bad HTML":{
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<html body>
					<a href="path/one">
						<span>Boot.dev></span>
					</a>
				</html body>
				`,
			expectedURLs: []string{"https://blog.boot.dev/path/one"},
		},
		"Invalid href URl":{
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<html>
					<body>
						<a href=":\\invalidURL">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
				`,
			expectedURLs: nil,
		},
		"handle invalid base URL":{
			inputURL: `:\\invalidBaseURL`,
			inputBody: `
				<html>
					<body>
						<a href="/path">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
				`,
			expectedURLs:      nil,
			errorContains: "couldn't parse base URL",
		},
	}

	for name, tc := range tests{
		t.Run(name, func(t *testing.T) {
			urls, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil  && !strings.Contains(err.Error(), tc.errorContains){
				t.Errorf("Test: %s FAIL unexpected error: %v", name, err)
				return
			}else if err != nil && tc.errorContains == ""{
				t.Errorf("Test %v FAIL: unexpected error: %v", name, err)
				return
			}else if err == nil && tc.errorContains != ""{
				t.Errorf("Test %v FAIL: expected error containing '%v', got none.", name, tc.errorContains)
				return
			}
			if !reflect.DeepEqual(urls, tc.expectedURLs){
				t.Errorf("Test: %s FAIL expected urls: %v, got: %v", name, tc.expectedURLs, urls)
				return
			}
		})
	}
}