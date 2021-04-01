-- phpMyAdmin SQL Dump
-- version 4.9.0.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Apr 01, 2021 at 04:26 PM
-- Server version: 10.3.15-MariaDB
-- PHP Version: 7.3.6
USE warehouse_db;

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `warehouse_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `wh_stock_history`
--

CREATE TABLE `wh_stock_history` (
  `id` bigint(20) NOT NULL,
  `sk_id` bigint(20) NOT NULL,
  `type` tinyint(4) NOT NULL,
  `amount` int(11) NOT NULL,
  `additional_info` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `wh_stock_history`
--

INSERT INTO `wh_stock_history` (`id`, `sk_id`, `type`, `amount`, `additional_info`, `created_at`) VALUES
(1, 1, 100, 12, '{}', '2021-04-01 11:05:15'),
(2, 1, 50, 1, '{}', '2021-04-01 11:46:18');

-- --------------------------------------------------------

--
-- Table structure for table `wh_stock_keeping`
--

CREATE TABLE `wh_stock_keeping` (
  `id` bigint(20) NOT NULL,
  `product_id` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  `quantity` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `wh_stock_keeping`
--

INSERT INTO `wh_stock_keeping` (`id`, `product_id`, `created_at`, `updated_at`, `quantity`) VALUES
(1, 101, '2021-04-01 11:04:22', '2021-04-01 11:42:38', 11);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `wh_stock_history`
--
ALTER TABLE `wh_stock_history`
  ADD PRIMARY KEY (`id`),
  ADD KEY `sk_id_idx` (`sk_id`);

--
-- Indexes for table `wh_stock_keeping`
--
ALTER TABLE `wh_stock_keeping`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_id` (`product_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `wh_stock_history`
--
ALTER TABLE `wh_stock_history`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `wh_stock_keeping`
--
ALTER TABLE `wh_stock_keeping`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `wh_stock_history`
--
ALTER TABLE `wh_stock_history`
  ADD CONSTRAINT `wh_stock_history_ibfk_1` FOREIGN KEY (`sk_id`) REFERENCES `wh_stock_keeping` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;