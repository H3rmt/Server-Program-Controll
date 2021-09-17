# created by PHP-Storm
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

