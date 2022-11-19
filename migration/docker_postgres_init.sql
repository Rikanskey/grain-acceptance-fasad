-- CREATE DATABASE grain_acceptance;

create type grain_variety as enum ('Fodder', 'Food');

create type process_type as enum ('Acceptance', 'Shipment');

create table if not exists grain_type
(
    id        serial
        primary key,
    name      text not null,
    gost_name text not null,
    sort      grain_variety
);

create table if not exists impurity_parameter
(
    id   serial
        primary key,
    name varchar(256) not null
);

create table if not exists company
(
    id       serial
        primary key,
    name     text    not null,
    postcode integer not null,
    state    text    not null,
    city     text    not null,
    street   text    not null,
    house    text    not null,
    office   text
);

create table if not exists part
(
    id   serial
        primary key,
    name char(256) not null
);

create table if not exists storage
(
    id           serial
        primary key,
    capacity     real not null,
    filled       real not null,
    grain_name   text,
    sort         grain_variety,
    grain_class  integer,
    grain_number integer
);

create table if not exists grain_card
(
    id           serial
        primary key,
    storage_id   integer
        references storage
            on update cascade on delete restrict,
    customer_id  integer
        references company
            on update cascade on delete restrict,
    sender_id    integer
        references company
            on update cascade on delete restrict,
    date         date not null,
    process_type process_type
);

create table if not exists consignment
(
    id                        serial
        primary key,
    grain_card_id             integer not null
        references grain_card
            on update cascade on delete restrict,
    model_transport           text,
    number_transport          text,
    first_name_driver         text,
    second_name_driver        text,
    middle_name_driver        text,
    trailer_number            text,
    grain_type_id             integer
        references grain_type
            on update cascade on delete restrict,
    moisture                  real    default 0.0,
    origin                    text,
    nature                    integer default 0,
    fall_number               real    default 0.0,
    color                     text,
    smell                     text,
    small_grains_percent      real    default 0.0,
    vitreous                  real    default 0.0,
    gluten                    real    default 0.0,
    gost_culture              real    default 0.0,
    filminess                 real    default 0.0,
    contamitation             real    default 0.0,
    type_subtype_composition  real    default 0.0,
    core                      real    default 0.0,
    bug                       real    default 0.0,
    additional_notes          text,
    accepted_gross_weight     real    default 0.0,
    accepted_transport_weight real    default 0.0,
    sent_gross_weight         real    default 0.0,
    sent_transport_weight     real    default 0.0
);

create table if not exists grain_classification
(
    id                 serial
        primary key,
    moisture           real    not null,
    grain_impurity     real    not null,
    volatiles_impurity real    not null,
    type               integer not null,
    subtype            integer not null,
    grain_type_id      integer
        references grain_type
            on update cascade on delete restrict
);

create table if not exists part_of_impurity_parameter
(
    id                    serial
        primary key,
    impurity_parameter_id integer
        references impurity_parameter
            on update cascade on delete restrict,
    part_id               integer
        references part
            on update cascade on delete restrict,
    value                 integer default 0.0 not null,
    consignment_id        integer
        references consignment
            on update cascade on delete restrict
);
