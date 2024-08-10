CREATE TABLE carts (
    cart_item_id SERIAL NOT NULL,
    user_id VARCHAR(100) NOT NULL,
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    CONSTRAINT carts_pkey PRIMARY KEY (cart_item_id),
    CONSTRAINT unique_user_product UNIQUE (user_id, product_id)
);

