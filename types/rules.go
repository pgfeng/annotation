package types

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"

	"github.com/pgfeng/annotation/pkg"
)

type Rules struct {
	tags []string
	hash string
}

func (s *Rules) GetName() string {
	return "Rules"
}

func (s *Rules) Copy() pkg.Type {
	if s == nil {
		return &Rules{}
	}
	return &Rules{
		tags: append([]string{}, s.tags...),
		hash: s.hash,
	}
}

func (s *Rules) InitValue(v string) {
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
