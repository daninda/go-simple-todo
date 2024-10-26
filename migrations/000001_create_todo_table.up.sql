create table todo (
  id serial primary key,
  title varchar(255) not null,
  description text,
  completed boolean not null
);