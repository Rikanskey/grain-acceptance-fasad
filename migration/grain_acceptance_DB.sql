create database grain_acceptance;

create type grain_variety as enum ('Fodder', 'Food');

create type process_type as enum ('Acceptance', 'Shipment');

create table grain_type
(
    id        serial
        constraint grain_type_pkey
            primary key,
    name      text not null,
    gost_name text not null,
    sort      grain_variety
);

create table impurity_parameter
(
    id   serial
        constraint impurity_parameter_pkey
            primary key,
    name varchar(256) not null
);

create table company
(
    id       serial
        constraint company_pkey
            primary key,
    name     text    not null,
    postcode integer not null,
    state    text    not null,
    city     text    not null,
    street   text    not null,
    house    text    not null,
    office   text
);

create table Part
(
    id   serial
        constraint part_pkey
            primary key,
    name char(256) not null
);

create table storage
(
    id           serial
        constraint storage_pkey
            primary key,
    capacity     real not null,
    filled       real not null,
    grain_name   text,
    sort         grain_variety,
    grain_class  integer,
    grain_number integer
);


create table grain_card
(
    id           serial
        constraint grain_card_pkey
            primary key,
    storage_id   integer
        constraint grain_card_storage_id_fkey
            references storage
            on update cascade on delete restrict,
    customer_id  integer
        constraint grain_card_customer_id_fkey
            references company
            on update cascade on delete restrict,
    sender_id    integer
        constraint grain_card_sender_id_fkey
            references company
            on update cascade on delete restrict,
    date         date not null,
    process_type process_type
);

create table consignment
(
    id                        serial
        constraint consignment_pkey
            primary key,
    grain_card_id             integer not null
        constraint consignment_grain_card_id_fkey
            references grain_card
            on update cascade on delete restrict,
    model_transport           text,
    number_transport          text,
    first_name_driver         text,
    second_name_driver        text,
    middle_name_driver        text,
    trailer_number            text,
    grain_type_id             integer
        constraint consignment_grain_type_id_fkey
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

create table impurity_in_consignment_grain
(
    impurity_parameter_id integer not null
        constraint impurity_in_consignment_grain_impurity_parameter_id_fkey
            references impurity_parameter
            on update cascade on delete restrict,
    consignment_id        integer
        constraint impurity_in_consignment_grain_consignment_id_fkey
            references consignment
            on update cascade on delete restrict
);


create table grain_classification
(
    id                 serial
        constraint grain_classification_pkey
            primary key,
    moisture           real    not null,
    grain_impurity     real    not null,
    volatiles_impurity real    not null,
    type               integer not null,
    subtype            integer not null,
    grain_type_id      integer
        constraint grain_classification_grain_type_id_fkey
            references grain_type
            on update cascade on delete restrict
);

create table part_of_impurity_parameter
(
    id                    serial
        constraint part_of_impurity_parameter_pkey
            primary key,
    impurity_parameter_id integer
        constraint part_of_impurity_parameter_impurity_parameter_id_fkey
            references impurity_parameter
            on update cascade on delete restrict,
    part_id               integer
        constraint part_of_impurity_parameter_part_id_fkey
            references Part
            on update cascade on delete restrict,
    value                 integer default 0.0 not null
);


INSERT INTO public.Part (id, name) VALUES (1, 'Минеральная примесь                                                                                                                                                                                                                                             ');
INSERT INTO public.Part (id, name) VALUES (2, 'Органическая примесь                                                                                                                                                                                                                                            ');
INSERT INTO public.Part (id, name) VALUES (3, 'Сорная примесь                                                                                                                                                                                                                                                  ');
INSERT INTO public.Part (id, name) VALUES (4, 'Испорченные зерна                                                                                                                                                                                                                                               ');
INSERT INTO public.Part (id, name) VALUES (5, 'Проход 1.0/1.5                                                                                                                                                                                                                                                  ');
INSERT INTO public.Part (id, name) VALUES (6, 'Галька                                                                                                                                                                                                                                                          ');
INSERT INTO public.Part (id, name) VALUES (7, 'Овсюг                                                                                                                                                                                                                                                           ');
INSERT INTO public.Part (id, name) VALUES (8, 'Овес                                                                                                                                                                                                                                                            ');
INSERT INTO public.Part (id, name) VALUES (9, 'Спорынья                                                                                                                                                                                                                                                        ');
INSERT INTO public.Part (id, name) VALUES (10, 'Головня                                                                                                                                                                                                                                                         ');
INSERT INTO public.Part (id, name) VALUES (11, 'Битые зерна                                                                                                                                                                                                                                                     ');
INSERT INTO public.Part (id, name) VALUES (12, 'Изъеденные зерна                                                                                                                                                                                                                                                ');
INSERT INTO public.Part (id, name) VALUES (13, 'Щуплые зерна                                                                                                                                                                                                                                                    ');
INSERT INTO public.Part (id, name) VALUES (14, 'Проросшие зерна                                                                                                                                                                                                                                                 ');
INSERT INTO public.Part (id, name) VALUES (15, 'Поврежденных зерен                                                                                                                                                                                                                                              ');
INSERT INTO public.Part (id, name) VALUES (16, 'Давленных зерен                                                                                                                                                                                                                                                 ');
INSERT INTO public.Part (id, name) VALUES (17, 'Обрушенных зерен                                                                                                                                                                                                                                                ');
INSERT INTO public.Part (id, name) VALUES (18, 'Зеленых                                                                                                                                                                                                                                                         ');
INSERT INTO public.Part (id, name) VALUES (19, 'Проход                                                                                                                                                                                                                                                          ');
INSERT INTO public.Part (id, name) VALUES (20, 'Пшеница                                                                                                                                                                                                                                                         ');
INSERT INTO public.Part (id, name) VALUES (21, 'Кукуруза                                                                                                                                                                                                                                                        ');
INSERT INTO public.Part (id, name) VALUES (22, 'Рожь                                                                                                                                                                                                                                                            ');
INSERT INTO public.Part (id, name) VALUES (23, 'Овес                                                                                                                                                                                                                                                            ');
INSERT INTO public.Part (id, name) VALUES (24, 'Ячмень                                                                                                                                                                                                                                                          ');


INSERT INTO public.company (id, name, postcode, state, city, street, house, office) VALUES (1, 'ООО "Фасад-Комплект"', 309500, 'Белгородская область', 'Старый Оскол', 'улица Конева', '47', '6');
INSERT INTO public.company (id, name, postcode, state, city, street, house, office) VALUES (3, 'ООО "Торговый дом "Агромир"', 308007, 'Белгородская область', 'Белгород', 'улица Победы', '47', '5');
INSERT INTO public.company (id, name, postcode, state, city, street, house, office) VALUES (4, 'ЗАО Племзавод "Семеновский"', 425222, 'республика Марий Эл, Медведевский район', 'Кузнецово', 'улица Мира', '1', '0');
INSERT INTO public.company (id, name, postcode, state, city, street, house, office) VALUES (2, 'ООО "Русагро"', 308002, 'Белгородская область', 'Белгород', 'проспект Богдана Хмельницкого', '111', '0');

INSERT INTO public.grain_type (id, name, gost_name, sort) VALUES (1, 'Кукуруза кормовая', 'ГОСТ Р 53903-2010', 'Fodder');
INSERT INTO public.grain_type (id, name, gost_name, sort) VALUES (2, 'Кукуруза', 'ГОСТ 13634-90', 'Food');

INSERT INTO public.impurity_parameter (id, name) VALUES (1, 'Сорная примесь');
INSERT INTO public.impurity_parameter (id, name) VALUES (2, 'Зерновая примесь');

INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (1, 14, 5, 3, 1, 0, 1);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (2, 14, 10, 4, 2, 0, 1);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (3, 14, 15, 5, 3, 0, 1);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (4, 25, 5, 5, 1, 1, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (5, 25, 5, 5, 1, 2, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (6, 25, 5, 5, 1, 5, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (7, 25, 5, 5, 1, 6, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (8, 25, 5, 10, 2, 1, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (9, 25, 5, 10, 2, 2, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (10, 25, 5, 10, 2, 3, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (11, 25, 5, 10, 2, 4, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (12, 25, 5, 10, 2, 5, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (13, 25, 5, 10, 2, 6, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (14, 25, 5, 10, 2, 7, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (15, 25, 5, 10, 2, 8, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (16, 25, 5, 10, 2, 9, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (17, 25, 5, 15, 3, 1, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (18, 25, 5, 15, 3, 2, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (19, 25, 5, 15, 3, 3, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (20, 25, 5, 15, 3, 4, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (21, 25, 5, 15, 3, 5, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (22, 25, 5, 15, 3, 6, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (23, 25, 5, 15, 3, 7, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (24, 25, 5, 15, 3, 8, 2);
INSERT INTO public.grain_classification (id, moisture, grain_impurity, volatiles_impurity, type, subtype, grain_type_id) VALUES (25, 25, 5, 15, 3, 9, 2);


