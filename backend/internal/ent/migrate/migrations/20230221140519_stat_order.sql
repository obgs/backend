-- modify "stat_descriptions" table
ALTER TABLE "stat_descriptions" ADD COLUMN "order_number" bigint NOT NULL DEFAULT 0;

-- modify existing stats by adding order number for them
UPDATE "stat_descriptions"
  SET "order_number" = subquery.order_number
  FROM (
    SELECT
      ROW_NUMBER() OVER(PARTITION BY game_id ORDER BY id) AS order_number,
      id
    FROM stat_descriptions
    JOIN stat_description_game
    ON stat_description_game.stat_description_id = stat_descriptions.id
  ) as subquery
WHERE stat_descriptions.id = subquery.id;

-- drop default value for order_number
ALTER TABLE "stat_descriptions" ALTER COLUMN "order_number" DROP DEFAULT;
