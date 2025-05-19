// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	"encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Log string         // Always "log"
type State string       // Always "state"
type StateUpdate string // Always "state_update"

func (c Log) Default() Log                 { return "log" }
func (c State) Default() State             { return "state" }
func (c StateUpdate) Default() StateUpdate { return "state_update" }

func (c Log) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c State) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c StateUpdate) MarshalJSON() ([]byte, error) { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return json.Marshal(string(v))
}
