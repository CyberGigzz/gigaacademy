BEGIN;

INSERT INTO courses (name, description, created_at, updated_at)
SELECT 
    'Course ' || n AS name,
    'Description for Course ' || n AS description,
    CURRENT_TIMESTAMP AS created_at,
    CURRENT_TIMESTAMP AS updated_at
FROM generate_series(1, 20) AS n;

COMMIT;