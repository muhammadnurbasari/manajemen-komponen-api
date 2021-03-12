package auditTrailsUsecase

import (
	"bytes"
	"io"
	"io/ioutil"
	"manajemen-komponen-api/models/auditTrailsModel"
	"manajemen-komponen-api/services/auditTrails"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type auditTrailsUsecase struct {
	auditTrailsRepository auditTrails.AuditTrailsRepository
}

// NewAuditTrailsUsecase - will; create a new auditTrails Usecase that representation auditTrails.AuditTrailsUsecase interface
func NewAuditTrailsUsecase(auditTrailsRepo auditTrails.AuditTrailsRepository) auditTrails.AuditTrailsUsecase {
	return &auditTrailsUsecase{
		auditTrailsRepository: auditTrailsRepo,
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

//MiddlewareAuditTrail - usecase for middleware audit trails
func (auditUC *auditTrailsUsecase) MiddlewareAuditTrail(context *gin.Context) {
	methodName := context.Request.Method
	functionName := context.HandlerName()
	ipAddress := context.ClientIP()
	urlApi := context.Request.URL.EscapedPath()
	resCode := context.Writer.Status()
	reqBody := ""
	resBody := ""
	userId := context.GetHeader("userid")

	if urlApi == "/user/login" {
		userId = "0"
	}

	userIdInt, errUserId := strconv.Atoi(userId)

	if errUserId != nil {
		log.Error().Msg("MiddlewareAuditTrail errUserId : " + errUserId.Error())
		return
	}

	if methodName == http.MethodPost || methodName == http.MethodPut {
		var buf bytes.Buffer
		tee := io.TeeReader(context.Request.Body, &buf)
		body, errReadAll := ioutil.ReadAll(tee)
		if errReadAll != nil {
			log.Error().Msg("MiddlewareAuditTrail errReadAll : " + errReadAll.Error())
			return
		}

		context.Request.Body = ioutil.NopCloser(&buf)
		reqBody = string(body)
	}

	w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: context.Writer}
	context.Writer = w
	context.Next()

	resBody = w.body.String()

	data := auditTrailsModel.DataAuditTrails{
		UserId:       userIdInt,
		UrlApi:       urlApi,
		FunctionName: functionName,
		IpAddress:    ipAddress,
		MethodApi:    methodName,
		ResponseCode: int16(resCode),
		RequestBody:  reqBody,
		ResponseBody: resBody,
	}

	err := auditUC.auditTrailsRepository.InsertAuditTrails(&data)
	if err != nil {
		log.Error().Msg("MiddlewareAuditTrail err =" + err.Error())
		return
	}
}
