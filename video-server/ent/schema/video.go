package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
	"github.com/google/uuid"
)

// Video holds the schema definition for the Video entity.
type Video struct {
	ent.Schema
}

// Mixin video mixin
func (Video) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}

// Fields of the Video.
func (Video) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("name").NotEmpty(),
	}
}

// Edges of the Video.
func (Video) Edges() []ent.Edge {
	return []ent.Edge{
		// from
		edge.To("comments", Comment.Type).StorageKey(edge.Column("video_id")),
		// to
		edge.From("author", User.Type).Ref("videos").Unique(),
	}
}
