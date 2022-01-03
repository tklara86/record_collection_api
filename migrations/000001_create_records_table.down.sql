ALTER TABLE record_artists DROP CONSTRAINT fk_artists;
ALTER TABLE record_artists DROP CONSTRAINT fk_records;

ALTER TABLE record_genres DROP CONSTRAINT fk_records;
ALTER TABLE record_genres DROP CONSTRAINT fk_genres;

DROP TABLE IF EXISTS artists;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS records;
DROP TABLE IF EXISTS genres;
DROP TABLE IF EXISTS record_artists;
DROP TABLE IF EXISTS record_genres;



