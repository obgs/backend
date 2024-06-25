-- Create "game_versions" table
CREATE TABLE "game_versions" ("id" character varying NOT NULL, "version_number" bigint NOT NULL DEFAULT 1, "game_version_game" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "game_versions_games_game" FOREIGN KEY ("game_version_game") REFERENCES "games" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "stat_description_game_version" table
CREATE TABLE "stat_description_game_version" ("stat_description_id" character varying NOT NULL, "game_version_id" character varying NOT NULL, PRIMARY KEY ("stat_description_id", "game_version_id"), CONSTRAINT "stat_description_game_version_game_version_id" FOREIGN KEY ("game_version_id") REFERENCES "game_versions" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "stat_description_game_version_stat_description_id" FOREIGN KEY ("stat_description_id") REFERENCES "stat_descriptions" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Modify "matches" table
ALTER TABLE "matches" ADD COLUMN "game_version_matches" character varying NOT NULL DEFAULT '', DROP CONSTRAINT "matches_games_matches";
