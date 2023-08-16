-- set timezone
SET TIMEZONE="Asia/Dhaka";
-- create users table
CREATE TABLE "users"(
id UUID DEFAULT gen_random_uuid () PRIMARY KEY,
created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
updated_at TIMESTAMP WITH TIME ZONE  DEFAULT NOW (),
email VARCHAR(50) UNIQUE NOT NULL,
identifier VARCHAR(50) UNIQUE NOT NULL,
phone_number VARCHAR(12) NULL,
password VARCHAR(200) NOT NULL,
first_name VARCHAR(50),
last_name VARCHAR(50)
admin BOOLEAN DEFAULT FALSE NOT NULL
)

