// Code generated by ent, DO NOT EDIT.

package contact

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the contact type in the database.
	Label = "contact"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldVerifyCode holds the string denoting the verify_code field in the database.
	FieldVerifyCode = "verify_code"
	// FieldIsVerified holds the string denoting the is_verified field in the database.
	FieldIsVerified = "is_verified"
	// Table holds the table name of the contact in the database.
	Table = "contacts"
)

// Columns holds all SQL columns for contact fields.
var Columns = []string{
	FieldID,
	FieldFirstName,
	FieldLastName,
	FieldEmail,
	FieldPhone,
	FieldCategory,
	FieldMessage,
	FieldVerifyCode,
	FieldIsVerified,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	FirstNameValidator func(string) error
	// LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// CategoryValidator is a validator for the "category" field. It is called by the builders before save.
	CategoryValidator func(int) error
	// MessageValidator is a validator for the "message" field. It is called by the builders before save.
	MessageValidator func(string) error
	// VerifyCodeValidator is a validator for the "verify_code" field. It is called by the builders before save.
	VerifyCodeValidator func(string) error
	// DefaultIsVerified holds the default value on creation for the "is_verified" field.
	DefaultIsVerified bool
)

// OrderOption defines the ordering options for the Contact queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFirstName orders the results by the first_name field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
}

// ByMessage orders the results by the message field.
func ByMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMessage, opts...).ToFunc()
}

// ByVerifyCode orders the results by the verify_code field.
func ByVerifyCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVerifyCode, opts...).ToFunc()
}

// ByIsVerified orders the results by the is_verified field.
func ByIsVerified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsVerified, opts...).ToFunc()
}
