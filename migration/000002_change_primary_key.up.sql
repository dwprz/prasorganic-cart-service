ALTER TABLE carts DROP CONSTRAINT IF EXISTS unique_user_product;

DROP INDEX IF EXISTS unique_user_product CASCADE;

ALTER TABLE carts DROP COLUMN IF EXISTS cart_item_id CASCADE;

ALTER TABLE carts ADD PRIMARY KEY (user_id, product_id);