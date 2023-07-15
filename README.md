# meli_challenge_back

El proyecto fue desarrollado utilizando el lenguaje de programación Golang y basado en una base de datos transaccional de PostgreSQL.

Para ejecutar correctamente el proyecto en un entorno local, es necesario crear una base de datos en PostgreSQL. 

A continuación, se presentan los scripts necesarios para la creación de la base de datos:

-------------------------------------------------- 
--            crear la base de datos            --
--------------------------------------------------

CREATE DATABASE meli
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

-------------------------------------------------- 
-- crear el esquema en el cual vamos a trabajar --
--------------------------------------------------

CREATE SCHEMA challenge
    AUTHORIZATION postgres;

-------------------------------------------------- 
--               crear las tablas               --
--------------------------------------------------

CREATE SEQUENCE challenge.tbl_country_id_seq;

CREATE TABLE challenge.tbl_country
(
    id integer NOT NULL DEFAULT nextval('challenge.tbl_country_id_seq'),
    description character varying(50),
    PRIMARY KEY (id)
);

ALTER SEQUENCE challenge.tbl_country_id_seq
OWNED BY challenge.tbl_country.id;


CREATE SEQUENCE challenge.tbl_type_id_seq;

CREATE TABLE challenge.tbl_type
(
    id integer NOT NULL DEFAULT nextval('challenge.tbl_type_id_seq'),
    description character varying(50),
    PRIMARY KEY (id)
);

ALTER SEQUENCE challenge.tbl_type_id_seq
OWNED BY challenge.tbl_type.id;


CREATE SEQUENCE challenge.tbl_event_id_seq;

CREATE TABLE challenge.tbl_event
(
    id integer NOT NULL DEFAULT nextval('challenge.tbl_event_id_seq'),
    description character varying(200),
    created_at timestamp with time zone,
    type_id integer,
    country_id integer,
    PRIMARY KEY (id),
    FOREIGN KEY (type_id)
        REFERENCES challenge.tbl_type (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    FOREIGN KEY (country_id)
        REFERENCES challenge.tbl_country (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

ALTER SEQUENCE challenge.tbl_event_id_seq
OWNED BY challenge.tbl_event.id;

-------------------------------------------------- 
--           insertar datos iniciales           --
--------------------------------------------------

-- datos de pais

INSERT INTO challenge.tbl_country(description) VALUES ('argentina');
INSERT INTO challenge.tbl_country(description) VALUES ('brasil');
INSERT INTO challenge.tbl_country(description) VALUES ('colombia');

-- datos de tipo
-- de momento solo uno porque no se especificaron otros tipos

INSERT INTO challenge.tbl_type(description) VALUES ('alerta');



-------------------------------------------------- 
--              ejecutar proyecto               --
--------------------------------------------------
Para ejecutar el proyecto de manera local, es importante utilizar la estructura del siguiente comando:

  go run cmd/main.go db_name db_password db_username db_host db_port

Ejemplo de uso: 

  go run cmd/main.go meli mipassword postgres localhost 5432


-------------------------------------------------- 
--                      Nota                    --
--------------------------------------------------
  -  Para asegurarte de tener todas las dependencias necesarias, puedes utilizar el comando go mod download.
     Esto descargará automáticamente todas las dependencias especificadas en el archivo go.mod.
     
  -  En el directorio "postman_collection", encontrarás una colección de Postman que podrás utilizar para probar las diferentes funcionalidades del API. 

      1- get events all

      2 - get all filtered events

      3 - get events by description

      4- get events by countries

      5- get events by type

      6- get top-countries

      7- get countries

      8- get types

      9- set event

     


