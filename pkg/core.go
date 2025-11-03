package pkg

import (
	"go/ast"
	"go/token"
	"regexp"
	"strings"

	"golang.org/x/tools/go/packages"
)

// LoadPackageByPath Load package and parse annotations from function/method comments.
func LoadPackageByPath(s string, t []Type) *PackageFunctions {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes,
		Fset: token.NewFileSet(),
	}
	packageLists, err := packages.Load(cfg, s)
	if err != nil || len(packageLists) == 0 {
		return nil
	}
	var result PackageFunctions
	for _, pkg := range packageLists {
		pkgPath := pkg.PkgPath
		for _, f := range pkg.Syntax {
			// range over declarations to find function declarations
			for _, decl := range f.Decls {
				fd, ok := decl.(*ast.FuncDecl)
				if !ok {
					continue
				}
				// 构造 methodKey
				//
				var methodKey string
				methodKey = pkg.Name + "." + fd.Name.Name
				// 收集注释文本（使用 FuncDecl.Doc）
				var commentText string
				if fd.Doc != nil {
					parts := make([]string, 0, len(fd.Doc.List))
					for _, cg := range fd.Doc.List {
						text := strings.TrimSpace(cg.Text)
						if strings.HasPrefix(text, "//") {
							text = strings.TrimSpace(strings.TrimPrefix(text, "//"))
						} else if strings.HasPrefix(text, "/*") {
							text = strings.TrimSpace(strings.TrimPrefix(text, "/*"))
							text = strings.TrimSpace(strings.TrimSuffix(text, "*/"))
						}
						if text != "" {
							parts = append(parts, text)
						}
					}
					commentText = strings.Join(parts, "\n")
				}
				if commentText == "" {
					continue
				}
				// 对每个 handler 匹配注解并记录
				for _, handler := range t {
					name := handler.GetName()
					// 判断是否以 @name 开头
					re := regexp.MustCompile(`(?i)^\s*?@` + name)
					commentTextLines := strings.Split(commentText, "\n")
					var fn = PackageFunction{
						PackageName:  pkgPath,
						FunctionName: methodKey,
					}
					for _, comment := range commentTextLines {
						if len(re.FindStringSubmatch(comment)) == 0 {
							continue
						}
						comment = strings.TrimSpace(comment)
						if comment == "" {
							continue
						}
						at := NewAnnotation(handler.Copy(), comment)
						if at != nil {
							currentFunction := result.Get(fn)
							if currentFunction != nil {
								currentFunction.Annotations = append(currentFunction.Annotations, at)
							} else {
								fn.Annotations = []*Annotation{at}
								result = append(result, fn)
							}
						}
					}
				}
			}
		}
	}
	return &result
}
