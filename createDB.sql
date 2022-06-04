# -------------------------------- Auth --------------------------------

create schema Auth;
use Auth;

create table settings
(
    name  varchar(30) not null
        primary key,
    value text        not null,
    constraint settings_name_uindex
        unique (name)
);

create table users
(
    ID     int auto_increment
        primary key,
    name   varchar(60)          not null,
    passwd varchar(60)          not null,
    admin  tinyint(1) default 0 not null,
    constraint users_ID_uindex
        unique (ID),
    constraint users_name_uindex
        unique (name)
);

create table sessions
(
    ID          int auto_increment
        primary key,
    expire_date timestamp   not null,
    hash        varchar(60) not null,
    user_id     int         not null,
    constraint sessions_ID_uindex
        unique (ID),
    constraint sessions_users_fk
        foreign key (user_id) references users (ID)
);

create table user_programs_permissions
(
    user_id    int           null,
    program_id int           null,
    permission int default 0 not null,
    constraint user_programs_permissions_programs_fk
        foreign key (program_id) references Programs.programs (ID),
    constraint user_programs_permissions_users_fk
        foreign key (user_id) references users (ID)
);
# 0 = read, 1 = start, 2 = stop

# -------------------------------- Programs --------------------------------

create schema Programs;
use Programs;

create table programs
(
    ID              int auto_increment
        primary key,
    APIKey          varchar(30)          not null,
    Name            varchar(30)          not null,
    Description     text                 not null,
    Imagesource     text                 not null,
    Active          tinyint(1) default 0 null,
    StatechangeTime timestamp            null,
    constraint programs_ID_uindex
        unique (ID),
    constraint programs_APIKey_uindex
        unique (APIKey)
);

create table activity
(
    ID         int                                                     not null
        primary key,
    program_ID int                                                     not null,
    Type       enum ('Backgroundprocess', 'Process', 'Recive', 'Send') not null,
    Date       timestamp default CURRENT_TIMESTAMP                     not null on update CURRENT_TIMESTAMP,
    constraint activity_ID_uindex
        unique (ID),
    constraint activity_programs_fk
        foreign key (program_ID) references programs (ID)
);

create table logs
(
    ID         int auto_increment
        primary key,
    program_ID int                                          not null,
    Date       timestamp default CURRENT_TIMESTAMP          not null on update CURRENT_TIMESTAMP,
    Number     int                                          not null,
    Message    text                                         not null,
    Type       enum ('Low', 'Normal', 'Important', 'Error') not null,
    constraint logs_ID_uindex
        unique (ID),
    constraint logs_programs_fk
        foreign key (program_ID) references programs (ID)
);