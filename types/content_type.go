package types

import "github.com/pgfeng/annotation/pkg"

type ContentType struct {
	MediaType string
}

func (c *ContentType) Copy() pkg.Type {
	newContentType := &ContentType{
		MediaType: c.MediaType,
	}
	return newContentType
}
func (c *ContentType) GetName() string {
	return "ContentType"
}

// InitValue 解析：@ContentType application/json
func (c *ContentType) InitValue(v string) {
	c.MediaType = trimSpaces(v)
}
