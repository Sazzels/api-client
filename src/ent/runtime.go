// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api-client/src/ent/request"
	"api-client/src/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	requestFields := schema.Request{}.Fields()
	_ = requestFields
	// requestDescName is the schema descriptor for name field.
	requestDescName := requestFields[0].Descriptor()
	// request.DefaultName holds the default value on creation for the name field.
	request.DefaultName = requestDescName.Default.(string)
}
