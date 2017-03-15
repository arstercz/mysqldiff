CREATE TABLE `mysql_diff` (
      `id` int(10) NOT NULL AUTO_INCREMENT,
      `host` varchar(20) DEFAULT NULL,
      `port` int(10) DEFAULT NULL,
      `db` varchar(30) NOT NULL DEFAULT 'all',
      `tag` varchar(20) NOT NULL,
      `changes` text,
      `create_time` datetime DEFAULT NULL,
      PRIMARY KEY (`id`),
      KEY `idx_cretime` (`create_time`),
      KEY `idx_host_port` (`host`,`port`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
