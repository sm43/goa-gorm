// Code generated by goa v3.1.1, DO NOT EDIT.
//
// user HTTP client CLI support package
//
// Command:
// $ goa gen github.com/sm43/goa-gorm/design

package client

import (
	"encoding/json"
	"fmt"

	user "github.com/sm43/goa-gorm/gen/user"
)

// BuildAddPayload builds the payload for the user add endpoint from CLI flags.
func BuildAddPayload(userAddBody string) (*user.User, error) {
	var err error
	var body AddRequestBody
	{
		err = json.Unmarshal([]byte(userAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"id\": 6623806583676096193,\n      \"name\": \"pwv\"\n   }'")
		}
	}
	v := &user.User{
		ID:   body.ID,
		Name: body.Name,
	}

	return v, nil
}
