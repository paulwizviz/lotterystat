package ebzcli

import (
	"errors"
	"reflect"
	"testing"
)

var (
	successfulCases = func(t *testing.T) {
		testcases := []struct {
			name     string
			goos     string
			rawUrl   string
			wantCmd  string
			wantArgs []string
			wantErr  error
		}{
			{
				name:     "macOS",
				goos:     "darwin",
				rawUrl:   "http://localhost:8080",
				wantCmd:  "open",
				wantArgs: []string{"http://localhost:8080"},
				wantErr:  nil,
			},
			{
				name:     "Windows",
				goos:     "windows",
				rawUrl:   "http://localhost:8080",
				wantCmd:  "rundll32",
				wantArgs: []string{"url.dll,FileProtocolHandler", "http://localhost:8080"},
				wantErr:  nil,
			},
			{
				name:     "Linux",
				goos:     "linux",
				rawUrl:   "http://localhost:8080",
				wantCmd:  "xdg-open",
				wantArgs: []string{"http://localhost:8080"},
				wantErr:  nil,
			},
		}

		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				gotCmd, gotArgs, gotErr := getBrowserCommand(tc.goos, tc.rawUrl)
				if !errors.Is(gotErr, tc.wantErr) {
					t.Fatalf("Unmatch error. Want: %v Got: %v", tc.wantErr, gotErr)
				}
				if tc.wantCmd != gotCmd {
					t.Fatalf("Dissimilar commands. Want: %v Got: %v", tc.wantCmd, gotCmd)
				}
				if !reflect.DeepEqual(tc.wantArgs, gotArgs) {
					t.Fatalf("Dissimilar arguments. Want: %v Got: %v", tc.wantArgs, gotArgs)
				}
			})
		}
	}

	failedCases = func(t *testing.T) {
		testcases := []struct {
			name    string
			goos    string
			rawUrl  string
			wantErr error
		}{
			{
				name:    "Invalid URL incomplete schema",
				rawUrl:  "url",
				wantErr: ErrBadlyFormattedURL,
			},
			{
				name:    "Not localhost format",
				rawUrl:  "http://localhos:1234",
				wantErr: ErrBadlyFormattedURL,
			},
			{
				name:    "Missing port number",
				rawUrl:  "http://localhost",
				wantErr: ErrBadlyFormattedURL,
			},
			{
				name:    "Unsupported platform",
				goos:    "win32",
				rawUrl:  "http://localhost:1234",
				wantErr: ErrUnsupportPlatform,
			},
		}

		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				_, _, gotErr := getBrowserCommand(tc.goos, tc.rawUrl)
				if !errors.Is(gotErr, tc.wantErr) {
					t.Fatalf("URL error. Want: %v Got: %v", tc.wantErr, gotErr)
				}
			})
		}
	}
)

func TestGetBrowserCommand(t *testing.T) {
	t.Run("Success", successfulCases)
	t.Run("Failed", failedCases)
}
