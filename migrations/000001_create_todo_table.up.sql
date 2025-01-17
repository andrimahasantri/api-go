CREATE TABLE todo
(
    id SERIAL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    status SMALLINT NOT NULL,
    created_on TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    updated_on TIMESTAMP(0) WITHOUT TIME ZONE,
    deleted_on  TIMESTAMP(0) WITHOUT TIME ZONE,

    PRIMARY KEY (id)
);

-- export POSTGRESQL_URL='postgres://postgres:andri123@localhost:5432/rest?sslmode=disable'