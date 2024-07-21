-- create "groups" table
CREATE TABLE "groups" ("id" uuid NOT NULL, "name" character varying NOT NULL, "description" character varying NOT NULL DEFAULT '', "logo_url" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "groups_name_key" to table: "groups"
CREATE UNIQUE INDEX "groups_name_key" ON "groups" ("name");
-- create "group_memberships" table
CREATE TABLE "group_memberships" ("id" uuid NOT NULL, "role" character varying NOT NULL, "group_members" uuid NOT NULL, "user_group_memberships" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "group_memberships_groups_members" FOREIGN KEY ("group_members") REFERENCES "groups" ("id") ON DELETE NO ACTION, CONSTRAINT "group_memberships_users_group_memberships" FOREIGN KEY ("user_group_memberships") REFERENCES "users" ("id") ON DELETE NO ACTION);
-- create "group_settings" table
CREATE TABLE "group_settings" ("id" uuid NOT NULL, "visibility" character varying NOT NULL DEFAULT 'PUBLIC', "join_policy" character varying NOT NULL DEFAULT 'OPEN', "minimum_role_to_invite" character varying NULL, "group_settings" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "group_settings_groups_settings" FOREIGN KEY ("group_settings") REFERENCES "groups" ("id") ON DELETE SET NULL);
-- create index "group_settings_group_settings_key" to table: "group_settings"
CREATE UNIQUE INDEX "group_settings_group_settings_key" ON "group_settings" ("group_settings");
