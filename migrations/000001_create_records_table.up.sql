

CREATE TABLE "records" (
                           "record_id" SERIAL PRIMARY KEY,
                           "title" varchar NOT NULL,
                           "label" varchar NOT NULL,
                           "year" varchar NOT NULL,
                           "cover" varchar NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now()),
                           "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "genres" (
                          "genre_id" SERIAL PRIMARY KEY,
                          "genre_name" varchar,
                          "created_at" timestamptz NOT NULL DEFAULT (now()),
                          "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "artists" (
                           "artist_id" SERIAL PRIMARY KEY,
                           "name" varchar,
                           "created_at" timestamptz NOT NULL DEFAULT (now()),
                           "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "record_genres" (
                                 "id" SERIAL PRIMARY KEY,
                                 "record_id" int,
                                 "genre_id" int,
                                 "created_at" timestamptz NOT NULL DEFAULT (now()),
                                 "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "record_artists" (
                                  "id" SERIAL PRIMARY KEY,
                                  "artist_id" int,
                                  "record_id" int,
                                  "created_at" timestamptz NOT NULL DEFAULT (now()),
                                  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "record_genres" ADD FOREIGN KEY ("record_id") REFERENCES "records" ("record_id");

ALTER TABLE "record_genres" ADD FOREIGN KEY ("genre_id") REFERENCES "genres" ("genre_id");

ALTER TABLE "record_artists" ADD FOREIGN KEY ("artist_id") REFERENCES "artists" ("artist_id");

ALTER TABLE "record_artists" ADD FOREIGN KEY ("record_id") REFERENCES "records" ("record_id");

CREATE INDEX ON "records" ("title");

CREATE INDEX ON "genres" ("genre_name");

CREATE INDEX ON "artists" ("name");

CREATE INDEX ON "record_genres" ("record_id");

CREATE INDEX ON "record_genres" ("genre_id");

CREATE INDEX ON "record_genres" ("record_id", "genre_id");

CREATE INDEX ON "record_artists" ("artist_id");

CREATE INDEX ON "record_artists" ("record_id");

CREATE INDEX ON "record_artists" ("artist_id", "record_id");
