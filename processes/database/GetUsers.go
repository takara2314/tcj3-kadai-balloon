package database

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func GetUsers(ctx *context.Context, client *firestore.Client, users *map[string][]interface{}) error {
	iter := client.Collection("users").Documents(*ctx)
	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		profile := doc.Data()
		var lineId string = profile["student_num"].(string)

		(*users)[lineId] = append(
			(*users)[lineId],
			profile["line_id"],
			profile["line_name"],
			profile["student_num"],
			profile["class"],
		)
	}

	return nil
}
