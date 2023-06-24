package main

import (
	"database/sql"
	"reflect"
	"strings"

	_ "modernc.org/sqlite"
)

func listDrivers() {
	for _, driver := range sql.Drivers() {
		Printfln("Driver: %v", driver)
	}
}

func openDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite", "products.db")
	if err == nil {
		Printfln("Opened database")
	}
	return
}

func InsertNewCategoryPrepare(db *sql.DB) (query *sql.Stmt) {
	query, _ = db.Prepare("INSERT INTO Categories (Name) VALUES (?)")
	return query
}

func ChangeProductCategoryPrepare(db *sql.DB) (query *sql.Stmt) {
	query, _ = db.Prepare("UPDATE Products SET Category = ? WHERE Id = ?")
	return query
}

func scanIntoStruct(rows *sql.Rows, target interface{}) (results interface{}, err error) {
	targetVal := reflect.ValueOf(target)
	if targetVal.Kind() == reflect.Ptr {
		targetVal = targetVal.Elem()
	}
	if targetVal.Kind() != reflect.Struct {
		return
	}
	colNames, _ := rows.Columns()
	colTypes, _ := rows.ColumnTypes()
	references := []interface{}{}
	fieldVal := reflect.Value{}
	var placeholder interface{}

	for i, colName := range colNames {
		colNameParts := strings.Split(colName, ".")
		fieldVal = targetVal.FieldByName(colNameParts[0])
		if fieldVal.IsValid() && fieldVal.Kind() == reflect.Struct && len(colNameParts) > 1 {
			var namePart string
			for _, namePart = range colNameParts[1:] {
				compFunction := matchColName(namePart)
				fieldVal = fieldVal.FieldByNameFunc(compFunction)
			}
		}
		if !fieldVal.IsValid() || !colTypes[i].ScanType().ConvertibleTo(fieldVal.Type()) {
			references = append(references, &placeholder)
		} else if fieldVal.Kind() != reflect.Ptr && fieldVal.CanAddr() {
			fieldVal = fieldVal.Addr()
			references = append(references, fieldVal.Interface())
		}
	}

	resultSlice := reflect.MakeSlice(reflect.SliceOf(targetVal.Type()), 0, 10)
	for rows.Next() {
		err = rows.Scan(references...)
		if err == nil {
			resultSlice = reflect.Append(resultSlice, targetVal)
		} else {
			break
		}
	}
	results = resultSlice.Interface()

	return
}

func matchColName(colName string) func(string) bool {
	return func(fieldName string) bool {
		return strings.EqualFold(colName, fieldName)
	}
}
