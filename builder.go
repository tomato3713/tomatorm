package tomatorm

import "fmt"

type Builder interface {
	fmt.Stringer
	Build() (string, []any)
}
