package types

import (
	"strconv"
	"strings"

	"github.com/pgfeng/annotation/pkg"
)

type param struct {
	Name       string
	IsRequired bool
	Default    string
	Summary    string
	Type       string
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
func (p *param) New(paramType ParamType) {
	p.ParamType = ParseParamType(paramType)
}
func (p *param) ToMap() map[string]string {
	return map[string]string{
		"name":        p.Name,
		"is_required": strconv.FormatBool(p.IsRequired),
		"default":     p.Default,
		"summary":     p.Summary,
		"paramType":   string(p.ParamType),
		"type":        p.Type,
	}
}

func (p *param) GetName() string {
	return "Param"
}
func (p *param) Copy() pkg.Type {
	return &param{
		Name:       p.Name,
		IsRequired: p.IsRequired,
		Default:    p.Default,
		Summary:    p.Summary,
		ParamType:  p.ParamType,
		Type:       p.Type,
	}
}

// InitValue 解析：@Param name="名称", type="query", required=true, default="默认值", summary="参数简介"
func (p *param) InitValue(v string) {
	// 先设置默认值
	p.IsRequired = false
	p.Default = ""
	p.Summary = ""
	p.Type = "string"
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
	if pType, ok := m["paramType"]; ok {
		p.ParamType = ParseParamType(ParamType(pType))
	}
	if t, ok := m["type"]; ok {
		p.Type = t
	}
}
func (ps *param) getInstance() {
	switch ps.ParamType {
	case ParamTypePath:
		p := &PathParam{}
		p.Name = ps.Name
		p.IsRequired = ps.IsRequired
		p.Default = ps.Default
		p.Summary = ps.Summary
		return
	case ParamTypeQuery:
		p := &QueryParam{}
		p.Name = ps.Name
		p.IsRequired = ps.IsRequired
		p.Default = ps.Default
		p.Summary = ps.Summary
		return
	case ParamTypeHeader:
		p := &HeaderParam{}
		p.Name = ps.Name
		p.IsRequired = ps.IsRequired
		p.Default = ps.Default
		p.Summary = ps.Summary
		return
	case ParamTypeCookie:
		p := &CookieParam{}
		p.Name = ps.Name
		p.IsRequired = ps.IsRequired
		p.Default = ps.Default
		p.Summary = ps.Summary
		return
	case ParamTypeForm:
		p := &FormParam{}
		p.Name = ps.Name
		p.IsRequired = ps.IsRequired
		p.Default = ps.Default
		p.Summary = ps.Summary
		return
	case ParamTypeFile:
		p := &FileParam{}
		p.Name = ps.Name
		p.IsRequired = ps.IsRequired
		p.Default = ps.Default
		p.Summary = ps.Summary
		return
	case ParamTypeBody:
		p := &BodyParam{}
		p.Name = ps.Name
		p.IsRequired = ps.IsRequired
		p.Default = ps.Default
		p.Summary = ps.Summary
		return
	default:
		return
	}
}
