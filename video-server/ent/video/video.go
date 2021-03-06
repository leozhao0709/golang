// Code generated by entc, DO NOT EDIT.

package video

import (
	"time"
)

const (
	// Label holds the string label denoting the video type in the database.
	Label = "video"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"

	// Table holds the table name of the video in the database.
	Table = "videos"
	// CommentsTable is the table the holds the comments relation/edge.
	CommentsTable = "comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "comments"
	// CommentsColumn is the table column denoting the comments relation/edge.
	CommentsColumn = "video_id"
	// AuthorTable is the table the holds the author relation/edge.
	AuthorTable = "videos"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "author_id"
)

// Columns holds all SQL columns for video fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Video type.
var ForeignKeys = []string{
	"author_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the create_time field.
	DefaultCreateTime func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
