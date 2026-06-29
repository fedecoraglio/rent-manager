CREATE TABLE tenants (
                         ten_id BIGINT NOT NULL AUTO_INCREMENT,

                         cou_id BIGINT NULL,
                         sta_id BIGINT NULL,

                         ten_name VARCHAR(255) NOT NULL,
                         ten_email VARCHAR(255) NULL,
                         ten_phone VARCHAR(50) NULL,
                         ten_document_number VARCHAR(100) NULL,

                         ten_city VARCHAR(150) NULL,
                         ten_street VARCHAR(255) NULL,
                         ten_street_number VARCHAR(50) NULL,
                         ten_floor VARCHAR(50) NULL,
                         ten_apartment VARCHAR(50) NULL,
                         ten_postal_code VARCHAR(20) NULL,

                         ten_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         ten_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                         PRIMARY KEY (ten_id),

                         UNIQUE KEY uk_tenants_document_number (ten_document_number),

                         INDEX idx_tenants_country_id (cou_id),
                         INDEX idx_tenants_state_id (sta_id),

                         CONSTRAINT fk_tenants_country
                             FOREIGN KEY (cou_id) REFERENCES countries(cou_id),

                         CONSTRAINT fk_tenants_state
                             FOREIGN KEY (sta_id) REFERENCES states(sta_id)
);