drop table if exists roles cascade;
drop table if exists organizations cascade;
drop table if exists projects cascade;
drop table if exists accounts cascade;
drop table if exists account_organization cascade;
drop table if exists account_project cascade;
drop table if exists toolkits cascade;

create table roles
(
    id   serial not null,
    name text   not null,
    primary key (id)
);

create table organizations
(
    id   serial not null,
    name text   not null,
    primary key (id)
);

create table projects
(
    id              serial  not null,
    name            text    not null,
    organization_id integer not null,
    toolkit_id      integer not null,
    repository_url  text,
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
    id              serial  not null,
    account_id      integer not null,
    organization_id integer not null,
    role_id         integer not null,
    primary key (id)
);

create table account_project
(
    id                      serial  not null,
    account_organization_id integer not null,
    project_id              integer not null,
    role_id                 integer not null,
    primary key (id)
);

create table toolkits
(
    id             serial not null,
    name           text   not null,
    repository_url text   not null,
    image_url      text   not null,
    primary key (id)
);

insert into roles(name)
values ('admin'),
       ('member'),
       ('guest');

insert into organizations(name)
values ('The Database Organization');

insert into accounts(uuid)
values ('1234');

insert into account_organization(account_id, organization_id, role_id)
values (1, 1, 1);

insert into account_project(account_organization_id, project_id, role_id)
values (1, 1, 1);

insert into toolkits(name, repository_url,
                     image_url)
values ('The Database Toolkit', 'https://github.com/packlify/toolkit.git',
        'https://avatars.githubusercontent.com/u/152694722?s=200&v=4');

insert into projects(name, organization_id, toolkit_id, repository_url)
values ('The Database Project', 1, 1, 'https://github.com/packlify/project.git');