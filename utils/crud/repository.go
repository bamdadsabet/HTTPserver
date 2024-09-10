package crud

import (
	"database/sql"
	"http-server/helper"
	"http-server/utils/helper"

	"github.com/qustavo/dotsql"
)

type RepositoryStruct struct {
	Db *sql.DB
	QueryFile *dotsql.DotSql
}

type RepositoryInterface[TRsp, TReq any] interface {
	Save(val TReq) (U, TRsp)
	Update(id string, val TReq) (bool, error)
	DeleteById(id string) (bool, error)
	FindAll() ([]TRsp, error)
	FindById(id string) (TRsp, error)
}

func NewRepository [TRsp, TReq any](db *sql.DB) RepositoryInterface[TRsp, TReq] {
	dot, err := dotsql.LoadFromFile("domain/employee/query.sql")
	helper.CHeckErr(err, "loading sql file")
	return RepositoryStruct{db, dot}
}
func (t RepositoryStruct) Save(valInfoStruct) (EmployeeStruct, error) {
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

	createdItem, err := helper.ScanRow(row)

	if(err != nil) {}

	return createdEmployee, nil
}

func (t RepositoryStruct) Update(id string, valInfoStruct) (bool, error) {
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

func (t RepositoryStruct) DeleteById(id string) (bool, error) {
	res, err := t.QueryFile.Query(t.Db, "delete-employee-by-id", id)

	// FIXME 
	fmt.Println("res is", res)
	
	if err != nil {
		return false, err
	}

	return true, nil
}

func (t RepositoryStruct) FindById(id string) (EmployeeStruct, error) {
	var res Struct

	row, err := t.QueryFile.QueryRow(t.Db, "get-employee-by-id", id)

	if err != nil {
		returnStruct{}, err
	}

	if err := row.Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Email,
		&res.Floor,
		&res.PhoneExtension,
		&res.Role, 
		&res.BranchId,  
		&res.CreatedAt, 
		&res.UpdatedAt); err != nil {
		returnStruct{}, err
	}

	return res, nil
}

func (t RepositoryStruct) FindAll() ([]EmployeeStruct, error) {
	var res []Struct

	rows, err := t.QueryFile.Query(t.Db, "get-all-employee")
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		var resItmStruct

		if err := rows.Scan(
			&resItm.Id,
			&resItm.FirstName,
			&resItm.LastName,
			&resItm.Email,
			&resItm.Floor,
			&resItm.PhoneExtension,
			&resItm.Role,
			&resItm.BranchId,  
			&resItm.CreatedAt, 
			&resItm.UpdatedAt); err != nil {
			log.Fatal(err)
			return nil, err
		}
		res = append(res, resItm)
	}

	return res, nil
}