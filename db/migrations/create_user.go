package main

import (
	"github.com/go-rel/rel"
)

func MigrateCreateUser(schema *rel.Schema) {
	schema.CreateTable("users", func(t *rel.Table) {
		t.ID("id", rel.Primary(true), rel.Unique(true))
		t.DateTime("created_at")
		t.DateTime("updated_at")
		t.String("username", rel.Limit(20), rel.Required(true), rel.Unique(true))
		t.String("password", rel.Limit(100), rel.Required(true))
		t.String("email", rel.Limit(100), rel.Required(true))
		t.String("first_name", rel.Limit(100), rel.Required(true))
		t.String("last_name", rel.Limit(100), rel.Required(true))
		t.Int("role_id", rel.Required(true), rel.Unsigned(true))

		t.ForeignKey("role_id", "roles", "id")
	})
}

func RollbackCreateUser(schema *rel.Schema) {
	schema.DropTable("users")
}
