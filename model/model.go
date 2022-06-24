package model

import (
	"os"

	"github.com/4nkitd/gobackup/archive"
	"github.com/4nkitd/gobackup/compressor"
	"github.com/4nkitd/gobackup/config"
	"github.com/4nkitd/gobackup/database"
	"github.com/4nkitd/gobackup/encryptor"
	"github.com/4nkitd/gobackup/logger"
	"github.com/4nkitd/gobackup/storage"
)

// Model class
type Model struct {
	Config config.ModelConfig
}

// Perform model
func (ctx Model) Perform() {
	logger.Info("======== " + ctx.Config.Name + " ========")
	logger.Info("WorkDir:", ctx.Config.DumpPath+"\n")

	defer func() {
		if r := recover(); r != nil {
			ctx.cleanup()
		}

		ctx.cleanup()
	}()

	err := database.Run(ctx.Config)
	if err != nil {
		logger.Error(err)
		return
	}

	if ctx.Config.Archive != nil {
		err = archive.Run(ctx.Config)
		if err != nil {
			logger.Error(err)
			return
		}
	}

	archivePath, err := compressor.Run(ctx.Config)
	if err != nil {
		logger.Error(err)
		return
	}

	archivePath, err = encryptor.Run(archivePath, ctx.Config)
	if err != nil {
		logger.Error(err)
		return
	}

	err = storage.Run(ctx.Config, archivePath)
	if err != nil {
		logger.Error(err)
		return
	}

}

// Cleanup model temp files
func (ctx Model) cleanup() {
	logger.Info("Cleanup temp: " + ctx.Config.TempPath + "/\n")
	err := os.RemoveAll(ctx.Config.TempPath)
	if err != nil {
		logger.Error("Cleanup temp dir "+ctx.Config.TempPath+" error:", err)
	}
	logger.Info("======= End " + ctx.Config.Name + " =======\n\n")
}
