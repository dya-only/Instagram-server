package controllers

import (
	"go-template/ent/comment"
	"go-template/models"
	"go-template/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateComment(c *fiber.Ctx) error {
	var body models.Comment

	if err := c.BodyParser(&body); err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	post, err := utils.DbConn.Comment.
		Create().
		SetAuthor(body.Author).
		SetPostid(body.PostId).
		SetContent(body.Content).
		SetLikes(0).
		Save(ctx)

	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"body":    post,
	})
}

func GetAllComment(c *fiber.Ctx) error {
	comments, err := utils.DbConn.Comment.Query().All(ctx)

	if comments == nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   "comments is nil",
		})
	}
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"body":    comments,
	})
}

func GetCommentByPost(c *fiber.Ctx) error {
	postid, _ := c.ParamsInt("postid")

	comments, err := utils.DbConn.Comment.Query().Where(comment.Postid(int(postid))).All(ctx)

	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"body":    comments,
	})
}

func UpdateLikes(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	comment, err := utils.DbConn.Comment.Get(ctx, id)

	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	updated, err := utils.DbConn.Comment.
		UpdateOneID(id).
		SetLikes(comment.Likes + 1).
		Save(ctx)

	if comment == nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   "Not found comment",
		})
	}
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"body":    updated,
	})
}
