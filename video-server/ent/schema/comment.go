package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Mixin comment mixin
func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Text("content").NotEmpty(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).Ref("comments").Unique(),
		edge.From("video", Video.Type).Ref("comments").Unique(),
	}
}
