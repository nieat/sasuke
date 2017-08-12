CREATE DATABASE SASUKE_TEST
update user set authentication_string=PASSWORD("password") where User='root';
create user 'test_user'@'localhost' identified by 'password';
grant all privileges on *.* to test_user@"%" identified by 'password';
CREATE TABLE SASUKE_TEST.articles (id INT , content VARCHAR(255) ,user_id INTEGER , created_at DATETIME,updated_at DATETIME);
CREATE TABLE SASUKE_TEST.users (id INT ,first_naem VARCHAR(20) ,last_name VARCHAR(20) ,created_at DATETIME,updated_at DATETIME);
