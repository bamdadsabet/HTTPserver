package branch

import (
	"database/sql"
	"fmt"
	"http-server/helper"
	"log"
	"time"

	"github.com/qustavo/dotsql"
)

type BranchRepositoryInterface interface {
	Save(val BranchInfoStruct) (BranchStruct, error)
	Update(id string, val BranchInfoStruct) (bool, error)
	DeleteById(id string) (bool, error)
	FindAll() ([]BranchStruct, error)
	FindById(id string) (BranchStruct, error)
}

type BranchRepositoryStruct struct {
	Db *sql.DB
	QueryFile *dotsql.DotSql
}

func NewRepository (db *sql.DB) BranchRepositoryInterface {
	dot, err := dotsql.LoadFromFile("domain/branch/query.sql")
	helper.CHeckErr(err, "loading sql file")
	return BranchRepositoryStruct {db, dot}
}

func (t BranchRepositoryStruct) Save(val BranchInfoStruct) (BranchStruct, error) {
	row, err := t.QueryFile.QueryRow(
		t.Db,
		"create-branch", 
		val.Name, 
		val.Address, 
		val.PhoneNumber, 
	)

	var createdBranch BranchStruct

	if err := row.Scan(
		&createdBranch.Id,
		&createdBranch.Address,
		&createdBranch.PhoneNumber,
		&createdBranch.CreatedAt,
		&createdBranch.UpdatedAt); err != nil {
		fmt.Println(err)
		return BranchStruct{}, err
	}

	if err != nil {
		return BranchStruct{}, err
	}

	return createdBranch, nil
}

func (t BranchRepositoryStruct) Update(id string, val BranchInfoStruct) (bool, error) {
	res, err := t.QueryFile.Query(
		t.Db,
		"update-branch",
		id ,
		val.Name, 
		val.Address, 
		val.PhoneNumber, 
		time.Now(),
	)

	fmt.Println(res)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (t BranchRepositoryStruct) DeleteById(id string) (bool, error) {
	res, err := t.QueryFile.Query(t.Db, "delete-branch-by-id", id)

	// FIXME 
	fmt.Println("res is", res)
	
	if err != nil {
		return false, err
	}

	return true, nil
}

func (t BranchRepositoryStruct) FindById(id string) (BranchStruct, error) {
	var res BranchStruct

	row, err := t.QueryFile.QueryRow(t.Db, "get-branch-by-id", id)

	if err != nil {
		return BranchStruct{}, err
	}

	if err := row.Scan(
		&res.Id,
		&res.Name,
		&res.Address,
		&res.PhoneNumber,
		&res.CreatedAt, 
		&res.UpdatedAt); err != nil {
		return BranchStruct{}, err
	}

	return res, nil
}

func (t BranchRepositoryStruct) FindAll() ([]BranchStruct, error) {
	var res []BranchStruct

	rows, err := t.QueryFile.Query(t.Db, "get-all-branch")
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		var resItm BranchStruct

		if err := rows.Scan(
			&resItm.Id,
			&resItm.Name,
			&resItm.Address,
			&resItm.PhoneNumber, 
			&resItm.CreatedAt, 
			&resItm.UpdatedAt); err != nil {
			log.Fatal(err)
			return nil, err
		}
		res = append(res, resItm)
	}

	return res, nil
}