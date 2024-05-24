package repository

import (
	"context"
	"errors"
	"fmt"
	

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"main.go/leaves/model"
)

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *employeeRepository {
	return &employeeRepository{
		db: db,
	}
}

func (repo employeeRepository) FindByUserId(ctx context.Context) (*[]model.Emp, error) {
	var employee []model.Emp

	// db := repo.WithContext(ctx)
	// defer cancel()

	if result := repo.db.Find(&employee); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &[]model.Emp{}, nil
		}
		// return &model.Emp{}, ae.InternalServerError("InternalServerError","something went wrong",
		// 	fmt.Errorf("something went wrong %v", result.Error))

		return &[]model.Emp{}, errors.New("error occured")
	}

	return &employee, nil
}

func (repo employeeRepository) FetchLeavesByEmpId(ctx context.Context) (*[]model.Leave, error) {
	var leave []model.Leave

	if result := repo.db.Order("emp_id").Find(&leave); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &[]model.Leave{}, nil
		}

		return &[]model.Leave{}, errors.New("error occurred")
	}

	return &leave, nil

}

func (repo employeeRepository) Create(ctx context.Context, c *model.Leave) error {
	dbCtx := repo.db.WithContext(ctx)
	fmt.Println(c)
	result := dbCtx.Clauses(clause.OnConflict{DoNothing: true}).Create(c)

	if result.Error != nil {
		// return ae.UnProcessableError("CustomerCreationFailed", "Customer creation failed due to unknown reason", result.Error)
		return errors.New("error occured")
	}

	if result.RowsAffected == 0 {
		// return ae.UnProcessableError("CustomerAlreadyExist", "Customer already exist. Duplicate record", nil)
		return errors.New("error occurred, duplicate records")
	}
	return nil
}

// func (repo employeeRepository) DeleteByStartDate(ctx context.Context, c *model.Leave) error {
	
// 	dbCtx := repo.db.WithContext(ctx)
// 	result := dbCtx.Clauses(clause.OnConflict{DoNothing: true}).Delete(c)

// 	if result.Error != nil {
// 		// return ae.UnProcessableError("CustomerCreationFailed", "Customer creation failed due to unknown reason", result.Error)
// 		return errors.New("error occured")
// 	}

// 	if result.RowsAffected == 0 {
// 		// return ae.UnProcessableError("CustomerAlreadyExist", "Customer already exist. Duplicate record", nil)
// 		return errors.New("error occurred, duplicate records")
// 	}
// 	return nil


// }


