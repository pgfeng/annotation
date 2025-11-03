package types

import "github.com/pgfeng/annotation/pkg"

type Description struct {
	Text string
}

func (d *Description) ToMap() map[string]string {
	return map[string]string{
		"text": d.Text,
	}
}

func (d *Description) GetName() string {
	return "Description"
}

// InitValue Parse：@Description this is a description text
func (d *Description) InitValue(v string) {
	d.Text = v
}
func (d *Description) Copy() pkg.Type {
	return &Description{
		Text: d.Text,
	}
}
