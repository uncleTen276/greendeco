-- set timezone
SET TIMEZONE="Asia/Dhaka";

CREATE EXTENSION IF NOT EXISTS pg_trgm;
SELECT * FROM pg_extension;
-- create users table
CREATE TABLE IF NOT EXISTS "users"(
id UUID DEFAULT gen_random_uuid () PRIMARY KEY,
email VARCHAR(50) UNIQUE NOT NULL,
identifier VARCHAR(50) UNIQUE NOT NULL,
password VARCHAR(200) NOT NULL,
first_name VARCHAR(50),
last_name VARCHAR(50),
phone_number VARCHAR(12),
avatar TEXT,
admin BOOLEAN DEFAULT FALSE NOT NULL,
created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
updated_at TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS "roles"(
id  UUID DEFAULT gen_random_uuid () PRIMARY KEY,
name  VARCHAR(50) UNIQUE NOT NULL,
created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
updated_at TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS "user_role"(
user_id UUID  NOT NULL,
role_id UUID  NOT NULL,
created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
updated_at TIMESTAMP DEFAULT current_timestamp,
PRIMARY KEY (user_id, role_id),
FOREIGN KEY(role_id) REFERENCES roles (id),
FOREIGN KEY(user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS "categories"(
id  UUID DEFAULT gen_random_uuid () PRIMARY KEY,
name VARCHAR(200) UNIQUE NOT NULL, 
created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
updated_at TIMESTAMP DEFAULT current_timestamp 
);

CREATE TABLE IF NOT EXISTS "products"(
id  UUID DEFAULT gen_random_uuid () PRIMARY KEY,
category_id UUID NOT NULL,
name VARCHAR(200) UNIQUE NOT NULL,
is_publish BOOLEAN DEFAULT FALSE,
available BOOLEAN DEFAULT FALSE,
size VARCHAR(50) NOT NULL,
type VARCHAR(10) NOT NULL,
images TEXT[] NOT NULL,
description TEXT,
detail TEXT NOT NULL,
light VARCHAR(50) NOT NULL,
difficulty VARCHAR(50) NOT NULL,
warter VARCHAR(200) NOT NULL,
qr_image TEXT,
created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
updated_at TIMESTAMP DEFAULT current_timestamp,
FOREIGN KEY(category_id) REFERENCES categories(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "recommends"(
product_id UUID NOT NULL,
recommend_product UUID NOT NULL, 
created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
updated_at TIMESTAMP DEFAULT current_timestamp,
PRIMARY KEY (product_id, recommend_product),
FOREIGN KEY(product_id) REFERENCES products (id) ON DELETE CASCADE,
FOREIGN KEY(recommend_product) REFERENCES products (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "variants"(
id  UUID DEFAULT gen_random_uuid () PRIMARY KEY,
available BOOLEAN,
product UUID NOT NULL,
name VARCHAR(200) UNIQUE NOT NULL,
color VARCHAR(50), 
price DECIMAL NOT NULL,
currency VARCHAR(10) NOT NULL,
image TEXT NOT NULL,
description TEXT NOT NULL,   
created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
updated_at TIMESTAMP DEFAULT current_timestamp,
FOREIGN KEY(product) REFERENCES products (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "default_product_variant"(
product_id UUID NOT NULL UNIQUE, 
variant_id UUID NOT NULL UNIQUE,
created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
updated_at TIMESTAMP DEFAULT current_timestamp,
PRIMARY KEY (product_id, variant_id),
FOREIGN KEY(product_id) REFERENCES products (id) ON DELETE CASCADE,
FOREIGN KEY(variant_id) REFERENCES variants (id) ON DELETE CASCADE
);

CREATE VIEW "published_products" AS 
SELECT products.id, products.category_id, products.name,
    products.available , products.size, products.type,
    products.images, products.description, products.detail,
    products.light, products.difficulty, products.warter, variants.price ,   products.created_at, default_product_variant.variant_id
FROM products, variants,default_product_variant 
WHERE default_product_variant.product_id = products.id 
    AND default_product_variant.variant_id = variants.id 
    AND products.is_publish = true;



-- CREATE ADMIN Account
-- password 1234567890
INSERT INTO "users" (email,identifier,password,first_name,last_name, phone_number,admin) VALUES ('admin@gmail.com','admin@gmail.com','$2a$10$xIgiGxp0THwDy1R8uxko..t3O8s9aeikqk9olnJCLLI/92FUbtFey','','admin','+844785976','true');

