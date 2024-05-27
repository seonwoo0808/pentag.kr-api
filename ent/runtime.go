// Code generated by ent, DO NOT EDIT.

package ent

import (
	"pentag.kr/api-server/ent/contact"
	"pentag.kr/api-server/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	contactFields := schema.Contact{}.Fields()
	_ = contactFields
	// contactDescFirstName is the schema descriptor for first_name field.
	contactDescFirstName := contactFields[0].Descriptor()
	// contact.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	contact.FirstNameValidator = func() func(string) error {
		validators := contactDescFirstName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(first_name string) error {
			for _, fn := range fns {
				if err := fn(first_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// contactDescLastName is the schema descriptor for last_name field.
	contactDescLastName := contactFields[1].Descriptor()
	// contact.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	contact.LastNameValidator = func() func(string) error {
		validators := contactDescLastName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(last_name string) error {
			for _, fn := range fns {
				if err := fn(last_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// contactDescEmail is the schema descriptor for email field.
	contactDescEmail := contactFields[2].Descriptor()
	// contact.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	contact.EmailValidator = func() func(string) error {
		validators := contactDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// contactDescPhone is the schema descriptor for phone field.
	contactDescPhone := contactFields[3].Descriptor()
	// contact.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	contact.PhoneValidator = func() func(string) error {
		validators := contactDescPhone.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(phone string) error {
			for _, fn := range fns {
				if err := fn(phone); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// contactDescCategory is the schema descriptor for category field.
	contactDescCategory := contactFields[4].Descriptor()
	// contact.CategoryValidator is a validator for the "category" field. It is called by the builders before save.
	contact.CategoryValidator = contactDescCategory.Validators[0].(func(int) error)
	// contactDescMessage is the schema descriptor for message field.
	contactDescMessage := contactFields[5].Descriptor()
	// contact.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	contact.MessageValidator = func() func(string) error {
		validators := contactDescMessage.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(message string) error {
			for _, fn := range fns {
				if err := fn(message); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// contactDescVerifyCode is the schema descriptor for verify_code field.
	contactDescVerifyCode := contactFields[6].Descriptor()
	// contact.VerifyCodeValidator is a validator for the "verify_code" field. It is called by the builders before save.
	contact.VerifyCodeValidator = func() func(string) error {
		validators := contactDescVerifyCode.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(verify_code string) error {
			for _, fn := range fns {
				if err := fn(verify_code); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// contactDescIsVerified is the schema descriptor for is_verified field.
	contactDescIsVerified := contactFields[7].Descriptor()
	// contact.DefaultIsVerified holds the default value on creation for the is_verified field.
	contact.DefaultIsVerified = contactDescIsVerified.Default.(bool)
}