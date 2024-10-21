-- name: GetGroupsOfMember :many
SELECT * FROM groups JOIN member_groups ON groups.id = member_groups.group_id WHERE member_groups.member_id = ?;

-- name: GetGroupsAll :many
SELECT * FROM groups;

-- name: GetItemsOfGroup :many
SELECT * FROM items WHERE group_id = ? ORDER BY timestamp DESC;

-- name: AddItemToGroup :one
INSERT INTO items (name, timestamp, price, group_id, author_id) VALUES (?, ?, ?, ?, ?) RETURNING id;

-- name: GetUserByUsernameAndPassword :one
SELECT id, username FROM members WHERE username = ? AND password = ?;

-- name: AddUser :one
INSERT INTO members (username, displayName, password) VALUES (?, ?, ?) RETURNING id;

-- name: AddGroup :one
INSERT INTO groups (name) VALUES (?) RETURNING id;

-- name: GetMembersOfGroup :many
SELECT id, username, displayName FROM members JOIN member_groups ON member_groups.member_id = members.id WHERE group_id = ?;

-- name: AddMemberToGroup :exec
INSERT INTO member_groups (group_id, member_id) VALUES (?, ?);
