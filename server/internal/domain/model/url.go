package domain

import (
	"fmt"

	"github.com/cespare/xxhash/v2"

	jerrors "github.com/sota0121/janusly/server/internal/domain/errors"
)

type ShortenedURL struct {
	OriginalURL string `json:"original_url"`
	TinyURL     string `json:"tiny_url"`
	Hash        string `json:"hash"`
	Salt        string `json:"salt"`
}

var ErrOriginalURLIsRequired = &jerrors.JanuslyError{
	Code:    "URL_ORIGINAL_URL_IS_REQUIRED",
	Message: "Original URL is required",
}

// GenerateTinyURL generates tiny URL from original URL.
func (s *ShortenedURL) GenerateTinyURL() error {
	// Original URL is required.
	if s.OriginalURL == "" {
		return ErrOriginalURLIsRequired
	}

	// Calculate hash value from original URL.
	srcURL := fmt.Sprintf("%s%s", s.OriginalURL, s.Salt)
	hash := xxhash.Sum64([]byte(srcURL))
	s.Hash = fmt.Sprintf("%x", hash)

	// Currently, we use hash value as tiny URL.
	s.TinyURL = s.Hash
	return nil
}
