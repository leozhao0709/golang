// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/leozhao0709/golang/video-server/ent/comment"
	"github.com/leozhao0709/golang/video-server/ent/predicate"
	"github.com/leozhao0709/golang/video-server/ent/user"
	"github.com/leozhao0709/golang/video-server/ent/video"
)

// VideoUpdate is the builder for updating Video entities.
type VideoUpdate struct {
	config
	hooks      []Hook
	mutation   *VideoMutation
	predicates []predicate.Video
}

// Where adds a new predicate for the builder.
func (vu *VideoUpdate) Where(ps ...predicate.Video) *VideoUpdate {
	vu.predicates = append(vu.predicates, ps...)
	return vu
}

// SetName sets the name field.
func (vu *VideoUpdate) SetName(s string) *VideoUpdate {
	vu.mutation.SetName(s)
	return vu
}

// AddCommentIDs adds the comments edge to Comment by ids.
func (vu *VideoUpdate) AddCommentIDs(ids ...int) *VideoUpdate {
	vu.mutation.AddCommentIDs(ids...)
	return vu
}

// AddComments adds the comments edges to Comment.
func (vu *VideoUpdate) AddComments(c ...*Comment) *VideoUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return vu.AddCommentIDs(ids...)
}

// SetAuthorID sets the author edge to User by id.
func (vu *VideoUpdate) SetAuthorID(id int) *VideoUpdate {
	vu.mutation.SetAuthorID(id)
	return vu
}

// SetNillableAuthorID sets the author edge to User by id if the given value is not nil.
func (vu *VideoUpdate) SetNillableAuthorID(id *int) *VideoUpdate {
	if id != nil {
		vu = vu.SetAuthorID(*id)
	}
	return vu
}

// SetAuthor sets the author edge to User.
func (vu *VideoUpdate) SetAuthor(u *User) *VideoUpdate {
	return vu.SetAuthorID(u.ID)
}

// Mutation returns the VideoMutation object of the builder.
func (vu *VideoUpdate) Mutation() *VideoMutation {
	return vu.mutation
}

// ClearComments clears all "comments" edges to type Comment.
func (vu *VideoUpdate) ClearComments() *VideoUpdate {
	vu.mutation.ClearComments()
	return vu
}

// RemoveCommentIDs removes the comments edge to Comment by ids.
func (vu *VideoUpdate) RemoveCommentIDs(ids ...int) *VideoUpdate {
	vu.mutation.RemoveCommentIDs(ids...)
	return vu
}

// RemoveComments removes comments edges to Comment.
func (vu *VideoUpdate) RemoveComments(c ...*Comment) *VideoUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return vu.RemoveCommentIDs(ids...)
}

// ClearAuthor clears the "author" edge to type User.
func (vu *VideoUpdate) ClearAuthor() *VideoUpdate {
	vu.mutation.ClearAuthor()
	return vu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (vu *VideoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vu.hooks) == 0 {
		if err = vu.check(); err != nil {
			return 0, err
		}
		affected, err = vu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vu.check(); err != nil {
				return 0, err
			}
			vu.mutation = mutation
			affected, err = vu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vu.hooks) - 1; i >= 0; i-- {
			mut = vu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VideoUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VideoUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VideoUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vu *VideoUpdate) check() error {
	if v, ok := vu.mutation.Name(); ok {
		if err := video.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (vu *VideoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   video.Table,
			Columns: video.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: video.FieldID,
			},
		},
	}
	if ps := vu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: video.FieldName,
		})
	}
	if vu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.CommentsTable,
			Columns: []string{video.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !vu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.CommentsTable,
			Columns: []string{video.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.CommentsTable,
			Columns: []string{video.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vu.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   video.AuthorTable,
			Columns: []string{video.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   video.AuthorTable,
			Columns: []string{video.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{video.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// VideoUpdateOne is the builder for updating a single Video entity.
type VideoUpdateOne struct {
	config
	hooks    []Hook
	mutation *VideoMutation
}

// SetName sets the name field.
func (vuo *VideoUpdateOne) SetName(s string) *VideoUpdateOne {
	vuo.mutation.SetName(s)
	return vuo
}

// AddCommentIDs adds the comments edge to Comment by ids.
func (vuo *VideoUpdateOne) AddCommentIDs(ids ...int) *VideoUpdateOne {
	vuo.mutation.AddCommentIDs(ids...)
	return vuo
}

// AddComments adds the comments edges to Comment.
func (vuo *VideoUpdateOne) AddComments(c ...*Comment) *VideoUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return vuo.AddCommentIDs(ids...)
}

// SetAuthorID sets the author edge to User by id.
func (vuo *VideoUpdateOne) SetAuthorID(id int) *VideoUpdateOne {
	vuo.mutation.SetAuthorID(id)
	return vuo
}

// SetNillableAuthorID sets the author edge to User by id if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableAuthorID(id *int) *VideoUpdateOne {
	if id != nil {
		vuo = vuo.SetAuthorID(*id)
	}
	return vuo
}

// SetAuthor sets the author edge to User.
func (vuo *VideoUpdateOne) SetAuthor(u *User) *VideoUpdateOne {
	return vuo.SetAuthorID(u.ID)
}

// Mutation returns the VideoMutation object of the builder.
func (vuo *VideoUpdateOne) Mutation() *VideoMutation {
	return vuo.mutation
}

// ClearComments clears all "comments" edges to type Comment.
func (vuo *VideoUpdateOne) ClearComments() *VideoUpdateOne {
	vuo.mutation.ClearComments()
	return vuo
}

// RemoveCommentIDs removes the comments edge to Comment by ids.
func (vuo *VideoUpdateOne) RemoveCommentIDs(ids ...int) *VideoUpdateOne {
	vuo.mutation.RemoveCommentIDs(ids...)
	return vuo
}

// RemoveComments removes comments edges to Comment.
func (vuo *VideoUpdateOne) RemoveComments(c ...*Comment) *VideoUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return vuo.RemoveCommentIDs(ids...)
}

// ClearAuthor clears the "author" edge to type User.
func (vuo *VideoUpdateOne) ClearAuthor() *VideoUpdateOne {
	vuo.mutation.ClearAuthor()
	return vuo
}

// Save executes the query and returns the updated entity.
func (vuo *VideoUpdateOne) Save(ctx context.Context) (*Video, error) {
	var (
		err  error
		node *Video
	)
	if len(vuo.hooks) == 0 {
		if err = vuo.check(); err != nil {
			return nil, err
		}
		node, err = vuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vuo.check(); err != nil {
				return nil, err
			}
			vuo.mutation = mutation
			node, err = vuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vuo.hooks) - 1; i >= 0; i-- {
			mut = vuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VideoUpdateOne) SaveX(ctx context.Context) *Video {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VideoUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VideoUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuo *VideoUpdateOne) check() error {
	if v, ok := vuo.mutation.Name(); ok {
		if err := video.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (vuo *VideoUpdateOne) sqlSave(ctx context.Context) (_node *Video, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   video.Table,
			Columns: video.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: video.FieldID,
			},
		},
	}
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Video.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := vuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: video.FieldName,
		})
	}
	if vuo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.CommentsTable,
			Columns: []string{video.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !vuo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.CommentsTable,
			Columns: []string{video.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.CommentsTable,
			Columns: []string{video.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vuo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   video.AuthorTable,
			Columns: []string{video.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   video.AuthorTable,
			Columns: []string{video.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Video{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{video.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}