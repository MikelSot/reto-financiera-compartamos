-- Documentacion grafica
-- https://dbdiagram.io/d/reto-financiera-6619eed803593b6b61e2a3ea

CREATE TYPE "gender" AS ENUM (
  'M',
  'F'
);

CREATE TABLE "customers"
(
    "id"         serial PRIMARY KEY,
    "first_name" varchar(50)       NOT NULL,
    "last_name"  varchar(50)       NOT NULL,
    "dni"        varchar(8) UNIQUE NOT NULL,
    "birth_date" date,
    "gender"     gender,
    "password"   varchar(250),
    "email"      varchar(150) UNIQUE,
    "is_staff"   bool      DEFAULT false,
    "picture"    varchar(250),
    "nickname"   varchar(100),
    "created_at" timestamp DEFAULT now(),
    "updated_at" timestamp,
    "deleted_at" timestamp
);

CREATE TABLE "cities"
(
    "id"          serial PRIMARY KEY,
    "name"        varchar(100) NOT NULL,
    "postal_Code" varchar(20),
    "created_at"  timestamp DEFAULT now(),
    "updated_at"  timestamp,
    "deleted_at"  timestamp
);

CREATE TABLE "city_customers"
(
    "id"          serial PRIMARY KEY,
    "customer_id" integer NOT NULL,
    "city_id"     integer NOT NULL,
    "created_at"  timestamp DEFAULT now(),
    "updated_at"  timestamp,
    "deleted_at"  timestamp
);

COMMENT
ON TABLE "customers" IS '
Tabla de clientes
Se usara tambien como usuario
';

COMMENT
ON COLUMN "customers"."first_name" IS 'Nombre del cliente';

COMMENT
ON COLUMN "customers"."last_name" IS 'Apellido del cliente';

COMMENT
ON COLUMN "customers"."dni" IS 'DNI';

COMMENT
ON COLUMN "customers"."birth_date" IS 'Fecha de nacimiento';

COMMENT
ON COLUMN "customers"."gender" IS 'Sexo';

COMMENT
ON COLUMN "customers"."email" IS 'Correo';

COMMENT
ON COLUMN "customers"."is_staff" IS 'flag para validar si es usuario admin';

COMMENT
ON COLUMN "customers"."picture" IS 'Imagen del cliente';

COMMENT
ON COLUMN "customers"."created_at" IS 'Fecha de creacion del registro';

COMMENT
ON COLUMN "customers"."updated_at" IS 'Fecha de actualizacion del registro';

COMMENT
ON COLUMN "customers"."deleted_at" IS 'Fecha de eliminacion del registro';

COMMENT
ON TABLE "cities" IS 'Tabla de ciudades';

COMMENT
ON COLUMN "cities"."name" IS 'Nombre de la Ciudad';

COMMENT
ON COLUMN "cities"."postal_Code" IS 'Codigo postal';

COMMENT
ON COLUMN "cities"."created_at" IS 'Fecha de creacion del registro';

COMMENT
ON COLUMN "cities"."updated_at" IS 'Fecha de actualizacion del registro';

COMMENT
ON COLUMN "cities"."deleted_at" IS 'Fecha de eliminacion del registro';

COMMENT
ON TABLE "city_customers" IS 'Tabla de relacion muchos a muchos de customers y cities';

COMMENT
ON COLUMN "city_customers"."customer_id" IS 'id del cliente';

COMMENT
ON COLUMN "city_customers"."city_id" IS 'id de la ciudad';

COMMENT
ON COLUMN "city_customers"."created_at" IS 'Fecha de creacion del registro';

COMMENT
ON COLUMN "city_customers"."updated_at" IS 'Fecha de actualizacion del registro';

COMMENT
ON COLUMN "city_customers"."deleted_at" IS 'Fecha de eliminacion del registro';

ALTER TABLE "city_customers"
    ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "city_customers"
    ADD FOREIGN KEY ("city_id") REFERENCES "cities" ("id");


-- Insertar datos de prueba

INSERT INTO customers ("first_name", "last_name", "dni", "birth_date", "gender", "is_staff")
VALUES ('Juan', 'Pérez', '12345678', '1990-05-15', 'M', true),
       ('María', 'Gómez', '23456789', '1985-09-20', 'F', false),
       ('Pedro', 'Rodríguez', '34567890', '1978-03-10', 'M', false),
       ('Laura', 'López', '45678901', '1995-11-25', 'F', false),
       ('Carlos', 'Martínez', '56789012', '1980-07-18', 'M', false),
       ('Ana', 'Fernández', '67890123', '1992-02-08', 'F', false),
       ('Javier', 'Sánchez', '78901234', '1973-06-30', 'M', false),
       ('Elena', 'Díaz', '89012345', '1987-08-12', 'F', false),
       ('Sara', 'Ruiz', '90123456', '1982-04-05', 'F', false),
       ('David', 'García', '01234567', '1975-10-29', 'M', false);

INSERT INTO cities ("name", "postal_Code")
VALUES ('Ciudad de México', '01000'),
       ('Buenos Aires', 'B1228'),
       ('Madrid', '28001'),
       ('Lima', '00051'),
       ('Santiago', '8320000'),
       ('Bogotá', '110221'),
       ('Caracas', '1060'),
       ('Quito', '170135'),
       ('Montevideo', '11000'),
       ('Asunción', '1100');

INSERT INTO city_customers ("customer_id", "city_id")
VALUES (1, 1),
       (2, 2),
       (3, 3),
       (4, 4),
       (5, 5),
       (6, 6),
       (7, 7),
       (8, 8),
       (9, 9),
       (10, 10);


