CREATE TABLE domain (
    id serial primary key,
    domain varchar(30),
    latency integer,
    available boolean,
    last_update timestamp
);
