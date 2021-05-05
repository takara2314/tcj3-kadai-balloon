package database

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func GetKadais(ctx *context.Context, client *firestore.Client, kadais *map[string][]interface{}) error {
	iter := client.Collection("kadais").Documents(*ctx)
	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		detail := doc.Data()
		var kadaiId string = detail["kadai_id"].(string)

		(*kadais)[kadaiId] = append(
			(*kadais)[kadaiId],
			detail["class"],
			detail["subject"],
			detail["title"],
			detail["due"],
			detail["creator"],
			detail["editor"],
			detail["a_schedule_id"],
			detail["b_schedule_id"],
		)
	}

	return nil
}
