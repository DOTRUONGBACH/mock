package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"

	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{field.UUID("id", uuid.UUID{}).Annotations(entproto.Field(1)).Default(uuid.New).Annotations(entgql.OrderField("ID")),
		field.String("email").Unique().NotEmpty().Annotations(entgql.OrderField("EMAIL")),
		field.String("password").NotEmpty().Annotations(entgql.OrderField("PASSWORD")),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
