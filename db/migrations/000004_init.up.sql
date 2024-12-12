CREATE TYPE pet_type AS ENUM ('Perro', 'Gato', 'Ave', 'Reptil', 'Otro');

ALTER TABLE refugee 
ALTER COLUMN type TYPE pet_type USING type::text::pet_type;

CREATE TABLE IF NOT EXISTS cities (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

ALTER TABLE shelter
ADD COLUMN city INT REFERENCES cities(id);