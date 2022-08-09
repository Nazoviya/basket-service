CREATE TABLE "users" (
  "user_id" bigserial PRIMARY KEY,
  "user_name" varchar NOT NULL,
  "user_orders_count" bigint NOT NULL,
  "user_total_purchase" decimal NOT NULL
);

CREATE TABLE "baskets" (
  "basket_id" bigserial PRIMARY KEY,
  "basket_price" real NOT NULL,
  "user_id" bigint NOT NULL
);

CREATE TABLE "userbasket" (
  "product_id" bigserial PRIMARY KEY,
  "product_name" varchar NOT NULL,
  "product_price" real NOT NULL,
  "product_vat" smallint NOT NULL
);

CREATE TABLE "total_basket" (
  "price" real NOT NULL,
  "vat" real NOT NULL,
  "total_price" real NOT NULL,
  "discount" real NOT NULL
);

CREATE TABLE "orders" (
  "order_id" bigserial PRIMARY KEY,
  "order_date" bigint NOT NULL,
  "basket_id" bigint NOT NULL
);

CREATE TABLE "products" (
  "product_id" bigserial PRIMARY KEY,
  "product_name" varchar NOT NULL,
  "product_price" real NOT NULL,
  "product_vat" smallint NOT NULL
);

CREATE TABLE "order_details" (
  "order_dts_id" bigserial PRIMARY KEY,
  "quantity" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "order_id" bigint NOT NULL
);

ALTER TABLE "baskets" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "orders" ADD FOREIGN KEY ("basket_id") REFERENCES "baskets" ("basket_id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("order_id");

INSERT INTO "public"."products" ("product_id", "product_name", "product_price", "product_vat") VALUES
(1, 'Laptop', 2200, 18),
(2, 'Smartphone', 1350.99, 18),
(3, 'Tea', 40, 1),
(4, 'Sugar', 12, 1),
(5, 'Milk', 14.52, 1),
(6, 'Egg', 10.25, 1),
(7, 'Soap', 26, 8),
(8, 'Tooth paste', 35.75, 8),
(9, 'Paper towel', 116.2, 8),
(10, 'Keyboard', 280, 18);