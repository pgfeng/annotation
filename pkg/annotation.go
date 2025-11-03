package pkg

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strings"
)

// Type 接口必须与 annotation_type 下具体类型的方法签名一致
type Type interface {
	GetName() string
	InitValue(v string)
	Copy() Type
}

type Annotation struct {
	Name     string
	Value    string
	Instance Type
}

func (a Annotation) GetName() string {
	return a.Name
}

func (a Annotation) GetValue() string {
	return a.Value
}

type PackageFunctions []PackageFunction

func (f PackageFunctions) Get(fn PackageFunction) *PackageFunction {
	for i := range f {
		if f[i].FunctionName == fn.FunctionName && f[i].PackageName == fn.PackageName {
			return &f[i]
		}
	}
	return nil
}
func (f PackageFunctions) Filter(p Type) *PackageFunctions {
	// 获取只包含对应Type的方法
	var result PackageFunctions
	for i := 0; i < len(f); i++ {
		for ii := 0; ii < len(f[i].Annotations); ii++ {
			if reflect.TypeOf(f[i].Annotations[ii].Instance) == reflect.TypeOf(p) {
				fn := result.Get(f[i])
				if fn != nil {
					fn.Annotations = append(fn.Annotations, f[i].Annotations[ii])
				} else {
					result = append(result, PackageFunction{
						PackageName:  f[i].PackageName,
						FunctionName: f[i].FunctionName,
						Annotations: []*Annotation{
							f[i].Annotations[ii],
						},
					})
				}
			}
		}
	}
	return &result
}

// GetImportSpecs Get AST Import Spec
func (f PackageFunctions) GetImportSpecs() (importSpecs []ast.Spec) {
	for _, fn := range f {
		fnSpec := fn.GetImportSpec()
		if fnSpec == nil || fnSpec.Path == nil {
			continue
		}
		fnPath := fnSpec.Path.Value

		exists := false
		for _, s := range importSpecs {
			if is, ok := s.(*ast.ImportSpec); ok && is.Path != nil {
				if is.Path.Value == fnPath {
					exists = true
					break
				}
			}
		}
		if !exists {
			importSpecs = append(importSpecs, fnSpec)
		}
	}
	return importSpecs
}

type PackageFunction struct {
	PackageName  string
	FunctionName string
	Annotations  []*Annotation
}

func (f *PackageFunction) Find(p Type) *Type {
	for i := range f.Annotations {
		if reflect.DeepEqual(f.Annotations[i].Instance, p) {
			return &f.Annotations[i].Instance
		}
	}
	return nil
}

func (f *PackageFunction) GetImportSpec() *ast.ImportSpec {
	return &ast.ImportSpec{
		Path: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%s"`, f.PackageName)},
	}
}

func NewAnnotation(t Type, annotationTxt string) *Annotation {
	parts := strings.Fields(strings.TrimSpace(annotationTxt))
	if len(parts) == 0 {
		return nil
	}

	name := strings.TrimPrefix(parts[0], "@")
	if !strings.EqualFold(name, t.GetName()) {
		return nil
	}

	value := ""
	if len(parts) > 1 {
		value = strings.Join(parts[1:], " ")
	}

	t.InitValue(value)

	return &Annotation{
		Name:     name,
		Value:    value,
		Instance: t,
	}
}
