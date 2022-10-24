-- find and save all foreign keys into a temporary table
SELECT conrelid::regclass AS table_name,
       conname AS foreign_key,
       pg_get_constraintdef(oid) AS constraint_def
INTO ids_with_types_fks
FROM   pg_constraint
WHERE  contype = 'f'
AND    connamespace = 'public'::regnamespace
ORDER  BY conrelid::regclass::text, contype DESC;

-- remove all foreign keys
DO 'DECLARE r record;
    BEGIN
        FOR r IN SELECT table_name, foreign_key
                 FROM ids_with_types_fks
        LOOP
            EXECUTE ''ALTER TABLE '' || r.table_name || '' DROP CONSTRAINT ''|| r.foreign_key || '';'';
        END LOOP;
    END';

-- modify "group_applications" table
ALTER TABLE "group_applications" ALTER COLUMN "group_id" TYPE character varying, ALTER COLUMN "group_membership_application_id" TYPE character varying;
-- modify "group_membership_applications" table
ALTER TABLE "group_membership_applications" ALTER COLUMN "id" TYPE character varying;
-- modify "group_memberships" table
ALTER TABLE "group_memberships" ALTER COLUMN "id" TYPE character varying, ALTER COLUMN "group_members" TYPE character varying, ALTER COLUMN "user_group_memberships" TYPE character varying;
-- modify "group_settings" table
ALTER TABLE "group_settings" ALTER COLUMN "id" TYPE character varying, ALTER COLUMN "group_settings" TYPE character varying;
-- modify "groups" table
ALTER TABLE "groups" ALTER COLUMN "id" TYPE character varying;
-- modify "player_supervision_request_approvals" table
ALTER TABLE "player_supervision_request_approvals" ALTER COLUMN "id" TYPE character varying, ALTER COLUMN "player_supervision_request_approvals" TYPE character varying, ALTER COLUMN "user_supervision_request_approvals" TYPE character varying;
-- modify "player_supervision_requests" table
ALTER TABLE "player_supervision_requests" ALTER COLUMN "id" TYPE character varying, ALTER COLUMN "player_supervision_requests" TYPE character varying, ALTER COLUMN "user_sent_supervision_requests" TYPE character varying;
-- modify "players" table
ALTER TABLE "players" ALTER COLUMN "id" TYPE character varying, ALTER COLUMN "user_main_player" TYPE character varying;
-- modify "user_group_membership_applications" table
ALTER TABLE "user_group_membership_applications" ALTER COLUMN "user_id" TYPE character varying, ALTER COLUMN "group_membership_application_id" TYPE character varying;
-- modify "user_players" table
ALTER TABLE "user_players" ALTER COLUMN "user_id" TYPE character varying, ALTER COLUMN "player_id" TYPE character varying;
-- modify "users" table
ALTER TABLE "users" ALTER COLUMN "id" TYPE character varying;

-- add the foreign keys back
DO 'DECLARE r record;
    BEGIN
        FOR r IN SELECT table_name, foreign_key, constraint_def
                 FROM ids_with_types_fks
        LOOP
            EXECUTE ''ALTER TABLE '' || r.table_name || '' ADD CONSTRAINT ''|| r.foreign_key || '' '' || r.constraint_def || '';'';
        END LOOP;
    END';

-- remove the temporary table
DROP TABLE ids_with_types_fks;
