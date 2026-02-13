create table event
(
    events_id serial
        constraint event_events_id_fk
            references events,
    id        integer     not null
        constraint event_pk
            primary key,
    messages  varchar(60) not null
);

alter table event
    owner to postgres;

