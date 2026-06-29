CREATE TABLE roles (
                       rol_id BIGINT NOT NULL AUTO_INCREMENT,
                       rol_code VARCHAR(50) NOT NULL,
                       rol_name VARCHAR(100) NOT NULL,
                       rol_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

                       PRIMARY KEY (rol_id),
                       UNIQUE KEY uk_roles_code (rol_code)
);

INSERT INTO roles (rol_code, rol_name)
VALUES
    ('admin', 'Administrator'),
    ('manager', 'Manager');