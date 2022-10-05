create table employees
(
    id                character varying(64)  not null primary key default uuid_generate_v4(),
    first_name        character varying(100) not null,
    last_name         character varying(100),
    salary            numeric(8, 2)          not null             default 0,
    commission_pct    numeric(2, 2),
    hire_date         date                   not null             default now(),
    created_date      timestamp              not null             default now(),
    created_by        character varying(100) not null,
    last_updated_date timestamp,
    last_updated_by   character varying(100)
);
