CREATE TYPE "user_role" AS ENUM (
  'super',
  'owner',
  'employee'
);

CREATE TYPE "product_type" AS ENUM (
  'super',
  'owner',
  'employee'
);

CREATE TYPE "order_status" AS ENUM (
  'super',
  'owner',
  'employee'
);

CREATE TYPE "conversation_stage" AS ENUM (
  'super',
  'owner',
  'employee'
);

CREATE TABLE "users" (
  "id" int PRIMARY KEY,
  "username" varchar UNIQUE,
  "password" varchar,
  "role" user_role,
  "active" bool DEFAULT false,
  "created_at" timestamp DEFAULT (now()),
  "store" int
);

CREATE TABLE "stores" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "wa_token" varchar,
  "active" bool,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "product_categories" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "created_at" timestamp DEFAULT (now()),
  "store" int
);

CREATE TABLE "products" (
  "id" int PRIMARY KEY,
  "type" product_type,
  "name" varchar,
  "price" float,
  "created_at" timestamp DEFAULT (now()),
  "store" int,
  "category" int
);

CREATE TABLE "product_variants" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "price" float,
  "created_at" timestamp DEFAULT (now()),
  "product" int
);

CREATE TABLE "product_components" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "price" float,
  "created_at" timestamp DEFAULT (now()),
  "product" int
);

CREATE TABLE "orders" (
  "id" int PRIMARY KEY,
  "status" order_status,
  "price" float,
  "created_at" timestamp DEFAULT (now()),
  "store" int
);

CREATE TABLE "order_items" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "details" text,
  "price" float,
  "created_at" timestamp DEFAULT (now()),
  "order" int,
  "product" int
);

CREATE TABLE "customers" (
  "id" int PRIMARY KEY,
  "wa_id" varchar,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "conversations" (
  "id" int PRIMARY KEY,
  "created_at" timestamp DEFAULT (now()),
  "customer" int,
  "store" int
);

ALTER TABLE "stores" ADD FOREIGN KEY ("id") REFERENCES "users" ("store");

ALTER TABLE "product_categories" ADD FOREIGN KEY ("store") REFERENCES "stores" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("store") REFERENCES "stores" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("category") REFERENCES "product_categories" ("id");

ALTER TABLE "product_variants" ADD FOREIGN KEY ("product") REFERENCES "products" ("id");

ALTER TABLE "product_components" ADD FOREIGN KEY ("product") REFERENCES "products" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("store") REFERENCES "stores" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("product") REFERENCES "products" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order") REFERENCES "orders" ("id");

ALTER TABLE "conversations" ADD FOREIGN KEY ("customer") REFERENCES "customers" ("id");

ALTER TABLE "conversations" ADD FOREIGN KEY ("store") REFERENCES "stores" ("id");
