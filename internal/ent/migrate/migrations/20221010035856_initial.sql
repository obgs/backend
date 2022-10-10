-- create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "name" character varying NOT NULL DEFAULT '', "email" character varying NOT NULL, "password" character varying NOT NULL, "avatar_url" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
