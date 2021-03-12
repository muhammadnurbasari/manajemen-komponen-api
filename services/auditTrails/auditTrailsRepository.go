package auditTrails

import "manajemen-komponen-api/models/auditTrailsModel"

// AuditTrailsRepository - repository audittrails interface
type AuditTrailsRepository interface {
	InsertAuditTrails(data *auditTrailsModel.DataAuditTrails) error
}
