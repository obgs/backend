-- modify "stat_descriptions" table
ALTER TABLE "stat_descriptions" ADD COLUMN "metadata" character varying NULL;
ALTER TABLE "stat_descriptions" DROP COLUMN "possible_values";
