#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "create table article (
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    title varchar(50) NOT NULL,
    body varchar(1000)
    );"
$CMD_MYSQL -e  "insert into article values (1, '記事1', '記事1です。');"
$CMD_MYSQL -e  "insert into article values (2, '記事2', '記事2です。');"
