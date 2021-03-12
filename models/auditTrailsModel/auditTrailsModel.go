package auditTrailsModel

import "github.com/jinzhu/gorm"

type AuditTrails struct {
	gorm.Model
	UserID       int    `gorm:"column:user_id"`
	UrlApi       string `gorm:"column:url_api"`
	FunctionName string `gorm:"column:function_name"`
	IpAddress    string `gorm:"column:ip_address"`
	MethodApi    string `gorm:"column:method_api"`
	ResponseCode int16  `gorm:"column:response_code"`
	RequestBody  string `gorm:"column:request_body"`
	ResponseBody string `gorm:"column:response_body"`
}

// DataAuditTrails - parameter to insert data audit_trails
type DataAuditTrails struct {
	UserId       int
	UrlApi       string
	FunctionName string
	IpAddress    string
	MethodApi    string
	ResponseCode int16
	RequestBody  string
	ResponseBody string
}
