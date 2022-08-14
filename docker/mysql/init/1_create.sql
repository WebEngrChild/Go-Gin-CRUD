CREATE DATABASE golang;
use golang;

CREATE TABLE companies (
                        id TINYINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL
);

CREATE TABLE departments (
                        id TINYINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
                        name VARCHAR (255) NOT NULL
);

CREATE TABLE persons (
                         id TINYINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
                         name VARCHAR (255) NOT NULL,
                         gender BOOLEAN,
                         birthday DATE,
                         phone VARCHAR (255)
);

CREATE TABLE branches (
                          id TINYINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
                          name VARCHAR (255) NOT NULL
);

CREATE TABLE employees (
                        id TINYINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
                        person_id TINYINT NOT NULL,
                        company_id TINYINT NOT NULL,
                        department_id TINYINT NOT NULL,
                        branch_id TINYINT NOT NULL,
                        FOREIGN KEY (person_id) REFERENCES persons(id),
                        FOREIGN KEY (company_id) REFERENCES companies(id),
                        FOREIGN KEY (department_id) REFERENCES departments(id),
                        FOREIGN KEY (branch_id) REFERENCES branches(id)
);

INSERT INTO persons (name, gender, birthday, phone)
VALUES ('鈴木一郎', 0, '1993-12-31', '090-9999-8888'),
       ('田中太郎', 0, '1989-1-20', '080-7777-8888'),
       ('佐藤優作', 0, '1978-6-16', '080-6666-8888'),
       ('加藤美玲', 1, '2000-5-2', '080-5555-8888'),
       ('山本美月', 1, '1999-8-8', '080-4444-8888'),
       ('井上由美', 1, '1978-6-18', '080-3333-8888');

INSERT INTO companies (name)
VALUES ('アップル株式会社'),
       ('オレンジ株式会社'),
       ('パイナップル株式会社'),
       ('メロン株式会社'),
       ('グレープ株式会社'),
       ('バナナ株式会社');

INSERT INTO departments (name)
VALUES ('営業部'),
       ('マーケティング部'),
       ('経理財務部'),
       ('研究開発部'),
       ('研究開発部'),
       ('製造部');

INSERT INTO branches (name)
VALUES ('東京'),
       ('大阪'),
       ('愛知'),
       ('福岡'),
       ('横浜'),
       ('仙台');

INSERT INTO employees (person_id, company_id, department_id, branch_id)
VALUES (1, 1, 1, 2),
       (2, 5, 2, 6),
       (3, 1, 4, 3),
       (4, 4, 4, 3),
       (5, 4, 3, 1),
       (6, 1, 3, 1);