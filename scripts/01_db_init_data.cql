INSERT INTO hydromon.nodes (id, description, manual_reading) VALUES
('lujan-1', 'Nodo instalado debajo de un puente', false);

INSERT INTO hydromon.states (node_id, name, photos_per_reading, reading_interval, lower_limit, upper_limit, previous_state, next_state) VALUES
('lujan-1', '1', 1, 3600, -999, 1, null, '2');

INSERT INTO hydromon.states (node_id, name, photos_per_reading, reading_interval, lower_limit, upper_limit, previous_state, next_state) VALUES
('lujan-1', '2', 1, 2700, 1, 2, '1', '3');

INSERT INTO hydromon.states (node_id, name, photos_per_reading, reading_interval, lower_limit, upper_limit, previous_state, next_state) VALUES
('lujan-1', '3', 1, 1800, 2, 999, '2', null);

INSERT INTO hydromon.configurations (node_id, configuration) VALUES
('lujan-1', '{"1":{"interval":25,"upperLimit":1,"lowerLimit":-999,"picturesNum":1,"next":"2","prev":""},"2":{"interval":10,"upperLimit":2,"lowerLimit":1,"picturesNum":1,"next":"3","prev":"1"},"3":{"interval":5,"upperLimit":999,"lowerLimit":2,"picturesNum":1,"next":"","prev":"2"}}');