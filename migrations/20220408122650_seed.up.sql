INSERT INTO accounts (name, updated_at, created_at) 
SELECT 'test' as name, NOW() as updated_at, NOW() as created_at
WHERE NOT EXISTS (
    SELECT id FROM accounts WHERE name='admin' AND id=1 
);
INSERT INTO users (email, account_id, updated_at, created_at) 
SELECT 'info@examples.org' as email, 1 as account_id, NOW() as updated_at, NOW() as created_at
WHERE NOT EXISTS (
    SELECT id FROM users WHERE email='info@examples.org' AND id=1 
);

INSERT INTO examples (name, account_id, user_id, updated_at, created_at) 
SELECT 'test' as name, 1 as account_id, 1 as user_id, NOW() as updated_at, NOW() as created_at
WHERE NOT EXISTS (
    SELECT id FROM examples WHERE name = 'test' AND id=1 
);

INSERT INTO examples (name, account_id, user_id, updated_at, created_at) 
SELECT 'test2' as name, 1 as account_id, 1 as user_id, NOW() as updated_at, NOW() as created_at
WHERE NOT EXISTS (
    SELECT id FROM examples WHERE name = 'test2' AND id=2
);
