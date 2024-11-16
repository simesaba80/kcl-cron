create table if not exists users(
id serial primary key,
uid varchar(100) not null unique,
fitbit_user_id varchar(10),
name varchar(20) not null,
sex varchar(5) not null,
height int,
weight int,
age int,
job varchar(50),
access_token varchar(250),
refresh_token varchar(250),
created_at timestamp default current_timestamp
);

create table if not exists activities(
id serial primary key,
user_id varchar(100) not null,
steps int,
calories int,
date date not null,
created_at timestamp default current_timestamp
);

create table if not exists sleep(
id serial primary key,
user_id varchar(100) not null,
minutes int,
deep_sleep int,
light_sleep int,
rem_sleep int,
wake int,
date date not null,
created_at timestamp default current_timestamp  
);

create table if not exists meal(
id serial primary key,
user_id varchar(100) not null,
meal_name varchar(100),
calories int,
protein real,
fat real,
carbohydrates real,
salt real,
calcium real,
date date not null,
created_at timestamp default current_timestamp  
);