package employee

import (
	"database/sql"
	"fmt"
	"http-server/utils/helper"
	"time"

	"github.com/qustavo/dotsql"
)

type EmployeeRepositoryInterface interface {
	Save(val EmployeeInfoStruct) (EmployeeStruct, error)
	Update(id string, val EmployeeInfoStruct) (bool, error)
	DeleteById(id string) (bool, error)
	FindAll() ([]EmployeeStruct, error)
	FindById(id string) (EmployeeStruct, error)
}

type EmployeeRepositoryStruct struct {
	Db *sql.DB
	QueryFile *dotsql.DotSql
}

func NewRepository (db *sql.DB) EmployeeRepositoryInterface {
	dot, err := dotsql.LoadFromFile("domain/employee/query.sql")
	helper.CHeckErr(err, "loading sql file")
	return EmployeeRepositoryStruct {db, dot}
}

func (t EmployeeRepositoryStruct) Save(val EmployeeInfoStruct) (EmployeeStruct, error) {
	row, err := t.QueryFile.QueryRow(
		t.Db,
		"create-employee", 
		val.FirstName, 
		val.LastName, 
		val.Email, 
		val.Floor, 
		val.PhoneExtension, 
		val.Role, 
		val.BranchId,
	)

	var createdEmployee EmployeeStruct

	if err := row.Scan(
		&createdEmployee.Id,
		&createdEmployee.FirstName,
		&createdEmployee.LastName,
		&createdEmployee.Email,
		&createdEmployee.Floor,
		&createdEmployee.PhoneExtension,
		&createdEmployee.Role,
		&createdEmployee.BranchId,
		&createdEmployee.CreatedAt,
		&createdEmployee.UpdatedAt); err != nil {
		fmt.Println(err)
		return EmployeeStruct{}, err
	}
	fmt.Println("createdEmployee", createdEmployee)

	

	if err != nil {
		return EmployeeStruct{}, err
	}

	return createdEmployee, nil
}

func (t EmployeeRepositoryStruct) Update(id string, val EmployeeInfoStruct) (bool, error) {
	res, err := t.QueryFile.Query(
		t.Db,
		"update-employee",
		id ,
		val.FirstName, 
		val.LastName, 
		val.Email, 
		val.Floor, 
		val.PhoneExtension, 
		val.Role, 
		val.BranchId,
		time.Now(),
	)

	fmt.Println(res)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (t EmployeeRepositoryStruct) DeleteById(id string) (bool, error) {
	res, err := t.QueryFile.Query(t.Db, "delete-employee-by-id", id)

	// FIXME 
	fmt.Println("res is", res)
	
	if err != nil {
		return false, err
	}

	return true, nil
}

func (t EmployeeRepositoryStruct) FindById(id string) (EmployeeStruct, error) {

	row, err := t.QueryFile.QueryRow(t.Db, "get-employee-by-id", id)

	if err != nil {
		return EmployeeStruct{}, err
	}

	fmt.Println("row", row)
	x, e := helper.ScanRow[EmployeeStruct](row)

	fmt.Println(x, e)
	
// 	bb := []interface{}{
//     &res.Id,
//     &res.FirstName,
//     &res.LastName,
//     &res.Email,
//     &res.Floor,
//     &res.PhoneExtension,
//     &res.Role,
//     &res.BranchId,
//     &res.CreatedAt,
//     &res.UpdatedAt,
// }

// 	fmt.Println(bb)
// 	if err := row.Scan(
// 		&res.Id,
// 		&res.FirstName,
// 		&res.LastName,
// 		&res.Email,
// 		&res.Floor,
// 		&res.PhoneExtension,
// 		&res.Role, 
// 		&res.BranchId,  
// 		&res.CreatedAt, 
// 		&res.UpdatedAt); err != nil {
// 		return EmployeeStruct{}, err
// 	}

	return x, nil
}

func (t EmployeeRepositoryStruct) FindAll() ([]EmployeeStruct, error) {
	var res []EmployeeStruct

	rows, err := t.QueryFile.Query(t.Db, "get-all-employee")
	defer rows.Close()

	if err != nil {
		return res, err
	}

	r, d := helper.ScanRows[EmployeeStruct](rows)

	fmt.Println(d)

	return r, nil
}


