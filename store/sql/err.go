package sql

import "errors"

// 常见错误错误
var (
	ErrPKEmpty        = errors.New("pk is empty")
	ErrQueryEmpty     = errors.New("query is empty")
	ErrConditionEmpty = errors.New("condition is empty")
	ErrNotPtr         = errors.New("not valid pointer")
	ErrNotSlice       = errors.New("not valid slice")
	ErrEmptyResult    = errors.New("empty result")
)
