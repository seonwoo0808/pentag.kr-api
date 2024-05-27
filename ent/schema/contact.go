package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Contact holds the schema definition for the Contact entity.
type Contact struct {
	ent.Schema
}

// Fields of the Contact.
func (Contact) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").
			NotEmpty().
			MaxLen(50),
		field.String("last_name").
			NotEmpty().
			MaxLen(50),
		field.String("email").
			NotEmpty().
			MaxLen(500),
		field.String("phone").
			NotEmpty().
			MaxLen(30),
		field.Int("category").
			Range(0, 2),
		field.String("message").
			NotEmpty().
			MaxLen(5000),
		field.String("verify_code").
			NotEmpty().
			MaxLen(10),
		field.Bool("is_verified").
			Default(false),
	}
}

// Edges of the Contact.
func (Contact) Edges() []ent.Edge {
	return nil
}
