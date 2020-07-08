-- phpMyAdmin SQL Dump
-- version 5.0.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 08 Jul 2020 pada 06.07
-- Versi server: 10.4.11-MariaDB
-- Versi PHP: 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `task_4`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `food`
--

CREATE TABLE `food` (
  `IDFood` int(11) NOT NULL,
  `Type` varchar(30) DEFAULT NULL,
  `Name` int(5) DEFAULT NULL,
  `URL_image` varchar(50) DEFAULT NULL,
  `Width_image` varchar(50) DEFAULT NULL,
  `Height_image` varchar(50) DEFAULT NULL,
  `URL_thumbnail` varchar(5) DEFAULT NULL,
  `Width_thumbnail` varchar(60) DEFAULT NULL,
  `height_thumbnail` varchar(25) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Struktur dari tabel `superpower`
--

CREATE TABLE `superpower` (
  `IDSuper` int(11) NOT NULL,
  `Name` varchar(30) DEFAULT NULL,
  `Age` int(5) DEFAULT NULL,
  `Task1` varchar(50) DEFAULT NULL,
  `Task2` varchar(50) DEFAULT NULL,
  `Task3` varchar(50) DEFAULT NULL,
  `Status` varchar(5) DEFAULT NULL,
  `Pointer` varchar(60) DEFAULT NULL,
  `Title` varchar(25) DEFAULT NULL,
  `Detail` varchar(30) DEFAULT NULL,
  `SecretIdentity` varchar(30) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `food`
--
ALTER TABLE `food`
  ADD PRIMARY KEY (`IDFood`);

--
-- Indeks untuk tabel `superpower`
--
ALTER TABLE `superpower`
  ADD PRIMARY KEY (`IDSuper`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `superpower`
--
ALTER TABLE `superpower`
  MODIFY `IDSuper` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=30;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
