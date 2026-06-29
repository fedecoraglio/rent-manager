CREATE TABLE users (
       usr_id BIGINT NOT NULL AUTO_INCREMENT,
       rol_id BIGINT NOT NULL,

       usr_name VARCHAR(255) NOT NULL,
       usr_email VARCHAR(255) NOT NULL,
       usr_password_hash VARCHAR(255) NOT NULL,

       usr_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       usr_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

       PRIMARY KEY (usr_id),
       UNIQUE KEY uk_users_email (usr_email),
       INDEX idx_users_role_id (rol_id),

       CONSTRAINT fk_users_role
           FOREIGN KEY (rol_id) REFERENCES roles(rol_id)
);

INSERT INTO users (usr_id, rol_id, usr_name, usr_email, usr_password_hash, usr_created_at, usr_updated_at)
VALUES (1, 1, 'Federico Coraglio', 'federicocoraglio@gmail.com', '$2a$12$7JR0iW6sjVUd2D.zk25ioelxRujf7wzIEKz3.wjV2nfFmb9YmSKKO', '2026-06-09 10:13:00', '2026-06-09 10:13:00');
