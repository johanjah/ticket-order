package main

import (
	"github.com/go-rel/rel"
)

func MigrateCreatePaymentType(schema *rel.Schema) {
	schema.CreateTable("payment_types", func(t *rel.Table) {
		t.ID("id", rel.Primary(true), rel.Unique(true))
		t.DateTime("created_at")
		t.DateTime("updated_at")
		t.String("payment_name", rel.Limit(100), rel.Required(true), rel.Unique(true))
		t.String("payment_description", rel.Limit(200))
	})
}

func RollbackCreatePaymentType(schema *rel.Schema) {
	schema.DropTable("payment_types")
}
