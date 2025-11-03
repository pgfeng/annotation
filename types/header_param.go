package types

import "github.com/pgfeng/annotation/pkg"

type HeaderParam struct {
	Name       string
	IsRequired bool
	Default    string
	Summary    string
}

func (p *HeaderParam) GetName() string {
	return "HeaderParam"
}
func (p *HeaderParam) Copy() pkg.Type {
	return &HeaderParam{
		Name:       p.Name,
		IsRequired: p.IsRequired,
		Default:    p.Default,
		Summary:    p.Summary,
	}
}

// InitValue 解析：@QueryParam name="名称", required=true, default="默认值", summary="参数简介"
// 支持值用双引号并且忽略引号内的逗号
func (p *HeaderParam) InitValue(v string) {
	// 先设置默认值
	param := Param{
		ParamType: ParamTypeHeader,
	}
	param.InitValue(v)
	p.Name = param.Name
	p.IsRequired = param.IsRequired
	p.Default = param.Default
	p.Summary = param.Summary
}
