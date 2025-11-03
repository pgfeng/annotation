package types

import (
	"strconv"

	"github.com/pgfeng/annotation/pkg"
)

type PathParam struct {
	Name       string
	IsRequired bool
	Default    string
	Summary    string
	Type       string
}

func (p *PathParam) ToMap() map[string]string {
	return map[string]string{
		"name":        p.Name,
		"is_required": strconv.FormatBool(p.IsRequired),
		"default":     p.Default,
		"summary":     p.Summary,
		"type":        p.Type,
	}
}

func (p *PathParam) GetName() string {
	return "PathParam"
}
func (p *PathParam) Copy() pkg.Type {
	return &PathParam{
		Name:       p.Name,
		IsRequired: p.IsRequired,
		Default:    p.Default,
		Summary:    p.Summary,
		Type:       p.Type,
	}
}

// InitValue 解析：@PathParam name="名称", required=true, default="默认值", summary="参数简介"
func (p *PathParam) InitValue(v string) {
	// 先设置默认值
	param := param{
		ParamType: ParamTypePath,
	}
	param.InitValue(v)
	p.Name = param.Name
	p.IsRequired = param.IsRequired
	p.Default = param.Default
	p.Summary = param.Summary
	p.Type = param.Type
}
