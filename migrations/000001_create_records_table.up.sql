CREATE TABLE IF NOT EXISTS records (
                                       record_id BIGSERIAL PRIMARY KEY,
                                       title text NOT NULL,
                                       label text NOT NULL,
                                       year integer NOT NULL,
                                       cover text NOT NULL,
                                       created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                                       updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS genres (
                                      genre_id BIGSERIAL PRIMARY KEY,
                                      genre_name text NOT NULL,
                                      created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                                      updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);


CREATE TABLE IF NOT EXISTS artists (
                                       artist_id BIGSERIAL PRIMARY KEY,
                                       first_name text NOT NULL,
                                       last_name text NOT NULL,
                                       created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                                       updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);


CREATE TABLE IF NOT EXISTS record_genres (
                                             id BIGSERIAL PRIMARY KEY,
                                             record_id INTEGER,
                                             genre_id INTEGER,
                                             created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                                             updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                                             CONSTRAINT fk_records
                                             FOREIGN KEY (record_id) REFERENCES records(record_id) ON DELETE CASCADE,
                                             CONSTRAINT fk_genres
                                             FOREIGN KEY (genre_id) REFERENCES genres(genre_id) ON DELETE CASCADE

);


CREATE TABLE IF NOT EXISTS record_artists (
                                              id BIGSERIAL PRIMARY KEY,
                                              artist_id INTEGER,
                                              record_id INTEGER,
                                              created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                                              updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                                              CONSTRAINT fk_artists
                                              FOREIGN KEY (artist_id) REFERENCES artists(artist_id) ON DELETE CASCADE,
                                              CONSTRAINT fk_records
                                              FOREIGN KEY (record_id) REFERENCES records(record_id) ON DELETE CASCADE

);



