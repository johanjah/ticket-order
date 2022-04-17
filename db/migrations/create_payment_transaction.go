package main

import (
	"github.com/go-rel/rel"
)

func MigrateCreatePaymentTransaction(schema *rel.Schema) {
	schema.CreateTable("payment_transactions", func(t *rel.Table) {
		t.ID("id", rel.Primary(true), rel.Unique(true))
		t.DateTime("created_at")
		t.DateTime("updated_at")
		t.Decimal("total_pay", rel.Required(true))
		t.DateTime("transaction_date", rel.Required(true))
		t.Int("user_id", rel.Required(true), rel.Unsigned(true))
		t.Int("payment_type_id", rel.Required(true), rel.Unsigned(true))

		t.ForeignKey("user_id", "users", "id")
		t.ForeignKey("payment_type_id", "payment_types", "id")
	})
}

func RollbackCreatePaymentTransaction(schema *rel.Schema) {
	schema.DropTable("payment_transaction")
}
