CREATE TABLE owners (
    own_id BIGINT NOT NULL AUTO_INCREMENT,
    own_name VARCHAR(255) NOT NULL,
    own_email VARCHAR(255) NULL,
    own_phone VARCHAR(50) NULL,
    own_document_number VARCHAR(100) NULL,
    own_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    own_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (own_id),
    UNIQUE KEY uk_owners_document_number (own_document_number)
);
