-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE "users" (
  "id"              SERIAL PRIMARY KEY       NOT NULL,
  "google_id"       TEXT                     NOT NULL,
  "email"           TEXT                     NOT NULL,
  "role"            INT                      NOT NULL DEFAULT 1,
  "created_at"      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  "display_name"    TEXT,
  "profile_picture" TEXT,
  "token"           TEXT
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE "users";
