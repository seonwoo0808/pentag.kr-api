// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"pentag.kr/api-server/ent/contact"
	"pentag.kr/api-server/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeContact = "Contact"
)

// ContactMutation represents an operation that mutates the Contact nodes in the graph.
type ContactMutation struct {
	config
	op            Op
	typ           string
	id            *int
	first_name    *string
	last_name     *string
	email         *string
	phone         *string
	category      *int
	addcategory   *int
	message       *string
	verify_code   *string
	is_verified   *bool
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Contact, error)
	predicates    []predicate.Contact
}

var _ ent.Mutation = (*ContactMutation)(nil)

// contactOption allows management of the mutation configuration using functional options.
type contactOption func(*ContactMutation)

// newContactMutation creates new mutation for the Contact entity.
func newContactMutation(c config, op Op, opts ...contactOption) *ContactMutation {
	m := &ContactMutation{
		config:        c,
		op:            op,
		typ:           TypeContact,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withContactID sets the ID field of the mutation.
func withContactID(id int) contactOption {
	return func(m *ContactMutation) {
		var (
			err   error
			once  sync.Once
			value *Contact
		)
		m.oldValue = func(ctx context.Context) (*Contact, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Contact.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withContact sets the old Contact of the mutation.
func withContact(node *Contact) contactOption {
	return func(m *ContactMutation) {
		m.oldValue = func(context.Context) (*Contact, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ContactMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ContactMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ContactMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ContactMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Contact.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetFirstName sets the "first_name" field.
func (m *ContactMutation) SetFirstName(s string) {
	m.first_name = &s
}

// FirstName returns the value of the "first_name" field in the mutation.
func (m *ContactMutation) FirstName() (r string, exists bool) {
	v := m.first_name
	if v == nil {
		return
	}
	return *v, true
}

// OldFirstName returns the old "first_name" field's value of the Contact entity.
// If the Contact object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ContactMutation) OldFirstName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFirstName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFirstName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFirstName: %w", err)
	}
	return oldValue.FirstName, nil
}

// ResetFirstName resets all changes to the "first_name" field.
func (m *ContactMutation) ResetFirstName() {
	m.first_name = nil
}

// SetLastName sets the "last_name" field.
func (m *ContactMutation) SetLastName(s string) {
	m.last_name = &s
}

// LastName returns the value of the "last_name" field in the mutation.
func (m *ContactMutation) LastName() (r string, exists bool) {
	v := m.last_name
	if v == nil {
		return
	}
	return *v, true
}

// OldLastName returns the old "last_name" field's value of the Contact entity.
// If the Contact object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ContactMutation) OldLastName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastName: %w", err)
	}
	return oldValue.LastName, nil
}

// ResetLastName resets all changes to the "last_name" field.
func (m *ContactMutation) ResetLastName() {
	m.last_name = nil
}

// SetEmail sets the "email" field.
func (m *ContactMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *ContactMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the Contact entity.
// If the Contact object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ContactMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *ContactMutation) ResetEmail() {
	m.email = nil
}

// SetPhone sets the "phone" field.
func (m *ContactMutation) SetPhone(s string) {
	m.phone = &s
}

// Phone returns the value of the "phone" field in the mutation.
func (m *ContactMutation) Phone() (r string, exists bool) {
	v := m.phone
	if v == nil {
		return
	}
	return *v, true
}

// OldPhone returns the old "phone" field's value of the Contact entity.
// If the Contact object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ContactMutation) OldPhone(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPhone is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPhone requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPhone: %w", err)
	}
	return oldValue.Phone, nil
}

// ResetPhone resets all changes to the "phone" field.
func (m *ContactMutation) ResetPhone() {
	m.phone = nil
}

// SetCategory sets the "category" field.
func (m *ContactMutation) SetCategory(i int) {
	m.category = &i
	m.addcategory = nil
}

// Category returns the value of the "category" field in the mutation.
func (m *ContactMutation) Category() (r int, exists bool) {
	v := m.category
	if v == nil {
		return
	}
	return *v, true
}

// OldCategory returns the old "category" field's value of the Contact entity.
// If the Contact object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ContactMutation) OldCategory(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCategory is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCategory requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCategory: %w", err)
	}
	return oldValue.Category, nil
}

// AddCategory adds i to the "category" field.
func (m *ContactMutation) AddCategory(i int) {
	if m.addcategory != nil {
		*m.addcategory += i
	} else {
		m.addcategory = &i
	}
}

// AddedCategory returns the value that was added to the "category" field in this mutation.
func (m *ContactMutation) AddedCategory() (r int, exists bool) {
	v := m.addcategory
	if v == nil {
		return
	}
	return *v, true
}

// ResetCategory resets all changes to the "category" field.
func (m *ContactMutation) ResetCategory() {
	m.category = nil
	m.addcategory = nil
}

// SetMessage sets the "message" field.
func (m *ContactMutation) SetMessage(s string) {
	m.message = &s
}

// Message returns the value of the "message" field in the mutation.
func (m *ContactMutation) Message() (r string, exists bool) {
	v := m.message
	if v == nil {
		return
	}
	return *v, true
}

// OldMessage returns the old "message" field's value of the Contact entity.
// If the Contact object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ContactMutation) OldMessage(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMessage is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMessage requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMessage: %w", err)
	}
	return oldValue.Message, nil
}

// ResetMessage resets all changes to the "message" field.
func (m *ContactMutation) ResetMessage() {
	m.message = nil
}

// SetVerifyCode sets the "verify_code" field.
func (m *ContactMutation) SetVerifyCode(s string) {
	m.verify_code = &s
}

// VerifyCode returns the value of the "verify_code" field in the mutation.
func (m *ContactMutation) VerifyCode() (r string, exists bool) {
	v := m.verify_code
	if v == nil {
		return
	}
	return *v, true
}

// OldVerifyCode returns the old "verify_code" field's value of the Contact entity.
// If the Contact object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ContactMutation) OldVerifyCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldVerifyCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldVerifyCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldVerifyCode: %w", err)
	}
	return oldValue.VerifyCode, nil
}

// ResetVerifyCode resets all changes to the "verify_code" field.
func (m *ContactMutation) ResetVerifyCode() {
	m.verify_code = nil
}

// SetIsVerified sets the "is_verified" field.
func (m *ContactMutation) SetIsVerified(b bool) {
	m.is_verified = &b
}

// IsVerified returns the value of the "is_verified" field in the mutation.
func (m *ContactMutation) IsVerified() (r bool, exists bool) {
	v := m.is_verified
	if v == nil {
		return
	}
	return *v, true
}

// OldIsVerified returns the old "is_verified" field's value of the Contact entity.
// If the Contact object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ContactMutation) OldIsVerified(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsVerified is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsVerified requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsVerified: %w", err)
	}
	return oldValue.IsVerified, nil
}

// ResetIsVerified resets all changes to the "is_verified" field.
func (m *ContactMutation) ResetIsVerified() {
	m.is_verified = nil
}

// Where appends a list predicates to the ContactMutation builder.
func (m *ContactMutation) Where(ps ...predicate.Contact) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ContactMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ContactMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Contact, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ContactMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ContactMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Contact).
func (m *ContactMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ContactMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.first_name != nil {
		fields = append(fields, contact.FieldFirstName)
	}
	if m.last_name != nil {
		fields = append(fields, contact.FieldLastName)
	}
	if m.email != nil {
		fields = append(fields, contact.FieldEmail)
	}
	if m.phone != nil {
		fields = append(fields, contact.FieldPhone)
	}
	if m.category != nil {
		fields = append(fields, contact.FieldCategory)
	}
	if m.message != nil {
		fields = append(fields, contact.FieldMessage)
	}
	if m.verify_code != nil {
		fields = append(fields, contact.FieldVerifyCode)
	}
	if m.is_verified != nil {
		fields = append(fields, contact.FieldIsVerified)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ContactMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case contact.FieldFirstName:
		return m.FirstName()
	case contact.FieldLastName:
		return m.LastName()
	case contact.FieldEmail:
		return m.Email()
	case contact.FieldPhone:
		return m.Phone()
	case contact.FieldCategory:
		return m.Category()
	case contact.FieldMessage:
		return m.Message()
	case contact.FieldVerifyCode:
		return m.VerifyCode()
	case contact.FieldIsVerified:
		return m.IsVerified()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ContactMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case contact.FieldFirstName:
		return m.OldFirstName(ctx)
	case contact.FieldLastName:
		return m.OldLastName(ctx)
	case contact.FieldEmail:
		return m.OldEmail(ctx)
	case contact.FieldPhone:
		return m.OldPhone(ctx)
	case contact.FieldCategory:
		return m.OldCategory(ctx)
	case contact.FieldMessage:
		return m.OldMessage(ctx)
	case contact.FieldVerifyCode:
		return m.OldVerifyCode(ctx)
	case contact.FieldIsVerified:
		return m.OldIsVerified(ctx)
	}
	return nil, fmt.Errorf("unknown Contact field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ContactMutation) SetField(name string, value ent.Value) error {
	switch name {
	case contact.FieldFirstName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFirstName(v)
		return nil
	case contact.FieldLastName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastName(v)
		return nil
	case contact.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case contact.FieldPhone:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPhone(v)
		return nil
	case contact.FieldCategory:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCategory(v)
		return nil
	case contact.FieldMessage:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMessage(v)
		return nil
	case contact.FieldVerifyCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetVerifyCode(v)
		return nil
	case contact.FieldIsVerified:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsVerified(v)
		return nil
	}
	return fmt.Errorf("unknown Contact field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ContactMutation) AddedFields() []string {
	var fields []string
	if m.addcategory != nil {
		fields = append(fields, contact.FieldCategory)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ContactMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case contact.FieldCategory:
		return m.AddedCategory()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ContactMutation) AddField(name string, value ent.Value) error {
	switch name {
	case contact.FieldCategory:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCategory(v)
		return nil
	}
	return fmt.Errorf("unknown Contact numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ContactMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ContactMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ContactMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Contact nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ContactMutation) ResetField(name string) error {
	switch name {
	case contact.FieldFirstName:
		m.ResetFirstName()
		return nil
	case contact.FieldLastName:
		m.ResetLastName()
		return nil
	case contact.FieldEmail:
		m.ResetEmail()
		return nil
	case contact.FieldPhone:
		m.ResetPhone()
		return nil
	case contact.FieldCategory:
		m.ResetCategory()
		return nil
	case contact.FieldMessage:
		m.ResetMessage()
		return nil
	case contact.FieldVerifyCode:
		m.ResetVerifyCode()
		return nil
	case contact.FieldIsVerified:
		m.ResetIsVerified()
		return nil
	}
	return fmt.Errorf("unknown Contact field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ContactMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ContactMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ContactMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ContactMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ContactMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ContactMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ContactMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Contact unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ContactMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Contact edge %s", name)
}
