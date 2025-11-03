package types

import "github.com/pgfeng/annotation/pkg"

type Accept struct {
	MediaTypes []string
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

// InitValue 解析：@Accept application/json, text/html
func (a *Accept) InitValue(v string) {
	// 按逗号分割并去除空格
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
	// 添加最后一个部分
	if start < len(v) {
		part := v[start:]
		mediaTypes = append(mediaTypes, trimSpaces(part))
	}
	a.MediaTypes = mediaTypes
}
