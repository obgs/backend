-- create "players" table
CREATE TABLE "players" ("id" uuid NOT NULL, "name" character varying NOT NULL DEFAULT '', "user_main_player" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "players_users_main_player" FOREIGN KEY ("user_main_player") REFERENCES "users" ("id") ON DELETE SET NULL);
-- create index "players_user_main_player_key" to table: "players"
CREATE UNIQUE INDEX "players_user_main_player_key" ON "players" ("user_main_player");
-- create "user_players" table
CREATE TABLE "user_players" ("user_id" uuid NOT NULL, "player_id" uuid NOT NULL, PRIMARY KEY ("user_id", "player_id"), CONSTRAINT "user_players_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE, CONSTRAINT "user_players_player_id" FOREIGN KEY ("player_id") REFERENCES "players" ("id") ON DELETE CASCADE);


-- create a player for each user
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
INSERT INTO players (id, user_main_player) SELECT uuid_generate_v4(), id FROM users;
