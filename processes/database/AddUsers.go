package database

import (
	"context"

	"cloud.google.com/go/firestore"
)

func AddUsers(ctx *context.Context, client *firestore.Client, users *map[string][]interface{}, profile *map[string]interface{}) error {
	_, err := client.Collection("users").
		Doc((*profile)["line_id"].(string)).
		Set(*ctx, *profile)

	if err != nil {
		return err
	}

	var lineId string = (*profile)["line_id"].(string)

	(*users)[lineId] = append(
		(*users)[lineId],
		(*profile)["line_id"],
		(*profile)["line_name"],
		(*profile)["student_num"],
		(*profile)["class"],
	)

	return nil
}
