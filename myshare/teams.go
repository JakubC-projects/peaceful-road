package myshare

import (
	"context"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
)

func (c *Client) GetOrgs(ctx context.Context, u peacefulroad.User) ([]peacefulroad.Org, error) {
	var res response[[]peacefulroad.Org]

	err := c.get(ctx, u.Token, "/Profile/Organisations", &res)
	return res.Data, err
}
