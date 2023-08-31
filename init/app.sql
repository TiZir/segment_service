USE avito;

CREATE TABLE `segment`
(
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (`name`)
);

INSERT INTO `segment`(name) VALUES ("AVITO_VOICE_MESSAGES");
INSERT INTO `segment`(name) VALUES ("AVITO_PERFORMANCE_VAS");
INSERT INTO `segment`(name) VALUES ("AVITO_DISCOUNT_30");
INSERT INTO `segment`(name) VALUES ("AVITO_DISCOUNT_50");

CREATE TABLE `compliance`
(
    id_user INT NOT NULL,
    name_segment VARCHAR(255) NOT NULL,
    UNIQUE KEY `ukey_compliance_id_name` (`id_user`,`name_segment`),
    FOREIGN KEY (`name_segment`) REFERENCES `segment` (`name`) ON DELETE CASCADE
);

INSERT INTO `compliance`(id_user, name_segment) VALUES (1000, "AVITO_VOICE_MESSAGES");
INSERT INTO `compliance`(id_user, name_segment)  VALUES (1000, "AVITO_PERFORMANCE_VAS");
INSERT INTO `compliance`(id_user, name_segment)  VALUES (1000, "AVITO_DISCOUNT_30");
INSERT INTO `compliance`(id_user, name_segment)  VALUES (1002, "AVITO_VOICE_MESSAGES");
INSERT INTO `compliance`(id_user, name_segment)  VALUES (1002, "AVITO_PERFORMANCE_VAS");
INSERT INTO `compliance`(id_user, name_segment)  VALUES (1002, "AVITO_DISCOUNT_50");

CREATE TABLE `history`
(
    id_user INT NOT NULL,
    name_segment VARCHAR(255) NOT NULL,
    ts_add TIMESTAMP NOT NULL,
    ts_del TIMESTAMP DEFAULT NULL,
    UNIQUE KEY `ukey_history_id_name` (`id_user`,`name_segment`)
);

INSERT INTO `history` VALUES (1000, "AVITO_DISCOUNT_30", CURRENT_TIMESTAMP, NULL);
INSERT INTO `history` VALUES (1002, "AVITO_VOICE_MESSAGES", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);