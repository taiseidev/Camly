#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"

# Verify MySQL connection
if ! $CMD_MYSQL -e "SELECT 1" > /dev/null 2>&1; then
    echo "Failed to connect to MySQL. Please check your credentials and connection."
    exit 1
fi

# users テーブルの作成と初期データ挿入
$CMD_MYSQL -e "
START TRANSACTION;

CREATE TABLE IF NOT EXISTS users (
    id INT(10) AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name VARCHAR(50),
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    status ENUM('active', 'inactive', 'suspended') NOT NULL DEFAULT 'active',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    INDEX idx_email (email)
);

INSERT INTO users (id, name, email, password_hash, status) VALUES
    (1, 'John Doe', 'johndoe@example.com', '$2a$10$EXAMPLE_HASH_1', 'active'),
    (2, 'Jane Smith', 'janesmith@example.com', '$2a$10$EXAMPLE_HASH_2', 'active');

COMMIT;
"

echo "Database setup complete."
