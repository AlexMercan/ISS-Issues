package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"server/controller"
	"server/ent"
	"server/middleware"
	"server/repository"

	"server/ent/migrate"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

func Open(connString string) *ent.Client {
	client, err := sql.Open("pgx", connString)
	if err != nil {
		log.Fatal(err)
	}
	driver := entsql.OpenDB(dialect.Postgres, client)
	return ent.NewClient(ent.Driver(driver))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connString := os.Getenv("POSTGRES_URL")
	client := Open(connString)

	defer client.Close()

	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	userAuthRepo := repository.CreateUserAuthRepository(client)
	userAuthController := controller.CreateUserAuthController(userAuthRepo, []byte("verysecretkey"))
	userRepository := repository.CreateUserRepository(client)
	userController := controller.CreateUserController(userRepository)
	issueRepository := repository.CreateIssueRepository(client)
	issueController := controller.CreateIssueController(issueRepository)
	issueTagRepository := repository.CreateIssueTagRepository(client)
	issueTagController := controller.CreateIssueTagController(issueTagRepository)

	r := chi.NewRouter()
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	r.Route("/auth", func(r chi.Router) {
		r.Use(chiMiddleware.RequestID)
		r.Use(chiMiddleware.RealIP)
		r.Use(chiMiddleware.Logger)
		r.Use(chiMiddleware.Recoverer)
		r.Post("/login", userAuthController.Login)
		r.Post("/register", userAuthController.Register)
	})
	r.Route("/api", func(r chi.Router) {
		r.Use(chiMiddleware.RequestID)
		r.Use(chiMiddleware.RealIP)
		r.Use(chiMiddleware.Logger)
		r.Use(chiMiddleware.Recoverer)
		r.Use(middleware.AuthMiddleware)
		r.Get("/issues", issueController.GetIssues)
        r.Get("/user", userController.GetUser)
		r.Post("/issues/{issueId}", issueController.UpdateIssue)
		r.Get("/issues/{issueId}", issueController.GetIssue)
		r.Post("/issues", issueController.SaveIssue)
		r.Get("/issuetags", issueTagController.GetAll)
	})

	err = http.ListenAndServe(":5000", r)
	if err != nil {
		panic(err)
	}
}
