package database

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/4nkitd/gobackup/helper"
	"github.com/4nkitd/gobackup/logger"
)

// PostgreSQL database
//
// type: postgresql
// host: localhost
// port: 5432
// database: test
// username:
// password:
type PostgreSQL struct {
	Base
	host        string
	port        string
	database    string
	username    string
	password    string
	dumpCommand string
}

func (ctx PostgreSQL) perform() (err error) {
	viper := ctx.viper
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 5432)

	ctx.host = viper.GetString("host")
	ctx.port = viper.GetString("port")
	ctx.database = viper.GetString("database")
	ctx.username = viper.GetString("username")
	ctx.password = viper.GetString("password")

	if err = ctx.prepare(); err != nil {
		return
	}

	err = ctx.dump()
	return
}

func (ctx *PostgreSQL) prepare() (err error) {
	// mysqldump command
	dumpArgs := []string{}
	if len(ctx.database) == 0 {
		return fmt.Errorf("PostgreSQL database config is required")
	}
	if len(ctx.host) > 0 {
		dumpArgs = append(dumpArgs, "--host="+ctx.host)
	}
	if len(ctx.port) > 0 {
		dumpArgs = append(dumpArgs, "--port="+ctx.port)
	}
	if len(ctx.username) > 0 {
		dumpArgs = append(dumpArgs, "--username="+ctx.username)
	}

	ctx.dumpCommand = "pg_dump " + strings.Join(dumpArgs, " ") + " " + ctx.database

	return nil
}

func (ctx *PostgreSQL) dump() error {
	dumpFilePath := path.Join(ctx.dumpPath, ctx.database+".sql")
	logger.Info("-> Dumping PostgreSQL...")
	if len(ctx.password) > 0 {
		os.Setenv("PGPASSWORD", ctx.password)
	}
	_, err := helper.Exec(ctx.dumpCommand, "-f", dumpFilePath)
	if err != nil {
		return err
	}
	logger.Info("dump path:", dumpFilePath)
	return nil
}
