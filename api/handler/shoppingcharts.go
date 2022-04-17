package handler

import (
	"errors"
	"fmt"
	"github.com/go-rel/gin-example/shoppingcharts"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"go.uber.org/zap"
)

// ShoppingChart for shoppingcharts endpoints.
type ShoppingChart struct {
	repository    rel.Repository
	shoppingChart shoppingcharts.Service
}

// Index handle GET /.
func (t ShoppingChart) Index(c *gin.Context) {
	var (
		result []shoppingcharts.ShoppingChart
		filter = shoppingcharts.Filter{}
	)

	t.shoppingChart.Search(c, &result, filter)
	render(c, result, 200)
}

// Create handle POST /
func (t ShoppingChart) Create(c *gin.Context) {
	var (
		shoppingChart shoppingcharts.ShoppingChart
	)

	if err := c.ShouldBindJSON(&shoppingChart); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(c, ErrBadRequest, 400)
		return
	}

	if err := t.shoppingChart.Create(c, &shoppingChart); err != nil {
		render(c, err, 422)
		return
	}

	c.Header("Location", fmt.Sprint(c.Request.RequestURI, "/", shoppingChart.ID))
	render(c, shoppingChart, 201)
}

// Show handle GET /{ID}
func (t ShoppingChart) Show(c *gin.Context) {
	var (
		shoppingChart = c.MustGet(loadKey).(shoppingcharts.ShoppingChart)
	)

	render(c, shoppingChart, 200)
}

// Update handle PATCH /{ID}
func (t ShoppingChart) Update(c *gin.Context) {
	var (
		shoppingChart = c.MustGet(loadKey).(shoppingcharts.ShoppingChart)
		changes       = rel.NewChangeset(&shoppingChart)
	)

	if err := c.ShouldBindJSON(&shoppingChart); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(c, ErrBadRequest, 400)
		return
	}

	if err := t.shoppingChart.Update(c, &shoppingChart, changes); err != nil {
		render(c, err, 422)
		return
	}

	render(c, shoppingChart, 200)
}

// Destroy handle DELETE /{ID}
func (t ShoppingChart) Destroy(c *gin.Context) {
	var (
		shoppingChart = c.MustGet(loadKey).(shoppingcharts.ShoppingChart)
	)

	t.shoppingChart.Delete(c, &shoppingChart)
	render(c, nil, 204)
}

// Clear handle DELETE /
func (t ShoppingChart) Clear(c *gin.Context) {
	t.shoppingChart.Clear(c)
	render(c, nil, 204)
}

// Load is middleware that loads roless to context.
func (t ShoppingChart) Load(c *gin.Context) {
	var (
		id, _         = strconv.Atoi(c.Param("ID"))
		shoppingChart shoppingcharts.ShoppingChart
	)

	if err := t.repository.Find(c, &shoppingChart, where.Eq("id", id)); err != nil {
		if errors.Is(err, rel.ErrNotFound) {
			render(c, err, 404)
			c.Abort()
			return
		}
		panic(err)
	}

	c.Set(loadKey, shoppingChart)
	c.Next()
}

// Mount handlers to router group.
func (t ShoppingChart) Mount(router *gin.RouterGroup) {
	router.GET("/", t.Index)
	router.POST("/", t.Create)
	router.GET("/:ID", t.Load, t.Show)
	router.PATCH("/:ID", t.Load, t.Update)
	router.DELETE("/:ID", t.Load, t.Destroy)
	router.DELETE("/", t.Clear)
}

// NewShoppingChart handler.
func NewShoppingChart(repository rel.Repository, shoppingChart shoppingcharts.Service) ShoppingChart {
	return ShoppingChart{
		repository:    repository,
		shoppingChart: shoppingChart,
	}
}
