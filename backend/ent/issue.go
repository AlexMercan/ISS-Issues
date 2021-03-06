// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"server/ent/issue"
	"server/ent/user"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Issue is the model entity for the Issue schema.
type Issue struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Status holds the value of the "status" field.
	Status issue.Status `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID int `json:"owner_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IssueQuery when eager-loading is set.
	Edges IssueEdges `json:"edges"`
}

// IssueEdges holds the relations/edges for other nodes in the graph.
type IssueEdges struct {
	// IssueCreator holds the value of the issueCreator edge.
	IssueCreator *User `json:"issueCreator,omitempty"`
	// AssignedTags holds the value of the assignedTags edge.
	AssignedTags []*IssueTag `json:"assignedTags,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// IssueCreatorOrErr returns the IssueCreator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e IssueEdges) IssueCreatorOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.IssueCreator == nil {
			// The edge issueCreator was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.IssueCreator, nil
	}
	return nil, &NotLoadedError{edge: "issueCreator"}
}

// AssignedTagsOrErr returns the AssignedTags value or an error if the edge
// was not loaded in eager-loading.
func (e IssueEdges) AssignedTagsOrErr() ([]*IssueTag, error) {
	if e.loadedTypes[1] {
		return e.AssignedTags, nil
	}
	return nil, &NotLoadedError{edge: "assignedTags"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Issue) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case issue.FieldID, issue.FieldOwnerID:
			values[i] = new(sql.NullInt64)
		case issue.FieldName, issue.FieldDescription, issue.FieldStatus:
			values[i] = new(sql.NullString)
		case issue.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Issue", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Issue fields.
func (i *Issue) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case issue.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case issue.FieldName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[j])
			} else if value.Valid {
				i.Name = value.String
			}
		case issue.FieldDescription:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[j])
			} else if value.Valid {
				i.Description = value.String
			}
		case issue.FieldStatus:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[j])
			} else if value.Valid {
				i.Status = issue.Status(value.String)
			}
		case issue.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case issue.FieldOwnerID:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[j])
			} else if value.Valid {
				i.OwnerID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryIssueCreator queries the "issueCreator" edge of the Issue entity.
func (i *Issue) QueryIssueCreator() *UserQuery {
	return (&IssueClient{config: i.config}).QueryIssueCreator(i)
}

// QueryAssignedTags queries the "assignedTags" edge of the Issue entity.
func (i *Issue) QueryAssignedTags() *IssueTagQuery {
	return (&IssueClient{config: i.config}).QueryAssignedTags(i)
}

// Update returns a builder for updating this Issue.
// Note that you need to call Issue.Unwrap() before calling this method if this Issue
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Issue) Update() *IssueUpdateOne {
	return (&IssueClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Issue entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Issue) Unwrap() *Issue {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Issue is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Issue) String() string {
	var builder strings.Builder
	builder.WriteString("Issue(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteString(", name=")
	builder.WriteString(i.Name)
	builder.WriteString(", description=")
	builder.WriteString(i.Description)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", i.Status))
	builder.WriteString(", created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", owner_id=")
	builder.WriteString(fmt.Sprintf("%v", i.OwnerID))
	builder.WriteByte(')')
	return builder.String()
}

// Issues is a parsable slice of Issue.
type Issues []*Issue

func (i Issues) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
