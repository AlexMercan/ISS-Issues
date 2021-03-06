// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"server/ent/issuetag"
	"server/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IssueTagDelete is the builder for deleting a IssueTag entity.
type IssueTagDelete struct {
	config
	hooks    []Hook
	mutation *IssueTagMutation
}

// Where appends a list predicates to the IssueTagDelete builder.
func (itd *IssueTagDelete) Where(ps ...predicate.IssueTag) *IssueTagDelete {
	itd.mutation.Where(ps...)
	return itd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (itd *IssueTagDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(itd.hooks) == 0 {
		affected, err = itd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IssueTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			itd.mutation = mutation
			affected, err = itd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(itd.hooks) - 1; i >= 0; i-- {
			if itd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = itd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, itd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (itd *IssueTagDelete) ExecX(ctx context.Context) int {
	n, err := itd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (itd *IssueTagDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: issuetag.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: issuetag.FieldID,
			},
		},
	}
	if ps := itd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, itd.driver, _spec)
}

// IssueTagDeleteOne is the builder for deleting a single IssueTag entity.
type IssueTagDeleteOne struct {
	itd *IssueTagDelete
}

// Exec executes the deletion query.
func (itdo *IssueTagDeleteOne) Exec(ctx context.Context) error {
	n, err := itdo.itd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{issuetag.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (itdo *IssueTagDeleteOne) ExecX(ctx context.Context) {
	itdo.itd.ExecX(ctx)
}
