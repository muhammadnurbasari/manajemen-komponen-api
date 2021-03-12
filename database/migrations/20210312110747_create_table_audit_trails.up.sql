CREATE TABLE IF NOT EXISTS audit_trails(
    id BIGINT NOT NULL AUTO_INCREMENT,
    users_id INT NOT NULL,
    url_api VARCHAR(255) NOT NULL,
    function_name VARCHAR(255) NOT NULL,
    ip_address VARCHAR(100) NOT NULL,
    method_api ENUM('GET','POST','UPDATE','DELETE','PATCH') NOT NULL,
    response_code SMALLINT NOT NULL,
    request_body LONGTEXT,
    response_body LONGTEXT NOT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id)
);