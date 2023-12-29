CREATE TABLE users
(
    id          bigserial PRIMARY KEY ,
    created_at  timestamp with time zone default now(),
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone,
    name        text,
    email       text,
    phone       text,
    password    text,
    is_active   boolean                  default true,
    avatar_path text
);