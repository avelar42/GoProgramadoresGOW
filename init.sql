CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS programadores (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  apelido TEXT UNIQUE NOT NULL,
  nome TEXT NOT NULL,
  nascimento DATE NOT NULL,
  stack TEXT[] NOT NULL
);