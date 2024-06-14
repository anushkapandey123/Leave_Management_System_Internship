package repository

import (
	"context"
	"errors"

	// "fmt"
	"reflect"

	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"main.go/leaves/model"
)



type leaveRepository struct {
	db *gorm.DB
}

func NewLeaveRepository(db *gorm.DB) *leaveRepository {
	return &leaveRepository{
		db: db,
	}
}

func (repo leaveRepository) FindNameByUserId(ctx context.Context, id int) (model.Emp, error) {
	var employee model.Emp

	if result := repo.db.Where("id = ?", id).Find(&employee); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.Emp{}, nil
		}

		return model.Emp{}, errors.New("error occured")
	}

	return employee, nil
}

func (repo leaveRepository) FetchLeavesByEmpId(ctx context.Context) (*[]model.Leave, error) {
	var leave []model.Leave

	if result := repo.db.Order("emp_id").Find(&leave); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &[]model.Leave{}, nil
		}

		return &[]model.Leave{}, errors.New("error occurred")
	}

	return &leave, nil

}

func (repo leaveRepository) Create(ctx context.Context, c *model.Leave) error {
	dbCtx := repo.db.WithContext(ctx)
	result := dbCtx.Clauses(clause.OnConflict{DoNothing: true}).Create(c)

	if result.Error != nil {

		return errors.New("error occured")
	}

	if result.RowsAffected == 0 {

		return errors.New("error occurred, duplicate records")
	}
	return nil
}

func (repo leaveRepository) FindLeave(ctx context.Context, empid int, sdate time.Time, edate time.Time) (bool, error) {

	var leave model.Leave

	if result := repo.db.Where("start_date = ?", sdate).Where("end_date = ?", edate).Where("emp_id = ?", empid).Find(&leave); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, errors.New("some error occurred")

	}

	var emptyLeave model.Leave

	// record not found
	if reflect.DeepEqual(leave, emptyLeave) {
		return false, nil
	}

	return true, nil
}

func (repo leaveRepository) Delete(ctx context.Context, c *model.Leave, sdate time.Time, edate time.Time) error {

	dbCtx := repo.db.WithContext(ctx)
	result := dbCtx.Clauses(clause.OnConflict{DoNothing: true}).Where("start_date = ?", sdate).Where("end_date = ?", edate).Delete(c)

	if result.Error != nil {

		return errors.New("error occured")
	}

	if result.RowsAffected == 0 {

		return errors.New("error occurred, duplicate records")
	}
	return nil

}


func (repo leaveRepository) CheckForOverlappingLeaves(ctx context.Context, date time.Time, empid int) (bool, error) {
	var leave model.Leave

	if result := repo.db.Where("start_date <= ?", date).Where("end_date >= ?", date).Where("emp_id = ?", empid).Find(&leave); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, errors.New("some error occurred")

	}

	var emptyLeave model.Leave
	
	// record not found
	if reflect.DeepEqual(leave, emptyLeave) {
		return false, nil
	}

	return true, nil
}

func (repo leaveRepository) FetchLeavesByEmailId(ctx context.Context, email any) (*[]model.Leave, error) {

	var leave []model.Leave

	if result := repo.db.Where("email = ?", email).Find(&leave); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &[]model.Leave{}, nil
		}

		return &[]model.Leave{}, errors.New("error occurred")
	}

	return &leave, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (repo leaveRepository) FindLeaveByEmailId(ctx context.Context, email any, sdate time.Time, edate time.Time) (bool, error) {

	var leave model.Leave

	if result := repo.db.Where("start_date = ?", sdate).Where("end_date = ?", edate).Where("email = ?", email).Find(&leave); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, errors.New("some error occurred")

	}

	var emptyLeave model.Leave

	// record not found
	if reflect.DeepEqual(leave, emptyLeave) {
		return false, nil
	}

	return true, nil
}

func (repo leaveRepository) GetLatestLeave(ctx context.Context, email any) (model.Leave, error) {

	var leave model.Leave

	if result := repo.db.Order("start_date desc").Where("email = ?", email).First(&leave); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return leave, nil
		}

		return leave, errors.New("some error occurred")

	}

	return leave, nil

}

func (repo leaveRepository) CheckForOverlappingLeavesNew(ctx context.Context, date time.Time, email any) (bool, error) {
	var leave model.Leave

	if result := repo.db.Where("start_date <= ?", date).Where("end_date >= ?", date).Where("email = ?", email).Find(&leave); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, errors.New("some error occurred")

	}

	var emptyLeave model.Leave

	// record not found
	if reflect.DeepEqual(leave, emptyLeave) {
		return false, nil
	}

	return true, nil
}

func (repo leaveRepository) FindNameByEmail(ctx context.Context, email any) (model.User, error) {

	var user model.User

	if result := repo.db.Where("email = ?", email).Find(&user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.User{}, nil
		}

		return model.User{}, errors.New("error occured")
	}

	return user, nil


}
