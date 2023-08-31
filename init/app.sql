USE avito;

CREATE TABLE `segment`
(
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (`name`)
);

CREATE TABLE `compliance`
(
    id_user INT NOT NULL,
    name_segment VARCHAR(255) NOT NULL,
    UNIQUE KEY `ukey_compliance_id_name` (`id_user`,`name_segment`),
    FOREIGN KEY (`name_segment`) REFERENCES `segment` (`name`) ON DELETE CASCADE
);

CREATE TABLE `history`
(
    id_user INT NOT NULL,
    name_segment VARCHAR(255) NOT NULL,
    ts_add TIMESTAMP NOT NULL,
    ts_del TIMESTAMP DEFAULT NULL,
    UNIQUE KEY `ukey_history_id_name` (`id_user`,`name_segment`)
);