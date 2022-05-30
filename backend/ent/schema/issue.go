package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Issue holds the schema definition for the Issue entity.
type Issue struct {
	ent.Schema
}

// Fields of the Issue.
func (Issue) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MinLen(4).
			NotEmpty().
			MaxLen(25),
		field.String("description").
			MaxLen(256).
			NotEmpty(),
		field.Enum("status").
			Values("Open", "Closed"),
		field.Time("created_at").Default(time.Now),
		field.Int("owner_id"),
	}
}

// Edges of the Issue.
func (Issue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("issueCreator", User.Type).
			Ref("issuesCreated").
			Required().
			Field("owner_id").
			Unique(),
		edge.To("assignedTags", IssueTag.Type),
	}
}
