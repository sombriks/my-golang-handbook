-- name: List :many
select * from todos where lower(description) like lower(concat('%',?,'%'));

-- name: Find :one
select * from todos where id = ?;

-- name: Insert :one
insert into todos (description,done) values (?,?) returning *;

-- name: Update :one
update todos set description = ?, done = ?,
                 updated = CURRENT_TIMESTAMP
             where id = ? returning *
