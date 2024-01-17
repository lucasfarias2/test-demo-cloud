drop table if exists todos cascade;

create table accounts
(
    id            serial  not null,
    name          text    not null,
    admin_user_id integer not null,
    primary key (id)
);

create table users
(
    id   serial not null,
    name text   not null,
    email text not null,
    primary key (id)
);

insert into users(name, email) values ('user1', 'lucas@asd.com');