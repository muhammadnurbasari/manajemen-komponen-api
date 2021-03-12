package logErrorsUsecase

import (
	"manajemen-komponen-api/models/logErrorsModel"
	"manajemen-komponen-api/services/logErrors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type logErrorsUsecase struct {
	logErrorsRepository logErrors.LogErrorsRepository
}

// NewLogErrorsUsecase - will create an object representation logErrors.LogErrorsUsecases interface
func NewLogErrorsUsecase(logErrorsRepo logErrors.LogErrorsRepository) logErrors.LogErrorsUsecase {
	return &logErrorsUsecase{logErrorsRepository: logErrorsRepo}
}

// CheckIsError - method for insert to log_errors table
func (logErrorsUC *logErrorsUsecase) CheckIsError(err error, userId string, moduleName string) bool {
	userIdInt, errUserId := strconv.Atoi(userId)
	if errUserId != nil {
		return true
	}

	if err != nil {
		log.Error().Msg(err.Error())
		logErrorsData := logErrorsModel.DataLogErrors{
			UserId:           userIdInt,
			ModuleName:       moduleName,
			ErrorDescription: err.Error(),
		}

		errInsertLog := logErrorsUC.logErrorsRepository.InsertLogErrors(&logErrorsData)
		if errInsertLog != nil {
			log.Error().Msg(errInsertLog.Error())
			return true
		}
		return true
	}

	return false
}
