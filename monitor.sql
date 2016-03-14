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
-- Table structure for table `monitor`
--

DROP TABLE IF EXISTS `monitor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `monitor` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `value` varchar(3200) NOT NULL DEFAULT '',
  `agent_id` varchar(255) NOT NULL DEFAULT '',
  `job_name` varchar(255) NOT NULL DEFAULT '',
  `index` int(11) NOT NULL DEFAULT '0',
  `job_state` varchar(255) NOT NULL DEFAULT '',
  `updated` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `monitor`
--

LOCK TABLES `monitor` WRITE;
/*!40000 ALTER TABLE `monitor` DISABLE KEYS */;
INSERT INTO `monitor` VALUES (1,'{\n	    \"value\": {\n	        \"properties\": {\n	            \"logging\": {\n	                \"max_log_file_size\": \"\"\n	            }\n	        },\n	        \"job\": {\n	            \"name\": \"postgres\",\n	            \"release\": \"\",\n	            \"template\": \"postgres\",\n	            \"version\": \"5\",\n	            \"sha1\": \"0228430151570ea1e6cbb0c26c79b5ceaa83ef4d\",\n	            \"blobstore_id\": \"f26d465e-55fb-4efa-9b6f-3c03e3c06bba\",\n	            \"templates\": [\n	                {\n	                    \"name\": \"postgres\",\n	                    \"version\": \"5\",\n	                    \"sha1\": \"0228430151570ea1e6cbb0c26c79b5ceaa83ef4d\",\n	                    \"blobstore_id\": \"f26d465e-55fb-4efa-9b6f-3c03e3c06bba\"\n	                }\n	            ]\n	        },\n	        \"packages\": {\n	            \"common\": {\n	                \"name\": \"common\",\n	                \"version\": \"6.1\",\n	                \"sha1\": \"41bbc068713cc4a2de2433a3b8d8f03ce399c442\",\n	                \"blobstore_id\": \"17e30dc6-434b-4d5a-5cda-6203395d0333\"\n	            },\n	            \"postgres\": {\n	                \"name\": \"postgres\",\n	                \"version\": \"5.1\",\n	                \"sha1\": \"92a88bb74c02c9fe8459590ef695dbc6f4d413f0\",\n	                \"blobstore_id\": \"5780870b-09f3-46d8-6bb2-b96f52ed3529\"\n	            }\n	        },\n	        \"configuration_hash\": \" \",\n	        \"networks\": {\n	            \"cf12\": {\n	                \"cloud_properties\": {\n	                    \"net_id\": \"53f22403-b386-4158-a30f-07f0b8cc2ea7\"\n	                },\n	                \"default\": [\n	                    \"dns\",\n	                    \"gateway\"\n	                ],\n	                \"dns\": [\n	                    \"192.168.138.11\",\n	                    \"10.10.245.59\"\n	                ],\n	                \"dns_record_name\": \"0.postgres.cf12.cf-release.microbosh\",\n	                \"ip\": \"192.168.138.22\",\n	                \"netmask\": \"255.255.255.0\"\n	            }\n	        },\n	        \"resource_pool\": {\n	            \"cloud_properties\": {\n	                \"availability_zone\": \"bigdata\",\n	                \"instance_type\": \"flavor_737\"\n	            },\n	            \"name\": \"normal\",\n	            \"stemcell\": {\n	                \"name\": \"bosh-openstack-kvm-ubuntu-lucid-go_agent\",\n	                \"version\": \"2719\"\n	            }\n	        },\n	        \"deployment\": \"cf-release\",\n	        \"index\": 0,\n	        \"persistent_disk\": 30720,\n	        \"persistent_disk_pool\": {\n	            \"cloud_properties\": {},\n	            \"disk_size\": 30720,\n	            \"name\": \"e629ad1c-48ea-44ca-82fe-02a023c50819\"\n	        },\n	        \"rendered_templates_archive\": {\n	            \"sha1\": \"36f90aaa68ef2ad08e3e8a3b2e2d01490e8ad02a\",\n	            \"blobstore_id\": \"a43cbd05-4917-47dc-968d-f316063d078b\"\n	        },\n	        \"agent_id\": \"6b3733ee-7b6f-4d65-9b68-d3d08ddba267\",\n	        \"bosh_protocol\": \"1\",\n	        \"job_state\": \"running\",\n	        \"vitals\": {\n	            \"cpu\": {\n	                \"sys\": \"0.0\",\n	                \"user\": \"0.1\",\n	                \"wait\": \"0.0\"\n	            },\n	            \"disk\": {\n	                \"ephemeral\": {\n	                    \"inode_percent\": \"0\",\n	                    \"percent\": \"1\"\n	                },\n	                \"persistent\": {\n	       ','6b3733ee-7b6f-4d65-9b68-d3d08ddba267','postgres',0,'running','2016-03-02 16:17:11');
/*!40000 ALTER TABLE `monitor` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-03-14 10:38:45
