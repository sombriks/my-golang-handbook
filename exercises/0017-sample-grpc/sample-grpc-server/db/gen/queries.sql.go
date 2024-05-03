// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package gen

import (
	"context"
)

const find = `-- name: Find :one
select id, description, done, created, updated from todos where id = ?
`

func (q *Queries) Find(ctx context.Context, id int64) (Todo, error) {
	row := q.db.QueryRowContext(ctx, find, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Done,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const insert = `-- name: Insert :one
insert into todos (description,done) values (?,?) returning id, description, done, created, updated
`

type InsertParams struct {
	Description string
	Done        bool
}

func (q *Queries) Insert(ctx context.Context, arg InsertParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, insert, arg.Description, arg.Done)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Done,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const list = `-- name: List :many
select id, description, done, created, updated from todos where lower(description) like lower(concat('%',?,'%'))
`

func (q *Queries) List(ctx context.Context, concat interface{}) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, list, concat)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Description,
			&i.Done,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const update = `-- name: Update :one
update todos set description = ?, done = ?,
                 updated = CURRENT_TIMESTAMP
             where id = ? returning id, description, done, created, updated
`

type UpdateParams struct {
	Description string
	Done        bool
	ID          int64
}

func (q *Queries) Update(ctx context.Context, arg UpdateParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, update, arg.Description, arg.Done, arg.ID)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Done,
		&i.Created,
		&i.Updated,
	)
	return i, err
}
