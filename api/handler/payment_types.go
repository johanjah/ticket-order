package handler

import (
	"errors"
	"fmt"
	payment_type "github.com/go-rel/gin-example/payment-type"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"go.uber.org/zap"
)

// PaymentType for paymentTypes endpoints.
type PaymentType struct {
	repository  rel.Repository
	paymentType payment_type.Service
}

// Index handle GET /.
func (t PaymentType) Index(c *gin.Context) {
	var (
		result []payment_type.PaymentType
		filter = payment_type.Filter{
			PaymentName:        c.Query("payment_name"),
			PaymentDescription: c.Query("payment_description"),
		}
	)

	t.paymentType.Search(c, &result, filter)
	render(c, result, 200)
}

// Create handle POST /
func (t PaymentType) Create(c *gin.Context) {
	var (
		paymentType payment_type.PaymentType
	)

	if err := c.ShouldBindJSON(&paymentType); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(c, ErrBadRequest, 400)
		return
	}

	if err := t.paymentType.Create(c, &paymentType); err != nil {
		render(c, err, 422)
		return
	}

	c.Header("Location", fmt.Sprint(c.Request.RequestURI, "/", paymentType.ID))
	render(c, paymentType, 201)
}

// Show handle GET /{ID}
func (t PaymentType) Show(c *gin.Context) {
	var (
		paymentType = c.MustGet(loadKey).(payment_type.PaymentType)
	)

	render(c, paymentType, 200)
}

// Update handle PATCH /{ID}
func (t PaymentType) Update(c *gin.Context) {
	var (
		paymentType = c.MustGet(loadKey).(payment_type.PaymentType)
		changes     = rel.NewChangeset(&paymentType)
	)

	if err := c.ShouldBindJSON(&paymentType); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(c, ErrBadRequest, 400)
		return
	}

	if err := t.paymentType.Update(c, &paymentType, changes); err != nil {
		render(c, err, 422)
		return
	}

	render(c, paymentType, 200)
}

// Destroy handle DELETE /{ID}
func (t PaymentType) Destroy(c *gin.Context) {
	var (
		paymentType = c.MustGet(loadKey).(payment_type.PaymentType)
	)

	t.paymentType.Delete(c, &paymentType)
	render(c, nil, 204)
}

// Clear handle DELETE /
func (t PaymentType) Clear(c *gin.Context) {
	t.paymentType.Clear(c)
	render(c, nil, 204)
}

// Load is middleware that loads roless to context.
func (t PaymentType) Load(c *gin.Context) {
	var (
		id, _       = strconv.Atoi(c.Param("ID"))
		paymentType payment_type.PaymentType
	)

	if err := t.repository.Find(c, &paymentType, where.Eq("id", id)); err != nil {
		if errors.Is(err, rel.ErrNotFound) {
			render(c, err, 404)
			c.Abort()
			return
		}
		panic(err)
	}

	c.Set(loadKey, paymentType)
	c.Next()
}

// Mount handlers to router group.
func (t PaymentType) Mount(router *gin.RouterGroup) {
	router.GET("/", t.Index)
	router.POST("/", t.Create)
	router.GET("/:ID", t.Load, t.Show)
	router.PATCH("/:ID", t.Load, t.Update)
	router.DELETE("/:ID", t.Load, t.Destroy)
	router.DELETE("/", t.Clear)
}

// NewPaymentType handler.
func NewPaymentType(repository rel.Repository, paymentType payment_type.Service) PaymentType {
	return PaymentType{
		repository:  repository,
		paymentType: paymentType,
	}
}
