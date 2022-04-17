package main

import (
	"github.com/go-rel/rel"
)

func MigrateCreateShoppingChart(schema *rel.Schema) {
	schema.CreateTable("shopping_charts", func(t *rel.Table) {
		t.ID("id", rel.Primary(true), rel.Unique(true))
		t.DateTime("created_at")
		t.DateTime("updated_at")
		t.Int("ticket_count", rel.Required(true))
		t.Int("user_id", rel.Required(true), rel.Unsigned(true))
		t.Int("event_id", rel.Required(true), rel.Unsigned(true))

		t.ForeignKey("user_id", "user", "id")
		t.ForeignKey("event_id", "event", "id")
	})
}

func RollbackCreateShoppingChart(schema *rel.Schema) {
	schema.DropTable("shopping_charts")
}
