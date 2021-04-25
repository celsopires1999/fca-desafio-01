USE desafio-01;
CREATE TABLE fcmodule (
    id integer not null,
    name varchar(100)
);

SET character_set_client = utf8;
SET character_set_connection = utf8;
SET character_set_results = utf8;
SET collation_connection = utf8_general_ci;

INSERT INTO fcmodule (id, name) VALUES (1, 'Docker');
INSERT INTO fcmodule (id, name) VALUES (2, 'Padrões e técnicas avançadas com Git e Gitbub');
INSERT INTO fcmodule (id, name) VALUES (3, 'Integração contínua');
INSERT INTO fcmodule (id, name) VALUES (4, 'Kubernets');
INSERT INTO fcmodule (id, name) VALUES (5, 'Service Mesh com Istio');
