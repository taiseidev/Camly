#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"

# articles テーブルの作成と初期データ挿入
$CMD_MYSQL -e "CREATE TABLE IF NOT EXISTS article (
    id INT(10) AUTO_INCREMENT NOT NULL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    body VARCHAR(1000)
);"
$CMD_MYSQL -e "INSERT INTO article (id, title, body) VALUES
    (1, '記事1', '記事1です。'),
    (2, '記事2', '記事2です。');"

# users テーブルの作成と初期データ挿入
$CMD_MYSQL -e "CREATE TABLE IF NOT EXISTS users (
    id INT(10) AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE
);"
$CMD_MYSQL -e "INSERT INTO users (id, name, email) VALUES
    (1, 'John Doe', 'johndoe@example.com'),
    (2, 'Jane Smith', 'janesmith@example.com');"

echo "Database setup complete."
