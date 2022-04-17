package main

import (
	"context"
	"fmt"
	"github.com/go-rel/gin-example/config"
	"os"

	"github.com/go-rel/migration"
	"github.com/go-rel/mysql"
	"github.com/go-rel/rel"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// load config
	config.LoadEnv()

	// create var
	var (
		ctx = context.TODO()

		dsn = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("MYSQL_USERNAME"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_DATABASE"))
		repo = rel.New(mysql.MustOpen(dsn))
		m    = migration.New(repo)
	)

	// Register migrations
	counter := 1
	m.Register(counter, MigrateCreatePaymentType, RollbackCreatePaymentType)
	counter++
	m.Register(counter, MigrateCreateRole, RollbackCreateRole)
	counter++
	m.Register(counter, MigrateCreateUser, RollbackCreateUser)
	counter++
	m.Register(counter, MigrateCreateEvent, RollbackCreateEvent)
	counter++
	m.Register(counter, MigrateCreateShoppingChart, RollbackCreateShoppingChart)
	counter++
	m.Register(counter, MigrateCreatePaymentTransaction, RollbackCreatePaymentTransaction)
	counter++
	m.Register(counter, MigrateCreateDetailTransaction, RollbackCreateDetailTransaction)

	// Run migrations
	m.Migrate(ctx)
	// OR: ( do rollback() while counter < 0
	//for counter > 0 {
	//	m.Rollback(ctx)
	//	counter--
	//}
}
