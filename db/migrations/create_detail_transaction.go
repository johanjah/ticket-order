package main

import (
	"github.com/go-rel/rel"
)

func MigrateCreateDetailTransaction(schema *rel.Schema) {
	schema.CreateTable("detail_transactions", func(t *rel.Table) {
		t.ID("id", rel.Primary(true), rel.Unique(true))
		t.DateTime("created_at")
		t.DateTime("updated_at")
		t.Decimal("price", rel.Required(true))
		t.Decimal("total_price", rel.Required(true))
		t.Int("ticket_count", rel.Required(true))
		t.Int("payment_transaction_id", rel.Required(true), rel.Unsigned(true))

		t.ForeignKey("payment_transaction_id", "payment_transactions", "id")
	})
}

func RollbackCreateDetailTransaction(schema *rel.Schema) {
	schema.DropTable("detail_transaction")
}
