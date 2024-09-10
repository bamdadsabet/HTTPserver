package helper

import (
	"database/sql"
	"reflect"
)

type columnInfo struct {
	Name string
	Type reflect.Type
	Addr interface{}
}

func getScanArgs(rv reflect.Value) ([]interface{}, []columnInfo) {
	columns := getColumns(rv)
	scanArgs := make([]interface{}, len(columns))

	for i, col := range columns {
		if col.Addr != nil {
			scanArgs[i] = col.Addr
		} else {
			ptr := reflect.New(col.Type)
			col.Addr = ptr.Interface()
			scanArgs[i] = ptr.Interface()
		}
	}
	
	return scanArgs, columns
}

func ScanRow[T any](r *sql.Row) (T, error) {
	var result T
	rv := reflect.ValueOf(&result).Elem()
	scanArgs, columns := getScanArgs(rv)

	if err := r.Scan(scanArgs...); err != nil {
		return result, err
	}

	setValuesFromScanArgs(columns, rv, scanArgs)

	return result, nil
}

func ScanRows[T any](r *sql.Rows) ([]T, error) {

	var res []T

	var resItem T
	rv := reflect.ValueOf(&resItem).Elem()

	defer r.Close()

	for r.Next() {
		scanArgs, columns := getScanArgs(rv)

		if err := r.Scan(scanArgs...); err != nil {
			return res, err
		}

		setValuesFromScanArgs(columns, rv, scanArgs)

		res = append(res, resItem)
	}

	return res, nil
}

func getColumns(rv reflect.Value) []columnInfo {
	var columns []columnInfo
	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		fv := rv.Field(i)

		if field.Anonymous && fv.Kind() == reflect.Struct {
			columns = append(columns, getColumns(fv)...)
		} else {
			columns = append(columns, columnInfo{
				Name: field.Name,
				Type: field.Type,
				Addr: nil,
			})
		}
	}

	return columns
}

func setValuesFromScanArgs(columns []columnInfo, rv reflect.Value, scanArgs []interface{}) {
	for i, col := range columns {
		value := reflect.ValueOf(scanArgs[i])
		if value.IsValid() {
			rv.FieldByName(col.Name).Set(value.Elem())
		}
	}
}