package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("avatar"),
		field.String("email").Unique(),
		field.String("name"),
		field.String("username").Unique(),
		field.String("password"),
		field.String("bookmarks"),
		field.String("likes"),
		field.String("follower"),
		field.String("following"),
		field.Time("create_at").Default(time.Now),
		field.Time("update_at").UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
