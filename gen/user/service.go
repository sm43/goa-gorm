// Code generated by goa v3.1.1, DO NOT EDIT.
//
// user service
//
// Command:
// $ goa gen github.com/sm43/goa-gorm/design

package user

import (
	"context"
)

// The user service gives user details.
type Service interface {
	// Add new user and return its ID.
	Add(context.Context, *User) (res *User, err error)
	// List all users
	List(context.Context) (res []*StoredUser, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "user"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"add", "list"}

// User is the payload type of the user service add method.
type User struct {
	// ID of a user
	ID *uint64
	// Name of person
	Name *string
}

// A StoredUser describes a user retrieved by the storage service.
type StoredUser struct {
	// ID is the unique id of the blog.
	ID uint64
	// Name of person
	Name string
}
