create table todos (
  id serial primary key,
  title varchar not null,
  status integer not null,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now()
);
