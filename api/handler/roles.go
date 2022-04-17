package handler

import (
	"errors"
	"fmt"
	"github.com/go-rel/gin-example/roles"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"go.uber.org/zap"
)

type ctx int

const (
	loadKey string = "todosLoadKey"
)

// Role for roles endpoints.
type Role struct {
	repository rel.Repository
	role       roles.Service
}

// Index handle GET /.
func (t Role) Index(c *gin.Context) {
	var (
		result []roles.Role
		filter = roles.Filter{
			Keyword: c.Query("role_name"),
		}
	)

	t.role.Search(c, &result, filter)
	render(c, result, 200)
}

// Create handle POST /
func (t Role) Create(c *gin.Context) {
	var (
		role roles.Role
	)

	if err := c.ShouldBindJSON(&role); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(c, ErrBadRequest, 400)
		return
	}

	if err := t.role.Create(c, &role); err != nil {
		render(c, err, 422)
		return
	}

	c.Header("Location", fmt.Sprint(c.Request.RequestURI, "/", role.ID))
	render(c, role, 201)
}

// Show handle GET /{ID}
func (t Role) Show(c *gin.Context) {
	var (
		role = c.MustGet(loadKey).(roles.Role)
	)

	render(c, role, 200)
}

// Update handle PATCH /{ID}
func (t Role) Update(c *gin.Context) {
	var (
		role    = c.MustGet(loadKey).(roles.Role)
		changes = rel.NewChangeset(&role)
	)

	if err := c.ShouldBindJSON(&role); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(c, ErrBadRequest, 400)
		return
	}

	if err := t.role.Update(c, &role, changes); err != nil {
		render(c, err, 422)
		return
	}

	render(c, role, 200)
}

// Destroy handle DELETE /{ID}
func (t Role) Destroy(c *gin.Context) {
	var (
		role = c.MustGet(loadKey).(roles.Role)
	)

	t.role.Delete(c, &role)
	render(c, nil, 204)
}

// Clear handle DELETE /
func (t Role) Clear(c *gin.Context) {
	t.role.Clear(c)
	render(c, nil, 204)
}

// Load is middleware that loads roless to context.
func (t Role) Load(c *gin.Context) {
	var (
		id, _ = strconv.Atoi(c.Param("ID"))
		role  roles.Role
	)

	if err := t.repository.Find(c, &role, where.Eq("id", id)); err != nil {
		if errors.Is(err, rel.ErrNotFound) {
			render(c, err, 404)
			c.Abort()
			return
		}
		panic(err)
	}

	c.Set(loadKey, role)
	c.Next()
}

// Mount handlers to router group.
func (t Role) Mount(router *gin.RouterGroup) {
	router.GET("/", t.Index)
	router.POST("/", t.Create)
	router.GET("/:ID", t.Load, t.Show)
	router.PATCH("/:ID", t.Load, t.Update)
	router.DELETE("/:ID", t.Load, t.Destroy)
	router.DELETE("/", t.Clear)
}

// NewRole handler.
func NewRole(repository rel.Repository, role roles.Service) Role {
	return Role{
		repository: repository,
		role:       role,
	}
}
