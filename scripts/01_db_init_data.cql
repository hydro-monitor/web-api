INSERT INTO hydromon.nodes (id, description, manual_reading) VALUES
('lujan-1', 'Nodo instalado debajo de un puente', false);

INSERT INTO hydromon.configurations (node_id, configuration) VALUES
('lujan-1', '{"1":{"interval":25,"upperLimit":1,"lowerLimit":-999,"picturesNum":1,"next":"2","prev":""},"2":{"interval":10,"upperLimit":2,"lowerLimit":1,"picturesNum":1,"next":"3","prev":"1"},"3":{"interval":5,"upperLimit":999,"lowerLimit":2,"picturesNum":1,"next":"","prev":"2"}}');