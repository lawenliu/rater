package models

// PostersForm definiton.
type MediasForm struct {
	Name     string `form:"name" valid:"Required"`
	Dtype    string `form:"dtype"`
}
