CREATE TABLE IF NOT EXISTS log_errors(
    id BIGINT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    module_name varchar(255) NOT NULL,
    error_description varchar(500) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id)
);