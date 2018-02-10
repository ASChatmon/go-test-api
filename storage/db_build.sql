CREATE DATABASE `cpu_metrics` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `cpu_metrics`;
CREATE TABLE `cpu` (
  `cpu_id` int(11) NOT NULL AUTO_INCREMENT,
  `cpu` int(11) DEFAULT NULL,
  `vender_id` varchar(45) DEFAULT NULL,
  `family` varchar(45) DEFAULT NULL,
  `model` varchar(45) DEFAULT NULL,
  `stepping` int(11) DEFAULT '0',
  `physical_id` varchar(45) DEFAULT NULL,
  `core_id` varchar(45) DEFAULT NULL,
  `cores` int(11) DEFAULT '0',
  `model_name` varchar(45) DEFAULT NULL,
  `mhz` int(11) DEFAULT '0',
  `cache_size` int(11) DEFAULT '0',
  `timestamp` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`cpu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;
CREATE TABLE `disk` (
  `disk_id` int(11) NOT NULL AUTO_INCREMENT,
  `fstype` varchar(45) DEFAULT NULL,
  `total` bigint(20) DEFAULT '0',
  `free` bigint(20) DEFAULT '0',
  `used` bigint(20) DEFAULT '0',
  `used_percent` float DEFAULT '0',
  `inodes_total` bigint(20) DEFAULT '0',
  `inodes_used` bigint(20) DEFAULT '0',
  `inodes_free` bigint(20) DEFAULT '0',
  `inodes_used_percent` float DEFAULT '0',
  `timestamp` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`disk_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
CREATE TABLE `memory` (
  `memory_id` int(11) NOT NULL AUTO_INCREMENT,
  `total` bigint(40) DEFAULT '0',
  `available` bigint(20) DEFAULT '0',
  `used` bigint(20) DEFAULT '0',
  `percent_used` float DEFAULT '0',
  `active` bigint(20) DEFAULT '0',
  `inactive` bigint(20) DEFAULT '0',
  `wired` bigint(20) DEFAULT '0',
  `buffers` bigint(20) DEFAULT '0',
  `timestamp` varchar(45)  DEFAULT NULL,
  PRIMARY KEY (`memory_id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;
