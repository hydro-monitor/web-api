INSERT INTO hydromon.nodes (id, description, configuration, state, manual_reading) VALUES
('lujan-1', 'Nodo instalado debajo de un puente', 'configuracion-1', '1', false);

INSERT INTO hydromon.configurations (id, description, states) VALUES
('configuracion-1', 'Configuracion para nodo 1', ['1', '2', '3']);

INSERT INTO hydromon.states (id, photos_per_reading, seconds_between_readings, water_level_limit_for_previous_state,
water_level_limit_for_next_state, previous_state, next_state) VALUES
('1', 1, 3600, -999, 1, null, '2');

INSERT INTO hydromon.states (id, photos_per_reading, seconds_between_readings, water_level_limit_for_previous_state,
water_level_limit_for_next_state, previous_state, next_state) VALUES
('2', 1, 2700, 1, 2, '1', '3');

INSERT INTO hydromon.states (id, photos_per_reading, seconds_between_readings, water_level_limit_for_previous_state,
water_level_limit_for_next_state, previous_state, next_state) VALUES
('3', 1, 1800, 2, 999, '2', null);