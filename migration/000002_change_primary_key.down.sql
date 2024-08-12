ALTER TABLE carts DROP CONSTRAINT IF EXISTS carts_pkey;

DROP INDEX IF EXISTS carts_pkey CASCADE;

ALTER TABLE carts ADD COLUMN cart_item_id SERIAL PRIMARY KEY;

CREATE UNIQUE INDEX unique_user_product ON carts USING btree (user_id, product_id);