package controllers

import (
	"context"
	"go-template/middlewares"

	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

func Users(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/user", func(c *fiber.Ctx) error {
		return CreateUser(c)
	})

	api.Get("/user", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return GetAllUser(c)
	})

	api.Get("/user/:id", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return GetUserById(c)
	})

	api.Get("/user/name/:name", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return GetUserByName(c)
	})

	api.Patch("/user/avatar", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return UpdateAvatar(c)
	})

}

func Auth(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/auth/by-pass", func(c *fiber.Ctx) error {
		return LoginByPassword(c)
	})

	api.Post("/auth/by-token", func(c *fiber.Ctx) error {
		return LoginByToken(c)
	})
}

func Posts(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/post", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return CreatePost(c)
	})

	api.Get("/post", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return GetAllPost(c)
	})

	api.Get("/post/:id", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return GetPostById(c)
	})

	api.Get("/post/only/:id", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return GetPostMe(c)
	})

	api.Get("/post/recommend/:next", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return GetPostRecommend(c)
	})

	api.Patch("/post/like/:id/:userid", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return AddLike(c)
	})

	api.Delete("/post/like/:id/:userid", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return RemoveLike(c)
	})
}

func Comments(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/comment", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return CreateComment(c)
	})

	api.Get("/comment", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return GetAllComment(c)
	})

	api.Get("/comment/:postid", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return GetCommentByPost(c)
	})

	api.Get("/comment/likes/:id", middlewares.JwtGuard, func(c *fiber.Ctx) error {
		return UpdateLikes(c)
	})
}

func Files(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/uploads/avatar", func(c *fiber.Ctx) error {
		return GetDefaultAvatar(c)
	})
	api.Get("/uploads/post", func(c *fiber.Ctx) error {
		return GetDefaultPost(c)
	})

	api.Get("/uploads/avatar/:file", func(c *fiber.Ctx) error {
		return GetAvatarImg(c)
	})

	api.Get("uploads/post/:file", func(c *fiber.Ctx) error {
		return GetPostImg(c)
	})
}
