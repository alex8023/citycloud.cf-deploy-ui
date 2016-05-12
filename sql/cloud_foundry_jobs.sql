-- MySQL dump 10.13  Distrib 5.6.30, for debian-linux-gnu (x86_64)
--
-- Host: 127.0.0.1    Database: cf_deploy_ui
-- ------------------------------------------------------
-- Server version	5.6.30-0ubuntu0.14.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `cloud_foundry_jobs`
--

DROP TABLE IF EXISTS `cloud_foundry_jobs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cloud_foundry_jobs` (
  `name` varchar(255) NOT NULL DEFAULT '',
  `job_name` varchar(255) NOT NULL,
  `resources_pool` varchar(255) NOT NULL DEFAULT '',
  `instances` int(11) NOT NULL DEFAULT '1',
  `static_ips` varchar(2048) NOT NULL DEFAULT '',
  PRIMARY KEY (`job_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cloud_foundry_jobs`
--

LOCK TABLES `cloud_foundry_jobs` WRITE;
/*!40000 ALTER TABLE `cloud_foundry_jobs` DISABLE KEYS */;
INSERT INTO `cloud_foundry_jobs` VALUES ('cloud_controller_clock','cloud_controller_clock','normal2',1,'192.168.129.112'),('cloud_controller_ng','cloud_controller_ng','normal2',1,'192.168.129.110'),('cloud_controller_worker','cloud_controller_worker','normal2',1,'192.168.129.111'),('dea_next','dea_next','normal2',1,'192.168.129.115'),('etcd','etcd','normal2',1,'192.168.129.106'),('gorouter','gorouter','normal2',1,'192.168.129.101'),('haproxy111111111','haproxy','normal2',1,'192.168.129.100'),('hm9000','hm9000','normal2',1,'192.168.129.113'),('loggregator','loggregator','normal2',1,'192.168.129.107'),('loggregator_traffic','loggregator_traffic','normal2',1,'192.168.129.108'),('nats','nats','normal2',1,'192.168.129.104'),('nfs','nfs','normal2',1,'192.168.129.103'),('postgres','postgres','normal2',1,'192.168.129.102'),('stats','stats','normal2',1,'192.168.129.114'),('syslog_aggregator','syslog_aggregator','normal2',1,'192.168.129.105'),('uaa','uaa','normal2',1,'192.168.129.109');
/*!40000 ALTER TABLE `cloud_foundry_jobs` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-05-12 14:38:47
