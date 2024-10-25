package main

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type Activity struct {
    ID int `json:"id"`
    Title string `json:"title"`
    Category string `json:"category"`
    Description string `json:"description"`
    ActivityDate time.Time `json:"activity_date"`
    Status string `json:"status"`
    CreatedAt time.Time `json:"created_at"`
}

func InitDB() (*sql.DB, error) {
    dns := "user=postgres.jmilgpvsyrktrjqfzekq password=muharromi12 host=aws-0-ap-southeast-1.pooler.supabase.com port=6543 dbname=postgres";
    db, err := sql.Open("postgres", dns)
    if err != nil {
        return nil,err
    }

    err = db.Ping()
    if err != nil {
        db.Close()
        return nil,err
    }
    return db, nil
}

func main() {
    db, err := InitDB()
    if  err != nil {
        panic(err)
    }
    defer db.Close()

    app := fiber.New()

    app.Listen(":8081")
}