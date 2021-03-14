package globalDTO

// BadRequest - response bad request
type BadRequest struct {
	Status    int    `json:"status" example:"400"`
	Messagges string `json:"messages" example:"MISSING PARAMETERS"`
}

// ForbiddenRes - response forbidden request
type ForbiddenRes struct {
	Status    int    `json:"status" example:"403"`
	Messagges string `json:"messages" example:"Invalid Token"`
}

// UnAuthorized - response unauthorized request
type UnAuthorized struct {
	Status    int    `json:"status" example:"401"`
	Messagges string `json:"messages" example:"unauthorized user"`
}

// InternalServerErr - response internal serve error request
type InternalServerErr struct {
	Status    int    `json:"status" example:"500"`
	Messagges string `json:"messages" example:"Please Contact admin, Server is Having Error"`
}
