CREATE TABLE IF NOT EXISTS public.items(
    id VARCHAR(36) PRIMARY KEY,
    description text NOT NULL,
    price decimal(19, 2) NOT NULL
);