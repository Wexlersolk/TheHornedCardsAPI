CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS thehorned_cards_table (
    "id" uuid  PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "group_name" varchar NOT NULL,
    "card_hint" varchar,
    "display_word" varchar NOT NULL,
    "hidden_word" varchar NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);


CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS thehorned_groups_table (
    "id" uuid  PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "group_name" varchar NOT NULL,
    "group_info" varchar
);