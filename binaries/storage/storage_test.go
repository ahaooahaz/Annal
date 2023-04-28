package storage

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/DATA-DOG/go-sqlmock"
)

func EXPECT_InsertSQL(mock sqlmock.Sqlmock, table string, cols []string, values []interface{}, inIdx int) {
	if len(values) <= 0 {
		return
	}
	fieldTags := make(map[string]string)

	vt := reflect.TypeOf(values[0])
	if vt.Kind() == reflect.Ptr {
		vt = vt.Elem()
	}
	if vt.Kind() == reflect.Struct {
		numFields := 0
		for i := 0; i < vt.NumField(); i++ {
			if unicode.IsUpper(rune(vt.Field(i).Name[0])) {
				numFields++
				fieldTags[vt.Field(i).Name] = vt.Field(i).Tag.Get("db")
			}
		}
	}

	fieldValues := make([]map[string]driver.Value, len(values))
	for i, value := range values {
		vv := reflect.ValueOf(value)
		if vv.Kind() == reflect.Ptr {
			vv = vv.Elem()
		}

		if vv.Kind() == reflect.Struct {
			fieldValues[i] = make(map[string]driver.Value)
			for k, v := range fieldTags {
				fieldValues[i][v] = vv.FieldByName(k).Interface()
			}
		}
	}

	mark := make([]string, len(cols))
	for i := 0; i < len(cols); i++ {
		mark[i] = "?"
	}

	marks := make([]string, len(values))
	for i := 0; i < len(values); i++ {
		marks[i] = fmt.Sprintf("(%s)", strings.Join(mark, ", "))
	}

	args := []driver.Value{}
	for _, v := range fieldValues {
		for _, col := range cols {
			args = append(args, v[col])
		}
	}

	SQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES %v", table, strings.Join(cols, ", "), strings.Join(marks, ", "))
	fmt.Printf("command: %v, args: %v\n", SQL, args)
	mock.ExpectExec(SQL).WithArgs(args...).WillReturnResult(sqlmock.NewResult(1, int64(len(values))))
}
