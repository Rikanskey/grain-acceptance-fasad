package runner

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"grain-acceptance/internal/app"
	"grain-acceptance/internal/app/commands"
	"grain-acceptance/internal/app/queries"
	"grain-acceptance/internal/config"
	"grain-acceptance/internal/dao"
	"grain-acceptance/internal/port/http"
	"grain-acceptance/internal/server"
	"log"
)

func Start(configDir string) {
	cfg := newConfig(configDir)
	db := initDB(cfg)
	application := newApplication(db)
	startServer(cfg, application)
}

func newConfig(configDir string) *config.Config {
	cfg, err := config.New(configDir)
	if err != nil {
		log.Panicln(err)
	}

	return cfg
}

func initDB(cfg *config.Config) *sql.DB {
	//dbInfo := fmt.Sprintf("port=5432 user=%s password=%s dbname=%s sslmode=disable",
	//	cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Name)
	dbInfo := fmt.Sprintf("postgresql://%s:%s@postgres/%s?sslmode=disable",
		cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Name)
	db, err := sql.Open("postgres", dbInfo)
	//conn, er : = pgx.Connect(context.Background(), dbInfo)
	//defer conn.Close(context.Background())
	//if conn != nil && er == nil {
	//	fmt.Println("wew")
	//}
	//logrus.Print()

	if err != nil {
		fmt.Println("DB err")
	}
	err = db.Ping()
	if err != nil {
		logrus.WithError(err).Println("No connection")
	}
	//defer db.Close()
	return db
}

func newApplication(db *sql.DB) app.Application {
	//cardRepository := dao.NewCardRepository(db)
	consignmentRepository := dao.NewConsignmentRepository(db)
	fmt.Println(consignmentRepository)

	return app.Application{
		Commands: app.Commands{
			UpdateConsignmentAnalyse:         commands.NewUpdateConsignmentAnalyseHandler(consignmentRepository),
			UpdateConsignmentGrossWeight:     commands.NewUpdateConsignmentGrossWeightHandler(consignmentRepository),
			UpdateConsignmentTransportWeight: commands.NewUpdateConsignmentTransportWeightHandler(consignmentRepository),
		},
		Queries: app.Queries{
			GetConsignment: queries.NewGetConsignmentHandler(consignmentRepository),
		},
	}
}

func startServer(cfg *config.Config, application app.Application) {
	logrus.Info(fmt.Sprintf("Starting HTTP server on address: %s", cfg.HTTP.Port))

	httpServer := server.New(cfg, http.NewHandler(application))

	err := httpServer.Run()

	logrus.WithError(err).Fatal("HTTP server stopped")
}
