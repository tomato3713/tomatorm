package tomatorm

import (
	"fmt"
	"strings"
)

type SelectBuilder struct {
	args []any

	table string
	cols  []string

	groupbyCols []string
	orderbyCols []string
}

var _ Builder = new(SelectBuilder)

func (sb *SelectBuilder) Select(cols ...string) *SelectBuilder {
	sb.cols = cols
	return sb
}

func (sb *SelectBuilder) From(table string) *SelectBuilder {
	sb.table = table
	return sb
}

func (sb *SelectBuilder) GroupBy(col ...string) *SelectBuilder {
	sb.groupbyCols = append(sb.groupbyCols, col...)
	return sb
}

func (sb *SelectBuilder) OrderBy(col ...string) *SelectBuilder {
	sb.orderbyCols = append(sb.orderbyCols, col...)
	return sb
}

func (sb *SelectBuilder) Build() (sql string, args []any) {
	b := strings.Builder{}

	b.WriteString("SELECT ")
	for idx, col := range sb.cols {
		if idx == len(sb.cols)-1 {
			b.WriteString(col + " ")
		} else {
			b.WriteString(col + ", ")
		}
	}
	b.WriteString("FROM " + sb.table)

	if len(sb.orderbyCols) > 0 {
		b.WriteString(" ORDER BY ")
		for idx, col := range sb.orderbyCols {
			if idx == len(sb.orderbyCols)-1 {
				b.WriteString(col)
			} else {
				b.WriteString(col + ", ")
			}
		}
	}

	if len(sb.groupbyCols) > 0 {
		b.WriteString(" GROUP BY ")
		for idx, col := range sb.groupbyCols {
			if idx == len(sb.groupbyCols)-1 {
				b.WriteString(col)
			} else {
				b.WriteString(col + ", ")
			}
		}
	}

	b.WriteString(";")
	return b.String(), sb.args
}

func (sb *SelectBuilder) String() string {
	s, args := sb.Build()
	b := strings.Builder{}
	b.WriteString(s)
	if len(args) > 0 {
		b.WriteString("[")
		for idx, arg := range args {
			if idx == len(args)-1 {
				b.WriteString(fmt.Sprintf("%v", arg))
			} else {
				b.WriteString(fmt.Sprintf("%v , ", arg))
			}
		}
		b.WriteString("]")
	}
	return s
}
