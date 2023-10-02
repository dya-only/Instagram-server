package controllers

import "github.com/gofiber/fiber/v2"

func GetAvatarImg(c *fiber.Ctx) error {
	file := c.Params("file")

	return c.SendFile("public/uploads/avatars/" + file)
}

func GetPostImg(c *fiber.Ctx) error {
	file := c.Params("file")

	return c.SendFile("public/uploads/posts/" + file)
}

func GetDefaultAvatar(c *fiber.Ctx) error {
	return c.SendFile("public/uploads/avatars/default.jpg")
}

func GetDefaultPost(c *fiber.Ctx) error {
	return c.SendFile("public/uploads/posts/default.png")
}
