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
-- Table structure for table `template`
--

DROP TABLE IF EXISTS `template`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `template` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `template_file` varchar(255) NOT NULL DEFAULT '',
  `description` varchar(255) NOT NULL DEFAULT '',
  `sid` bigint(20) NOT NULL DEFAULT '0',
  `file_type` varchar(255) NOT NULL DEFAULT '',
  `target_file` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `template`
--

LOCK TABLES `template` WRITE;
/*!40000 ALTER TABLE `template` DISABLE KEYS */;
INSERT INTO `template` VALUES (11,'jdk','/home/ubuntu/customservice/cfdemo/jdk1.8.0_40','jdk1.8.0_40',7,'War',''),(12,'tomcat+cfWeb','/home/ubuntu/customservice/cfdemo/tomcat','tomcat+cfdemo',7,'War',''),(13,'example','/home/ubuntu/deploy/example/example.war','example.war 应用部署包',8,'War',''),(14,'manifest','/home/ubuntu/customservice/example/manifest.yml','应用部署配置文件，描述应用部署需求与配置',8,'Template','/home/ubuntu/customservice/example/manifest.yml'),(15,'cfdemo-start','/home/ubuntu/customservice/cfdemo/cfdemo-start.sh','启动脚本',7,'War',''),(16,'cfdemo-stop','/home/ubuntu/customservice/cfdemo/cfdemo-restart.sh','重启脚本',7,'War',''),(17,'cfdemo-stop','/home/ubuntu/customservice/cfdemo/cfdemo-stop.sh','停止任务脚本',7,'War',''),(18,'jdk7','/home/ubuntu/customservice/jdk7','cfweb运行时环境jdk7',9,'War',''),(19,'cfweb','/home/ubuntu/customservice/cfweb/tomcat6','cfweb程序，部署于tomcat下',9,'War',''),(20,'cfweb-start','/home/ubuntu/customservice/cfweb/cfweb-start.sh','cfweb startup',9,'War',''),(21,'cfweb-restart','/home/ubuntu/customservice/cfweb/cfweb-restart.sh','cfweb restart',9,'War',''),(22,'cfweb-stop','/home/ubuntu/customservice/cfweb/cfweb-stop.sh','cfweb stop',9,'War',''),(23,'es_runner','/home/ubuntu/customservice/es/logdb','es主程序',10,'War',''),(24,'jdk7','/home/ubuntu/customservice/jdk7','es运行时环境',10,'War',''),(25,'csfsmp','/home/ubuntu/customservice/cfweb/templates/link.js','csfsmp地址配置',9,'Template','/home/ubuntu/customservice/cfweb/tomcat6/webapps/cfWeb/link.js'),(26,'es_start','/home/ubuntu/customservice/es/es_start.sh','es启动',10,'War',''),(27,'es_restart.sh','/home/ubuntu/customservice/es/es_restart.sh','es重启',10,'War',''),(28,'es_stop','/home/ubuntu/customservice/es/es_stop.sh','es stop',10,'War',''),(29,'es_connection','/home/ubuntu/customservice/cfweb/templates/esconnect.properties','es配置',9,'Template','/home/ubuntu/customservice/cfweb/tomcat6/webapps/cfWeb/WEB-INF/classes/esconnect.properties'),(30,'lognode','/home/ubuntu/customservice/cfweb/templates/lognode.properties','lognode配置',9,'Template','/home/ubuntu/customservice/cfweb/tomcat6/webapps/cfWeb/WEB-INF/classes/lognode.properties'),(31,'lognode','/home/ubuntu/customservice/lognode/lognode','lognode主程序',11,'War',''),(32,'es_connection','/home/ubuntu/customservice/lognode/templates/esconnect.properties','es链接配置',11,'Template','/home/ubuntu/customservice/lognode/lognode/config/esconnect.properties'),(33,'ccdb','/home/ubuntu/customservice/lognode/templates/datasource.properties','ccdb配置',11,'Template','/home/ubuntu/customservice/lognode/lognode/config/datasource.properties'),(34,'pattern_config','/home/ubuntu/customservice/lognode/templates/pattern_config.properties','日志分析配置',11,'Template','/home/ubuntu/customservice/lognode/lognode/config/pattern_config.properties'),(35,'mysql_env','/home/ubuntu/customservice/mysql/mysql_env.sh','mysql配置',12,'War',''),(36,'cfweb-database','/home/ubuntu/customservice/cfweb/templates/app-data.xml','cfweb datastore',9,'Template','/home/ubuntu/customservice/cfweb/tomcat6/webapps/cfWeb/WEB-INF/classes/app-data.xml'),(37,'cfweb_dev','/home/ubuntu/customservice/mysql/cfweb_dev.sql','cfweb数据导入',12,'War',''),(38,'service_broker','/home/ubuntu/customservice/mysql/servicebroker.sql','servicebroker导入脚本',12,'War',''),(39,'csf','/home/ubuntu/customservice/mysql/csf.sql','csf脚本',12,'War',''),(40,'csfservice','/home/ubuntu/customservice/mysql/csfservice.sql','csfservice',12,'War',''),(41,'cf-mysql-service-broker','/home/ubuntu/customservice/cf-mysql-service-broker','cf-mysql-service-broker 主程序',13,'War',''),(42,'manifest','/home/ubuntu/customservice/cf-mysql-service-broker/manifest.yml','应用配置文件',13,'Template','/home/ubuntu/customservice/cf-mysql-service-broker/manifest.yml'),(43,'sql-setting','/home/ubuntu/customservice/cf-mysql-service-broker/templates/sql-settings.yaml','service 配置文件',13,'Template','/home/ubuntu/customservice/cf-mysql-service-broker/cf-mysql-service-broker-0.0.1-SNAPSHOT/WEB-INF/classes/sql-settings.yaml'),(44,'cfweb_dev','/home/ubuntu/customservice/mysql/cfweb_dev_template.sql','cfweb_dev模板文件',12,'Template','/home/ubuntu/customservice/mysql/cfweb_dev.sql'),(45,'jdk7','/home/ubuntu/customservice/jdk7','lognode 运行环境',11,'War',''),(46,'mysql_env_template','/home/ubuntu/customservice/mysql/mysql_env_template.sh','mysql_env 配置文件',12,'Template','/home/ubuntu/customservice/mysql/mysql_env.sh'),(47,'source','/home/ubuntu/customservice/mysql/sources _template','mysql 安装源配置文件',12,'Template','/home/ubuntu/customservice/mysql/sources'),(48,'source_list','/home/ubuntu/customservice/mysql/sources','更新源文件',12,'War',''),(49,'lognode_start','/home/ubuntu/customservice/lognode/lognode_start.sh','启动脚本',11,'War',''),(50,'lognode_restart','/home/ubuntu/customservice/lognode/lognode_restart.sh','重启脚本',11,'War',''),(51,'lognode_stop','/home/ubuntu/customservice/lognode/lognode_stop.sh','关闭脚本',11,'War',''),(52,'patterns','/home/ubuntu/customservice/lognode/patterns','lognode日志分析patterns',11,'War','');
/*!40000 ALTER TABLE `template` ENABLE KEYS */;
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
