ALTER TABLE refugee 
ALTER COLUMN size TYPE TEXT USING size::text;

DROP TYPE pet_size;