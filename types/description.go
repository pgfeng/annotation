package types

import "github.com/pgfeng/annotation/pkg"

type Description struct {
	Text string
}

func (d *Description) GetName() string {
	return "Description"
}

// InitValue 解析：@Description 详细描述内容
func (d *Description) InitValue(v string) {
	d.Text = v
}
func (d *Description) Copy() pkg.Type {
	return &Description{
		Text: d.Text,
	}
}
