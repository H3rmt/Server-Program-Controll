create schema Programs;
use Programs;

create table programs
(
    ID              int(13)     not null
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
    ID         int(13)                                                 not null
        primary key,
    program_ID int(13)                                                 not null,
    Type       enum ('Backgroundprocess', 'Process', 'Recive', 'Send') not null,
    Date       timestamp default current_timestamp()                   not null on update current_timestamp(),
    constraint activity_ibfk_1
        foreign key (program_ID) references programs (ID)
);

create index program_ID
    on activity (program_ID);

create table logs
(
    ID         int(13) auto_increment
        primary key,
    program_ID int(13)                                      not null,
    Date       timestamp default current_timestamp()        not null on update current_timestamp(),
    Number     int                                          not null,
    Message    text                                         not null,
    Type       enum ('Low', 'Normal', 'Important', 'Error') not null,
    constraint logs_ibfk_1
        foreign key (program_ID) references programs (ID)
);

create index program_ID
    on logs (program_ID);

create table settings
(
    Name  varchar(30) not null,
    Value text        not null,
    constraint Settings_Name_uindex
        unique (Name)
);

# initial password = authorise

INSERT INTO settings (Name, Value) VALUES ('adminCookie', '');
INSERT INTO settings (Name, Value) VALUES ('password', '20569225230b1abc60cff5c8cd4c990024841f733d7bf22b53a46b30bb53e8b0');
INSERT INTO settings (Name, Value) VALUES ('timeout', '86400');

create user Website
    identified by '/6uM8qlYUm*NFCef';
grant create, insert, select, update on Programs.* to Website;

create user go
    identified by 'e73EG6dP2f8F2dAx';
grant create, insert, select, update on Programs.* to go;