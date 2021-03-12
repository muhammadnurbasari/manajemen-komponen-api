package logErrorsModel

import "github.com/jinzhu/gorm"

type (
	// LogErrors - object table log_errors
	LogErrors struct {
		gorm.Model
		UserId           int    `gorm:"column:user_id"`
		ModuleName       string `gorm:"column:module_name"`
		ErrorDescription string `gorm:"column:error_description"`
	}

	// DataLogErrors - parameters to insert data log_errors
	DataLogErrors struct {
		UserId           int    `json:"user_id"`
		ModuleName       string `json:"module_name"`
		ErrorDescription string `json:"error_description"`
	}
)
