package main

import (
	"fmt"
)

type Column struct {
	Name    string
	Type    string
	Size    int
	Null    bool
	Primary bool
	Unique  bool
}

type Table struct {
	Name    string
	Columns *[]Column
}

func main() {

	var columns []Column

	columns = append(columns, Column{
		Name:    "Id",
		Type:    "Integer",
		Size:    0,
		Null:    false,
		Primary: true,
		Unique:  true,
	})

	columns = append(columns, Column{
		Name:    "Name",
		Type:    "VARCHAR",
		Size:    256,
		Null:    false,
		Primary: false,
		Unique:  false,
	})

	table := Table{
		Name:    "TEST",
		Columns: &columns,
	}

	sql := "CREATE TABLE " + table.Name + " ();"
	fmt.Println(idx, column)

}
