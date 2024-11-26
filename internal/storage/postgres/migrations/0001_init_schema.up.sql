create table if not exists
    songs (
        "id" uuid primary key,
        "name" text not null,
        "group" text not null,
        "text" text[] not null default '{}',
        "release_date" date not null,
        "link" text not null default ''
    );