DROP TABLE IF EXISTS organization_departments;

ALTER TABLE organization ADD COLUMN departments TEXT[];

DROP TABLE IF EXISTS departments;