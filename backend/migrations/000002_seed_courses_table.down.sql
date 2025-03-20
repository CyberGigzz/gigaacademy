BEGIN;

DELETE FROM courses WHERE name LIKE 'Course [0-9]%';

COMMIT;