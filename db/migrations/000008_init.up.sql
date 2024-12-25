CREATE TYPE pet_size AS ENUM ('pequeño', 'mediano', 'grande', 'Otro');

ALTER TABLE refugee 
ALTER COLUMN size TYPE pet_size USING size::text::pet_size;