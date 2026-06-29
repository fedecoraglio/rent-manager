CREATE TABLE states (
                        sta_id BIGINT NOT NULL AUTO_INCREMENT,
                        cou_id BIGINT NOT NULL,
                        sta_code VARCHAR(20) NULL,
                        sta_name VARCHAR(150) NOT NULL,
                        sta_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        sta_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                        PRIMARY KEY (sta_id),
                        UNIQUE KEY uk_states_country_name (cou_id, sta_name),
                        UNIQUE KEY uk_states_country_code (cou_id, sta_code),
                        INDEX idx_states_country_id (cou_id),

                        CONSTRAINT fk_states_country
                            FOREIGN KEY (cou_id) REFERENCES countries(cou_id)
);

INSERT INTO countries (cou_code, cou_name) VALUES ('AR','Argentina');
INSERT INTO states (cou_id,sta_code,sta_name)
VALUES
    (1, 'C',  'Ciudad Autónoma de Buenos Aires'),
    (1, 'B',  'Buenos Aires'),
    (1, 'K',  'Catamarca'),
    (1, 'H',  'Chaco'),
    (1, 'U',  'Chubut'),
    (1, 'X',  'Córdoba'),
    (1, 'W',  'Corrientes'),
    (1, 'E',  'Entre Ríos'),
    (1, 'P',  'Formosa'),
    (1, 'Y',  'Jujuy'),
    (1, 'L',  'La Pampa'),
    (1, 'F',  'La Rioja'),
    (1, 'M',  'Mendoza'),
    (1, 'N',  'Misiones'),
    (1, 'Q',  'Neuquén'),
    (1, 'R',  'Río Negro'),
    (1, 'A',  'Salta'),
    (1, 'J',  'San Juan'),
    (1, 'D',  'San Luis'),
    (1, 'Z',  'Santa Cruz'),
    (1, 'S',  'Santa Fe'),
    (1, 'G',  'Santiago del Estero'),
    (1, 'V',  'Tierra del Fuego'),
    (1, 'T',  'Tucumán');