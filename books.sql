-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 08, 2023 at 08:25 AM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_crud`
--

-- --------------------------------------------------------

--
-- Table structure for table `books`
--

CREATE TABLE `books` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `judul` longtext DEFAULT NULL,
  `nama_penulis` longtext DEFAULT NULL,
  `tahun` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `books`
--

INSERT INTO `books` (`id`, `created_at`, `updated_at`, `deleted_at`, `judul`, `nama_penulis`, `tahun`) VALUES
(1, '2023-04-08 14:12:10.620', '2023-04-08 14:12:24.751', NULL, 'To Kill a Mockingbird', 'Harper Lee', '1960'),
(2, '2023-04-08 14:13:24.110', '2023-04-08 14:13:24.110', NULL, '1984', 'George Orwell', '1949'),
(3, '2023-04-08 14:13:53.064', '2023-04-08 14:13:53.064', NULL, 'The Great Gatsby', 'F. Scott Fitzgerald', '1925'),
(4, '2023-04-08 14:14:47.809', '2023-04-08 14:14:47.809', NULL, 'Moby Dick', 'Herman Melville', '1851'),
(5, '2023-04-08 14:15:47.866', '2023-04-08 14:15:47.866', NULL, 'War and Peace', 'Leo Tolstoy', '1869'),
(6, '2023-04-08 14:17:49.786', '2023-04-08 14:17:49.786', NULL, 'The Odyssey', 'Homer', '8th century BC'),
(7, '2023-04-08 14:18:43.862', '2023-04-08 14:18:43.862', NULL, 'The Adventures of Huckleberry Finn', 'Mark Twain', '1884'),
(8, '2023-04-08 14:19:20.298', '2023-04-08 14:19:20.298', NULL, 'Jane Eyre', 'Charlotte Bronte', '1847'),
(9, '2023-04-08 14:19:45.439', '2023-04-08 14:19:45.439', NULL, 'Animal Farm', 'George Orwell', '1945'),
(10, '2023-04-08 14:20:39.003', '2023-04-08 14:20:39.003', NULL, 'Wuthering Heights', 'Emily Bronte', '1847'),
(11, '2023-04-08 14:21:06.665', '2023-04-08 14:21:06.665', NULL, 'Catch-22', 'Joseph Heller', '1961'),
(12, '2023-04-08 14:21:34.674', '2023-04-08 14:21:34.674', NULL, 'Heart of Darkness', 'Joseph Conrad', '1899'),
(13, '2023-04-08 14:22:02.450', '2023-04-08 14:22:02.450', NULL, 'The Picture of Dorian Gray', 'Oscar Wilde', '1890'),
(14, '2023-04-08 14:22:25.625', '2023-04-08 14:22:25.625', NULL, 'The Stranger', 'Albert Camus', '1942'),
(15, '2023-04-08 14:22:52.857', '2023-04-08 14:22:52.857', NULL, 'Brave New World', 'Aldous Huxley', '1932'),
(16, '2023-04-08 14:23:14.975', '2023-04-08 14:23:14.975', NULL, 'The Canterbury Tales', 'Geoffrey Chaucer', '1387'),
(17, '2023-04-08 14:23:42.771', '2023-04-08 14:23:42.771', NULL, 'The Sun Also Rises', 'Ernest Hemingway', '1926'),
(18, '2023-04-08 14:24:11.760', '2023-04-08 14:24:11.760', NULL, 'Crime and Punishment', 'Fyodor Dostoevsky', '1866'),
(19, '2023-04-08 14:24:40.577', '2023-04-08 14:24:40.577', NULL, 'The Brothers Karamazov', 'Fyodor Dostoevsky', '1880');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `books`
--
ALTER TABLE `books`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_books_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `books`
--
ALTER TABLE `books`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=20;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
