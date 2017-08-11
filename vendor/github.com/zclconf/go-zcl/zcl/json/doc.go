// Package json is the JSON parser for ZCL. It parses JSON files and returns
// implementations of the core ZCL structural interfaces in terms of the
// JSON data inside.
//
// This is not a generic JSON parser. Instead, it deals with the mapping from
// the JSON information model to the ZCL information model, using a number
// of hard-coded structural conventions.
package json
