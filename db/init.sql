drop table if exists todos cascade;

create table organizations
(
    id            serial not null,
    name          text   not null,
    admin_user_id text   not null,
    primary key (id)
);

create table projects
(
    id              serial  not null,
    name            text    not null,
    organization_id integer not null
);

insert into organizations(name, admin_user_id)
values ('The Database Organization', '0000');

insert into projects(name, organization_id)
values ('The Database Project', 1);