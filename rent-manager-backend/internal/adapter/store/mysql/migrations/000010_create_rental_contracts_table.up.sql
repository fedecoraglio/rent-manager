CREATE TABLE interest_calculation_types (
    ict_id BIGINT NOT NULL AUTO_INCREMENT,
    ict_code VARCHAR(50) NOT NULL,
    ict_name VARCHAR(100) NOT NULL,

    PRIMARY KEY (ict_id),
    UNIQUE KEY uk_interest_calculation_types_code (ict_code)
);

INSERT INTO interest_calculation_types (ict_code, ict_name)
VALUES
    ('daily_from_due_day_next_day', 'Daily interest starting the day after the due date'),
    ('daily_from_month_first_day', 'Daily interest starting from the first day of the month');


CREATE TABLE rent_adjustment_types (
       rat_id BIGINT NOT NULL AUTO_INCREMENT,
       rat_code VARCHAR(50) NOT NULL,
       rat_name VARCHAR(100) NOT NULL,

       PRIMARY KEY (rat_id),
       UNIQUE KEY uk_rent_adjustment_types_code (rat_code)
);

INSERT INTO rent_adjustment_types (rat_code, rat_name)
VALUES
    ('ipc_argentina', 'Argentine CPI');

CREATE TABLE rental_contracts (
      rco_id BIGINT NOT NULL AUTO_INCREMENT,
      pro_id BIGINT NOT NULL,
      ten_id BIGINT NOT NULL,
      cst_id BIGINT NOT NULL,
      ict_id BIGINT NOT NULL,

      rco_start_date DATE NOT NULL,
      rco_end_date DATE NOT NULL,
      rco_total_payments INT NOT NULL,
      rco_monthly_amount DECIMAL(12,2) NOT NULL,
      rco_deposit_amount DECIMAL(12,2) NOT NULL DEFAULT 0,
      rco_currency VARCHAR(10) NOT NULL DEFAULT 'ARS',

      rco_due_day INT NOT NULL DEFAULT 10,
      rco_daily_interest_percentage DECIMAL(5,2) NOT NULL DEFAULT 0,

      rco_notes TEXT NULL,

      rat_id BIGINT NOT NULL,
      rco_adjustment_frequency_months INT NOT NULL DEFAULT 4,

      rco_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      rco_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

      PRIMARY KEY (rco_id),

      INDEX idx_rental_contracts_property_id (pro_id),
      INDEX idx_rental_contracts_tenant_id (ten_id),
      INDEX idx_rental_contracts_status_id (cst_id),
      INDEX idx_rental_contracts_interest_type_id (ict_id),
      INDEX idx_rental_contracts_adjustment_type_id (rat_id),

      CONSTRAINT fk_rental_contracts_adjustment_type
          FOREIGN KEY (rat_id) REFERENCES rent_adjustment_types(rat_id),

      CONSTRAINT fk_rental_contracts_property
          FOREIGN KEY (pro_id) REFERENCES properties(pro_id),

      CONSTRAINT fk_rental_contracts_tenant
          FOREIGN KEY (ten_id) REFERENCES tenants(ten_id),

      CONSTRAINT fk_rental_contracts_status
          FOREIGN KEY (cst_id) REFERENCES contract_statuses(cst_id),

      CONSTRAINT fk_rental_contracts_interest_type
          FOREIGN KEY (ict_id) REFERENCES interest_calculation_types(ict_id)
);