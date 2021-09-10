// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package dcrdb

import (
	"context"
)

const getCipherBatch = `-- name: GetCipherBatch :one
SELECT epoch_id, data FROM decryptor.cipher_batch
WHERE epoch_id = $1
`

func (q *Queries) GetCipherBatch(ctx context.Context, epochID int64) (DecryptorCipherBatch, error) {
	row := q.db.QueryRow(ctx, getCipherBatch, epochID)
	var i DecryptorCipherBatch
	err := row.Scan(&i.EpochID, &i.Data)
	return i, err
}
