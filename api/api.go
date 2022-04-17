package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/go-rel/gin-example/api/middleware"
	"github.com/go-rel/gin-example/events"
	payment_type "github.com/go-rel/gin-example/payment-type"
	"github.com/go-rel/gin-example/roles"
	"github.com/go-rel/gin-example/shoppingcharts"
	"github.com/go-rel/gin-example/users"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-contrib/sessions/cookie"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/go-rel/gin-example/api/handler"
	"github.com/go-rel/gin-example/scores"
	"github.com/go-rel/gin-example/todos"
	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

// New api.
func New(repository rel.Repository) *gin.Engine {
	var (
		logger, _      = zap.NewProduction()
		router         = gin.New()
		scores         = scores.New(repository)
		todos          = todos.New(repository, scores)
		healthzHandler = handler.NewHealthz()
		todosHandler   = handler.NewTodos(repository, todos)
		scoreHandler   = handler.NewScore(repository)

		role         = roles.New(repository)
		rolesHandler = handler.NewRole(repository, role)

		user        = users.New(repository)
		userHandler = handler.NewUser(repository, user)

		paymentType        = payment_type.New(repository)
		paymentTypeHandler = handler.NewPaymentType(repository, paymentType)

		event        = events.New(repository)
		eventHandler = handler.NewEvent(repository, event)

		shoppingChart        = shoppingcharts.New(repository)
		shoppingChartHandler = handler.NewShoppingChart(repository, shoppingChart)
	)

	healthzHandler.Add("database", repository)

	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(requestid.New())
	router.Use(cors.Default())

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.Use(middleware.Authentication())

	healthzHandler.Mount(router.Group("/healthz"))
	todosHandler.Mount(router.Group("/todos"))
	scoreHandler.Mount(router.Group("/score"))
	rolesHandler.Mount(router.Group("/roles"))
	userHandler.Mount(router.Group("/users"))
	paymentTypeHandler.Mount(router.Group("/paymentTypes"))
	eventHandler.Mount(router.Group("/events"))
	shoppingChartHandler.Mount(router.Group("/shoppingCharts"))

	return router
}
