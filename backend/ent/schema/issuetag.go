package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// IssueTag holds the schema definition for the IssueTag entity.
type IssueTag struct {
	ent.Schema
}

// Fields of the IssueTag.
func (IssueTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MinLen(1).
			MaxLen(15).
			Unique(),
	}
}

// Edges of the IssueTag.
func (IssueTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("assignedIssues", Issue.Type).
			Ref("assignedTags"),
	}
}
