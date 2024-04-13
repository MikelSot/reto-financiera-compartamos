package bootstrap

import (
	"context"

	"github.com/joho/godotenv"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"

	"github.com/MikelSot/reto-financiera-compartamos/infrastructure/handler"
	"github.com/MikelSot/reto-financiera-compartamos/model"
)

func Run(boot []byte) {
	_ = godotenv.Load()

	ctx := context.Background()
	applicationName := getApplicationName()

	db := newDatabase(ctx, applicationName)
	ginEntry := newGinEntry(boot)

	ginEntry.Bootstrap(ctx)

	logger := newLogger()

	api := ginEntry.Router

	handler.InitRoutes(model.RouterSpecification{
		Api:    api,
		Logger: logger,
		DB:     db,
	})

	rkentry.GlobalAppCtx.WaitForShutdownSig()
	ginEntry.Interrupt(ctx)
}
