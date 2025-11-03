package types

import (
	"strconv"

	"github.com/pgfeng/annotation/pkg"
)

type FormParam struct {
	Name       string
	IsRequired bool
	Default    string
	Summary    string
	Type       string
}

func (p *FormParam) ToMap() map[string]string {
	return map[string]string{
		"name":        p.Name,
		"is_required": strconv.FormatBool(p.IsRequired),
		"default":     p.Default,
		"summary":     p.Summary,
		"type":        p.Type,
	}
}

func (p *FormParam) GetName() string {
	return "FormParam"
}
func (p *FormParam) Copy() pkg.Type {
	return &FormParam{
		Name:       p.Name,
		IsRequired: p.IsRequired,
		Default:    p.Default,
		Summary:    p.Summary,
		Type:       p.Type,
	}
}

// InitValue 解析：@FormParam name="名称", required=true, default="默认值", summary="参数简介"
// 支持值用双引号并且忽略引号内的逗号
func (p *FormParam) InitValue(v string) {
	// 先设置默认值
	param := param{
		ParamType: ParamTypeForm,
	}
	param.InitValue(v)
	p.Name = param.Name
	p.IsRequired = param.IsRequired
	p.Default = param.Default
	p.Summary = param.Summary
	p.Type = param.Type
}
