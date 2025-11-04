package types

import "github.com/pgfeng/annotation/pkg"

type Group struct {
	Group string
}

func (c *Group) ToMap() map[string]string {
	return map[string]string{
		"Content-Type": c.Group,
	}
}

func (c *Group) Copy() pkg.Type {
	newGroup := &Group{
		Group: c.Group,
	}
	return newGroup
}
func (c *Group) GetName() string {
	return "Group"
}

// InitValue Parse：@Group /api
func (c *Group) InitValue(v string) {
	c.Group = trimSpaces(v)
}
