package runtime

import (
	"fmt"

	"github.com/dchest/uniuri"
)

const (
	AnnotationDomain = "ai.tensorchord.domain"
)

const (
	// stdLen is a standard length of uniuri string to achive ~95 bits of entropy.
	stdLen = 16
)

// StdChars is a set of standard characters allowed in uniuri string.
var stdChars = []byte("abcdefghijklmnopqrstuvwxyz0123456789")

func makeDomain(name, baseDomain string) (string, error) {
	if baseDomain == "" {
		return "", fmt.Errorf("baseDomain is required")
	}

	if name == "" {
		return "", fmt.Errorf("name is required")
	}

	hash := uniuri.NewLenChars(stdLen, stdChars)

	return fmt.Sprintf("%s-%s.%s",
		name, hash, baseDomain), nil
}