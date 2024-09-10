--name: create-employee
INSERT INTO employee 
(first_name, last_name, email, floor, phone_extension, role, branch_id) 
VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *;

--name: get-all-employee
SELECT * 
FROM employee;

--name: get-employee-by-id
SELECT * 
FROM employee 
WHERE id = $1;

--name: delete-employee-by-id 
DELETE FROM employee 
WHERE id = $1;

--name: update-employee
UPDATE employee 
SET first_name = $2, last_name = $3, email = $4, floor = $5, phone_extension = $6, role = $7, branch_id = $8, updated_at = $9
WHERE id = $1 RETURNING *;

 
