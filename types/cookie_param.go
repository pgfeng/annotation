package types

import "github.com/pgfeng/annotation/pkg"

type CookieParam struct {
	Name       string
	IsRequired bool
	Default    string
	Summary    string
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
	}
}

// InitValue 解析：@CookieParam name="名称", required=true, default="默认值", summary="参数简介"
func (p *CookieParam) InitValue(v string) {
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
