create table candles
(
    id            bigint generated by default as identity primary key,
    created_at    timestamp with time zone,
    updated_at    timestamp with time zone,
    deleted_at    timestamp with time zone,
    time          timestamp with time zone,
    open          numeric,
    high          numeric,
    low           numeric,
    close         numeric,
    volume        numeric,
    amount        numeric,
    market_id     bigint  constraint fk_candles_market   references markets,
    resolution_id bigint  constraint fk_candles_resolution   references resolutions
);

alter table candles
    owner to postgres;

create index idx_candles_market_resolution_time
    on candles (resolution_id asc, market_id asc, time desc);

