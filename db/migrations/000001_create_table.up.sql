create table if not exists users(
id serial primary key,
uid varchar(100) not null unique,
name varchar(20) not null,
sex varchar(5) not null,
height int,
weight int,
age int,
job varchar(50),
created_at timestamp default current_timestamp
);

create table if not exists momentum(
id serial primary key,
user_id varchar(100) not null,
steps int,
calories int,
distance int,
max_heart_rate int,
min_heart_rate int,
date date not null,
created_at timestamp default current_timestamp
);

create table if not exists sleep(
id serial primary key,
user_id varchar(100) not null,
hours int,
started_at timestamp,
ended_at timestamp,
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
calorie_per_100g int,
date date not null,
created_at timestamp default current_timestamp  
);