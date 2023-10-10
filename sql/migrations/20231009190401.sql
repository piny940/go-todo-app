create table if not exists users (
  id serial primary key,
  email varchar not null unique,
  password varchar not null,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now()
);

create table if not exists todos (
  id serial primary key,
  title varchar not null,
  status integer not null,
  user_id integer not null references users(id),
  created_at timestamp not null default now(),
  updated_at timestamp not null default now()
);

create index if not exists todos_user_id_index on todos(user_id);
