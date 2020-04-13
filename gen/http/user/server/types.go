// Code generated by goa v3.1.1, DO NOT EDIT.
//
// user HTTP server types
//
// Command:
// $ goa gen github.com/sm43/goa-gorm/design

package server

import (
	"unicode/utf8"

	user "github.com/sm43/goa-gorm/gen/user"
	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "user" service "add" endpoint HTTP request
// body.
type AddRequestBody struct {
	// ID of a user
	ID *uint64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of person
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// AddResponseBody is the type of the "user" service "add" endpoint HTTP
// response body.
type AddResponseBody struct {
	// ID of a user
	ID *uint64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of person
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// ListResponseBody is the type of the "user" service "list" endpoint HTTP
// response body.
type ListResponseBody []*StoredUserResponse

// StoredUserResponse is used to define fields on response body types.
type StoredUserResponse struct {
	// ID is the unique id of the blog.
	ID uint64 `form:"id" json:"id" xml:"id"`
	// Name of person
	Name string `form:"name" json:"name" xml:"name"`
}

// NewAddResponseBody builds the HTTP response body from the result of the
// "add" endpoint of the "user" service.
func NewAddResponseBody(res *user.User) *AddResponseBody {
	body := &AddResponseBody{
		ID:   res.ID,
		Name: res.Name,
	}
	return body
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "user" service.
func NewListResponseBody(res []*user.StoredUser) ListResponseBody {
	body := make([]*StoredUserResponse, len(res))
	for i, val := range res {
		body[i] = marshalUserStoredUserToStoredUserResponse(val)
	}
	return body
}

// NewAddUser builds a user service add endpoint payload.
func NewAddUser(body *AddRequestBody) *user.User {
	v := &user.User{
		ID:   body.ID,
		Name: body.Name,
	}

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	return
}
