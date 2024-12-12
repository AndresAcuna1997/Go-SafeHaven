CREATE TABLE organization (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  deparments TEXT[],
  createdAt DATE NOT NULL
);

CREATE TABLE shelter (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  address TEXT NOT NULL,
  refugeesCount INT,
  contactPhone TEXT,
  contactEmail TEXT,
  createdAt DATE NOT NULL
);

CREATE TABLE refugee (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  type TEXT NOT NULL,
  size TEXT,
  age INT,
  additionalInfo JSON,
  pictures JSON,
  createdAt DATE NOT NULL
);