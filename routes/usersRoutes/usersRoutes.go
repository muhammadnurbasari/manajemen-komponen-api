package usersRoutes

import (
	"manajemen-komponen-api/modules/users/usersDelivery"
	"manajemen-komponen-api/modules/users/usersRepository"
	"manajemen-komponen-api/modules/users/usersUsecase"
	"manajemen-komponen-api/services/auditTrails/auditTrailsRepository"
	"manajemen-komponen-api/services/auditTrails/auditTrailsUsecase"
	"manajemen-komponen-api/services/logErrors/logErrorsRepository"
	"manajemen-komponen-api/services/logErrors/logErrorsUsecase"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UsersRoutes - list routes users
func UsersRoutes(r *gin.Engine, Db *gorm.DB) {
	auditTrailsRepo := auditTrailsRepository.NewAuditTrailsRepository(Db)
	auditTrailsUC := auditTrailsUsecase.NewAuditTrailsUsecase(auditTrailsRepo)

	logErrorsRepo := logErrorsRepository.NewLogErrorsRepository(Db)
	logErrorsUC := logErrorsUsecase.NewLogErrorsUsecase(logErrorsRepo)

	usersRepo := usersRepository.NewUsersRepository(Db)
	usersUC := usersUsecase.NewUsersUsecase(usersRepo)
	usersDelivery.NewUserHTTPHandler(r, usersUC, auditTrailsUC, logErrorsUC)

}
