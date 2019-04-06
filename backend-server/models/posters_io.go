package models

// PostersForm definiton.
type PostersForm struct {
	Name     string `form:"name" valid:"Required"`
	Title    string `form:"title"`
	Type     string `form:"type"`
	Category string `form:"category"`
	Dtype    string `form:"dtype"`
	Content  string `form:"content"`
	ReferUrl string `form:"refer_url"`
	Status   string `form:"status"`
}
