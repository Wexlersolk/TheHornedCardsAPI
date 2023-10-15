CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS thehorned_groups_table (
    "id" uuid  PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "group_name" varchar NOT NULL,
    "group_info" varchar
);