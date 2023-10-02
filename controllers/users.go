package controllers

import (
	"bytes"
	"go-template/ent/user"
	_user "go-template/ent/user"
	"go-template/models"
	"go-template/utils"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
)

var S3Client *s3.Client

func CreateUser(c *fiber.Ctx) error {
	var body models.User

	if err := c.BodyParser(&body); err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   "Failed to parse body",
		})
	}

	if exists, _ := utils.DbConn.User.Query().Where(_user.NameEQ(body.Name)).Exist(ctx); exists {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   "Name already exists",
		})
	}

	user, err := utils.DbConn.User.
		Create().
		SetAvatar("default.jpg").
		SetEmail(body.Email).
		SetUsername(body.UserName).
		SetName(body.Name).
		SetPassword(body.Password).
		SetBookmarks("[]").
		SetLikes("[]").
		SetFollower("[]").
		SetFollowing("[]").
		SetCreateAt(time.Now()).
		SetUpdateAt(time.Now()).
		Save(ctx)

	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"body":    user,
	})
}

func GetAllUser(c *fiber.Ctx) error {
	users, err := utils.DbConn.User.Query().All(ctx)

	if users == nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   "users is nil",
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
		"body":    users,
	})
}

func GetUserById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user, err := utils.DbConn.User.Get(ctx, id)

	if user == nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   "Not found user",
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
		"body":    user,
	})
}

func GetUserByName(c *fiber.Ctx) error {
	name := c.Params("name")
	user, err := utils.DbConn.User.Query().Where(user.Username(name)).All(ctx)

	if len(user) < 1 {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   "Not found user",
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
		"body":    user[0],
	})
}

func UpdateAvatar(c *fiber.Ctx) error {
	file, _ := c.FormFile("avatar")
	var body models.Update
	now := time.Now()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	client := s3.NewFromConfig(cfg)

	if err := c.BodyParser(&body); err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	if file == nil {
		user, _ := utils.DbConn.User.
			UpdateOneID(int(body.Id)).
			SetAvatar("default.jpg").
			Save(ctx)

		return c.JSON(fiber.Map{
			"success": true,
			"body":    user,
		})
	}

	fileBytes := make([]byte, file.Size)
	fileReader, _ := file.Open()
	_, _ = fileReader.Read(fileBytes)

	s3File := &s3.PutObjectInput{
		Bucket: aws.String("insta-clone-s3-bucket"),
		Key:    aws.String(now.Format("2006-01-02") + file.Filename),
		Body:   bytes.NewReader(fileBytes),
		ContentType: aws.String(utils.GetContentType(file.Filename)),
	}
	_, err = client.PutObject(ctx, s3File)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	// Save image on Disk
	// err := c.SaveFile(file, "public/uploads/avatars/"+now.Format("2006-01-02")+file.Filename)
	// if err != nil || file == nil {
	// 	return c.JSON(fiber.Map{
	// 		"success": false,
	// 		"error":   err,
	// 	})
	// }

	user, _ := utils.DbConn.User.
		UpdateOneID(int(body.Id)).
		SetAvatar(now.Format("2006-01-02") + file.Filename).
		Save(ctx)

	return c.JSON(fiber.Map{
		"success": true,
		"body":    user,
	})
}
