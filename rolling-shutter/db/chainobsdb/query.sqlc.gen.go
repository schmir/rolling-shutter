// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package chainobsdb

import (
	"context"
)

const getChainCollator = `-- name: GetChainCollator :one
SELECT activation_block_number, collator FROM chain_collator
WHERE activation_block_number <= $1
ORDER BY activation_block_number DESC LIMIT 1
`

func (q *Queries) GetChainCollator(ctx context.Context, activationBlockNumber int64) (ChainCollator, error) {
	row := q.db.QueryRow(ctx, getChainCollator, activationBlockNumber)
	var i ChainCollator
	err := row.Scan(&i.ActivationBlockNumber, &i.Collator)
	return i, err
}

const getEventSyncProgress = `-- name: GetEventSyncProgress :one
SELECT next_block_number, next_log_index FROM event_sync_progress LIMIT 1
`

type GetEventSyncProgressRow struct {
	NextBlockNumber int32
	NextLogIndex    int32
}

func (q *Queries) GetEventSyncProgress(ctx context.Context) (GetEventSyncProgressRow, error) {
	row := q.db.QueryRow(ctx, getEventSyncProgress)
	var i GetEventSyncProgressRow
	err := row.Scan(&i.NextBlockNumber, &i.NextLogIndex)
	return i, err
}

const getKeyperSet = `-- name: GetKeyperSet :one
SELECT keyper_config_index, activation_block_number, keypers, threshold FROM keyper_set
WHERE activation_block_number <= $1
ORDER BY activation_block_number DESC LIMIT 1
`

func (q *Queries) GetKeyperSet(ctx context.Context, activationBlockNumber int64) (KeyperSet, error) {
	row := q.db.QueryRow(ctx, getKeyperSet, activationBlockNumber)
	var i KeyperSet
	err := row.Scan(
		&i.KeyperConfigIndex,
		&i.ActivationBlockNumber,
		&i.Keypers,
		&i.Threshold,
	)
	return i, err
}

const getKeyperSetByKeyperConfigIndex = `-- name: GetKeyperSetByKeyperConfigIndex :one
SELECT keyper_config_index, activation_block_number, keypers, threshold FROM keyper_set WHERE keyper_config_index=$1
`

func (q *Queries) GetKeyperSetByKeyperConfigIndex(ctx context.Context, keyperConfigIndex int64) (KeyperSet, error) {
	row := q.db.QueryRow(ctx, getKeyperSetByKeyperConfigIndex, keyperConfigIndex)
	var i KeyperSet
	err := row.Scan(
		&i.KeyperConfigIndex,
		&i.ActivationBlockNumber,
		&i.Keypers,
		&i.Threshold,
	)
	return i, err
}

const getNextBlockNumber = `-- name: GetNextBlockNumber :one
SELECT next_block_number from event_sync_progress LIMIT 1
`

func (q *Queries) GetNextBlockNumber(ctx context.Context) (int32, error) {
	row := q.db.QueryRow(ctx, getNextBlockNumber)
	var next_block_number int32
	err := row.Scan(&next_block_number)
	return next_block_number, err
}

const insertChainCollator = `-- name: InsertChainCollator :exec
INSERT INTO chain_collator (activation_block_number, collator)
VALUES ($1, $2)
`

type InsertChainCollatorParams struct {
	ActivationBlockNumber int64
	Collator              string
}

func (q *Queries) InsertChainCollator(ctx context.Context, arg InsertChainCollatorParams) error {
	_, err := q.db.Exec(ctx, insertChainCollator, arg.ActivationBlockNumber, arg.Collator)
	return err
}

const insertKeyperSet = `-- name: InsertKeyperSet :exec
INSERT INTO keyper_set (
    keyper_config_index,
    activation_block_number,
    keypers,
    threshold
) VALUES (
    $1, $2, $3, $4
)
`

type InsertKeyperSetParams struct {
	KeyperConfigIndex     int64
	ActivationBlockNumber int64
	Keypers               []string
	Threshold             int32
}

func (q *Queries) InsertKeyperSet(ctx context.Context, arg InsertKeyperSetParams) error {
	_, err := q.db.Exec(ctx, insertKeyperSet,
		arg.KeyperConfigIndex,
		arg.ActivationBlockNumber,
		arg.Keypers,
		arg.Threshold,
	)
	return err
}

const updateEventSyncProgress = `-- name: UpdateEventSyncProgress :exec
INSERT INTO event_sync_progress (next_block_number, next_log_index)
VALUES ($1, $2)
ON CONFLICT (id) DO UPDATE
    SET next_block_number = $1,
        next_log_index = $2
`

type UpdateEventSyncProgressParams struct {
	NextBlockNumber int32
	NextLogIndex    int32
}

func (q *Queries) UpdateEventSyncProgress(ctx context.Context, arg UpdateEventSyncProgressParams) error {
	_, err := q.db.Exec(ctx, updateEventSyncProgress, arg.NextBlockNumber, arg.NextLogIndex)
	return err
}
