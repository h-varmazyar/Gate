create table public.posts
(
    id              bigserial primary key,
    created_at      timestamp with time zone,
    updated_at      timestamp with time zone,
    deleted_at      timestamp with time zone,
    posted_at       timestamp with time zone not null,
    content         varchar(4094),
    parent_id       bigint                   not null,
    sender_username varchar(64),
    provider        varchar(32),
    tags            varchar(32)[],
    sentiment       double precision default null,
    like_count      bigint,
    retwit_count    bigint,
    comment_count   bigint,
    quote_count     bigint,
    type            varchar(32)
);

create index idx_posts_deleted_at
    on public.posts (deleted_at);

