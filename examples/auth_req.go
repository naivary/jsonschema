package examples


// AuthRequest is setting all the needed
// fields to authenticate a user
//
// +jsonschema:meta:ID=test-id
// +jsonschema:meta:Draft=2-07
type AuthRequest struct {
    // Email of the user trying
    // to authenticate
    //
    // +jsonschema:validation:Maximum=3
    // +openapi:validation:M=3
    Email string `json:"email"`

    // Password is the raw password of
    // the user trying to authenticate
    //
    // +jsonschema:validation:Minimum=2
    Password string `json:"password"`
}
