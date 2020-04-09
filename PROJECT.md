# projetc

export GO111MODULE=on

go mod -vendor nor go mod -sync

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TYPE super_type AS ENUM (
	'Heroes',
	'Villain',
	'pixel');