package types

import (
	"strconv"

	"github.com/pgfeng/annotation/pkg"
)

type CookieParam struct {
	Name       string
	IsRequired bool
	Default    string
	Summary    string
	Type       string
}

func (p *CookieParam) ToMap() map[string]string {
	return map[string]string{
		"name":        p.Name,
		"is_required": strconv.FormatBool(p.IsRequired),
		"default":     p.Default,
		"summary":     p.Summary,
		"type":        p.Type,
	}
}
func (p *CookieParam) GetName() string {
	return "CookieParam"
}
func (p *CookieParam) Copy() pkg.Type {
	return &CookieParam{
		Name:       p.Name,
		IsRequired: p.IsRequired,
		Default:    p.Default,
		Summary:    p.Summary,
		Type:       p.Type,
	}
}

// InitValue Parse：@CookieParam name="名称", required=true, default="默认值", summary="参数简介"
func (p *CookieParam) InitValue(v string) {
	// Set default values first
	param := param{
		ParamType: ParamTypeHeader,
	}
	param.InitValue(v)
	p.Name = param.Name
	p.IsRequired = param.IsRequired
	p.Default = param.Default
	p.Summary = param.Summary
	p.Type = param.Type
}
