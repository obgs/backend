-- modify "stat_descriptions" table
ALTER TABLE "stat_descriptions" ADD COLUMN "possible_values" jsonb NULL;
