package main

import (
	"go-fiber-postgres/storage"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// wil help to talk to database
type Repository struct {
	DB *gorm.DB
}

// Decoding needed as golang itself doesn't understand json
type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

// api contract
func (r *Repository) SetupRoutes(app *fiber.App) {
	// api := app.Group("/api")
	// api.Post("/create_books", r.CreateBook) //calling a method called createBook
	// api.Delete("/delete_book/:id", r.DeleteBook)
	// api.Get("/get_books/:id", r.GetBookByID)
	// api.GetAll("/books", r.GetBooks)

}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "coud not create book"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book has been added"})
	return nil
}

func (r *Repository) DeleteBook(id string) {

}

func (r *Repository) GetBookByID(id string) {

}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config) //from env

	if err != nil {
		log.Fatal("Could not load the database")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()  //kind of express
	r.SetupRoutes(app)  //struct methods
	app.Listen(":8080") //start server on port

}
