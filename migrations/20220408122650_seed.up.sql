INSERT INTO accounts (name, updated_at, created_at) 
SELECT 'test' as name, NOW() as updated_at, NOW() as created_at
WHERE NOT EXISTS (
    SELECT id FROM accounts WHERE name='admin' 
);
INSERT INTO users (email, account_id, updated_at, created_at) 
SELECT 'info@examples.org' as email, 1 as account_id, NOW() as updated_at, NOW() as created_at
WHERE NOT EXISTS (
    SELECT id FROM users WHERE email='info@examples.org' 
);

INSERT INTO users (email, account_id, globaladmin, updated_at, created_at) 
SELECT 'admin@examples.org' as email, 1 as account_id, true as globaladmin, NOW() as updated_at, NOW() as created_at
WHERE NOT EXISTS (
    SELECT id FROM users WHERE email='admin@examples.org' 
);

INSERT INTO examples (name, account_id, user_id, updated_at, created_at) 
SELECT 'test' as name, 1 as account_id, 1 as user_id, NOW() as updated_at, NOW() as created_at
WHERE NOT EXISTS (
    SELECT id FROM examples WHERE name = 'test' 
);

INSERT INTO examples (name, account_id, user_id, updated_at, created_at) 
SELECT 'test2' as name, 1 as account_id, 1 as user_id, NOW() as updated_at, NOW() as created_at
WHERE NOT EXISTS (
    SELECT id FROM examples WHERE name = 'test2'
);
