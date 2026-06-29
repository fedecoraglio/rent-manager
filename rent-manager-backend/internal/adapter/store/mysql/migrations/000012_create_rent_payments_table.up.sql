CREATE TABLE rent_payments (
                               rpa_id BIGINT NOT NULL AUTO_INCREMENT,
                               rco_id BIGINT NOT NULL,

                               rpa_period DATE NOT NULL,
                               rpa_due_date DATE NOT NULL,
                               rpa_payment_date DATE NULL,

                               rpa_base_amount DECIMAL(12,2) NOT NULL,
                               rpa_suggested_adjustment_percentage DECIMAL(8,4) NULL,
                               rpa_applied_adjustment_percentage DECIMAL(8,4) NULL,

                               rpa_suggested_interest_amount DECIMAL(12,2) NULL,
                               rpa_applied_interest_amount DECIMAL(12,2) NULL,

                               rpa_total_amount DECIMAL(12,2) NOT NULL,
                               rpa_paid_amount DECIMAL(12,2) NULL,

                               rpa_is_paid BOOLEAN NOT NULL DEFAULT FALSE,
                               rpa_notes TEXT NULL,

                               rpa_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               rpa_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                               PRIMARY KEY (rpa_id),

                               UNIQUE KEY uk_rent_payments_contract_period (rco_id, rpa_period),

                               INDEX idx_rent_payments_contract_id (rco_id),
                               INDEX idx_rent_payments_period (rpa_period),
                               INDEX idx_rent_payments_due_date (rpa_due_date),
                               INDEX idx_rent_payments_is_paid (rpa_is_paid),

                               CONSTRAINT fk_rent_payments_contract
                                   FOREIGN KEY (rco_id) REFERENCES rental_contracts(rco_id)
);