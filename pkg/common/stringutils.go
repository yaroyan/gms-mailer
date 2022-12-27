package common

type Utils interface {
	IsEmpty(value string) bool
}

var StringUtils Utils = utils{}

type utils struct {
}

func (u utils) IsEmpty(value string) bool {
	return value == ""
}
