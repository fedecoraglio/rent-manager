CREATE TABLE countries (
                           cou_id BIGINT NOT NULL AUTO_INCREMENT,
                           cou_code VARCHAR(10) NOT NULL,
                           cou_name VARCHAR(150) NOT NULL,
                           cou_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           cou_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                           PRIMARY KEY (cou_id),
                           UNIQUE KEY uk_countries_code (cou_code),
                           UNIQUE KEY uk_countries_name (cou_name)
);