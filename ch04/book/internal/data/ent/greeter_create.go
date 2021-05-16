// Code generated by entc, DO NOT EDIT.

package ent

import (
	"book/internal/data/ent/greeter"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GreeterCreate is the builder for creating a Greeter entity.
type GreeterCreate struct {
	config
	mutation *GreeterMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (gc *GreeterCreate) SetName(s string) *GreeterCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (gc *GreeterCreate) SetNillableName(s *string) *GreeterCreate {
	if s != nil {
		gc.SetName(*s)
	}
	return gc
}

// SetDescription sets the "description" field.
func (gc *GreeterCreate) SetDescription(s string) *GreeterCreate {
	gc.mutation.SetDescription(s)
	return gc
}

// SetCreatedAt sets the "created_at" field.
func (gc *GreeterCreate) SetCreatedAt(t time.Time) *GreeterCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GreeterCreate) SetNillableCreatedAt(t *time.Time) *GreeterCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// SetUpdatedAt sets the "updated_at" field.
func (gc *GreeterCreate) SetUpdatedAt(t time.Time) *GreeterCreate {
	gc.mutation.SetUpdatedAt(t)
	return gc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gc *GreeterCreate) SetNillableUpdatedAt(t *time.Time) *GreeterCreate {
	if t != nil {
		gc.SetUpdatedAt(*t)
	}
	return gc
}

// SetID sets the "id" field.
func (gc *GreeterCreate) SetID(i int64) *GreeterCreate {
	gc.mutation.SetID(i)
	return gc
}

// Mutation returns the GreeterMutation object of the builder.
func (gc *GreeterCreate) Mutation() *GreeterMutation {
	return gc.mutation
}

// Save creates the Greeter in the database.
func (gc *GreeterCreate) Save(ctx context.Context) (*Greeter, error) {
	var (
		err  error
		node *Greeter
	)
	gc.defaults()
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GreeterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			node, err = gc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GreeterCreate) SaveX(ctx context.Context) *Greeter {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (gc *GreeterCreate) defaults() {
	if _, ok := gc.mutation.Name(); !ok {
		v := greeter.DefaultName
		gc.mutation.SetName(v)
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		v := greeter.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		v := greeter.DefaultUpdatedAt()
		gc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GreeterCreate) check() error {
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := gc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New("ent: missing required field \"description\"")}
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	return nil
}

func (gc *GreeterCreate) sqlSave(ctx context.Context) (*Greeter, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (gc *GreeterCreate) createSpec() (*Greeter, *sqlgraph.CreateSpec) {
	var (
		_node = &Greeter{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: greeter.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: greeter.FieldID,
			},
		}
	)
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: greeter.FieldName,
		})
		_node.Name = value
	}
	if value, ok := gc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: greeter.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: greeter.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: greeter.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// GreeterCreateBulk is the builder for creating many Greeter entities in bulk.
type GreeterCreateBulk struct {
	config
	builders []*GreeterCreate
}

// Save creates the Greeter entities in the database.
func (gcb *GreeterCreateBulk) Save(ctx context.Context) ([]*Greeter, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Greeter, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GreeterMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GreeterCreateBulk) SaveX(ctx context.Context) []*Greeter {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
