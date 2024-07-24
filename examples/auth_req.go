package examples

// AuthRequest is setting all the needed
// fields to authenticate a user
//
// +jsonschema:meta:ID=test-id
// +jsonschema:meta:Draft=2-07
type AuthRequest[T any] struct {
	User      User            `json:"user,omitempty"`
	User2     *User           `json:"user2,omitempty"`
	Bar       string          `json:"bar,omitempty"`
	Foo       []string        `json:"foo,omitempty"`
	FooBarMap map[string]User `json:"fooBarMap"`
	// Size of the request
	// +jsonschema:validation:Maximum=1024
	Size  int  `json:"size,omitempty"`
	Size2 *int `json:"size2,omitempty"`
	F     T    `json:"f,omitempty"`
}

type User struct {
	Email    string
	Password string
}
