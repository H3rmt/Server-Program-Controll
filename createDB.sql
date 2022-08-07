-- ------------------------------ Auth --------------------------------

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


-- ------------------------------ Programs --------------------------------

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
    ID         int auto_increment
        primary key,
    program_ID int                                              not null,
    Type       enum ('Background','Process', 'Receive', 'Send') not null,
    Date       timestamp default CURRENT_TIMESTAMP              not null on update CURRENT_TIMESTAMP,
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
    Message    text                                         not null,
    Type       enum ('Low', 'Normal', 'Important', 'Error') not null,
    constraint logs_ID_uindex
        unique (ID),
    constraint logs_programs_fk
        foreign key (program_ID) references programs (ID)
);


-- ------------------------------ Auth 2 --------------------------------

use Auth;

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
-- 0 = read, 1 = start, 2 = stop


-- ------------------------------ Users --------------------------------

create user Website
    identified by '/6uM8qlYUm*NFCef';
grant create, insert, select, update on Programs.* to Website;
grant create, delete, insert, select, update on Auth.* to Website;

create user go
    identified by 'e73EG6dP2f8F2dAx';
grant create, insert, select, update on Programs.* to go;
grant select on Auth.* to go;


-- ------------------------------ Defaults --------------------------------

use Auth;

insert into settings (name, value) values ('timeout', '15');

insert into users (name, passwd, admin) values ('test','$2y$10$ysQqYeBXff8whujHBFcO9.kjhRb9Jblaz33QFaaqvJNHXE/z4jGAi',1);

insert into Programs.programs (ID, APIKey, Name, Description, Imagesource, Active, StatechangeTime) values (1, '62ae447bd8e8d', 'Vertretungsplan Bot', 'Discord bot um Vertretungen für die Schule anzuzeigen', 'https://cdn.discordapp.com/avatars/937432970065756190/1520488765dd4bff0b6e1bd5da464977.webp?size=256', 1, '2022-07-29 12:06:25');
insert into Programs.programs (ID, APIKey, Name, Description, Imagesource, Active, StatechangeTime) values (2, '62b4dff47b797', 'Voice Channel Bot', 'Bot to create and manage voice Channels', 'https://cdn.discordapp.com/avatars/839511426061434881/1f5d348784c9c0f1c593e72df941129c.webp?size=256', 1, '2022-07-29 12:06:27');
