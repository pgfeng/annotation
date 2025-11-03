package types

import (
	"strings"

	"github.com/pgfeng/annotation/pkg"
)

type Method string

const (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	DELETE  Method = "DELETE"
	PATCH   Method = "PATCH"
	OPTIONS Method = "OPTIONS"
	HEAD    Method = "HEAD"
)

func (m Method) String() string {
	return string(m)
}

// Route 注解类型，表示路由信息
type Route struct {
	Path   string
	Method Method
}

// GetName GetType 返回注解类型
func (r *Route) GetName() string {
	return "Route"
}
func (r *Route) Copy() pkg.Type {
	return &Route{
		Path:   r.Path,
		Method: r.Method,
	}
}

// InitValue 解析：@route /api/v1/resource GET
// 否则将全部作为 Path 并把 Method 设为 GET（默认）
func (r *Route) InitValue(v string) {
	parts := strings.Fields(strings.TrimSpace(v))
	//fmt.Println(parts)
	if len(parts) == 0 {
		r.Path = ""
		r.Method = "GET"
		return
	}

	// 如果最后一个 token 看起来像 HTTP 方法，则取为 Method
	last := strings.ToUpper(parts[len(parts)-1])
	if isHTTPMethod(last) && len(parts) > 1 {
		r.Method = Method(last)
		r.Path = strings.Join(parts[:len(parts)-1], " ")
		return
	}

	// 否则全部作为 Path，默认 Method 为 GET
	r.Path = strings.Join(parts, " ")
	r.Method = "GET"
	//fmt.Println("RouteInit", r.Path)
}

func isHTTPMethod(s string) bool {
	switch s {
	case GET.String(), POST.String(), PUT.String(), DELETE.String(), PATCH.String(), OPTIONS.String(), HEAD.String():
		return true
	default:
		return false
	}
}
