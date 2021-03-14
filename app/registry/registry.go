package registry

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"manajemen-komponen-api/app/appServer"
	"manajemen-komponen-api/database/mysql"
	"manajemen-komponen-api/models/appServerModel"
)

type AppRegistry struct {
	Conn       *gorm.DB
	serverHttp *appServer.HttpHandler
}

//NewAppRegistry will return new object for App Registry
func NewAppRegistry() *AppRegistry {
	return &AppRegistry{}
}

func initializeEnv() (*appServerModel.SetConnDb, error) {
	moduleName := "registry.initializeEnv"
	log.Debug().Msg("Read file env . . .")
	err := godotenv.Load("config/.env")
	if err != nil {
		return nil, errors.New(moduleName + ".err : " + err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New(moduleName + ".error : port cant empty")
	}

	log.Debug().Msg("mapping config database . . .")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	// dbTimezone := os.Getenv("DB_TIME_ZONE")
	// dbSSL := os.Getenv("DB_SSL")
	maxIdle := os.Getenv("MAX_IDLE")
	maxConn := os.Getenv("MAX_CONN")

	myMaxIdle, errMaxIdle := strconv.Atoi(maxIdle)
	if errMaxIdle != nil {
		return nil, errors.New(moduleName + ".errMaxIdle : " + errMaxIdle.Error())
	}

	myMaxConn, errMaxConn := strconv.Atoi(maxConn)
	if errMaxConn != nil {
		return nil, errors.New(moduleName + ".errMaxConn : " + errMaxConn.Error())
	}

	log.Debug().Msg("set config database . . .")
	setConnDb := appServerModel.SetConnDb{
		DbHost: dbHost,
		DbPort: dbPort,
		DbUser: dbUser,
		DbPass: dbPass,
		DbName: dbName,
		// DbSSL:      dbSSL,
		// DbTimezone: dbTimezone,
		MaxIdle: myMaxIdle,
		MaxConn: myMaxConn,
	}

	return &setConnDb, nil
}

//StartServer will do the server initialization
func (reg *AppRegistry) StartServer() {
	printAsciiArt()
	log.Debug().Msg("prepare start server . . .")
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)
	log.Debug().Msg("initial database . . .")
	setConnDb, errInitialEnv := initializeEnv()
	if errInitialEnv != nil {
		log.Error().Msg(errInitialEnv.Error())
		return
	}

	//initial database
	conn, errConn := getDBConnection(setConnDb)
	if errConn != nil {
		log.Error().Msg(errConn.Error())
		return
	}

	reg.Conn = conn

	//close connection
	defer func() {
		log.Info().Msg("Close connection . . .")
		errConnClose := reg.Conn.Close()
		if errConnClose != nil {
			log.Error().Msg("Service manajemen-komponen-api.errConnClose : " + errConnClose.Error())
		}
	}()

	errApp := reg.initializeAppRegistry()
	if errApp != nil {
		log.Error().Msg(errApp.Error())
		return
	}

	//Run Swagger
	log.Info().Msg("Swagger run on /docs/swagger/index.html")
	reg.serverHttp.RunSwaggerMiddleware()

	//Run HTTP Server
	appVersion := "0.0.1"
	appPort := os.Getenv("PORT")
	log.Info().Msg("Last Update : " + time.Now().Format("2006-01-02 15:04:05"))
	log.Info().Msg("REST API Service Running version " + appVersion + " at port : " + appPort)
	if errHTTP := reg.serverHttp.RunHttpServer(); errHTTP != nil {
		log.Error().Msg(errHTTP.Error())
	}
}

func (reg *AppRegistry) initializeAppRegistry() error {
	//initial read file env
	log.Debug().Msg("prepare initial env . . .")
	port := os.Getenv("PORT")
	//initial handler
	errHandler := reg.initializeHandler(port)
	if errHandler != nil {
		return errHandler
	}

	//initial modules
	reg.initializeDomainModules()

	return nil
}

func (reg *AppRegistry) initializeHandler(appPort string) error {
	//Register HTTP Server Handler
	if appPort == "" {
		return errors.New("registry.error : port cant empty")
	}
	reg.serverHttp = appServer.NewHTTPHandler(":" + appPort)
	return nil
}

func getDBConnection(data *appServerModel.SetConnDb) (*gorm.DB, error) {
	log.Debug().Msg("data.DbHost : " + data.DbHost)
	log.Debug().Msg("data.DbUser : " + data.DbUser)
	log.Debug().Msg("data.DbPass : " + data.DbPass)
	log.Debug().Msg("data.DbName : " + data.DbName)
	// log.Debug().Msg("data.DbPort : " + data.DbPort)
	// log.Debug().Msg("data.DbSSL : " + data.DbSSL)
	log.Debug().Msg("data.DbTimezone : " + data.DbTimezone)
	conn, err := mysql.ConnMySQLORM(data.DbHost, data.DbPort, data.DbUser, data.DbPass,
		data.DbName, data.MaxIdle, data.MaxConn)
	if err != nil {
		return nil, errors.New("registry.getDBConnection.err : " + err.Error())
	}

	return conn, nil
}
