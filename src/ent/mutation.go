// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api-client/src/ent/predicate"
	"api-client/src/ent/request"
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeRequest = "Request"
)

// RequestMutation represents an operation that mutates the Request nodes in the graph.
type RequestMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Request, error)
	predicates    []predicate.Request
}

var _ ent.Mutation = (*RequestMutation)(nil)

// requestOption allows management of the mutation configuration using functional options.
type requestOption func(*RequestMutation)

// newRequestMutation creates new mutation for the Request entity.
func newRequestMutation(c config, op Op, opts ...requestOption) *RequestMutation {
	m := &RequestMutation{
		config:        c,
		op:            op,
		typ:           TypeRequest,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRequestID sets the ID field of the mutation.
func withRequestID(id int) requestOption {
	return func(m *RequestMutation) {
		var (
			err   error
			once  sync.Once
			value *Request
		)
		m.oldValue = func(ctx context.Context) (*Request, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Request.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRequest sets the old Request of the mutation.
func withRequest(node *Request) requestOption {
	return func(m *RequestMutation) {
		m.oldValue = func(context.Context) (*Request, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RequestMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RequestMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *RequestMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *RequestMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Request.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *RequestMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *RequestMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Request entity.
// If the Request object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RequestMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *RequestMutation) ResetName() {
	m.name = nil
}

// Where appends a list predicates to the RequestMutation builder.
func (m *RequestMutation) Where(ps ...predicate.Request) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the RequestMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *RequestMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Request, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *RequestMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *RequestMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Request).
func (m *RequestMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *RequestMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, request.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *RequestMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case request.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *RequestMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case request.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Request field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RequestMutation) SetField(name string, value ent.Value) error {
	switch name {
	case request.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Request field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *RequestMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *RequestMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RequestMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Request numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *RequestMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *RequestMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *RequestMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Request nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *RequestMutation) ResetField(name string) error {
	switch name {
	case request.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Request field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *RequestMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RequestMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RequestMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RequestMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RequestMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RequestMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RequestMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Request unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RequestMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Request edge %s", name)
}
