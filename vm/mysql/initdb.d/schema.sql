CREATE TABLE qandas
(
    id         INT       NOT NULL AUTO_INCREMENT primary key,
    question   TEXT      NOT NULL,
    answer     text,
    created_at TIMESTAMP NOT NULL
);
