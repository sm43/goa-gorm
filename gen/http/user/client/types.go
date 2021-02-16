// Code generated by goa v3.1.1, DO NOT EDIT.
//
// user HTTP client types
//
// Command:
// $ goa gen github.com/sm43/goa-gorm/design

package client

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

// ListResponseBody is the type of the "user" service "list" endpoint HTTP
// response body.
type ListResponseBody []*StoredUserResponse

// AddDbErrorResponseBody is the type of the "user" service "add" endpoint HTTP
// response body for the "db_error" error.
type AddDbErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ListDbErrorResponseBody is the type of the "user" service "list" endpoint
// HTTP response body for the "db_error" error.
type ListDbErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// StoredUserResponse is used to define fields on response body types.
type StoredUserResponse struct {
	// ID is the unique id of the blog.
	ID *uint64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of person
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// NewAddRequestBody builds the HTTP request body from the payload of the "add"
// endpoint of the "user" service.
func NewAddRequestBody(p *user.User) *AddRequestBody {
	body := &AddRequestBody{
		ID:   p.ID,
		Name: p.Name,
	}
	return body
}

// NewAddDbError builds a user service add endpoint db_error error.
func NewAddDbError(body *AddDbErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewListStoredUserOK builds a "user" service "list" endpoint result from a
// HTTP "OK" response.
func NewListStoredUserOK(body []*StoredUserResponse) []*user.StoredUser {
	v := make([]*user.StoredUser, len(body))
	for i, val := range body {
		v[i] = unmarshalStoredUserResponseToUserStoredUser(val)
	}
	return v
}

// NewListDbError builds a user service list endpoint db_error error.
func NewListDbError(body *ListDbErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateAddDbErrorResponseBody runs the validations defined on
// add_db_error_response_body
func ValidateAddDbErrorResponseBody(body *AddDbErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateListDbErrorResponseBody runs the validations defined on
// list_db_error_response_body
func ValidateListDbErrorResponseBody(body *ListDbErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateStoredUserResponse runs the validations defined on StoredUserResponse
func ValidateStoredUserResponse(body *StoredUserResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	return
}
