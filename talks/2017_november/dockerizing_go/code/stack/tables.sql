CREATE SCHEMA IF NOT EXISTS viennago AUTHORIZATION gopher;

CREATE TABLE IF NOT EXISTS viennago.Texts (
    txt TEXT NOT NULL
);

INSERT INTO viennago.Texts VALUES('Hello, Gophers!');
