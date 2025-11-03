package annotation

import (
	"github.com/pgfeng/annotation/pkg"
)

var ParseAnnotation = pkg.NewAnnotation

func LoadPackageByPath(s string, t []pkg.Type) *pkg.PackageFunctions {
	return pkg.LoadPackageByPath(s, t)
}
