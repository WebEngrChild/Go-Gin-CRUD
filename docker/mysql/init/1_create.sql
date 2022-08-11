CREATE DATABASE golang_db;
use golang_db;

CREATE TABLE companies (
                        id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL
);

CREATE TABLE departments (
                         id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
                         company_id TINYINT NOT NULL,
                         name VARCHAR (255) NOT NULL
);

CREATE TABLE grades (
                             name VARCHAR (255) NOT NULL
);

CREATE TABLE employees (
                            id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
                            department_id TINYINT NOT NULL,
                            position_id TINYINT NOT NULL,
                            person_id TINYINT NOT NULL
);