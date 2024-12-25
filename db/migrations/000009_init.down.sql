CREATE TYPE old_pet_size AS ENUM ('pequeño', 'mediano', 'grande', 'Otro');

ALTER TABLE refugee 
ALTER COLUMN size TYPE old_pet_size USING size::text::old_pet_size;

DROP TYPE pet_size;

ALTER TYPE old_pet_size RENAME TO pet_size;