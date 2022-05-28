# -------------------------------- Auth --------------------------------

create schema Auth;
use Auth;

create table settings
(
    Name  varchar(30) not null
        primary key,
    Value text        not null,
    constraint Settings_Name_uindex
        unique (Name)
);

create table users
(
    name   varchar(60)          not null,
    passwd varchar(60)          not null,
    ID     int auto_increment
        primary key,
    admin  tinyint(1) default 0 not null,
    constraint users_ID_uindexx
        unique (ID),
    constraint users_id_uindex
        unique (name)
);

create table sessions
(
    ID          int auto_increment
        primary key,
    expire_date timestamp   not null,
    hash        varchar(60) not null,
    user_id     int         not null,
    constraint sessions_id_uindex
        unique (hash),
    constraint sessions_users_ID_fk
        foreign key (user_id) references users (ID)
)
    auto_increment = 21;

create table user_programs_permissions
(
    user_id    int null,
    program_id int null,
    constraint user_programs_permissions_programs_ID_fk
        foreign key (program_id) references Programs.programs (ID),
    constraint user_programs_permissions_users_ID_fk
        foreign key (user_id) references users (ID)
);

# -------------------------------- Programs --------------------------------

create schema Programs;
use Programs;

create table programs
(
    ID              int         not null
        primary key,
    APIKey          varchar(30) not null,
    Name            varchar(30) not null,
    Description     text        not null,
    Imagesource     text        not null,
    Active          tinyint(1)  null,
    StatechangeTime timestamp   null,
    constraint APIKey
        unique (APIKey)
);

create table activity
(
    ID         int                                                     not null
        primary key,
    program_ID int                                                     not null,
    Type       enum ('Backgroundprocess', 'Process', 'Recive', 'Send') not null,
    Date       timestamp default CURRENT_TIMESTAMP                     not null on update CURRENT_TIMESTAMP,
    constraint activity_ibfk_1
        foreign key (program_ID) references programs (ID)
);

create index program_ID
    on activity (program_ID);

create table logs
(
    ID         int auto_increment
        primary key,
    program_ID int                                          not null,
    Date       timestamp default CURRENT_TIMESTAMP          not null on update CURRENT_TIMESTAMP,
    Number     int                                          not null,
    Message    text                                         not null,
    Type       enum ('Low', 'Normal', 'Important', 'Error') not null,
    constraint logs_ibfk_1
        foreign key (program_ID) references programs (ID)
);

create index program_ID
    on logs (program_ID);
