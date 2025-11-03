package main

import (
	"fmt"

	"github.com/pgfeng/annotation/pkg"
	"github.com/pgfeng/annotation/types"
)

func main() {

	route := "@Route     api/v1/users     GET"
	queryParam := "@QueryParam name=fieldName, required=true, default=默认值, summary=参数简介"

	annotation := pkg.NewAnnotation(&types.Route{}, route)
	if annotation == nil {
		fmt.Println("Failed to parse annotation")
		return
	}
	if qp, ok := annotation.Instance.(*types.Route); ok {
		fmt.Println("请求路径：", qp.Path)
		fmt.Println("请求方法：", qp.Method)
	}
	annotation = pkg.NewAnnotation(&types.QueryParam{}, queryParam)
	if annotation == nil {
		fmt.Println("Failed to parse annotation")
		return
	}
	fmt.Println(annotation)
	fmt.Println(annotation.Instance)

	// 访问具体类型字段，做指针断言（要求 pkg.Annotation.Instance 为接口类型）
	if qp, ok := annotation.Instance.(*types.QueryParam); ok {
		fmt.Println("参数名：", qp.Name)
		fmt.Println("是否必填：", qp.IsRequired)
		fmt.Println("默认值：", qp.Default)
		fmt.Println("简介：", qp.Summary)
	}
}
