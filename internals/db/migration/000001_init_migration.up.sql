CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
	"id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	"email" VARCHAR(100) UNIQUE NOT NULL,
	"first_name" VARCHAR(200) NOT NULL,
	"last_name" VARCHAR(200) NOT NULL,
	"password" VARCHAR(100) NOT NULL,
	"role" VARCHAR(50) DEFAULT 'user',
	"is_verified" BOOLEAN DEFAULT FALSE,
	"is_deleted" BOOLEAN DEFAULT FALSE,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP DEFAULT NULL
);

CREATE TABLE "urls" (
	"id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	"original_url" TEXT NOT NULL,
	"short_code" VARCHAR(60) UNIQUE NOT NULL,
	"click_count" INT DEFAULT 0,
	"is_active" BOOLEAN DEFAULT TRUE,
	"user_id" uuid REFERENCES users("id"),
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW())
);

CREATE TABLE "verification_codes" (
	"id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	"code" VARCHAR(255) UNIQUE NOT NULL,
	"is_active" BOOLEAN DEFAULT TRUE,
	"expires_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"user_id" uuid REFERENCES users("id") ON DELETE CASCADE,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW())
);

CREATE TABLE "password_reset" (
	"id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	"token" VARCHAR(255) UNIQUE NOT NULL,
	"is_active" BOOLEAN DEFAULT TRUE,
	"expires_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"user_id" uuid REFERENCES users("id") ON DELETE CASCADE,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW())
);


ALTER TABLE "urls" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "verification_codes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "password_reset" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE INDEX ON "urls" ("id");
CREATE INDEX ON "verification_codes" ("id");
CREATE INDEX ON "password_reset" ("id");