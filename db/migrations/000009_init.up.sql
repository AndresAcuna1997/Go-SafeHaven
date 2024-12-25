-- Crear el nuevo tipo enumerado con valores en capital case
CREATE TYPE new_pet_size AS ENUM ('Pequeño', 'Mediano', 'Grande', 'Otro');

-- Cambiar la columna size para que use el nuevo tipo enumerado
ALTER TABLE refugee 
ALTER COLUMN size TYPE new_pet_size USING size::text::new_pet_size;

-- Eliminar el tipo enumerado antiguo
DROP TYPE pet_size;

-- Renombrar el nuevo tipo enumerado al nombre original
ALTER TYPE new_pet_size RENAME TO pet_size;