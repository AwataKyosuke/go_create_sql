package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Column struct {
	Name          string
	Type          string
	Size          int
	NotNull       bool
	Default       string
	Primary       bool
	Unique        bool
	AutoIncrement bool
}

type Table struct {
	Name    string
	Columns *[]Column
}

func main() {

	var columns []Column

	columns = append(columns, Column{
		Name:          "id",
		Type:          "bigint unsigned",
		Primary:       true,
		AutoIncrement: true,
	})

	columns = append(columns, Column{
		Name:   "name",
		Type:   "varchar",
		Size:   256,
		Unique: true,
	})

	columns = append(columns, Column{
		Name:    "delete_flag",
		Type:    "tinyint",
		Default: "0",
	})

	table := Table{
		Name:    "test",
		Columns: &columns,
	}

	sql := "create table " + table.Name + " ( \n"

	for _, column := range *table.Columns {
		sql += column.Name + " " + column.Type

		if column.Size > 0 {
			sql += "(" + strconv.Itoa(column.Size) + ")"
		}

		if column.NotNull == false {
			sql += " not null"
		}

		if len(column.Default) > 0 {
			sql += " default " + column.Default
		}

		if column.Unique {
			sql += " unique"
		}

		if column.AutoIncrement {
			sql += " auto_increment"
		}

		if column.AutoIncrement || column.Primary {
			sql += " primary key"
		}

		sql += ", \n"
	}

	sql = strings.TrimRight(sql, "\n")

	sql = strings.TrimRight(sql, " ")

	sql = sql[:(len(sql) - 1)]

	sql += "\n"

	sql += ")"

	sql += ";"

	fmt.Println(sql)

}
