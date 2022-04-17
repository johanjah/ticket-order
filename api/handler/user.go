package handler

import (
	"errors"
	"fmt"
	"github.com/go-rel/gin-example/users"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"go.uber.org/zap"
)

// User for users endpoints.
type User struct {
	repository rel.Repository
	user       users.Service
}

// Index handle GET /.
func (t User) Index(c *gin.Context) {
	var (
		result []users.User
		filter = users.Filter{
			FirstName: c.Query("first_name"),
			LastName:  c.Query("last_name"),
			Email:     c.Query("email"),
		}
	)

	t.user.Search(c, &result, filter)
	render(c, result, 200)
}

// Create handle POST /
func (t User) Create(c *gin.Context) {
	var (
		user users.User
	)

	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(c, ErrBadRequest, 400)
		return
	}

	if err := t.user.Create(c, &user); err != nil {
		render(c, err, 422)
		return
	}

	c.Header("Location", fmt.Sprint(c.Request.RequestURI, "/", user.ID))
	render(c, user, 201)
}

// Show handle GET /{ID}
func (t User) Show(c *gin.Context) {
	var (
		user = c.MustGet(loadKey).(users.User)
	)

	render(c, user, 200)
}

// Update handle PATCH /{ID}
func (t User) Update(c *gin.Context) {
	var (
		user    = c.MustGet(loadKey).(users.User)
		changes = rel.NewChangeset(&user)
	)

	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(c, ErrBadRequest, 400)
		return
	}

	if err := t.user.Update(c, &user, changes); err != nil {
		render(c, err, 422)
		return
	}

	render(c, user, 200)
}

// Destroy handle DELETE /{ID}
func (t User) Destroy(c *gin.Context) {
	var (
		user = c.MustGet(loadKey).(users.User)
	)

	t.user.Delete(c, &user)
	render(c, nil, 204)
}

// Clear handle DELETE /
func (t User) Clear(c *gin.Context) {
	t.user.Clear(c)
	render(c, nil, 204)
}

// Load is middleware that loads roless to context.
func (t User) Load(c *gin.Context) {
	var (
		id, _ = strconv.Atoi(c.Param("ID"))
		user  users.User
	)

	if err := t.repository.Find(c, &user, where.Eq("id", id)); err != nil {
		if errors.Is(err, rel.ErrNotFound) {
			render(c, err, 404)
			c.Abort()
			return
		}
		panic(err)
	}

	c.Set(loadKey, user)
	c.Next()
}

// Mount handlers to router group.
func (t User) Mount(router *gin.RouterGroup) {
	router.GET("/", t.Index)
	router.POST("/", t.Create)
	router.GET("/:ID", t.Load, t.Show)
	router.PATCH("/:ID", t.Load, t.Update)
	router.DELETE("/:ID", t.Load, t.Destroy)
	router.DELETE("/", t.Clear)
}

// NewUser handler.
func NewUser(repository rel.Repository, user users.Service) User {
	return User{
		repository: repository,
		user:       user,
	}
}
