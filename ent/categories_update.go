// Code generated by ent, DO NOT EDIT.

package ent

import (
	"composeapp/ent/categories"
	"composeapp/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoriesUpdate is the builder for updating Categories entities.
type CategoriesUpdate struct {
	config
	hooks    []Hook
	mutation *CategoriesMutation
}

// Where appends a list predicates to the CategoriesUpdate builder.
func (cu *CategoriesUpdate) Where(ps ...predicate.Categories) *CategoriesUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetParentID sets the "parent_id" field.
func (cu *CategoriesUpdate) SetParentID(i int) *CategoriesUpdate {
	cu.mutation.ResetParentID()
	cu.mutation.SetParentID(i)
	return cu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cu *CategoriesUpdate) SetNillableParentID(i *int) *CategoriesUpdate {
	if i != nil {
		cu.SetParentID(*i)
	}
	return cu
}

// AddParentID adds i to the "parent_id" field.
func (cu *CategoriesUpdate) AddParentID(i int) *CategoriesUpdate {
	cu.mutation.AddParentID(i)
	return cu
}

// SetName sets the "name" field.
func (cu *CategoriesUpdate) SetName(s string) *CategoriesUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CategoriesUpdate) SetNillableName(s *string) *CategoriesUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetSystemName sets the "system_name" field.
func (cu *CategoriesUpdate) SetSystemName(s string) *CategoriesUpdate {
	cu.mutation.SetSystemName(s)
	return cu
}

// SetNillableSystemName sets the "system_name" field if the given value is not nil.
func (cu *CategoriesUpdate) SetNillableSystemName(s *string) *CategoriesUpdate {
	if s != nil {
		cu.SetSystemName(*s)
	}
	return cu
}

// SetDescription sets the "description" field.
func (cu *CategoriesUpdate) SetDescription(s string) *CategoriesUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *CategoriesUpdate) SetNillableDescription(s *string) *CategoriesUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// Mutation returns the CategoriesMutation object of the builder.
func (cu *CategoriesUpdate) Mutation() *CategoriesMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CategoriesUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CategoriesUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CategoriesUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CategoriesUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CategoriesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(categories.Table, categories.Columns, sqlgraph.NewFieldSpec(categories.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.ParentID(); ok {
		_spec.SetField(categories.FieldParentID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedParentID(); ok {
		_spec.AddField(categories.FieldParentID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(categories.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.SystemName(); ok {
		_spec.SetField(categories.FieldSystemName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(categories.FieldDescription, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{categories.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CategoriesUpdateOne is the builder for updating a single Categories entity.
type CategoriesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CategoriesMutation
}

// SetParentID sets the "parent_id" field.
func (cuo *CategoriesUpdateOne) SetParentID(i int) *CategoriesUpdateOne {
	cuo.mutation.ResetParentID()
	cuo.mutation.SetParentID(i)
	return cuo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cuo *CategoriesUpdateOne) SetNillableParentID(i *int) *CategoriesUpdateOne {
	if i != nil {
		cuo.SetParentID(*i)
	}
	return cuo
}

// AddParentID adds i to the "parent_id" field.
func (cuo *CategoriesUpdateOne) AddParentID(i int) *CategoriesUpdateOne {
	cuo.mutation.AddParentID(i)
	return cuo
}

// SetName sets the "name" field.
func (cuo *CategoriesUpdateOne) SetName(s string) *CategoriesUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CategoriesUpdateOne) SetNillableName(s *string) *CategoriesUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetSystemName sets the "system_name" field.
func (cuo *CategoriesUpdateOne) SetSystemName(s string) *CategoriesUpdateOne {
	cuo.mutation.SetSystemName(s)
	return cuo
}

// SetNillableSystemName sets the "system_name" field if the given value is not nil.
func (cuo *CategoriesUpdateOne) SetNillableSystemName(s *string) *CategoriesUpdateOne {
	if s != nil {
		cuo.SetSystemName(*s)
	}
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CategoriesUpdateOne) SetDescription(s string) *CategoriesUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *CategoriesUpdateOne) SetNillableDescription(s *string) *CategoriesUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// Mutation returns the CategoriesMutation object of the builder.
func (cuo *CategoriesUpdateOne) Mutation() *CategoriesMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CategoriesUpdate builder.
func (cuo *CategoriesUpdateOne) Where(ps ...predicate.Categories) *CategoriesUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CategoriesUpdateOne) Select(field string, fields ...string) *CategoriesUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Categories entity.
func (cuo *CategoriesUpdateOne) Save(ctx context.Context) (*Categories, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CategoriesUpdateOne) SaveX(ctx context.Context) *Categories {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CategoriesUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CategoriesUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CategoriesUpdateOne) sqlSave(ctx context.Context) (_node *Categories, err error) {
	_spec := sqlgraph.NewUpdateSpec(categories.Table, categories.Columns, sqlgraph.NewFieldSpec(categories.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Categories.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, categories.FieldID)
		for _, f := range fields {
			if !categories.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != categories.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.ParentID(); ok {
		_spec.SetField(categories.FieldParentID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedParentID(); ok {
		_spec.AddField(categories.FieldParentID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(categories.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.SystemName(); ok {
		_spec.SetField(categories.FieldSystemName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(categories.FieldDescription, field.TypeString, value)
	}
	_node = &Categories{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{categories.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}