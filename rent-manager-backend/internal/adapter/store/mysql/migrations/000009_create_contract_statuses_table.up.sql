CREATE TABLE contract_statuses (
    cst_id BIGINT NOT NULL AUTO_INCREMENT,
    cst_code VARCHAR(50) NOT NULL,
    cst_name VARCHAR(100) NOT NULL,

    PRIMARY KEY (cst_id),
    UNIQUE KEY uk_contract_statuses_code (cst_code)
);

INSERT INTO contract_statuses (cst_code, cst_name)
VALUES
('active', 'Active'),
('finished', 'Finished'),
('cancelled', 'Cancelled');
