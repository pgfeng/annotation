package types

import "github.com/pgfeng/annotation/pkg"

type Summary struct {
	Text string
}

func (s *Summary) GetName() string {
	return "Summary"
}

// InitValue 解析：@Summary 简要描述内容
func (s *Summary) InitValue(v string) {
	s.Text = v
}

func (s *Summary) Copy() pkg.Type {
	if s == nil {
		return &Summary{}
	}
	return &Summary{
		Text: s.Text,
	}
}
