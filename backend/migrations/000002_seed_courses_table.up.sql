BEGIN;

INSERT INTO courses (name, description, code, credits, start_date, end_date, max_capacity, status, created_at, updated_at)
SELECT 
    'Course ' || n AS name,
    'Description for Course ' || n AS description,
    'C' || LPAD(n::text, 3, '0') AS code, -- e.g., C001, C002, ..., C020
    (RANDOM() * 3 + 1)::int AS credits, -- Random credits between 1 and 4
    CURRENT_TIMESTAMP + (n || ' days')::interval AS start_date, -- Start date increments by n days
    CURRENT_TIMESTAMP + (n + 90 || ' days')::interval AS end_date, -- End date is 90 days after start
    (RANDOM() * 20 + 10)::int AS max_capacity, -- Random capacity between 10 and 30
    CASE 
        WHEN n % 3 = 0 THEN 'inactive'
        WHEN n % 3 = 1 THEN 'active'
        ELSE 'archived'
    END AS status,
    CURRENT_TIMESTAMP AS created_at,
    CURRENT_TIMESTAMP AS updated_at
FROM generate_series(1, 20) AS n;

COMMIT;