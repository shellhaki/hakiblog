create table if not exists blog_post(
    id serial primary key,
    title varchar(50),
    body text,
    tags text[],
    cover text,
    shares integer default 0,
    created_at timestamp with time zone default current_timestamp
);

create table if not exists comments(
    comment_id serial primary key,
    blog_id int references blog_post(id) on delete cascade,
    message varchar(500),
    created_at timestamp with time zone default current_timestamp
);

create table if not exists likes (
    blog_id int references blog_post(id) on delete cascade,
    count int,
    created_at timestamp with time zone default current_timestamp
);