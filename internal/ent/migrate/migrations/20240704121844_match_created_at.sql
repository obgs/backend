-- Modify "matches" table
ALTER TABLE "matches" ADD COLUMN "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP;
