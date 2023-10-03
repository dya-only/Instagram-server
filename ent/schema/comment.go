package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("author"),
		field.Int("postid"),
		field.String("username"),
		field.String("avatar"),
		field.String("content"),
		field.Int("likes"),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}
