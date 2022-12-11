-- create "games" table
CREATE TABLE "games" ("id" character varying NOT NULL, "name" character varying NOT NULL, "min_players" bigint NOT NULL DEFAULT 1, "max_players" bigint NOT NULL DEFAULT 10, "description" character varying NULL DEFAULT '', "boardgamegeek_url" character varying NULL, "user_games" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "games_users_games" FOREIGN KEY ("user_games") REFERENCES "users" ("id") ON DELETE NO ACTION);
-- create "game_favorites" table
CREATE TABLE "game_favorites" ("id" character varying NOT NULL, "game_favorites" character varying NOT NULL, "user_favorite_games" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "game_favorites_games_favorites" FOREIGN KEY ("game_favorites") REFERENCES "games" ("id") ON DELETE NO ACTION, CONSTRAINT "game_favorites_users_favorite_games" FOREIGN KEY ("user_favorite_games") REFERENCES "users" ("id") ON DELETE NO ACTION);
-- create "stat_descriptions" table
CREATE TABLE "stat_descriptions" ("id" character varying NOT NULL, "type" character varying NOT NULL, "name" character varying NOT NULL, "description" character varying NULL DEFAULT '', PRIMARY KEY ("id"));
-- create index "stat_descriptions_name_key" to table: "stat_descriptions"
CREATE UNIQUE INDEX "stat_descriptions_name_key" ON "stat_descriptions" ("name");
-- create "stat_description_game" table
CREATE TABLE "stat_description_game" ("stat_description_id" character varying NOT NULL, "game_id" character varying NOT NULL, PRIMARY KEY ("stat_description_id", "game_id"), CONSTRAINT "stat_description_game_stat_description_id" FOREIGN KEY ("stat_description_id") REFERENCES "stat_descriptions" ("id") ON DELETE CASCADE, CONSTRAINT "stat_description_game_game_id" FOREIGN KEY ("game_id") REFERENCES "games" ("id") ON DELETE CASCADE);
