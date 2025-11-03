package types

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"

	"github.com/pgfeng/annotation/pkg"
)

type Tags struct {
	tags []string
	hash string
}

func (s *Tags) GetName() string {
	return "Tags"
}

func (s *Tags) Copy() pkg.Type {
	if s == nil {
		return &Tags{}
	}
	return &Tags{
		tags: append([]string{}, s.tags...),
		hash: s.hash,
	}
}

func (s *Tags) InitValue(v string) {
	parts := strings.Fields(v)
	var finalParts []string
	for _, part := range parts {
		subParts := strings.Split(part, ",")
		for _, subPart := range subParts {
			trimmed := strings.TrimSpace(subPart)
			if trimmed != "" {
				finalParts = append(finalParts, trimmed)
			}
		}
	}
	sum := sha1.Sum([]byte(strings.Join(finalParts, ",")))
	s.hash = hex.EncodeToString(sum[:])
	s.tags = finalParts
}
