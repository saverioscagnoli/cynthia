package ds

import (
	"camilla/service/util"
	"fmt"
	"net/http"
)

func (r *routes) GetUser(userID Snowflake) (string, string) {
	return http.MethodGet, fmt.Sprintf("/users/%s", userID)
}

func (r *routes) GetMe() (string, string) {
	return http.MethodGet, "/users/@me"
}

func (c *ApiClient) GetUser(userID Snowflake) (*User, error) {
	_, endpoint := Routes.GetUser(userID)
	res, err := c.Get(endpoint)

	if err != nil {
		return nil, err
	}

	return util.Decode[User](res)
}
