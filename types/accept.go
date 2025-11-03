package types

import (
	"strings"

	"github.com/pgfeng/annotation/pkg"
)

type Accept struct {
	MediaTypes []string
}

func (a *Accept) ToMap() map[string]string {
	return map[string]string{
		"Accept": strings.Join(a.MediaTypes, ", "),
	}
}

func (a *Accept) Copy() pkg.Type {
	newAccept := &Accept{
		MediaTypes: make([]string, len(a.MediaTypes)),
	}
	copy(newAccept.MediaTypes, a.MediaTypes)
	return newAccept
}

func (a *Accept) GetName() string {
	return "Accept"
}

// InitValue Parse：@Accept application/json, text/html
func (a *Accept) InitValue(v string) {
	var mediaTypes []string
	start := 0
	inQuotes := false
	for i, char := range v {
		switch char {
		case '"':
			inQuotes = !inQuotes
		case ',':
			if !inQuotes {
				part := v[start:i]
				mediaTypes = append(mediaTypes, trimSpaces(part))
				start = i + 1
			}
		}
	}
	if start < len(v) {
		part := v[start:]
		mediaTypes = append(mediaTypes, trimSpaces(part))
	}
	a.MediaTypes = mediaTypes
}
