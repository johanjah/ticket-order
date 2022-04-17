package main

import (
	"github.com/go-rel/rel"
)

func MigrateCreateEvent(schema *rel.Schema) {
	schema.CreateTable("events", func(t *rel.Table) {
		t.ID("id", rel.Primary(true), rel.Unique(true))
		t.DateTime("created_at")
		t.DateTime("updated_at")
		t.String("event_name", rel.Limit(100), rel.Required(true), rel.Unique(true))
		t.String("event_description", rel.Limit(200))
		t.String("event_location", rel.Limit(200))
		t.DateTime("event_start_date", rel.Required(true))
		t.DateTime("event_end_date", rel.Required(true))
		t.Decimal("base_price", rel.Required(true))
	})
}

func RollbackCreateEvent(schema *rel.Schema) {
	schema.DropTable("events")
}
