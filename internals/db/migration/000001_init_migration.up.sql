CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
	"id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	"email" VARCHAR(100) UNIQUE NOT NULL,
	"first_name" VARCHAR(200) NOT NULL,
	"last_name" VARCHAR(200) NOT NULL,
	"password_hash" VARCHAR NOT NULL,
	"role" VARCHAR(50) DEFAULT 'user',
	"is_deleted" BOOLEAN DEFAULT FALSE,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP NULL
);

CREATE TABLE "urls" (
	"id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	"original_url" TEXT NOT NULL,
	"short_code" VARCHAR(20) UNIQUE NOT NULL,
	"click_count" INT DEFAULT 0,
	"is_active" BOOLEAN DEFAULT TRUE,
	"user_id" uuid REFERENCES users("id"),
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW())
);


ALTER TABLE "urls" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE INDEX ON "urls" ("id");