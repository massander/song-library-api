create table if not exists
    songs (
        "id" uuid primary key,
        "name" text not null,
        "group" text not null,
        "text" text[] not null default '{}',
        "releaseDate" date not null,
        "link" text not null default ''
    );