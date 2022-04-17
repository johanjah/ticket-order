package handler

import (
	"errors"
	"fmt"
	"github.com/go-rel/gin-example/events"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"go.uber.org/zap"
)

// Event for events endpoints.
type Event struct {
	repository rel.Repository
	event      events.Service
}

// Index handle GET /.
func (t Event) Index(c *gin.Context) {
	var (
		result []events.Event
		filter = events.Filter{
			EventName:        c.Query("event_name"),
			EventDescription: c.Query("event_description"),
		}
	)

	t.event.Search(c, &result, filter)
	render(c, result, 200)
}

// Create handle POST /
func (t Event) Create(c *gin.Context) {
	var (
		event events.Event
	)

	if err := c.ShouldBindJSON(&event); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(c, ErrBadRequest, 400)
		return
	}

	if err := t.event.Create(c, &event); err != nil {
		render(c, err, 422)
		return
	}

	c.Header("Location", fmt.Sprint(c.Request.RequestURI, "/", event.ID))
	render(c, event, 201)
}

// Show handle GET /{ID}
func (t Event) Show(c *gin.Context) {
	var (
		event = c.MustGet(loadKey).(events.Event)
	)

	render(c, event, 200)
}

// Update handle PATCH /{ID}
func (t Event) Update(c *gin.Context) {
	var (
		event   = c.MustGet(loadKey).(events.Event)
		changes = rel.NewChangeset(&event)
	)

	if err := c.ShouldBindJSON(&event); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(c, ErrBadRequest, 400)
		return
	}

	if err := t.event.Update(c, &event, changes); err != nil {
		render(c, err, 422)
		return
	}

	render(c, event, 200)
}

// Destroy handle DELETE /{ID}
func (t Event) Destroy(c *gin.Context) {
	var (
		event = c.MustGet(loadKey).(events.Event)
	)

	t.event.Delete(c, &event)
	render(c, nil, 204)
}

// Clear handle DELETE /
func (t Event) Clear(c *gin.Context) {
	t.event.Clear(c)
	render(c, nil, 204)
}

// Load is middleware that loads eventss to context.
func (t Event) Load(c *gin.Context) {
	var (
		id, _ = strconv.Atoi(c.Param("ID"))
		event events.Event
	)

	if err := t.repository.Find(c, &event, where.Eq("id", id)); err != nil {
		if errors.Is(err, rel.ErrNotFound) {
			render(c, err, 404)
			c.Abort()
			return
		}
		panic(err)
	}

	c.Set(loadKey, event)
	c.Next()
}

// Mount handlers to router group.
func (t Event) Mount(router *gin.RouterGroup) {
	router.GET("/", t.Index)
	router.POST("/", t.Create)
	router.GET("/:ID", t.Load, t.Show)
	router.PATCH("/:ID", t.Load, t.Update)
	router.DELETE("/:ID", t.Load, t.Destroy)
	router.DELETE("/", t.Clear)
}

// NewEvent handler.
func NewEvent(repository rel.Repository, event events.Service) Event {
	return Event{
		repository: repository,
		event:      event,
	}
}
