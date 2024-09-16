package main

import (
	"fmt"
	"go-fiber-postgres/models"
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
	// Initialize an empty Book struct to hold incoming JSON data
	book := Book{}

	// Try to parse the JSON request body and map it to the book struct
	// context.BodyParser automatically reads the body and unmarshals it into the provided struct
	err := context.BodyParser(&book)

	// If there is an error while parsing the request body, return a 422 Unprocessable Entity status
	// and respond with an error message

	if err != nil {
		// Respond with 422 status code and error message if parsing fails
		context.Status(http.StatusUnprocessableEntity).
			JSON(&fiber.Map{"message": "could not create book"})
		return err
	}

	// If parsing succeeds, return a 200 OK status and respond with a success message
	context.Status(http.StatusOK).
		JSON(&fiber.Map{"message": "book has been added"})

	// Return nil indicating the function executed successfully and no error occurred
	return nil
}

func (r *Repository) DeleteBook(id string) {

}

func (r *Repository) GetBookByID(id string) {

}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	// Define a slice to hold the books
	bookModels := []models.Books{}

	/*
		Why Slice is Best here:
		Dynamic Size: Slices can grow and shrink as needed, making them perfect for querying a database where you don’t know how many records will be returned.
		Sequential Data: Slices maintain the order of elements, which is often important when fetching data (e.g., sorting by title or published year).
		Optimized for Iteration: Slices can be easily iterated over using Go's for loops, making it easy to process each book after fetching them.
		ORM Compatibility: Go’s popular ORM libraries (e.g., GORM) expect a slice to populate multiple records, making it a natural fit.
	*/

	// Use the ORM to find all books and populate the bookModels slice
	err := r.DB.Find(&bookModels).Error // Pass the slice as a pointer

	// Check if there was an error during the database query
	if err != nil {
		// Log the detailed error internally (use a proper logging library in production)
		fmt.Println("Error fetching books from database:", err.Error()) // Example of internal logging

		// If there is an error, respond with a 400 Bad Request status and a generic error message
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"})
		return err // Returning the error for internal handling (e.g., middleware)
	}

	// If successful, respond with a 200 OK status and the fetched books data
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Fetched books successfully",
		"data":    bookModels,
	})

	return nil
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
