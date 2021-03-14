package logErrors

// LogErrorsUsecase - usecase log errors interface
type LogErrorsUsecase interface {
	CheckIsError(err error, userId string, moduleName string) bool
}
