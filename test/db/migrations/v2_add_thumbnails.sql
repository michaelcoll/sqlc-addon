CREATE TABLE thumbnails
(
    hash      TEXT PRIMARY KEY,
    height    INTEGER NOT NULL,
    width     INTEGER NOT NULL,
    thumbnail BLOB,

    CONSTRAINT hash_fk FOREIGN KEY (hash) REFERENCES photos (hash),
    CONSTRAINT thumbnails_unique UNIQUE (hash, height, width)
)
