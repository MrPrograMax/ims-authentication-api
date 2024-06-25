CREATE TABLE roles
(
    id          serial      not null unique primary key,
    name       varchar(255) not null
);

CREATE TABLE users
(
    id            serial       not null unique primary key,
    login         varchar(255) not null unique,
    password      varchar(255) not null unique,
    role_id int not null,

    CONSTRAINT fk_role
        FOREIGN KEY (role_id)
            REFERENCES roles(id)
);

INSERT INTO roles (name) VALUES ('admin');
INSERT INTO roles (name) VALUES ('manager');
INSERT INTO roles (name) VALUES ('worker');



