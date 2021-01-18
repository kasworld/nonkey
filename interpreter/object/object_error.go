package object

import (
	"fmt"

	"github.com/kasworld/nonkey/enum/objecttype"
)

func NewError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

func IsError(obj ObjectI) bool {
	if obj != nil {
		return obj.Type() == objecttype.ERROR
	}
	return false
}

// Error wraps string and implements ObjectI interface.
type Error struct {
	// Message contains the error-message we're wrapping
	Message string
}

// Type returns the type of this object.
func (e *Error) Type() objecttype.ObjectType {
	return objecttype.ERROR
}

// Inspect returns a string-representation of the given object.
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (e *Error) InvokeMethod(method string, env Environment, args ...ObjectI) ObjectI {

	//
	// There are no methods available upon a return-object.
	//
	// (The error-object is an implementation-detail.)
	//
	return nil
}

// ToInterface converts this object to a go-interface, which will allow
// it to be used naturally in our sprintf/printf primitives.
//
// It might also be helpful for embedded users.
func (e *Error) ToInterface() interface{} {
	return "<ERROR>"
}
