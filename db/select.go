package db

import (
	"strconv"
	"strings"

	"github.com/GoRustNet/xurl/str"
)

type Select struct {
	table  string
	fields string
	where  string
	order  string
	limit  int
	offset int
	sb     *strings.Builder
}

func SelectBuilder(table ...string) *Select {
	tableName := ""
	if len(table) > 0 {
		tableName = table[0]
	}
	return &Select{sb: &strings.Builder{}, limit: 0, offset: -1, table: tableName, fields: "*"}
}

func (s *Select) Table(table string) *Select {
	s.table = table
	return s
}

func (s *Select) Fields(fields string) *Select {
	s.fields = fields
	return s
}

func (s *Select) Where(where string) *Select {
	s.where = where
	return s
}

func (s *Select) Order(order string) *Select {
	s.order = order
	return s
}
func (s *Select) Limit(limit int) *Select {
	s.limit = limit
	return s
}
func (s *Select) Offset(offset int) *Select {
	s.offset = offset
	return s
}
func (s *Select) StringBuilder() *strings.Builder {
	return s.sb
}
func (s *Select) String() string {
	s.build()
	return s.StringBuilder().String()
}
func (s *Select) Build() string {
	return s.String()
}
func (s *Select) build() {
	s.sb.WriteString("SELECT ")
	s.sb.WriteString(s.fields)
	s.sb.WriteString(" FROM ")
	s.sb.WriteString(s.table)

	if str.IsNotEmpty(s.where) {
		s.sb.WriteString(" WHERE ")
		s.sb.WriteString(s.where)
	}

	if str.IsNotEmpty(s.order) {
		s.sb.WriteString(" ORDER BY ")
		s.sb.WriteString(s.order)
	}

	if s.limit > 0 {
		s.sb.WriteString(" LIMIT ")
		s.sb.WriteString(strconv.Itoa(s.limit))
	}

	if s.offset >= 0 {
		s.sb.WriteString(" OFFSET ")
		s.sb.WriteString(strconv.Itoa(s.offset))
	}
}
func (s *Select) CountFields(fields ...string) *Select {
	fieldName := "COUNT(*)"
	if len(fields) > 0 {
		fieldName = fields[0]
	}
	s.fields = fieldName
	return s
}
func (s *Select) ToCount(fields ...string) *Select {

	ss := &Select{
		table:  s.table,
		where:  s.where,
		sb:     &strings.Builder{},
		offset: -1,
	}
	ss.CountFields(fields...)
	return ss
}
