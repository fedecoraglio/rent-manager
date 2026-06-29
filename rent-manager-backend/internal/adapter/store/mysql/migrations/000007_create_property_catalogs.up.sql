CREATE TABLE property_types (
    pty_id BIGINT NOT NULL AUTO_INCREMENT,
    pty_code VARCHAR(50) NOT NULL,
    pty_name VARCHAR(100) NOT NULL,

    PRIMARY KEY (pty_id),
    UNIQUE KEY uk_property_types_code (pty_code)
);

INSERT INTO property_types (pty_code, pty_name)
VALUES
('apartment', 'Apartment'),
('house', 'House'),
('commercial', 'Commercial'),
('garage', 'Garage');

CREATE TABLE property_statuses (
    pst_id BIGINT NOT NULL AUTO_INCREMENT,
    pst_code VARCHAR(50) NOT NULL,
    pst_name VARCHAR(100) NOT NULL,

    PRIMARY KEY (pst_id),
    UNIQUE KEY uk_property_statuses_code (pst_code)
);

INSERT INTO property_statuses (pst_code, pst_name)
VALUES
('available', 'Available'),
('rented', 'Rented'),
('maintenance', 'Maintenance'),
('inactive', 'Inactive');
