package logErrorsRepository

import (
	"errors"
	"manajemen-komponen-api/models/logErrorsModel"
	"manajemen-komponen-api/services/logErrors"

	"github.com/jinzhu/gorm"
)

type sqlRepository struct {
	Conn *gorm.DB
}

// NewLogErrorsRepository - will create an object representation logErrors.LogErrorsRepository
func NewLogErrorsRepository(Conn *gorm.DB) logErrors.LogErrorsRepository {
	return &sqlRepository{Conn: Conn}
}

// InsertLogErrors - method for insert into table log_errors
func (db *sqlRepository) InsertLogErrors(data *logErrorsModel.DataLogErrors) error {
	logErrors := logErrorsModel.LogErrors{
		UserId:           data.UserId,
		ModuleName:       data.ModuleName,
		ErrorDescription: data.ErrorDescription,
	}

	err := db.Conn.Create(&logErrors).Error
	if err != nil {
		return errors.New("LogErrorsRepo.InsertLogErrors Err : " + err.Error())
	}

	return nil
}
