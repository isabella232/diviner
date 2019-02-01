// Copyright 2019 GRAIL, Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package diviner

import (
	"fmt"
	"sort"
	"strings"
)

// Kind represents the kind of a value.
type Kind int

const (
	Integer Kind = iota
	Real
	Str
)

func (k Kind) String() string {
	switch k {
	case Integer:
		return "integer"
	case Real:
		return "real"
	case Str:
		return "string"
	default:
		panic(k)
	}
}

// Value is the type of parameter values. Values must be
// directly comparable.
type Value interface {
	// String returns a textual description of the parameter value.
	String() string

	// Kind returns the kind of this value.
	Kind() Kind

	// Less returns true if the value is less than the provided value.
	// Less is defined only for values of the same type.
	Less(Value) bool

	// Float returns the floating point value of float-typed values.
	Float() float64

	// Int returns the integer value of integer-typed values.
	Int() int64

	// Str returns the string of string-typed values.
	Str() string
}

// Int is an integer-typed value.
type Int int64

// String implements Value.
func (v Int) String() string { return fmt.Sprint(int64(v)) }

// Kind implements Value.
func (Int) Kind() Kind { return Integer }

// Less implements Value.
func (v Int) Less(w Value) bool {
	return int64(v) < int64(w.(Int))
}

// Float implements Value.
func (Int) Float() float64 { panic("Float on Int") }

// Str implements Value.
func (Int) Str() string { panic("Str on Int") }

// Int implements Value.
func (v Int) Int() int64 { return int64(v) }

// Float is a float-typed value.
type Float float64

// String implements Value.
func (v Float) String() string { return fmt.Sprint(float64(v)) }

// Kind implements Value.
func (Float) Kind() Kind { return Real }

// Less implements Value.
func (v Float) Less(w Value) bool {
	return float64(v) < float64(w.(Float))
}

// Float implements Value.
func (v Float) Float() float64 { return float64(v) }

// Str implements Value.
func (Float) Str() string { panic("Str on Float") }

// Int implements Value.
func (Float) Int() int64 { panic("Int on Float") }

// String is a string-typed value.
type String string

// Less implements Value.
func (v String) Less(w Value) bool {
	return string(v) < string(w.(String))
}

// String implements Value.
func (v String) String() string { return string(v) }

// Kind implements Value.
func (String) Kind() Kind { return Str }

// Float implements Value.
func (String) Float() float64 { panic("Float on String") }

// Int implements Value.
func (String) Int() int64 { panic("Int on String") }

// Str implements Value.
func (v String) Str() string { return string(v) }

// Values is a set of named value, used as a concrete instantiation
// of a set of parameters.
type Values map[string]Value

// String returns a (stable) textual description of the value set.
func (v Values) String() string {
	keys := make([]string, 0, len(v))
	for key := range v {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	elems := make([]string, len(keys))
	for i, key := range keys {
		elems[i] = fmt.Sprintf("%s=%s", key, v[key])
	}
	return strings.Join(elems, ",")
}