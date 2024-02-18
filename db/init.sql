drop table if exists organizations cascade;
drop table if exists projects cascade;

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
    organization_id integer not null,
    toolkit_id      integer not null,
    primary key (id)
);

create table accounts
(
    id   serial not null,
    uuid text   not null,
    primary key (id)
);

create table account_organization
(
    account_id      integer not null,
    organization_id integer not null
);

create table toolkits
(
    id             serial not null,
    name           text   not null,
    repository_url text   not null,
    primary key (id)
);

insert into organizations(name, admin_user_id)
values ('The Database Organization', '1');

insert into accounts(uuid)
values ('1234');

insert into account_organization(account_id, organization_id)
values (1, 1);

insert into toolkits(name, repository_url)
values ('The Database Toolkit', 'https://github.com/packlify/toolkit.git');

insert into projects(name, organization_id, toolkit_id)
values ('The Database Project', 1, 1);