CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS thehorned_cards_table (
    "card_id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "group_id" int NOT NULL,
    "card_hint" varchar,
    "display_word" varchar NOT NULL,
    "hidden_word" varchar NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    FOREIGN KEY ("group_id") REFERENCES thehorned_groups_table("group_id")
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS thehorned_groups_table (
    "group_id" int PRIMARY KEY NOT NULL,
    "group_name" varchar NOT NULL
);
