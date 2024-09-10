-- name: create-branch
INSERT INTO branch 
(name, address, phone_number) 
VALUES($1, $2, $3) 
RETURNING *;

-- name: update-branch
UPDATE branch 
SET name=$2 , address=$3, phone_number=$4, updated_at=$5 
WHERE id=$1
RETURNING *;

-- name: delete-branch
DELETE FROM branch 
WHERE id=$1;

-- name: get-branch-list
SELECT * 
FROM branch;

-- name : get-branch
SELECT * 
FROM branch 
WHERE id=$1


