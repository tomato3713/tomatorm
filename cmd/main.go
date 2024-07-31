package main

import (
	"fmt"

	"github.com/tomato3713/tomatorm"
)

func main() {
	builder := tomatorm.NewSelectBuilder()
	query := builder.
		Select("id", "name").
		From("user").
		OrderBy("created_at", "id").
		GroupBy("gender")

	fmt.Println(query)
}
