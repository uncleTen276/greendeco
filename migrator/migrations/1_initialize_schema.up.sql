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

CREATE TABLE IF NOT EXISTS "products"(
id  UUID DEFAULT gen_random_uuid () PRIMARY KEY,
category_id UUID NOT NULL,
name VARCHAR(200) UNIQUE NOT NULL,
is_publish BOOLEAN DEFAULT FALSE,
size VARCHAR(50),
type VARCHAR(10),
images TEXT[],
description TEXT,
detail TEXT,
light VARCHAR(50),
difficulty VARCHAR(50),
warter VARCHAR(200),
qr_image TEXT,
created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
updated_at TIMESTAMP WITH TIME ZONE  DEFAULT NOW (),
FOREIGN KEY(category_id) REFERENCES categories(id)
);

CREATE TABLE IF NOT EXISTS "recommends"(
product_id UUID NOT NULL,
recommend_product UUID NOT NULL, 
created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
updated_at TIMESTAMP WITH TIME ZONE  DEFAULT NOW (),
PRIMARY KEY (product_id, recommend_product),
FOREIGN KEY(product_id) REFERENCES products (id),
FOREIGN KEY(recommend_product) REFERENCES products (id)
);


-- CREATE ADMIN Account
-- password 1234567890
INSERT INTO "users" (email,identifier,password,first_name,last_name, phone_number,admin) VALUES ('admin@gmail.com','admin@gmail.com','$2a$10$xIgiGxp0THwDy1R8uxko..t3O8s9aeikqk9olnJCLLI/92FUbtFey','','admin','+844785976','true');

