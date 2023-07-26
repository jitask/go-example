// You can edit this code!
// Click here and start typing.
package main

import (
	"context"
	"encoding/json"
	"fmt"
)

type ListPageVersionRequest struct {
	TeamUUID  string `json:"team_uuid"`
	SpaceUUID string `json:"space_uuid"`
}

func BindJSON[T any](ctx context.Context) (obj *T, err error) {
	obj = new(T)
	err = json.Unmarshal([]byte(`{"team_uuid": "aaaaaaaa"}`), obj)
	if err != nil {
		fmt.Println(err.Error())
	}

	return obj, nil
}

func main() {
	req, err := BindJSON[ListPageVersionRequest](context.TODO())
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(req.TeamUUID)
}
