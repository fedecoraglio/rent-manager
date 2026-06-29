CREATE TABLE properties (
                            pro_id BIGINT NOT NULL AUTO_INCREMENT,
                            own_id BIGINT NOT NULL,
                            pty_id BIGINT NOT NULL,
                            pst_id BIGINT NOT NULL,
                            cou_id BIGINT NOT NULL,
                            sta_id BIGINT NOT NULL,

                            pro_code VARCHAR(100) NOT NULL,
                            pro_title VARCHAR(255) NOT NULL,
                            pro_description TEXT NULL,

                            pro_street VARCHAR(255) NOT NULL,
                            pro_street_number VARCHAR(50) NULL,
                            pro_floor VARCHAR(50) NULL,
                            pro_apartment VARCHAR(50) NULL,
                            pro_city VARCHAR(100) NOT NULL,
                            pro_postal_code VARCHAR(50) NULL,

                            pro_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            pro_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                            PRIMARY KEY (pro_id),
                            UNIQUE KEY uk_properties_code (pro_code),

                            INDEX idx_properties_owner_id (own_id),
                            INDEX idx_properties_type_id (pty_id),
                            INDEX idx_properties_status_id (pst_id),
                            INDEX idx_properties_country_id (cou_id),
                            INDEX idx_properties_state_id (sta_id),

                            CONSTRAINT fk_properties_owner
                                FOREIGN KEY (own_id) REFERENCES owners(own_id),

                            CONSTRAINT fk_properties_type
                                FOREIGN KEY (pty_id) REFERENCES property_types(pty_id),

                            CONSTRAINT fk_properties_status
                                FOREIGN KEY (pst_id) REFERENCES property_statuses(pst_id),

                            CONSTRAINT fk_properties_country
                                FOREIGN KEY (cou_id) REFERENCES countries(cou_id),

                            CONSTRAINT fk_properties_state
                                FOREIGN KEY (sta_id) REFERENCES states(sta_id)
);