package postmodel

type FilterAll struct {
}

func NewFilterAll(f interface{}) *FilterAll {
	return f.(*FilterAll)
}

type FilterPublished struct {
	Published bool
}

func NewFilterPublished(f interface{}) *FilterPublished {
	return f.(*FilterPublished)
}

type FilterTitle struct {
	Id    string
	Title string
}

func NewFilterTitle(f interface{}) *FilterTitle {
	return f.(*FilterTitle)
}
