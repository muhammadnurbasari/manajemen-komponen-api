package auditTrails

import "github.com/gin-gonic/gin"

// AuditTrailsUsecase - usecase Audittrails Interface
type AuditTrailsUsecase interface {
	MiddlewareAuditTrail(context *gin.Context)
}
