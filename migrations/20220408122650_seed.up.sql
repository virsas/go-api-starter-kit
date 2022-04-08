INSERT INTO accounts (id, name, updated_at, created_at) 
SELECT 1 as id, 'test' as name, NOW() as updated_at, NOW() as created_at
WHERE NOT EXISTS (
    SELECT id FROM accounts WHERE name='admin' AND id=1 
);
INSERT INTO users (id, email, account_id, updated_at, created_at) 
SELECT 1 as id, 'info@examples.org' as email, 1 as account_id, NOW() as updated_at, NOW() as created_at
WHERE NOT EXISTS (
    SELECT id FROM users WHERE email='info@examples.org' AND id=1 
);

INSERT INTO examples (id, name, account_id, user_id, updated_at, created_at) 
SELECT 1 as id, 'test' as name, 1 as account_id, 1 as user_id, NOW() as updated_at, NOW() as created_at
WHERE NOT EXISTS (
    SELECT id FROM examples WHERE name = 'test' AND id=1 
);

INSERT INTO examples (id, name, account_id, user_id, updated_at, created_at) 
SELECT 2 as id, 'test2' as name, 1 as account_id, 1 as user_id, NOW() as updated_at, NOW() as created_at
WHERE NOT EXISTS (
    SELECT id FROM examples WHERE name = 'test2' AND id=2
);