-- MySQL dump 10.13  Distrib 5.6.28, for debian-linux-gnu (x86_64)
--
-- Host: 127.0.0.1    Database: cf_deploy_ui
-- ------------------------------------------------------
-- Server version	5.6.28-0ubuntu0.14.04.1

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
-- Table structure for table `component`
--

DROP TABLE IF EXISTS `component`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `component` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `value` varchar(255) NOT NULL DEFAULT '',
  `sid` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`,`sid`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `component`
--

LOCK TABLES `component` WRITE;
/*!40000 ALTER TABLE `component` DISABLE KEYS */;
INSERT INTO `component` VALUES (4,'csfsmp','10.162.2.146',9),(5,'csfsmpport','8080',9),(6,'mysql_ip','10.162.2.146',9),(7,'mysql_port','3306',9),(8,'mysql_user','root',9),(9,'mysql_passwd','123456',9),(10,'es_ip','10.162.2.147',9),(11,'es_port','9300',9),(12,'lognode_ip','10.162.2.147',9),(13,'lognode_port','5678',9),(14,'es_ip','10.162.2.147',11),(15,'es_port','9300',11),(16,'ccdb_ip','192.168.138.14',11),(17,'ccdb_port','5524',11),(18,'ccdb_user','ccadmin',11),(19,'ccdb_passwd','c1oudc0w',11),(20,'pattern_url','http://10.162.2.146:8080/cfWeb',11),(21,'manager_ip','10.162.2.146',13),(22,'manager_port','3306',13),(23,'manager_user','root',13),(24,'manager_passwd','123456',13),(25,'cf_target','api.ccipaas.com',12),(26,'service_ip','10.162.2.146',13),(27,'service_port','3306',13),(28,'service_user','root',13),(29,'service_password','123456',13),(30,'source_mirrors','mirrors.163.com',12),(31,'mysql_default_password','123456',12),(32,'cf_passwd','admin',12),(33,'mysql_service_guid','a7c5e16d-458f-4e43-9d2b-dca536ffc22a',12);
/*!40000 ALTER TABLE `component` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-04-22 17:29:39
