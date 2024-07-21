-- modify "group_membership_applications" table
ALTER TABLE "group_membership_applications" ADD COLUMN "group_applications" character varying NOT NULL, ADD COLUMN "user_group_membership_applications" character varying NOT NULL, ADD CONSTRAINT "group_membership_applications_groups_applications" FOREIGN KEY ("group_applications") REFERENCES "groups" ("id") ON DELETE NO ACTION, ADD CONSTRAINT "group_membership_applications__45cb98961f18fc117b600d1fca4af532" FOREIGN KEY ("user_group_membership_applications") REFERENCES "users" ("id") ON DELETE NO ACTION;