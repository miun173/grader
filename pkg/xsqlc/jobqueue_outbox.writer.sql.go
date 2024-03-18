// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: jobqueue_outbox.writer.sql

package xsqlc

import (
	"context"
)

const createOutboxItem = `-- name: CreateOutboxItem :one
INSERT INTO outbox_items (id, idempotent_key, "status", job_type, payload)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, "version"
`

type CreateOutboxItemParams struct {
	ID            string
	IdempotentKey string
	Status        string
	JobType       string
	Payload       string
}

type CreateOutboxItemRow struct {
	ID      string
	Version int32
}

func (q *Queries) CreateOutboxItem(ctx context.Context, arg CreateOutboxItemParams) (CreateOutboxItemRow, error) {
	row := q.db.QueryRowContext(ctx, createOutboxItem,
		arg.ID,
		arg.IdempotentKey,
		arg.Status,
		arg.JobType,
		arg.Payload,
	)
	var i CreateOutboxItemRow
	err := row.Scan(&i.ID, &i.Version)
	return i, err
}

const updateOutboxItem = `-- name: UpdateOutboxItem :one
UPDATE outbox_items
SET 
    "status" = $1,
    idempotent_key = $2,
    job_type = $3,
    payload = $4,
    "version" = "version" + 1
WHERE id = $5
    -- do optimistic locking
    AND "version" = "version"
    AND "version" = $6
RETURNING id, "version"
`

type UpdateOutboxItemParams struct {
	Status        string
	IdempotentKey string
	JobType       string
	Payload       string
	ID            string
	Version       int32
}

type UpdateOutboxItemRow struct {
	ID      string
	Version int32
}

func (q *Queries) UpdateOutboxItem(ctx context.Context, arg UpdateOutboxItemParams) (UpdateOutboxItemRow, error) {
	row := q.db.QueryRowContext(ctx, updateOutboxItem,
		arg.Status,
		arg.IdempotentKey,
		arg.JobType,
		arg.Payload,
		arg.ID,
		arg.Version,
	)
	var i UpdateOutboxItemRow
	err := row.Scan(&i.ID, &i.Version)
	return i, err
}
