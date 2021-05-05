package database

import (
	"context"

	"cloud.google.com/go/firestore"
)

func RemoveUser(ctx *context.Context, client *firestore.Client, users *map[string][]interface{}, lineId string) error {
	_, err := client.Collection("users").
		Doc(lineId).
		Delete(*ctx)

	if err != nil {
		return err
	}

	delete(*users, lineId)

	return nil
}
