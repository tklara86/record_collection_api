CREATE TABLE "records" (
                           "record_id" SERIAL PRIMARY KEY,
                           "title" varchar NOT NULL,
                           "release" varchar NOT NULL,
                           "cover" varchar NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now()),
                           "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "labels" (
                          "id" SERIAL PRIMARY KEY,
                          "name" varchar NOT NULL,
                          "created_at" timestamptz NOT NULL DEFAULT (now()),
                          "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "label_catalogue_number" (
                                          "id" SERIAL PRIMARY KEY,
                                          "label_id" int,
                                          "catalogue_number" varchar NOT NULL,
                                          "created_at" timestamptz NOT NULL DEFAULT (now()),
                                          "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "record_labels" (
                                 "id" SERIAL PRIMARY KEY,
                                 "record_id" int,
                                 "label_id" int,
                                 "created_at" timestamptz NOT NULL DEFAULT (now()),
                                 "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "genres" (
                          "genre_id" SERIAL PRIMARY KEY,
                          "genre_name" varchar,
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

CREATE TABLE "artists" (
                           "artist_id" SERIAL PRIMARY KEY,
                           "name" varchar,
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

CREATE TABLE "tracklist" (
                             "id" SERIAL PRIMARY KEY,
                             "position" varchar NOT NULL,
                             "title" varchar NOT NULL,
                             "duration" varchar NOT NULL,
                             "created_at" timestamptz NOT NULL DEFAULT (now()),
                             "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "record_tracklists" (
                                     "id" SERIAL PRIMARY KEY,
                                     "record_id" int,
                                     "tracklist_id" int,
                                     "created_at" timestamptz NOT NULL DEFAULT (now()),
                                     "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "record_tracklists_artists" (
                                             "id" SERIAL PRIMARY KEY,
                                             "record_tracklist_id" int,
                                             "artist_id" int,
                                             "created_at" timestamptz NOT NULL DEFAULT (now()),
                                             "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "record_images" (
                                 "id" SERIAL PRIMARY KEY,
                                 "record_id" int,
                                 "image" varchar NOT NULL
);

ALTER TABLE "label_catalogue_number" ADD FOREIGN KEY ("label_id") REFERENCES "labels" ("id");

ALTER TABLE "record_labels" ADD FOREIGN KEY ("record_id") REFERENCES "records" ("record_id");

ALTER TABLE "record_labels" ADD FOREIGN KEY ("label_id") REFERENCES "labels" ("id");

ALTER TABLE "record_genres" ADD FOREIGN KEY ("record_id") REFERENCES "records" ("record_id");

ALTER TABLE "record_genres" ADD FOREIGN KEY ("genre_id") REFERENCES "genres" ("genre_id");

ALTER TABLE "record_artists" ADD FOREIGN KEY ("artist_id") REFERENCES "artists" ("artist_id");

ALTER TABLE "record_artists" ADD FOREIGN KEY ("record_id") REFERENCES "records" ("record_id");

ALTER TABLE "record_tracklists" ADD FOREIGN KEY ("record_id") REFERENCES "records" ("record_id");

ALTER TABLE "record_tracklists" ADD FOREIGN KEY ("tracklist_id") REFERENCES "tracklist" ("id");

ALTER TABLE "record_tracklists_artists" ADD FOREIGN KEY ("record_tracklist_id") REFERENCES "record_tracklists" ("id");

ALTER TABLE "record_tracklists_artists" ADD FOREIGN KEY ("artist_id") REFERENCES "artists" ("artist_id");

ALTER TABLE "record_images" ADD FOREIGN KEY ("record_id") REFERENCES "records" ("record_id");

CREATE INDEX ON "records" ("title");

CREATE INDEX ON "labels" ("name");

CREATE INDEX ON "label_catalogue_number" ("catalogue_number");

CREATE INDEX ON "record_labels" ("record_id");

CREATE INDEX ON "record_labels" ("label_id");

CREATE INDEX ON "record_labels" ("record_id", "label_id");

CREATE INDEX ON "genres" ("genre_name");

CREATE INDEX ON "record_genres" ("record_id");

CREATE INDEX ON "record_genres" ("genre_id");

CREATE INDEX ON "record_genres" ("record_id", "genre_id");

CREATE INDEX ON "artists" ("name");

CREATE INDEX ON "record_artists" ("artist_id");

CREATE INDEX ON "record_artists" ("record_id");

CREATE INDEX ON "record_artists" ("artist_id", "record_id");

CREATE INDEX ON "tracklist" ("title");

CREATE INDEX ON "record_tracklists" ("record_id");

CREATE INDEX ON "record_tracklists" ("tracklist_id");

CREATE INDEX ON "record_tracklists" ("record_id", "tracklist_id");

CREATE INDEX ON "record_tracklists_artists" ("record_tracklist_id");

CREATE INDEX ON "record_tracklists_artists" ("artist_id");

CREATE INDEX ON "record_tracklists_artists" ("record_tracklist_id", "artist_id");

CREATE INDEX ON "record_images" ("image");
