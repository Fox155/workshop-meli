SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema prueba_go
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `prueba_go` ;

-- -----------------------------------------------------
-- Schema prueba_go
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `prueba_go` DEFAULT CHARACTER SET utf8 ;
USE `prueba_go` ;

-- -----------------------------------------------------
-- Table `prueba_go`.`Users`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `prueba_go`.`Users` ;

CREATE TABLE IF NOT EXISTS `prueba_go`.`Users` (
  `user_id` BIGINT NOT NULL,
  `dni` INT NOT NULL,
  `last_name` VARCHAR(45) NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`user_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `prueba_go`.`Purcharses`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `prueba_go`.`Purcharses` ;

CREATE TABLE IF NOT EXISTS `prueba_go`.`Purcharses` (
  `id` BIGINT NOT NULL,
  `image` TEXT NOT NULL,
  `Title` VARCHAR(45) NOT NULL,
  `Status` VARCHAR(45) NOT NULL,
  `Amount` FLOAT NOT NULL,
  `creation_date` DATE NOT NULL,
  `user_id` BIGINT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_Purcharses_Users_idx` (`user_id` ASC),
  CONSTRAINT `fk_Purcharses_Users`
    FOREIGN KEY (`user_id`)
    REFERENCES `prueba_go`.`Users` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
