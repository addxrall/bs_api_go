package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/addxrall/bs_api_go/db"
	"github.com/jackc/pgx/v5"
)

func main() {
	dbString := os.Getenv("GOOSE_DBSTRING")
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, dbString)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	books, err := queries.GetAllBooks(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(books)

}
