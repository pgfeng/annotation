package types

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"

	"github.com/pgfeng/annotation/pkg"
)

type Tags struct {
	Tags   []string
	Hashes []string // each tag's hash
	Hash   string   // overall hash
}

func (s *Tags) ToMap() map[string]string {
	return map[string]string{
		"tags":   strings.Join(s.Tags, ","),
		"hashes": strings.Join(s.Hashes, ","),
		"hash":   s.Hash,
	}
}

func (s *Tags) GetName() string {
	return "Tags"
}

func (s *Tags) Copy() pkg.Type {
	if s == nil {
		return &Tags{}
	}
	return &Tags{
		Tags:   append([]string{}, s.Tags...),
		Hashes: append([]string{}, s.Hashes...),
		Hash:   s.Hash,
	}
}

func (s *Tags) InitValue(v string) {
	parts := strings.Fields(v)
	var finalParts []string
	var hashes []string
	for _, part := range parts {
		subParts := strings.Split(part, ",")
		for _, subPart := range subParts {
			trimmed := strings.TrimSpace(subPart)
			if trimmed != "" {
				sum := sha1.Sum([]byte(trimmed))
				hashes = append(hashes, hex.EncodeToString(sum[:]))
				finalParts = append(finalParts, trimmed)
			}
		}
	}
	sum := sha1.Sum([]byte(strings.Join(finalParts, ",")))
	s.Hash = hex.EncodeToString(sum[:])
	s.Hashes = hashes
	s.Tags = finalParts
}
