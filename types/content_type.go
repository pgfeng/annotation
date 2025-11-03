package types

import "github.com/pgfeng/annotation/pkg"

type ContentType struct {
	MediaType string
}

func (c *ContentType) ToMap() map[string]string {
	return map[string]string{
		"Content-Type": c.MediaType,
	}
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

// InitValue Parse：@ContentType application/json
func (c *ContentType) InitValue(v string) {
	c.MediaType = trimSpaces(v)
}
