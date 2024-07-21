ALTER TABLE "matches" ADD CONSTRAINT "matches_game_versions_matches" FOREIGN KEY ("game_version_matches") REFERENCES "game_versions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;

DROP TABLE "stat_description_game";

ALTER TABLE "matches" DROP COLUMN "game_matches", ALTER COLUMN "game_version_matches" DROP DEFAULT;
