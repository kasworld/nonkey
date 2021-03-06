package object

import "github.com/kasworld/nonkey/enum/objecttype"

// Null wraps nothing and implements our ObjectI interface.
type Null struct{}

// Type returns the type of this object.
func (n *Null) Type() objecttype.ObjectType {
	return objecttype.NULL
}

// Inspect returns a string-representation of the given object.
func (n *Null) Inspect() string {
	return "null"
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (n *Null) InvokeMethod(method string, env Environment, args ...ObjectI) ObjectI {
	return nil
}

// ToInterface converts this object to a go-interface, which will allow
// it to be used naturally in our sprintf/printf primitives.
//
// It might also be helpful for embedded users.
func (n *Null) ToInterface() interface{} {
	return "<NULL>"
}
