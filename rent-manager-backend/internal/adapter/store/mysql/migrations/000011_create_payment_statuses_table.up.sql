CREATE TABLE payment_statuses (
    pas_id BIGINT NOT NULL AUTO_INCREMENT,
    pas_code VARCHAR(50) NOT NULL,
    pas_name VARCHAR(100) NOT NULL,

    PRIMARY KEY (pas_id),
    UNIQUE KEY uk_payment_statuses_code (pas_code)
);

INSERT INTO payment_statuses (pas_code, pas_name)
VALUES
('pending', 'Pending'),
('paid', 'Paid'),
('late', 'Late'),
('cancelled', 'Cancelled');
