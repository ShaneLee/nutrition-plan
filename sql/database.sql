CREATE SCHEMA `nutrition` ;


CREATE DATABASE  IF NOT EXISTS `nutrition`;
USE `nutrition`;

CREATE TABLE `nutrition`.`plan` (
  `meal_id` INT NOT NULL AUTO_INCREMENT,
  `time_submitted` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `day_no` INT NULL,
  `meal_no` INT NULL,
  `name` VARCHAR(255) NULL,
  `link` VARCHAR(255) NULL,
  `calories` INT NULL,
  PRIMARY KEY (`meal_id`));


