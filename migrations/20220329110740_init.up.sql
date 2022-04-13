CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL,
    name character varying(255) NOT NULL,
    created_at timestamp with time zone DEFAULT NOW(),
    updated_at timestamp with time zone,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL,
    email character varying(255) NOT NULL UNIQUE,
    locked boolean NOT NULL DEFAULT FALSE,
    account_id integer NOT NULL,
    created_at timestamp with time zone DEFAULT NOW(),
    updated_at timestamp with time zone,
    PRIMARY KEY (id),
    FOREIGN KEY (account_id) REFERENCES accounts(id)
);

CREATE TABLE IF NOT EXISTS examples (
    id SERIAL,
    name character varying(255) NOT NULL,
    user_id integer NOT NULL,
    account_id integer NOT NULL,
    created_at timestamp with time zone DEFAULT NOW(),
    updated_at timestamp with time zone,
    PRIMARY KEY (id),
    FOREIGN KEY (account_id) REFERENCES accounts(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);