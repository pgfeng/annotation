# Go Annotation Helper

A small helper library to parse and extract annotations / structured comments from Go source code.  
Designed to make it easy to read custom annotations (e.g. `@route`, `@accept`) and convert them into Go structures for generation, validation or runtime configuration.

## Features

- Parse Go source files and directories for annotation-style comments.
- Convert common annotation patterns into typed Go structures.
- Simple API for parsing from source string, file, or directory.
- Includes lightweight test helpers and examples.

## Installation

Install the package (replace with the actual import path if different):

```sh
go install github.com/pgfeng/annotation@latest
```

## Usage
Use the package build gin web framework route annotation:

### Example Annotation in Handler File
```go
// @Route /account/info get
// @Summary Get Current User Info
// @Description Get the information of the currently logged-in user.
// @FormParam name="username" required=true summary="Username"
// @FormParam name="password" required=true summary="Password"
// @Tags Account User
// @Rules AuthRequired
func Info(c *gin.Context) {
    c.JSON(200, gin.H{
    "status": false,
    "msg":    "Not logged in!",
    })
}
```

### Generating Routes File
```go
import (
	"context"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"log/slog"
	"os"
	"reflect"

	"github.com/pgfeng/annotation"
	"github.com/pgfeng/annotation/pkg"
	"github.com/pgfeng/annotation/types"
)

func getPackageFunctions() *pkg.PackageFunctions {
	return annotation.LoadPackageByPath("./backend/app/...", []pkg.Type{
		&types.Route{},
		&types.QueryParam{},
		&types.FormParam{},
		&types.FileParam{},
		&types.Summary{},
		&types.Description{},
		&types.PathParam{},
		&types.HeaderParam{},
		&types.BodyParam{},
		&types.CookieParam{},
		&types.Tags{},
		&types.Rules{},
		&types.Accept{},
		&types.ContentType{},
	})
}

var pkgFunks = getPackageFunctions()

func GenerateRouteFile() {
	funks := pkgFunks.Filter(&types.Route{})
	if funks == nil || len(*funks) == 0 {
		return
	}
	fset := token.NewFileSet()
	file := &ast.File{
		Name:  ast.NewIdent("route"),
		Decls: []ast.Decl{},
	}
	importSpecs := funks.GetImportSpecs()
	// Add Gin import
	importSpecs = append(importSpecs, &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: `"github.com/gin-gonic/gin"`,
		},
	})
	file.Decls = append(file.Decls, &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: importSpecs,
	})
	var funcBodyList []ast.Stmt
	for _, pf := range *funks {
		for i := 0; i < len(pf.Annotations); i++ {
			slog.Log(context.Background(), slog.LevelInfo, "Matched", *pf.Annotations[i])
			an := pf.Annotations[i]
			if an == nil {
				continue
			}
			anType := reflect.TypeOf(&types.Route{})
			if anType == nil {
				continue
			}
			if an.Instance == nil || reflect.TypeOf(an.Instance) != anType {
				continue
			}
			// Add gin route initialization statement
			routeInstance := an.Instance.(*types.Route)
			routePath := routeInstance.Path
			routeMethod := routeInstance.Method
			handlerName := pf.FunctionName
			// construct: router.GET("/path", handlerName)
			exprStmt := &ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   ast.NewIdent("router"),
						Sel: ast.NewIdent(string(routeMethod)),
					},
					Args: []ast.Expr{
						&ast.BasicLit{
							Kind:  token.STRING,
							Value: fmt.Sprintf(`"%s"`, routePath),
						},
						&ast.Ident{
							Name: handlerName,
						},
					},
				},
			}
			funcBodyList = append(funcBodyList, exprStmt)
		}
	}
	mainFunc := &ast.FuncDecl{
		Name: ast.NewIdent("InitRoutes"),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{ast.NewIdent("router")},
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("*gin"),
							Sel: ast.NewIdent("RouterGroup"),
						},
					},
				},
			},
			Results: nil,
		},
		Body: &ast.BlockStmt{
			List: funcBodyList,
		},
	}
	file.Decls = append(file.Decls, mainFunc)
	// Write to routes_gen.go
	out, err := os.Create("./backend/route/routes_gen.go")
	if err != nil {
		panic(err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	if err := printer.Fprint(out, fset, file); err != nil {
		panic(err)
	}
	slog.Log(context.Background(), slog.LevelInfo, "generated routes_gen.go")
}

```