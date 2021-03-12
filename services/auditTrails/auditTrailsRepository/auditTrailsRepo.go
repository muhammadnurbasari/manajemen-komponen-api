package auditTrailsRepository

import (
	"errors"
	"manajemen-komponen-api/models/auditTrailsModel"
	"manajemen-komponen-api/services/auditTrails"

	"github.com/jinzhu/gorm"
)

type sqlRepository struct {
	Conn *gorm.DB
}

// NewAauditTrailsRepository - will create object that representation auditTrails.AuditTrailsRepository interface
func NewAuditTrailsRepository(Conn *gorm.DB) auditTrails.AuditTrailsRepository {
	return &sqlRepository{Conn}
}

// InsertAuditTrails - insert into table audit_trails
func (db *sqlRepository) InsertAuditTrails(data *auditTrailsModel.DataAuditTrails) error {
	auditTrails := auditTrailsModel.AuditTrails{
		UserID:       data.UserId,
		UrlApi:       data.UrlApi,
		FunctionName: data.FunctionName,
		IpAddress:    data.IpAddress,
		MethodApi:    data.MethodApi,
		ResponseCode: data.ResponseCode,
		RequestBody:  data.RequestBody,
		ResponseBody: data.ResponseBody,
	}

	err := db.Conn.Create(&auditTrails).Error
	if err != nil {
		return errors.New("auditTrailsRepo.InsertAuditTrails err : " + err.Error())
	}

	return nil
}
