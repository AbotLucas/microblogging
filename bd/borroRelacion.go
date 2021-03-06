package bd

import (
	"context"
	"time"

	"github.com/abotlucas/microblogging/models"
)

/*BorroRelacion borra la relacion en la bd */
func BorroRelacion(rel models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, rel)
	if err != nil {
		return false, err
	}

	return true, nil

}
