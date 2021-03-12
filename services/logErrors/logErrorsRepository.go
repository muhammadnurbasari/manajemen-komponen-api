package logErrors

import "manajemen-komponen-api/models/logErrorsModel"

// LogErrorsRepository - repository log errors interface
type LogErrorsRepository interface {
	InsertLogErrors(data *logErrorsModel.DataLogErrors) error
}
