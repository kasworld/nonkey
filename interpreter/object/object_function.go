package object

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/kasworld/nonkey/enum/objecttype"
	"github.com/kasworld/nonkey/interpreter/ast"
	"github.com/kasworld/nonkey/interpreter/asti"
)

// Function wraps ast.Identifier array, ast.BlockStatement and Environment and implements ObjectI interface.
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Defaults   map[string]asti.ExpressionI
	Env        *Environment
}

// Type returns the type of this object.
func (f *Function) Type() objecttype.ObjectType {
	return objecttype.FUNCTION
}

// Inspect returns a string-representation of the given object.
func (f *Function) Inspect() string {
	var out bytes.Buffer
	parameters := make([]string, 0)
	for _, p := range f.Parameters {
		parameters = append(parameters, p.String())
	}
	fmt.Fprintf(&out, "fn(%v) {\n%v\n}", strings.Join(parameters, ", "), f.Body)
	return out.String()
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (f *Function) InvokeMethod(method string, env Environment, args ...ObjectI) ObjectI {
	if method == "methods" {
		static := []string{"methods"}
		dynamic := env.Names("function.")

		var names []string
		names = append(names, static...)
		for _, e := range dynamic {
			bits := strings.Split(e, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)

		result := make([]ObjectI, len(names))
		for i, txt := range names {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	return nil
}

// ToInterface converts this object to a go-interface, which will allow
// it to be used naturally in our sprintf/printf primitives.
//
// It might also be helpful for embedded users.
func (f *Function) ToInterface() interface{} {
	return "<FUNCTION>"
}
