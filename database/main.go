package main

import (
	"flag"
	"manajemen-komponen-api/database/migrate"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	err := godotenv.Load("config/.env")
	if err != nil {
		log.Error().Msg("Failed read configuration database")
		return
	}

	pathMigration := os.Getenv("PATH_MIGRATION")
	migrationDir := flag.String("migration-dir", pathMigration, "migration directory")

	upMigration := flag.Bool("up", false, "Up migration flag")
	downMigration := flag.Bool("down", false, "Down migration flag")
	versionMigration := flag.Bool("version", false, "Version of migration flag")

	newMigrationFile := flag.Bool("create", false, "Create new migration file")
	newMigrationFileName := flag.String("filename", "", "New migration file name")
	flag.Parse()

	if *newMigrationFile {
		if *newMigrationFileName == "" {
			log.Error().Msg("please specify migration file name with --filename")
			migrate.ShowHelp()
			return
		}

		//create new migration file
		err := migrate.CreateNewMigrationFile(*migrationDir, *newMigrationFileName)
		if err != nil {
			log.Error().Msg("failed to create migration file " + err.Error())
		}

		return
	}

	//check if at least up or down flag is specified
	if !(*upMigration || *downMigration || *versionMigration) {
		log.Error().Msg("please specify --up or --down for migration")
		migrate.ShowHelp()
		return
	}

	//check migration direction up/down
	if *upMigration && *downMigration {
		log.Warn().Msg("use --up or --down at once only")
		migrate.ShowHelp()
		return
	}

	//setting db config
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbDriver := os.Getenv("DB_DRIVER")
	migrationConf, errMigrationConf := migrate.NewMigrationConfig(*migrationDir, dbHost, dbPort, dbUser, dbPass, dbName, dbDriver)
	if errMigrationConf != nil {
		log.Error().Msg(errMigrationConf.Error())
		return
	}
	defer func() {
		errConnClose := migrationConf.Db.Close()
		if errConnClose != nil {
			log.Error().Msg("errConnClose : " + errConnClose.Error())
		}
	}()

	if *upMigration {
		err = migrate.MigrateUp(migrationConf)
		if err != nil {
			log.Error().Msg(err.Error())
			return
		}
	} else if *downMigration {
		err = migrate.MigrateDown(migrationConf)
		if err != nil {
			log.Error().Msg(err.Error())
			return
		}
	} else if *versionMigration {
		err = migrate.PrintMigrationVersion(migrationConf)
		if err != nil {
			log.Error().Msg(err.Error())
			return
		}
	}
}
