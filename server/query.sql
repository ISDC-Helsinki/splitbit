-- name: GetGroupsOfMember :many
SELECT * FROM groups JOIN member_groups ON groups.id = member_groups.group_id WHERE member_groups.member_id = ?;

-- name: GetActiveGroupsOfMember :many
SELECT * FROM groups JOIN member_groups ON groups.id = member_groups.group_id WHERE member_groups.member_id = ?;

-- name: GetActiveGroupsOfMemberAndAmountOwed :many
SELECT 
    g.id AS group_id,
    g.name AS group_name,
    mg.member_id,
    COALESCE((SELECT SUM(CASE 
                WHEN i.author_id = 1 THEN -i.price 
                ELSE i.price 
            END)
     FROM items i
     WHERE i.group_id = g.id), 0) AS net_amount
FROM 
    groups g
JOIN 
    member_groups mg 
    ON g.id = mg.group_id
WHERE 
    mg.member_id = ?;

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
SELECT 
    mg2.member_id, 
    m.username,
    m.displayName, -- Adding the member's displayName from the members table
    COUNT(mg1.group_id) AS common_group_count
FROM member_groups mg1
INNER JOIN member_groups mg2 ON mg1.group_id = mg2.group_id
INNER JOIN members m ON mg2.member_id = m.id  -- Join with members to get displayName
WHERE mg1.member_id = ? AND mg2.member_id != mg1.member_id
GROUP BY mg2.member_id, m.displayName;  -- Group by member_id and displayName



-- name: GetGroupByID :one
SELECT 
    id, 
    name 
FROM 
    groups 
WHERE 
    id = ?;

-- name: ArchiveGroup :exec
UPDATE groups SET is_archived = TRUE WHERE id = ?;

-- name: UnarchiveGroup :exec
UPDATE groups SET is_archived = FALSE WHERE id = ?;
