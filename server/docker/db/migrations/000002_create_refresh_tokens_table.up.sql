CREATE TABLE IF NOT EXISTS refresh_tokens (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT UNSIGNED NOT NULL,
    token_hash VARCHAR(255) NOT NULL UNIQUE,
    user_agent VARCHAR(255) DEFAULT NULL,
    ip_address VARCHAR(45) DEFAULT NULL,
    expires_at DATETIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_refresh_tokens_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_refresh_tokens_expires (expires_at),
    INDEX idx_refresh_tokens_user (user_id)
);
