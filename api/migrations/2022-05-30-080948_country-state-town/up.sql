CREATE TABLE countries (
    id serial NOT NULL,
    iso char(2) DEFAULT NULL,
    name varchar(80) DEFAULT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE states (
    id serial NOT NULL,
    name varchar(30) DEFAULT NULL,
    fk_country_id int DEFAULT NULL,
    slug varchar(255) DEFAULT NULL,
    PRIMARY KEY (id),
    FOREIGN key(fk_country_id) REFERENCES countries(id) on update cascade on delete restrict
);
CREATE TABLE towns (
    id_town serial NOT NULL,
    fk_state_id int NOT NULL,
    cod_town int NOT NULL,
    DC int NOT NULL,
    name varchar(100) NOT NULL,
    slug varchar(255),
    PRIMARY KEY(id_town, fk_state_id)
);
alter table states
add column id_state int,
add FOREIGN key(id_state) REFERENCES states(id) on update cascade on delete restrict