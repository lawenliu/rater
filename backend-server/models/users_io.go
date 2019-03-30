package models

// RegisterForm definiton.
type RegisterForm struct {
	Name     string `form:"name"     valid:"Required"`
	Password string `form:"password" valid:"Required"`
}

// LoginForm definiton.
type LoginForm struct {
	Name     string `form:"name"     valid:"Required"`
	Password string `form:"password" valid:"Required"`
}

// LoginInfo definiton.
type LoginInfo struct {
	Code     int   `json:"code"`
	UserInfo *User `json:"user"`
}

// LogoutForm defintion.
type LogoutForm struct {
	Name     string `form:"name"     valid:"Required"`
}

// PasswdForm definition.
type PasswdForm struct {
	Name     string `form:"name"     valid:"Required"`
	OldPass string `form:"old_password" valid:"Required"`
	NewPass string `form:"new_password" valid:"Required"`
}

// RolePutForm definiton.
type UserPutForm struct {
	ID       string `json:"id"`
	Name    string `json:"name"`
	Password string `json:"password"`
}

// UploadsForm definiton.
type UploadsForm struct {
	Name    string `json:"name"`
}
