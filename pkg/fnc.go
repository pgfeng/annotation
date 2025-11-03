package pkg

type Function struct {
	PackageName  string
	FunctionName string
}

type FncAnnotations map[Function][]*Annotation

// GetFunctionMap returns a map of functions to their annotations.
func (f PackageFunctions) GetFunctionMap() {
	var result FncAnnotations
	for i := range f {
		result[Function{
			PackageName:  f[i].PackageName,
			FunctionName: f[i].FunctionName,
		}] = f[i].Annotations
	}
}
