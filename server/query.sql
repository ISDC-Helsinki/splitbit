-- name: GetGroupsOfMember :many
SELECT * FROM groups JOIN member_groups ON groups.id = member_groups.group_id WHERE member_groups.member_id = ?;

-- name: GetGroupsAll :many
SELECT * FROM groups;

-- name: GetItemsOfGroup :many
SELECT * FROM items WHERE group_id = ? ORDER BY timestamp DESC;

-- name: AddItemToGroup :one
INSERT INTO items (name, timestamp, price, group_id, author_id, reimbursement) VALUES (?, ?, ?, ?, ?, ?) RETURNING id;

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

-- name: GetNetAmountForUserInGroup :one
SELECT 
    SUM(CASE 
            WHEN author_id = ? THEN -price 
            ELSE price 
        END) AS net_amount
FROM 
    items
WHERE 
    group_id = ?;

-- name: GetFriendsOfUser :many
SELECT mg2.member_id, COUNT(mg1.group_id) AS common_group_count
FROM member_groups mg1
INNER JOIN member_groups mg2 ON mg1.group_id = mg2.group_id
WHERE mg1.member_id = ? AND mg2.member_id != mg1.member_id
GROUP BY mg2.member_id;


-- name: GetGroupByID :one
SELECT 
    id, 
    name 
FROM 
    groups 
WHERE 
    id = ?;

