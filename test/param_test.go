package test

import (
	"testing"

	ann "github.com/pgfeng/annotation"
	"github.com/pgfeng/annotation/pkg"
	"github.com/pgfeng/annotation/types"
)

func TestParams(t *testing.T) {
	params := map[string]pkg.Type{
		`@QueryParam name="userId", required=true, default="123", summary="The ID of the user"`:              &types.QueryParam{},
		`@PathParam name="postId", required=false, default="456", summary="The ID of the post"`:              &types.PathParam{},
		`@FormParam name="username", required=true, default="guest", summary="The username of the user"`:     &types.FormParam{},
		`@FileParam name="profilePic", required=false, default="", summary="The profile picture file"`:       &types.FileParam{},
		`@HeaderParam name="AuthToken", required=true, default="", summary="The authentication token"`:       &types.HeaderParam{},
		`@CookieParam name="sessionId", required=false, default="", summary="The session ID from cookies"`:   &types.CookieParam{},
		`@BodyParam name="userData", required=true, default="", summary="The user data in the request body"`: &types.BodyParam{},
	}
	for p, tt := range params {
		t.Logf("Testing annotation: %s", p)
		annotation := ann.ParseAnnotation(tt, p)
		if annotation == nil {
			t.Errorf("Failed to parse annotation: %s", p)
			continue
		}
		t.Logf("got expected result %+v", annotation.Instance)

	}
}
