CREATE TABLE qandas
(
    id         INT       NOT NULL AUTO_INCREMENT primary key,
    question   TEXT      NOT NULL,
    answered   bool      not null default false,
    answer     text,
    created_at TIMESTAMP NOT NULL
);
