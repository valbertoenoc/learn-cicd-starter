package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPI(t *testing.T) {
	tests := []struct {
		key       string
		value     string
		expect    string
		expectErr string
	}{
		{
			key:       "Authorization",
			value:     "-",
			expectErr: "malformed authorization header",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestGetAPIKey case %d", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(tt.key, tt.value)

			output, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), tt.expectErr) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
				return
			}

			if output != tt.expect {
				t.Errorf("Unexpected: TestGetAPIKey:%s", output)
				return
			}

		})

	}

	fmt.Printf("tests: %v", tests)
}
