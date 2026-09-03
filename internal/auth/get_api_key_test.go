package auth

import (
	"testing"
	"net/http"
	"fmt"
	"strings"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		key			string
		value		string
		expect		string
		expectErr	string
	}{
		{
			expectErr:	"no authorization header",
		},
		{
			key:		"Authorization",
			expectErr:	"no authorization header",
		},
		{
			key:		"Authorization",
			value:		"-",
			expectErr:	"malformed authorization header",
		},
		{
			key:		"Authorization",
			value:		"Bearer xxxxxx",
			expectErr:	"malformed authorization header",
		},
		{
			key:		"Authorization",
			value:		"ApiKey xxxxxx",
			expect:		"xxxxxx",
			expectErr:	"not expecting an error",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestGetAPIKey Case #%v:", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(tt.key, tt.value)

			result, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), tt.expectErr) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
				return
			}

			if result != tt.expect {
				t.Errorf("Unexpected: TestGetAPIKey:%s", result)
				return
			}
		})
	}
}
