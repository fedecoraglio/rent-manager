CREATE TABLE inflation_indexes (
       ixi_id BIGINT AUTO_INCREMENT PRIMARY KEY,
       ixi_period DATE NOT NULL,
       ixi_percentage DECIMAL(10,4) NOT NULL,
       ixi_source VARCHAR(100) NULL,
       ixi_notes TEXT NULL,
       ixi_created_at DATETIME NOT NULL,
       ixi_updated_at DATETIME NOT NULL,

       UNIQUE KEY uk_inflation_indexes_period (ixi_period)
);