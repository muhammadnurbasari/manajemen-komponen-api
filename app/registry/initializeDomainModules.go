package registry

import (
	"manajemen-komponen-api/routes/usersRoutes"
)

//initializeDomainModules calls the domain module routes
//in folder routes/*
func (reg *AppRegistry) initializeDomainModules() {

	usersRoutes.UsersRoutes(reg.serverHttp.GetRouteEngine(), reg.Conn)
}
