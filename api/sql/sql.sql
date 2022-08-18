-- mysql database scripts
CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios (
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    passwd varchar(50) not null unique,
    created timestamp default current_timestamp()
) ENGINE=INNODB;
