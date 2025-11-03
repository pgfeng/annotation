package test

import (
	"testing"

	"github.com/pgfeng/annotation"
	"github.com/pgfeng/annotation/types"
)

func TestParseRouteAnnotations(t *testing.T) {
	// This is a placeholder to indicate that this file is part of the 'test' package.
	summarys := []string{
		"@Summary 获取用户列表 撒手",
		"@Summary 创建新用户",
		"@Summary 删除用户",
	}
	for _, summary := range summarys {
		annotation := annotation.ParseAnnotation(&types.Summary{}, summary)
		if annotation == nil {
			t.Errorf("Failed to parse annotation: %s", summary)
			continue
		}
		summaryInstance, ok := annotation.Instance.(*types.Summary)
		if !ok {
			t.Errorf("Failed to assert annotation instance to Summary type for annotation: %s", summary)
			continue
		} else {
			t.Logf("For input '%s', got expected result %+v", summary, summaryInstance)
		}
	}
	descriptions := []string{
		"@Description 该接口用于获取用户列表，支持分页和过滤功能。",
		"@Description 该接口用于创建一个新的用户，需要提供用户的基本信息。",
		"@Description 该接口用于删除指定的用户，操作不可逆，请谨慎使用。",
	}
	for _, description := range descriptions {
		annotation := annotation.ParseAnnotation(&types.Description{}, description)
		if annotation == nil {
			t.Errorf("Failed to parse annotation: %s", description)
			continue
		}
		descriptionInstance, ok := annotation.Instance.(*types.Description)
		if !ok {
			t.Errorf("Failed to assert annotation instance to Description type for annotation: %s", description)
			continue
		} else {
			t.Logf("For input '%s', got expected result %+v", description, descriptionInstance)
		}
	}
	tags := []string{
		"@Tags 用户管理, 用户操作",
		"@Tags 订单管理,订单查询",
		"@Tags 订单管理,订单查询",
	}
	for _, tag := range tags {
		annotation := annotation.ParseAnnotation(&types.Tags{}, tag)
		if annotation == nil {
			t.Errorf("Failed to parse annotation: %s", tag)
			continue
		}
		tagsInstance, ok := annotation.Instance.(*types.Tags)
		if !ok {
			t.Errorf("Failed to assert annotation instance to Tags type for annotation: %s", tag)
			continue
		} else {
			t.Logf("For input '%s', got expected result %+v", tag, tagsInstance)
		}
	}
	routes := []string{
		"@Route /api/v1/users GET",
		"@Route /api/v1/users POST",
		"@Route /api/v1/users OPTIONS",
		"@Route /api/{d+}/users OPTIONS",
		"@Route /api/v1/users",
		"@Route /api/v1/users   DELETE ",
		"@Route /api/v1/users   Put",
		"@Route /api/v1/users   patch",
		"@Route /api/v1/users   head",
		"@Route /api/v1/users",
	}

	for _, route := range routes {
		annotation := annotation.ParseAnnotation(&types.Route{}, route)
		if annotation == nil {
			t.Errorf("Failed to parse annotation: %s", route)
			continue
		}
		routeInstance, ok := annotation.Instance.(*types.Route)
		if !ok {
			t.Errorf("Failed to assert annotation instance to Route type for annotation: %s", route)
			continue
		} else {
			t.Logf("For input '%s', got expected result %+v", route, routeInstance)
		}
	}
}
