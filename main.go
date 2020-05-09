package main

import (
	"fmt"
	"strconv"
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
	Comment       string
}

type Table struct {
	Name    string
	Columns *[]Column
	Comment string
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
		Name:    "name",
		Type:    "varchar",
		Size:    256,
		Unique:  true,
		Comment: "名前",
	})

	columns = append(columns, Column{
		Name:    "delete_flag",
		Type:    "tinyint",
		Default: "0",
		Comment: "削除フラグ",
	})

	table := Table{
		Name:    "test",
		Columns: &columns,
		Comment: "テスト用テーブル",
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

		if len(column.Comment) > 0 {
			sql += " comment = " + column.Comment
		}

		sql += ", \n"
	}

	sql += ")"

	if len(table.Comment) > 0 {
		sql += " comment = " + table.Comment
	}

	sql += ";"

	fmt.Println(sql)

}
