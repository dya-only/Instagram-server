package controllers

import (
	"bytes"
	"encoding/json"
	"go-template/ent/post"
	"go-template/models"
	"go-template/utils"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
)

func CreatePost(c *fiber.Ctx) error {
	file, _ := c.FormFile("img")
	var body models.Post
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
			"err":     err,
		})
	}

	fileBytes := make([]byte, file.Size)
	fileReader, _ := file.Open()
	_, _ = fileReader.Read(fileBytes)

	s3File := &s3.PutObjectInput{
		Bucket:      aws.String("insta-clone-s3-bucket"),
		Key:         aws.String(now.Format("2006-01-02") + file.Filename),
		Body:        bytes.NewReader(fileBytes),
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
	// err := c.SaveFile(file, "public/uploads/posts/"+now.Format("2006-01-02")+file.Filename)
	// if err != nil || file == nil {
	// 	return c.JSON(fiber.Map{
	// 		"success": false,
	// 		"error":   err,
	// 	})
	// }

	post, _ := utils.DbConn.Post.
		Create().
		SetImg(now.Format("2006-01-02") + file.Filename).
		SetContent(body.Content).
		SetAuthor(int(body.Author)).
		SetLikes(0).
		SetCreateAt(time.Now()).
		SetUpdateAt(time.Now()).
		Save(ctx)

	return c.JSON(fiber.Map{
		"success": true,
		"body":    post,
	})
}

func GetAllPost(c *fiber.Ctx) error {
	posts, err := utils.DbConn.Post.Query().All(ctx)

	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"body":    posts,
	})
}

func GetPostById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	post, err := utils.DbConn.Post.Get(ctx, id)

	if post == nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
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
		"body":    post,
	})
}

func GetPostMe(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	posts, err := utils.DbConn.Post.Query().Where(post.Author(int(id))).All(ctx)

	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"body":    posts,
	})
}

func GetPostRecommend(c *fiber.Ctx) error {
	next, _ := c.ParamsInt("next")

	posts, err := utils.DbConn.Post.Query().Where().Offset(next * 10).Limit(10).All(ctx)

	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"body":    posts,
		"next":    next + 1,
	})
}

func AddLike(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	userid, _ := c.ParamsInt("userid")

	user, _ := utils.DbConn.User.Get(ctx, userid)
	var likes []int
	err := json.Unmarshal([]byte(user.Likes), &likes)

	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	likes = append(likes, id)
	likesArr, _ := json.Marshal(likes)
	user.
		Update().
		SetLikes(string(likesArr)).
		Save(ctx)

	_likes, _ := utils.DbConn.Post.Get(ctx, id)
	post, _ := utils.DbConn.Post.
		UpdateOneID(id).
		SetLikes(_likes.Likes + 1).
		Save(ctx)

	return c.JSON(fiber.Map{
		"success": true,
		"body":    post,
	})
}

func RemoveLike(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	userid, _ := c.ParamsInt("userid")

	user, _ := utils.DbConn.User.Get(ctx, userid)
	var _likes []int
	err := json.Unmarshal([]byte(user.Likes), &_likes)

	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	var likes []int
	for _, el := range _likes {
		if el != id {
			likes = append(likes, el)
		}
	}
	likesArr, _ := json.Marshal(likes)
	user.
		Update().
		SetLikes(string(likesArr)).
		Save(ctx)

	__likes, _ := utils.DbConn.Post.Get(ctx, id)
	post, _ := utils.DbConn.Post.
		UpdateOneID(id).
		SetLikes(__likes.Likes - 1).
		Save(ctx)

	return c.JSON(fiber.Map{
		"success": true,
		"body":    post,
	})
}

func IsLiked(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	postid, _ := c.ParamsInt("postid")

	user, err := utils.DbConn.User.Get(ctx, id)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   "Not Found User",
		})
	}

	var likes []int

	_ = json.Unmarshal([]byte(user.Likes), &likes)

	return c.JSON(fiber.Map{
		"success": true,
		"body":    slices.Contains(likes, postid),
	})
}

func AddBookmark(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	userid, _ := c.ParamsInt("userid")

	user, _ := utils.DbConn.User.Get(ctx, userid)
	var bookmarks []int
	err := json.Unmarshal([]byte(user.Bookmarks), &bookmarks)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	bookmarks = append(bookmarks, id)
	bookmarksArr, _ := json.Marshal(bookmarks)
	user.
		Update().
		SetBookmarks(string(bookmarksArr)).
		Save(ctx)

	return c.JSON(fiber.Map{
		"success": true,
		"body":    user,
	})
}

func RemoveBookmark(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	userid, _ := c.ParamsInt("userid")

	user, _ := utils.DbConn.User.Get(ctx, userid)
	var bookmarks []int
	err := json.Unmarshal([]byte(user.Bookmarks), &bookmarks)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	var _bookmarks []int
	for _, el := range bookmarks {
		if el != id {
			_bookmarks = append(_bookmarks, el)
		}
	}
	bookmarksArr, _ := json.Marshal(_bookmarks)
	user.
		Update().
		SetBookmarks(string(bookmarksArr)).
		Save(ctx)

	return c.JSON(fiber.Map{
		"success": true,
		"body":    user,
	})
}

func IsBookmarked(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	postid, _ := c.ParamsInt("postid")

	user, err := utils.DbConn.User.Get(ctx, id)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   "Not Found User",
		})
	}

	var bookmarks []int

	_ = json.Unmarshal([]byte(user.Bookmarks), &bookmarks)

	return c.JSON(fiber.Map{
		"success": true,
		"body":    slices.Contains(bookmarks, postid),
	})
}
