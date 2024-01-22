package main

import (
	"go-fiber-api/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strconv"
)

var posts []model.Post

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	posts = append(posts, model.Post{UserId: 1, Id: 1, Title: "sunt aut facere repellat provident occaecati excepturi optio reprehenderit", Body: "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"})
	posts = append(posts, model.Post{UserId: 1, Id: 2, Title: "qui est esse", Body: "est rerum tempore vitae\\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\\nqui aperiam non debitis possimus qui neque nisi nulla"})

	app.Get("/posts", getPosts)
	app.Get("/posts/:id", getPost)
	app.Post("/posts", createPost)
	app.Put("/posts/:id", updatePost)
	app.Delete("/posts/:id", deletePost)

	err := app.Listen(":8000")
	if err != nil {
		return
	}
}

func getPosts(c *fiber.Ctx) error {
	return c.JSON(posts)
}

func getPost(c *fiber.Ctx) error {
	postId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, post := range posts {
		if post.Id == postId {
			return c.JSON(post)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func createPost(c *fiber.Ctx) error {
	post := new(model.Post)
	err := c.BodyParser(post)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	posts = append(posts, *post)
	return c.JSON(post)
}

func updatePost(c *fiber.Ctx) error {
	postId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	postUpdate := new(model.Post)
	err = c.BodyParser(postUpdate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, post := range posts {
		if post.Id == postId {
			posts[i].Title = postUpdate.Title
			posts[i].Body = postUpdate.Body
			return c.JSON(posts[i])
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func deletePost(c *fiber.Ctx) error {
	postId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, post := range posts {
		if post.Id == postId {
			posts = append(posts[:i], posts[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}
