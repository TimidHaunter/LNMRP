CREATE TABLE `integer`  (
    `tinyint_a` tinyint UNSIGNED NOT NULL COMMENT '无符号tinyint显示长度为10',
    `tinyint_b` tinyint NOT NULL COMMENT '有符号tinyint',
    `tinyint_c` tinyint(3) UNSIGNED ZEROFILL NOT NULL COMMENT '无符号tinyint，显示长度3',
    `tinyint_d` tinyint(8) UNSIGNED ZEROFILL NOT NULL COMMENT '无符号tinyint，显示长度3'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '整数' ROW_FORMAT = Dynamic;

CREATE TABLE `float`  (
    `float` float(5,2) DEFAULT NULL,
    `double` double(5,2) DEFAULT NULL,
    `decimal` decimal(5,2) DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '浮点数' ROW_FORMAT = Dynamic;

INSERT INTO `float` (`float`, `double`, `decimal`) VALUES (1.23, 1.23, 1.23);
INSERT INTO `float` (`float`, `double`, `decimal`) VALUES (1.234, 1.234, 1.234);
INSERT INTO `float` (`float`, `double`, `decimal`) VALUES (1.2345, 1.2345, 1.2345);

CREATE TABLE `date` (
    `d` date DEFAULT NULL,
    `t` time DEFAULT NULL,
    `dt` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci;
INSERT INTO `date` (`d`, `t`, `dt`) VALUES (NOW(), NOW(), NOW());

CREATE TABLE `string` (
    `c` char(4) DEFAULT NULL,
    `vc` varchar(4) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci;
INSERT INTO `string` (`c`, `vc`) VALUES ('ab  ', 'ab  ');
SELECT length(c), length(vc) FROM `string`;
SELECT CONCAT(c, '+'), CONCAT(vc, '+') FROM `string`;

