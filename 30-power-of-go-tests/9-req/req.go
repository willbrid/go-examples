package req

import (
	"errors"
	"net/http"
)

var ErrRateLimit = errors.New("rate limit")

func Request(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		return ErrRateLimit
	}

	return nil
}
