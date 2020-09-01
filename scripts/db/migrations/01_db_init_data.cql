INSERT INTO hydromon.users (email, name, last_name, password, admin) VALUES
('john.doe@example.com', 'John', 'Doe', null, true);
INSERT INTO hydromon.users (email, name, last_name, password, admin) VALUES
('marcia.clark@example.com', 'Marcia', 'Clark', null, false);

INSERT INTO hydromon.nodes (id, description, manual_reading, "token") VALUES
('lujan-1', 'Nodo instalado debajo de un puente', false, null);

INSERT INTO hydromon.configurations (node_id, configuration) VALUES
('lujan-1', '{"1":{"interval":25,"upperLimit":1,"lowerLimit":-999,"picturesNum":1,"next":"2","prev":""},"2":{"interval":10,"upperLimit":2,"lowerLimit":1,"picturesNum":1,"next":"3","prev":"1"},"3":{"interval":5,"upperLimit":999,"lowerLimit":2,"picturesNum":1,"next":"","prev":"2"}}');