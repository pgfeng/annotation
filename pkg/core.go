package pkg

import (
	"fmt"
	"go/ast"
	"go/token"
	"regexp"
	"strings"

	"golang.org/x/tools/go/packages"
)

// LoadPackageByPath 加载包并从函数/方法的注释中解析注解。
// 返回值映射格式： map[methodKey]map[annotationName][]values
// methodKey 形如 "Type.Method"（方法）或 "Package.Func"（普通函数，使用包路径作为前缀）
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
			// 遍历声明，寻找函数声明
			for _, decl := range f.Decls {
				fd, ok := decl.(*ast.FuncDecl)
				if !ok {
					continue
				}
				// 构造 methodKey
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
	//for _, fn := range result {
	//	fmt.Println(fn)
	//	for _, at := range fn.Annotations {
	//		fmt.Println(at)
	//	}
	//}
	return &result
}
func ClonePointer[T any](src *T) *T {
	if src == nil {
		return nil
	}
	tmp := *src // 复制值
	return &tmp // 返回指向新值的指针
}

// 辅助：把 ast.Expr 转成简单的类型名，例如:
//
//	*ast.Ident -> ident.Name
//	*ast.StarExpr -> 递归取内部（去掉指针符号）
//	*ast.SelectorExpr -> sel.Sel.Name
func exprToString(e ast.Expr) string {
	switch v := e.(type) {
	case *ast.Ident:
		return v.Name
	case *ast.StarExpr:
		return exprToString(v.X)
	case *ast.SelectorExpr:
		// 例如 pkg.Type -> 返回 Type（去掉包前缀）
		return v.Sel.Name
	default:
		// 回退为 fmt.Sprintf 表示的形式（尽量避免）
		return fmt.Sprintf("%T", e)
	}
}
