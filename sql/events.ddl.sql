create table events
(
    id     serial
        constraint events_pk
            primary key,
    system varchar(50) not null,
    type   varchar(20) not null
);

comment on table events is 'Main Events Header';

comment on column events.id is 'Unique Identifier';

alter table events
    owner to postgres;

create index events_id_index
    on events (id);

