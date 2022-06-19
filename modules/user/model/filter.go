package usermodel

type FilterUserNameEmail struct {
	UserName string
	Email    string
}

func NewFilterUserNameEmail(f interface{}) *FilterUserNameEmail {
	return f.(*FilterUserNameEmail)
}

type FilterUserName struct {
	UserName string
}

func NewFilterUserName(f interface{}) *FilterUserName {
	return f.(*FilterUserName)
}
