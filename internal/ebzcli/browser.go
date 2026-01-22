package ebzcli

import (
	"errors"
	"fmt"
	"net/url"
	"os/exec"
	"regexp"
	"runtime"
)

var (
	ErrUnsupportPlatform = errors.New("unsupport platform")
	ErrUnableToStartApp  = errors.New("unable to start application")
	ErrBadlyFormattedURL = errors.New("badly formatted url")
)

var (
	localhostRegex = regexp.MustCompile(`^http://localhost:[0-9]{1,5}$`)
)

// raw url can only be http://localhost:<port number>
func getBrowserCommand(goos, rawUrl string) (string, []string, error) {
	if _, err := url.ParseRequestURI(rawUrl); err != nil {
		return "", nil, fmt.Errorf("%w:%w", ErrBadlyFormattedURL, err)
	}
	if !localhostRegex.MatchString(rawUrl) {
		return "", nil, fmt.Errorf("%w: not localhost", ErrBadlyFormattedURL)
	}
	switch goos {
	case "darwin":
		return "open", []string{rawUrl}, nil
	case "windows":
		return "rundll32", []string{"url.dll,FileProtocolHandler", rawUrl}, nil
	case "linux":
		return "xdg-open", []string{rawUrl}, nil
	default:
		return "", nil, fmt.Errorf("%w: Support for Windows, Linux, macOS only.", ErrUnsupportPlatform)
	}
}

func openBrowser(rawUrl string) error {
	goos := runtime.GOOS
	cmd, args, err := getBrowserCommand(goos, rawUrl)
	if err != nil {
		return err
	}
	if err := exec.Command(cmd, args...).Start(); err != nil {
		return fmt.Errorf("%w:%w", ErrUnableToStartApp, err)
	}
	return nil
}
