package examples

// AuthRequest is setting all the needed
// fields to authenticate a user
//
// +jsonschema:meta:ID=test-id
// +jsonschema:meta:Draft=2-07
type AuthRequest struct {
	User      User `json:"user,omitempty"`
	UserPtr   *User
	Bar       int             `json:"bar,omitempty"`
	Foo       []string        `json:"foo,omitempty"`
	FooBarMap map[string]User `json:"fooBarMap"`
	FooBar    string

	// Size of the request
	// +jsonschema:validation:Maximum=1024
	Size Size `json:"size,omitempty"`
}

type User struct {
	Email    string
	Password string
}

type Size int

type Alias = string

type Map map[string]string
