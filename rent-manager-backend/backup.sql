-- MySQL dump 10.13  Distrib 8.0.46, for Linux (x86_64)
--
-- Host: localhost    Database: rent-manager-db
-- ------------------------------------------------------
-- Server version	8.0.46

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `contract_statuses`
--

DROP TABLE IF EXISTS `contract_statuses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `contract_statuses` (
  `cst_id` bigint NOT NULL AUTO_INCREMENT,
  `cst_code` varchar(50) NOT NULL,
  `cst_name` varchar(100) NOT NULL,
  PRIMARY KEY (`cst_id`),
  UNIQUE KEY `uk_contract_statuses_code` (`cst_code`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `contract_statuses`
--

LOCK TABLES `contract_statuses` WRITE;
/*!40000 ALTER TABLE `contract_statuses` DISABLE KEYS */;
INSERT INTO `contract_statuses` VALUES (1,'active','Active'),(2,'finished','Finished'),(3,'cancelled','Cancelled');
/*!40000 ALTER TABLE `contract_statuses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `countries`
--

DROP TABLE IF EXISTS `countries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `countries` (
  `cou_id` bigint NOT NULL AUTO_INCREMENT,
  `cou_code` varchar(10) NOT NULL,
  `cou_name` varchar(150) NOT NULL,
  `cou_created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `cou_updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`cou_id`),
  UNIQUE KEY `uk_countries_code` (`cou_code`),
  UNIQUE KEY `uk_countries_name` (`cou_name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `countries`
--

LOCK TABLES `countries` WRITE;
/*!40000 ALTER TABLE `countries` DISABLE KEYS */;
INSERT INTO `countries` VALUES (1,'AR','Argentina','2026-06-21 09:46:06','2026-06-21 09:46:06');
/*!40000 ALTER TABLE `countries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `inflation_indexes`
--

DROP TABLE IF EXISTS `inflation_indexes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `inflation_indexes` (
  `ixi_id` bigint NOT NULL AUTO_INCREMENT,
  `ixi_period` date NOT NULL,
  `ixi_percentage` decimal(10,4) NOT NULL,
  `ixi_source` varchar(100) DEFAULT NULL,
  `ixi_notes` text,
  `ixi_created_at` datetime NOT NULL DEFAULT (now()),
  `ixi_updated_at` datetime NOT NULL DEFAULT (now()),
  PRIMARY KEY (`ixi_id`),
  UNIQUE KEY `uk_inflation_indexes_period` (`ixi_period`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inflation_indexes`
--

LOCK TABLES `inflation_indexes` WRITE;
/*!40000 ALTER TABLE `inflation_indexes` DISABLE KEYS */;
INSERT INTO `inflation_indexes` VALUES (1,'2025-12-01',2.8000,'INDEC','','2026-06-24 11:36:13','2026-06-24 11:36:13'),(2,'2026-01-01',2.9000,'INDEC','','2026-06-24 11:36:13','2026-06-24 11:36:13'),(3,'2026-02-01',2.9000,'INDEC','','2026-06-24 11:36:13','2026-06-24 11:36:13'),(4,'2026-03-01',3.4000,'INDEC','','2026-06-24 11:36:13','2026-06-24 11:36:13'),(5,'2026-04-01',2.6000,'INDEC','','2026-06-24 11:36:13','2026-06-24 11:36:13'),(6,'2026-05-01',2.1000,'INDEC','','2026-06-24 11:36:13','2026-06-24 11:36:13');
/*!40000 ALTER TABLE `inflation_indexes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `interest_calculation_types`
--

DROP TABLE IF EXISTS `interest_calculation_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `interest_calculation_types` (
  `ict_id` bigint NOT NULL AUTO_INCREMENT,
  `ict_code` varchar(50) NOT NULL,
  `ict_name` varchar(100) NOT NULL,
  PRIMARY KEY (`ict_id`),
  UNIQUE KEY `uk_interest_calculation_types_code` (`ict_code`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `interest_calculation_types`
--

LOCK TABLES `interest_calculation_types` WRITE;
/*!40000 ALTER TABLE `interest_calculation_types` DISABLE KEYS */;
INSERT INTO `interest_calculation_types` VALUES (1,'daily_from_due_day_next_day','Daily interest starting the day after the due date'),(2,'daily_from_month_first_day','Daily interest starting from the first day of the month');
/*!40000 ALTER TABLE `interest_calculation_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `owners`
--

DROP TABLE IF EXISTS `owners`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `owners` (
  `own_id` bigint NOT NULL AUTO_INCREMENT,
  `own_name` varchar(255) NOT NULL,
  `own_email` varchar(255) DEFAULT NULL,
  `own_phone` varchar(50) DEFAULT NULL,
  `own_document_number` varchar(100) DEFAULT NULL,
  `own_created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `own_updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`own_id`),
  UNIQUE KEY `uk_owners_document_number` (`own_document_number`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `owners`
--

LOCK TABLES `owners` WRITE;
/*!40000 ALTER TABLE `owners` DISABLE KEYS */;
INSERT INTO `owners` VALUES (1,'Evangelina Dellarossa','evangelinadellarossa@gmail.com','+5493512956046','30149872','2026-06-21 09:47:27','2026-06-21 09:47:27'),(2,'Federico Luis Coraglio','federicocoraglio@gmail.com','+5493515319536','30341317','2026-06-21 09:48:07','2026-06-21 09:48:07');
/*!40000 ALTER TABLE `owners` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment_statuses`
--

DROP TABLE IF EXISTS `payment_statuses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `payment_statuses` (
  `pas_id` bigint NOT NULL AUTO_INCREMENT,
  `pas_code` varchar(50) NOT NULL,
  `pas_name` varchar(100) NOT NULL,
  PRIMARY KEY (`pas_id`),
  UNIQUE KEY `uk_payment_statuses_code` (`pas_code`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment_statuses`
--

LOCK TABLES `payment_statuses` WRITE;
/*!40000 ALTER TABLE `payment_statuses` DISABLE KEYS */;
INSERT INTO `payment_statuses` VALUES (1,'pending','Pending'),(2,'paid','Paid'),(3,'late','Late'),(4,'cancelled','Cancelled');
/*!40000 ALTER TABLE `payment_statuses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `properties`
--

DROP TABLE IF EXISTS `properties`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `properties` (
  `pro_id` bigint NOT NULL AUTO_INCREMENT,
  `own_id` bigint NOT NULL,
  `pty_id` bigint NOT NULL,
  `pst_id` bigint NOT NULL,
  `cou_id` bigint NOT NULL,
  `sta_id` bigint NOT NULL,
  `pro_code` varchar(100) NOT NULL,
  `pro_title` varchar(255) NOT NULL,
  `pro_description` text,
  `pro_street` varchar(255) NOT NULL,
  `pro_street_number` varchar(50) DEFAULT NULL,
  `pro_floor` varchar(50) DEFAULT NULL,
  `pro_apartment` varchar(50) DEFAULT NULL,
  `pro_city` varchar(100) NOT NULL,
  `pro_postal_code` varchar(50) DEFAULT NULL,
  `pro_created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `pro_updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`pro_id`),
  UNIQUE KEY `uk_properties_code` (`pro_code`),
  KEY `idx_properties_owner_id` (`own_id`),
  KEY `idx_properties_type_id` (`pty_id`),
  KEY `idx_properties_status_id` (`pst_id`),
  KEY `idx_properties_country_id` (`cou_id`),
  KEY `idx_properties_state_id` (`sta_id`),
  CONSTRAINT `fk_properties_country` FOREIGN KEY (`cou_id`) REFERENCES `countries` (`cou_id`),
  CONSTRAINT `fk_properties_owner` FOREIGN KEY (`own_id`) REFERENCES `owners` (`own_id`),
  CONSTRAINT `fk_properties_state` FOREIGN KEY (`sta_id`) REFERENCES `states` (`sta_id`),
  CONSTRAINT `fk_properties_status` FOREIGN KEY (`pst_id`) REFERENCES `property_statuses` (`pst_id`),
  CONSTRAINT `fk_properties_type` FOREIGN KEY (`pty_id`) REFERENCES `property_types` (`pty_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `properties`
--

LOCK TABLES `properties` WRITE;
/*!40000 ALTER TABLE `properties` DISABLE KEYS */;
INSERT INTO `properties` VALUES (1,1,1,1,1,6,'LARIOJA1014','[Córdoba][Alberdi] Dpto La Rioja 1014','','La Rioja','1014','5','B','Córdoba','5000','2026-06-21 09:52:22','2026-06-21 09:52:22'),(2,2,1,1,1,6,'AVPUERR619','[Córdoba][Guemes] Dpto Av. Pueyrredon 619','','Avenida Pueyrredon','619','5','D','Córdoba','5000','2026-06-21 12:47:47','2026-06-21 12:48:38');
/*!40000 ALTER TABLE `properties` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `property_statuses`
--

DROP TABLE IF EXISTS `property_statuses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `property_statuses` (
  `pst_id` bigint NOT NULL AUTO_INCREMENT,
  `pst_code` varchar(50) NOT NULL,
  `pst_name` varchar(100) NOT NULL,
  PRIMARY KEY (`pst_id`),
  UNIQUE KEY `uk_property_statuses_code` (`pst_code`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `property_statuses`
--

LOCK TABLES `property_statuses` WRITE;
/*!40000 ALTER TABLE `property_statuses` DISABLE KEYS */;
INSERT INTO `property_statuses` VALUES (1,'available','Available'),(2,'rented','Rented'),(3,'maintenance','Maintenance'),(4,'inactive','Inactive');
/*!40000 ALTER TABLE `property_statuses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `property_types`
--

DROP TABLE IF EXISTS `property_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `property_types` (
  `pty_id` bigint NOT NULL AUTO_INCREMENT,
  `pty_code` varchar(50) NOT NULL,
  `pty_name` varchar(100) NOT NULL,
  PRIMARY KEY (`pty_id`),
  UNIQUE KEY `uk_property_types_code` (`pty_code`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `property_types`
--

LOCK TABLES `property_types` WRITE;
/*!40000 ALTER TABLE `property_types` DISABLE KEYS */;
INSERT INTO `property_types` VALUES (1,'apartment','Apartment'),(2,'house','House'),(3,'commercial','Commercial'),(4,'garage','Garage');
/*!40000 ALTER TABLE `property_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rent_adjustment_types`
--

DROP TABLE IF EXISTS `rent_adjustment_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rent_adjustment_types` (
  `rat_id` bigint NOT NULL AUTO_INCREMENT,
  `rat_code` varchar(50) NOT NULL,
  `rat_name` varchar(100) NOT NULL,
  PRIMARY KEY (`rat_id`),
  UNIQUE KEY `uk_rent_adjustment_types_code` (`rat_code`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rent_adjustment_types`
--

LOCK TABLES `rent_adjustment_types` WRITE;
/*!40000 ALTER TABLE `rent_adjustment_types` DISABLE KEYS */;
INSERT INTO `rent_adjustment_types` VALUES (1,'ipc_argentina','Argentine CPI');
/*!40000 ALTER TABLE `rent_adjustment_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rent_payments`
--

DROP TABLE IF EXISTS `rent_payments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rent_payments` (
  `rpa_id` bigint NOT NULL AUTO_INCREMENT,
  `rco_id` bigint NOT NULL,
  `rpa_period` date NOT NULL,
  `rpa_due_date` date NOT NULL,
  `rpa_payment_date` date DEFAULT NULL,
  `rpa_base_amount` decimal(12,2) NOT NULL,
  `rpa_suggested_adjustment_percentage` decimal(8,4) DEFAULT NULL,
  `rpa_applied_adjustment_percentage` decimal(8,4) DEFAULT NULL,
  `rpa_suggested_interest_amount` decimal(12,2) DEFAULT NULL,
  `rpa_applied_interest_amount` decimal(12,2) DEFAULT NULL,
  `rpa_total_amount` decimal(12,2) NOT NULL,
  `rpa_paid_amount` decimal(12,2) DEFAULT NULL,
  `rpa_is_paid` tinyint(1) NOT NULL DEFAULT '0',
  `rpa_notes` text,
  `rpa_created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `rpa_updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`rpa_id`),
  UNIQUE KEY `uk_rent_payments_contract_period` (`rco_id`,`rpa_period`),
  KEY `idx_rent_payments_contract_id` (`rco_id`),
  KEY `idx_rent_payments_period` (`rpa_period`),
  KEY `idx_rent_payments_due_date` (`rpa_due_date`),
  KEY `idx_rent_payments_is_paid` (`rpa_is_paid`),
  CONSTRAINT `fk_rent_payments_contract` FOREIGN KEY (`rco_id`) REFERENCES `rental_contracts` (`rco_id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rent_payments`
--

LOCK TABLES `rent_payments` WRITE;
/*!40000 ALTER TABLE `rent_payments` DISABLE KEYS */;
INSERT INTO `rent_payments` VALUES (1,1,'2024-12-01','2024-12-10','2024-12-10',280000.00,NULL,0.0000,0.00,0.00,280000.00,280000.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(2,1,'2025-01-01','2025-01-10','2025-01-10',280000.00,NULL,0.0000,0.00,0.00,280000.00,280000.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(3,1,'2025-02-01','2025-02-10','2025-02-10',280000.00,NULL,0.0000,0.00,0.00,280000.00,280000.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(4,1,'2025-03-01','2025-03-10','2025-03-10',280000.00,NULL,0.0000,0.00,0.00,280000.00,280000.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(5,1,'2025-04-01','2025-04-10','2025-04-10',280000.00,NULL,0.0000,0.00,0.00,280000.00,280000.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(6,1,'2025-05-01','2025-05-10','2025-05-10',280000.00,NULL,0.0000,0.00,0.00,280000.00,280000.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(7,1,'2025-06-01','2025-06-10','2025-06-10',280000.00,NULL,17.8571,0.00,0.00,330000.00,330000.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(8,1,'2025-07-01','2025-07-10','2025-07-10',280000.00,NULL,17.8571,0.00,0.00,330000.00,330000.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(9,1,'2025-08-01','2025-08-10','2025-08-10',280000.00,NULL,17.8571,0.00,0.00,330000.00,330000.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(10,1,'2025-09-01','2025-09-10','2025-09-10',280000.00,NULL,17.8571,0.00,0.00,330000.00,330000.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(11,1,'2025-10-01','2025-10-10','2025-10-10',280000.00,NULL,17.8571,0.00,0.00,330000.00,330000.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(12,1,'2025-11-01','2025-11-10','2025-11-10',280000.00,NULL,17.8571,0.00,0.00,330000.00,330000.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(13,1,'2025-12-01','2025-12-10','2025-12-10',280000.00,NULL,32.9246,0.00,0.00,372189.00,372189.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(14,1,'2026-01-01','2026-01-10','2026-01-10',280000.00,NULL,32.9246,0.00,0.00,372189.00,372189.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(15,1,'2026-02-01','2026-02-10','2026-02-10',280000.00,NULL,32.9246,0.00,0.00,372189.00,372189.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(16,1,'2026-03-01','2026-03-10','2026-03-10',280000.00,NULL,32.9246,0.00,0.00,372189.00,372189.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(17,1,'2026-04-01','2026-04-10','2026-04-10',280000.00,NULL,32.9246,0.00,0.00,372189.00,372189.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(18,1,'2026-05-01','2026-05-10','2026-05-10',280000.00,NULL,32.9246,0.00,0.00,372189.00,372189.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(19,1,'2026-06-01','2026-06-10','2026-06-10',280000.00,NULL,57.5000,0.00,0.00,441000.00,441000.00,1,'Historical payment import','2026-06-21 09:54:44','2026-06-21 09:54:44'),(20,2,'2026-03-01','2026-03-10','2026-03-10',450000.00,NULL,0.0000,0.00,0.00,450000.00,450000.00,1,'Historical payment import','2026-06-21 12:56:32','2026-06-21 12:56:32'),(21,2,'2026-04-01','2026-04-10','2026-04-10',450000.00,NULL,0.0000,0.00,0.00,450000.00,450000.00,1,'Historical payment import','2026-06-21 12:56:32','2026-06-21 12:56:32'),(22,2,'2026-05-01','2026-05-10','2026-05-10',450000.00,NULL,0.0000,0.00,0.00,450000.00,450000.00,1,'Historical payment import','2026-06-21 12:56:32','2026-06-21 12:56:32'),(23,2,'2026-06-01','2026-06-10','2026-06-10',450000.00,NULL,0.0000,0.00,0.00,450000.00,450000.00,1,'Historical payment import','2026-06-21 12:56:32','2026-06-21 12:56:32');
/*!40000 ALTER TABLE `rent_payments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rental_contracts`
--

DROP TABLE IF EXISTS `rental_contracts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rental_contracts` (
  `rco_id` bigint NOT NULL AUTO_INCREMENT,
  `pro_id` bigint NOT NULL,
  `ten_id` bigint NOT NULL,
  `cst_id` bigint NOT NULL,
  `ict_id` bigint NOT NULL,
  `rco_start_date` date NOT NULL,
  `rco_end_date` date NOT NULL,
  `rco_total_payments` int NOT NULL,
  `rco_monthly_amount` decimal(12,2) NOT NULL,
  `rco_deposit_amount` decimal(12,2) NOT NULL DEFAULT '0.00',
  `rco_currency` varchar(10) NOT NULL DEFAULT 'ARS',
  `rco_due_day` int NOT NULL DEFAULT '10',
  `rco_daily_interest_percentage` decimal(5,2) NOT NULL DEFAULT '0.00',
  `rco_notes` text,
  `rat_id` bigint NOT NULL,
  `rco_adjustment_frequency_months` int NOT NULL DEFAULT '4',
  `rco_created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `rco_updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`rco_id`),
  KEY `idx_rental_contracts_property_id` (`pro_id`),
  KEY `idx_rental_contracts_tenant_id` (`ten_id`),
  KEY `idx_rental_contracts_status_id` (`cst_id`),
  KEY `idx_rental_contracts_interest_type_id` (`ict_id`),
  KEY `idx_rental_contracts_adjustment_type_id` (`rat_id`),
  CONSTRAINT `fk_rental_contracts_adjustment_type` FOREIGN KEY (`rat_id`) REFERENCES `rent_adjustment_types` (`rat_id`),
  CONSTRAINT `fk_rental_contracts_interest_type` FOREIGN KEY (`ict_id`) REFERENCES `interest_calculation_types` (`ict_id`),
  CONSTRAINT `fk_rental_contracts_property` FOREIGN KEY (`pro_id`) REFERENCES `properties` (`pro_id`),
  CONSTRAINT `fk_rental_contracts_status` FOREIGN KEY (`cst_id`) REFERENCES `contract_statuses` (`cst_id`),
  CONSTRAINT `fk_rental_contracts_tenant` FOREIGN KEY (`ten_id`) REFERENCES `tenants` (`ten_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rental_contracts`
--

LOCK TABLES `rental_contracts` WRITE;
/*!40000 ALTER TABLE `rental_contracts` DISABLE KEYS */;
INSERT INTO `rental_contracts` VALUES (1,1,1,1,1,'2024-12-01','2027-11-30',36,280000.00,0.00,'ARS',10,5.00,'',1,6,'2026-06-21 09:53:40','2026-06-21 09:53:40'),(2,2,2,1,1,'2026-03-01','2028-02-29',24,450000.00,0.00,'ARS',10,5.00,'',1,4,'2026-06-21 12:50:21','2026-06-21 12:50:21');
/*!40000 ALTER TABLE `rental_contracts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
  `rol_id` bigint NOT NULL AUTO_INCREMENT,
  `rol_code` varchar(50) NOT NULL,
  `rol_name` varchar(100) NOT NULL,
  `rol_created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`rol_id`),
  UNIQUE KEY `uk_roles_code` (`rol_code`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'admin','Administrator','2026-06-21 09:46:07'),(2,'manager','Manager','2026-06-21 09:46:07');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schema_migrations`
--

DROP TABLE IF EXISTS `schema_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schema_migrations` (
  `version` bigint NOT NULL,
  `dirty` tinyint(1) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schema_migrations`
--

LOCK TABLES `schema_migrations` WRITE;
/*!40000 ALTER TABLE `schema_migrations` DISABLE KEYS */;
INSERT INTO `schema_migrations` VALUES (13,0);
/*!40000 ALTER TABLE `schema_migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `states`
--

DROP TABLE IF EXISTS `states`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `states` (
  `sta_id` bigint NOT NULL AUTO_INCREMENT,
  `cou_id` bigint NOT NULL,
  `sta_code` varchar(20) DEFAULT NULL,
  `sta_name` varchar(150) NOT NULL,
  `sta_created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `sta_updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`sta_id`),
  UNIQUE KEY `uk_states_country_name` (`cou_id`,`sta_name`),
  UNIQUE KEY `uk_states_country_code` (`cou_id`,`sta_code`),
  KEY `idx_states_country_id` (`cou_id`),
  CONSTRAINT `fk_states_country` FOREIGN KEY (`cou_id`) REFERENCES `countries` (`cou_id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `states`
--

LOCK TABLES `states` WRITE;
/*!40000 ALTER TABLE `states` DISABLE KEYS */;
INSERT INTO `states` VALUES (1,1,'C','Ciudad Autónoma de Buenos Aires','2026-06-21 09:46:06','2026-06-21 09:46:06'),(2,1,'B','Buenos Aires','2026-06-21 09:46:06','2026-06-21 09:46:06'),(3,1,'K','Catamarca','2026-06-21 09:46:06','2026-06-21 09:46:06'),(4,1,'H','Chaco','2026-06-21 09:46:06','2026-06-21 09:46:06'),(5,1,'U','Chubut','2026-06-21 09:46:06','2026-06-21 09:46:06'),(6,1,'X','Córdoba','2026-06-21 09:46:06','2026-06-21 09:46:06'),(7,1,'W','Corrientes','2026-06-21 09:46:06','2026-06-21 09:46:06'),(8,1,'E','Entre Ríos','2026-06-21 09:46:06','2026-06-21 09:46:06'),(9,1,'P','Formosa','2026-06-21 09:46:06','2026-06-21 09:46:06'),(10,1,'Y','Jujuy','2026-06-21 09:46:06','2026-06-21 09:46:06'),(11,1,'L','La Pampa','2026-06-21 09:46:06','2026-06-21 09:46:06'),(12,1,'F','La Rioja','2026-06-21 09:46:06','2026-06-21 09:46:06'),(13,1,'M','Mendoza','2026-06-21 09:46:06','2026-06-21 09:46:06'),(14,1,'N','Misiones','2026-06-21 09:46:06','2026-06-21 09:46:06'),(15,1,'Q','Neuquén','2026-06-21 09:46:06','2026-06-21 09:46:06'),(16,1,'R','Río Negro','2026-06-21 09:46:06','2026-06-21 09:46:06'),(17,1,'A','Salta','2026-06-21 09:46:06','2026-06-21 09:46:06'),(18,1,'J','San Juan','2026-06-21 09:46:06','2026-06-21 09:46:06'),(19,1,'D','San Luis','2026-06-21 09:46:06','2026-06-21 09:46:06'),(20,1,'Z','Santa Cruz','2026-06-21 09:46:06','2026-06-21 09:46:06'),(21,1,'S','Santa Fe','2026-06-21 09:46:06','2026-06-21 09:46:06'),(22,1,'G','Santiago del Estero','2026-06-21 09:46:06','2026-06-21 09:46:06'),(23,1,'V','Tierra del Fuego','2026-06-21 09:46:06','2026-06-21 09:46:06'),(24,1,'T','Tucumán','2026-06-21 09:46:06','2026-06-21 09:46:06');
/*!40000 ALTER TABLE `states` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tenants`
--

DROP TABLE IF EXISTS `tenants`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tenants` (
  `ten_id` bigint NOT NULL AUTO_INCREMENT,
  `cou_id` bigint DEFAULT NULL,
  `sta_id` bigint DEFAULT NULL,
  `ten_name` varchar(255) NOT NULL,
  `ten_email` varchar(255) DEFAULT NULL,
  `ten_phone` varchar(50) DEFAULT NULL,
  `ten_document_number` varchar(100) DEFAULT NULL,
  `ten_city` varchar(150) DEFAULT NULL,
  `ten_street` varchar(255) DEFAULT NULL,
  `ten_street_number` varchar(50) DEFAULT NULL,
  `ten_floor` varchar(50) DEFAULT NULL,
  `ten_apartment` varchar(50) DEFAULT NULL,
  `ten_postal_code` varchar(20) DEFAULT NULL,
  `ten_created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `ten_updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ten_id`),
  UNIQUE KEY `uk_tenants_document_number` (`ten_document_number`),
  KEY `idx_tenants_country_id` (`cou_id`),
  KEY `idx_tenants_state_id` (`sta_id`),
  CONSTRAINT `fk_tenants_country` FOREIGN KEY (`cou_id`) REFERENCES `countries` (`cou_id`),
  CONSTRAINT `fk_tenants_state` FOREIGN KEY (`sta_id`) REFERENCES `states` (`sta_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tenants`
--

LOCK TABLES `tenants` WRITE;
/*!40000 ALTER TABLE `tenants` DISABLE KEYS */;
INSERT INTO `tenants` VALUES (1,1,6,'Enzo Hernan Villarreal','villarrealenzohernan@gmail.com','+5493512780770','39690519','Córdoba','','','','','5000','2026-06-21 09:50:39','2026-06-21 09:51:00'),(2,1,3,'Gabriel Emmanuel Nievas','','','37307157','San Fernando del  Valle de Catamarc','Pje Cesar Carrizo','336','','','K1150','2026-06-21 12:45:15','2026-06-21 12:45:15');
/*!40000 ALTER TABLE `tenants` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `usr_id` bigint NOT NULL AUTO_INCREMENT,
  `rol_id` bigint NOT NULL,
  `usr_name` varchar(255) NOT NULL,
  `usr_email` varchar(255) NOT NULL,
  `usr_password_hash` varchar(255) NOT NULL,
  `usr_created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `usr_updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`usr_id`),
  UNIQUE KEY `uk_users_email` (`usr_email`),
  KEY `idx_users_role_id` (`rol_id`),
  CONSTRAINT `fk_users_role` FOREIGN KEY (`rol_id`) REFERENCES `roles` (`rol_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,1,'Federico Coraglio','federicocoraglio@gmail.com','$2a$12$7JR0iW6sjVUd2D.zk25ioelxRujf7wzIEKz3.wjV2nfFmb9YmSKKO','2026-06-09 10:13:00','2026-06-09 10:13:00');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-06-27 17:27:36
