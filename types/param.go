package types

import (
	"strconv"
	"strings"

	"github.com/pgfeng/annotation/pkg"
)

type Param struct {
	Name       string
	IsRequired bool
	Default    string
	Summary    string
	ParamType  ParamType
}
type ParamType string

const (
	ParamTypePath   ParamType = "path"
	ParamTypeQuery  ParamType = "query"
	ParamTypeHeader ParamType = "header"
	ParamTypeCookie ParamType = "cookie"
	ParamTypeBody   ParamType = "body"
	ParamTypeForm   ParamType = "form"
	ParamTypeFile   ParamType = "file"
)

func ParseParamType(s ParamType) ParamType {
	switch strings.ToLower(strings.TrimSpace(string(s))) {
	case string(ParamTypePath):
		return ParamTypePath
	case string(ParamTypeQuery):
		return ParamTypeQuery
	case string(ParamTypeHeader):
		return ParamTypeHeader
	case string(ParamTypeCookie):
		return ParamTypeCookie
	case string(ParamTypeBody):
		return ParamTypeBody
	case string(ParamTypeForm):
		return ParamTypeForm
	case string(ParamTypeFile):
		return ParamTypeFile
	default:
		return ParamTypeQuery
	}
}
func (p *Param) New(paramType ParamType) {
	p.ParamType = ParseParamType(paramType)
}
func (p *Param) GetName() string {
	return "Param"
}
func (p *Param) Copy() pkg.Type {
	return &Param{
		Name:       p.Name,
		IsRequired: p.IsRequired,
		Default:    p.Default,
		Summary:    p.Summary,
	}
}

// InitValue 解析：@Param name="名称", type="query", required=true, default="默认值", summary="参数简介"
func (p *Param) InitValue(v string) {
	// 先设置默认值
	p.IsRequired = false
	p.Default = ""
	p.Summary = ""
	if p.ParamType == "" {
		p.ParamType = ParseParamType(ParamType(v))
	}
	m := pkg.ParseKeyValues(v)
	if name, ok := m["name"]; ok {
		p.Name = name
	}
	if def, ok := m["default"]; ok {
		p.Default = def
	}
	if sum, ok := m["summary"]; ok {
		p.Summary = sum
	}
	if req, ok := m["required"]; ok {
		if b, err := strconv.ParseBool(req); err == nil {
			p.IsRequired = b
		}
	}
	if pType, ok := m["type"]; ok {
		p.ParamType = ParseParamType(ParamType(pType))
	}
}
func (p *Param) getInstance() {
	switch p.ParamType {
	case ParamTypePath:
		p := &PathParam{}
		p.Name = p.Name
		p.IsRequired = p.IsRequired
		p.Default = p.Default
		p.Summary = p.Summary
		return
	case ParamTypeQuery:
		p := &QueryParam{}
		p.Name = p.Name
		p.IsRequired = p.IsRequired
		p.Default = p.Default
		p.Summary = p.Summary
		return
	case ParamTypeHeader:
		p := &HeaderParam{}
		p.Name = p.Name
		p.IsRequired = p.IsRequired
		p.Default = p.Default
		p.Summary = p.Summary
		return
	case ParamTypeCookie:
		p := &CookieParam{}
		p.Name = p.Name
		p.IsRequired = p.IsRequired
		p.Default = p.Default
		p.Summary = p.Summary
		return
	case ParamTypeForm:
		p := &FormParam{}
		p.Name = p.Name
		p.IsRequired = p.IsRequired
		p.Default = p.Default
		p.Summary = p.Summary
		return
	case ParamTypeFile:
		p := &FileParam{}
		p.Name = p.Name
		p.IsRequired = p.IsRequired
		p.Default = p.Default
		p.Summary = p.Summary
		return
	case ParamTypeBody:
		p := &BodyParam{}
		p.Name = p.Name
		p.IsRequired = p.IsRequired
		p.Default = p.Default
		p.Summary = p.Summary
		return
	default:
		return
	}
}
