-- set timezone
SET TIMEZONE="Asia/Dhaka";
-- create users table
CREATE TABLE IF NOT EXISTS "users"(
id UUID DEFAULT gen_random_uuid () PRIMARY KEY,
email VARCHAR(50) UNIQUE NOT NULL,
identifier VARCHAR(50) UNIQUE NOT NULL,
password VARCHAR(200) NOT NULL,
first_name VARCHAR(50),
last_name VARCHAR(50),
phone_number VARCHAR(12),
avatar VARCHAR(200),
admin BOOLEAN DEFAULT FALSE NOT NULL,
created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
updated_at TIMESTAMP WITH TIME ZONE  DEFAULT NOW ()
);

CREATE TABLE IF NOT EXISTS "roles"(
id  UUID DEFAULT gen_random_uuid () PRIMARY KEY,
name  VARCHAR(50) UNIQUE NOT NULL,
created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
updated_at TIMESTAMP WITH TIME ZONE  DEFAULT NOW ()
);

CREATE TABLE IF NOT EXISTS "user_role"(
user_id UUID  NOT NULL,
role_id UUID  NOT NULL,
created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
updated_at TIMESTAMP WITH TIME ZONE  DEFAULT NOW (),
PRIMARY KEY (user_id, role_id),
FOREIGN KEY(role_id) REFERENCES roles (id),
FOREIGN KEY(user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS "categories"(
id  UUID DEFAULT gen_random_uuid () PRIMARY KEY,
name VARCHAR(200) UNIQUE NOT NULL, 
created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
updated_at TIMESTAMP WITH TIME ZONE  DEFAULT NOW ()
);
-- CREATE ADMIN Account
INSERT INTO "users" (email,identifier,password,first_name,last_name, phone_number,admin) VALUES ('admin@gmail.com','admin@gmail.com','1234567890','','admin','+844785976','true');

