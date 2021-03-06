package object

import "github.com/kasworld/nonkey/enum/objecttype"

// ReturnValue wraps ObjectI and implements ObjectI interface.
type ReturnValue struct {
	// Value is the object that is to be returned
	Value ObjectI
}

// Type returns the type of this object.
func (rv *ReturnValue) Type() objecttype.ObjectType {
	return objecttype.RETURN_VALUE
}

// Inspect returns a string-representation of the given object.
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (rv *ReturnValue) InvokeMethod(method string, env Environment, args ...ObjectI) ObjectI {

	//
	// There are no methods available upon a return-object.
	//
	// (The return-object is an implementation-detail.)
	//
	return nil
}

// ToInterface converts this object to a go-interface, which will allow
// it to be used naturally in our sprintf/printf primitives.
//
// It might also be helpful for embedded users.
func (rv *ReturnValue) ToInterface() interface{} {
	return "<RETURN_VALUE>"
}
