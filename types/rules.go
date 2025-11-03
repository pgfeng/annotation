package types

import (
	"strings"

	"github.com/pgfeng/annotation/pkg"
)

type Rules struct {
	Rules []string
}

func (s *Rules) ToMap() map[string]string {
	return map[string]string{
		"rules": strings.Join(s.Rules, ","),
	}
}

func (s *Rules) GetName() string {
	return "Rules"
}

func (s *Rules) Copy() pkg.Type {
	if s == nil {
		return &Rules{}
	}
	return &Rules{
		Rules: append([]string{}, s.Rules...),
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
	s.Rules = finalParts
}
