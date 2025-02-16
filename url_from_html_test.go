package main

import (
	"reflect"
	"testing"
)

func TestUrlFromHtml(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		inputBody     string
		expected	  []string
	}{
		{
			name:     "absolute and relative URLs",
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
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},

		{
			name:     "Deep absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<div>
					<a href="/path/one">
						<span>Boot.dev</span>
					</a>
					<div>
						<a href="https://other.com/path/one">
							<span>Boot.dev</span>
						</a>
					</div>
				</div>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "No URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<div>
					
				</div>
			</body>
		</html>
		`,
			expected: []string{},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual == nil {
				actual = []string{}
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}