package main

import (
	"github.com/go-rel/rel"
)

func MigrateCreateRole(schema *rel.Schema) {
	schema.CreateTable("roles", func(t *rel.Table) {
		t.ID("id", rel.Primary(true), rel.Unique(true))
		t.DateTime("created_at")
		t.DateTime("updated_at")
		t.String("role_name", rel.Limit(20), rel.Required(true))
	})
}

func RollbackCreateRole(schema *rel.Schema) {
	schema.DropTable("roles")
}
