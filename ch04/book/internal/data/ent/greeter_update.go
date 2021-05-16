// Code generated by entc, DO NOT EDIT.

package ent

import (
	"book/internal/data/ent/greeter"
	"book/internal/data/ent/predicate"
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GreeterUpdate is the builder for updating Greeter entities.
type GreeterUpdate struct {
	config
	hooks    []Hook
	mutation *GreeterMutation
}

// Where adds a new predicate for the GreeterUpdate builder.
func (gu *GreeterUpdate) Where(ps ...predicate.Greeter) *GreeterUpdate {
	gu.mutation.predicates = append(gu.mutation.predicates, ps...)
	return gu
}

// SetName sets the "name" field.
func (gu *GreeterUpdate) SetName(s string) *GreeterUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (gu *GreeterUpdate) SetNillableName(s *string) *GreeterUpdate {
	if s != nil {
		gu.SetName(*s)
	}
	return gu
}

// SetDescription sets the "description" field.
func (gu *GreeterUpdate) SetDescription(s string) *GreeterUpdate {
	gu.mutation.SetDescription(s)
	return gu
}

// SetCreatedAt sets the "created_at" field.
func (gu *GreeterUpdate) SetCreatedAt(t time.Time) *GreeterUpdate {
	gu.mutation.SetCreatedAt(t)
	return gu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gu *GreeterUpdate) SetNillableCreatedAt(t *time.Time) *GreeterUpdate {
	if t != nil {
		gu.SetCreatedAt(*t)
	}
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GreeterUpdate) SetUpdatedAt(t time.Time) *GreeterUpdate {
	gu.mutation.SetUpdatedAt(t)
	return gu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gu *GreeterUpdate) SetNillableUpdatedAt(t *time.Time) *GreeterUpdate {
	if t != nil {
		gu.SetUpdatedAt(*t)
	}
	return gu
}

// Mutation returns the GreeterMutation object of the builder.
func (gu *GreeterUpdate) Mutation() *GreeterMutation {
	return gu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GreeterUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gu.hooks) == 0 {
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GreeterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GreeterUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GreeterUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GreeterUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GreeterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   greeter.Table,
			Columns: greeter.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: greeter.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: greeter.FieldName,
		})
	}
	if value, ok := gu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: greeter.FieldDescription,
		})
	}
	if value, ok := gu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: greeter.FieldCreatedAt,
		})
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: greeter.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{greeter.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// GreeterUpdateOne is the builder for updating a single Greeter entity.
type GreeterUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GreeterMutation
}

// SetName sets the "name" field.
func (guo *GreeterUpdateOne) SetName(s string) *GreeterUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (guo *GreeterUpdateOne) SetNillableName(s *string) *GreeterUpdateOne {
	if s != nil {
		guo.SetName(*s)
	}
	return guo
}

// SetDescription sets the "description" field.
func (guo *GreeterUpdateOne) SetDescription(s string) *GreeterUpdateOne {
	guo.mutation.SetDescription(s)
	return guo
}

// SetCreatedAt sets the "created_at" field.
func (guo *GreeterUpdateOne) SetCreatedAt(t time.Time) *GreeterUpdateOne {
	guo.mutation.SetCreatedAt(t)
	return guo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (guo *GreeterUpdateOne) SetNillableCreatedAt(t *time.Time) *GreeterUpdateOne {
	if t != nil {
		guo.SetCreatedAt(*t)
	}
	return guo
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GreeterUpdateOne) SetUpdatedAt(t time.Time) *GreeterUpdateOne {
	guo.mutation.SetUpdatedAt(t)
	return guo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (guo *GreeterUpdateOne) SetNillableUpdatedAt(t *time.Time) *GreeterUpdateOne {
	if t != nil {
		guo.SetUpdatedAt(*t)
	}
	return guo
}

// Mutation returns the GreeterMutation object of the builder.
func (guo *GreeterUpdateOne) Mutation() *GreeterMutation {
	return guo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GreeterUpdateOne) Select(field string, fields ...string) *GreeterUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Greeter entity.
func (guo *GreeterUpdateOne) Save(ctx context.Context) (*Greeter, error) {
	var (
		err  error
		node *Greeter
	)
	if len(guo.hooks) == 0 {
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GreeterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			mut = guo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GreeterUpdateOne) SaveX(ctx context.Context) *Greeter {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GreeterUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GreeterUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GreeterUpdateOne) sqlSave(ctx context.Context) (_node *Greeter, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   greeter.Table,
			Columns: greeter.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: greeter.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Greeter.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, greeter.FieldID)
		for _, f := range fields {
			if !greeter.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != greeter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: greeter.FieldName,
		})
	}
	if value, ok := guo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: greeter.FieldDescription,
		})
	}
	if value, ok := guo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: greeter.FieldCreatedAt,
		})
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: greeter.FieldUpdatedAt,
		})
	}
	_node = &Greeter{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{greeter.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
