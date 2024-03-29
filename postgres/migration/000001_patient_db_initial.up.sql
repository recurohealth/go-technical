CREATE TYPE "enum_gender" AS ENUM (
    'M',
    'F',
    'U'
);

CREATE TABLE "patient" (
    "id" UUID PRIMARY KEY,
    "client_id" VARCHAR(255),
    "first_name" VARCHAR(64) NOT NULL,
    "last_name" VARCHAR(64) NOT NULL,
    "gender" enum_gender NOT NULL DEFAULT 'U',
    "date_of_birth" DATE NOT NULL,
    "phone_number" VARCHAR(15),
    "email" VARCHAR(64),
    "date_created" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "date_modified" TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE "address" (
    "id" uuid UNIQUE PRIMARY KEY,
    "city" varchar(64) NOT NULL,
    "address_line_one" varchar(255) NOT NULL,
    "address_line_two" varchar(255),
    "state" varchar(2) NOT NULL,
    "zip" varchar(10) NOT NULL,
    "country" varchar(64) NOT NULL,
    "date_created" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "date_modified" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "patient_id" UUID NOT NULL REFERENCES patient (id) ON DELETE CASCADE
);
