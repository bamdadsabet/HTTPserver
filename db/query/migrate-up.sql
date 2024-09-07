-- name: create-branch-table

CREATE TABLE IF NOT EXISTS "branch" (
  id SERIAL PRIMARY KEY,
  name varchar(50),
  address varchar(255),
  phone_number varchar(50),
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
);

-- name: create-user-table
CREATE TABLE IF NOT EXISTS "app_user" (
  id SERIAL PRIMARY KEY,
  username varchar(50),
  password varchar(255),
  branch_id int REFERENCES branch(id),
  created_at timestamp  default current_timestamp,
  updated_at timestamp default current_timestamp
);

-- name: create-permission-table
CREATE TABLE IF NOT EXISTS "permission" (
  id SERIAL PRIMARY KEY,
  value varchar(255),
  created_at timestamp default current_timestamp
);

-- name: create-permission_user-table
CREATE TABLE IF NOT EXISTS "permission_user" (
  id SERIAL PRIMARY KEY,
  user_id int REFERENCES app_user(id),
  permission_id int REFERENCES permission(id)
);

-- name: create-employee-table
CREATE TABLE IF NOT EXISTS "employee" (
  id SERIAL PRIMARY KEY,
  first_name varchar(50) ,
  last_name varchar(50) ,
  email varchar(255),
  floor smallint,
  phone_extension varchar(4),
  role varchar(155),
  branch_id int REFERENCES branch(id),
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
);

-- name: create-request-table
CREATE TABLE IF NOT EXISTS "request" (
  id SERIAL PRIMARY KEY,
  requester_id int REFERENCES app_user(id),
  employee_id int REFERENCES employee(id),
  status Boolean ,
  description varchar(255),
  created_at timestamp default current_timestamp
);



