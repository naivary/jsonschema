package examples

type AuthRequest struct {
    // Email of the user trying
    // to authenticate
    //
    // +jsonschema:validation:Max=3
    Email string `json:"email"`

    // Password is the raw password of
    // the user trying to authenticate
    //
    // +jsonschema:validation:Min=2
    Password string `json:"password"`
}
