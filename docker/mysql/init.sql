CREATE TABLE flerken.`dispositivo` (
                                       `id` bigint NOT NULL AUTO_INCREMENT,
                                       `id_usuario_app` bigint NOT NULL,
                                       `nome` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                       `fcm_token` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
                                       `created_at` datetime NOT NULL,
                                       `deleted_at` datetime DEFAULT NULL,
                                       `updated_at` datetime NOT NULL,
                                       `fcm_token_hash` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
                                       PRIMARY KEY (`id`),
                                       UNIQUE KEY `unique_constraint_dispositivo` (`fcm_token_hash`),
                                       KEY `dispositivo_id_usuario_app_IDX` (`id_usuario_app`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19487102 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

ALTER TABLE dispositivo ADD UNIQUE (fcm_token);


INSERT INTO flerken.dispositivo (id_usuario_app,nome,fcm_token,created_at,deleted_at,updated_at,fcm_token_hash) VALUES
                                                                                                                    (7,NULL,'c6BX1nk5Y6s:APA91bEb_D8KuCGC_ZTI2pvWIJJZdrymGwrwlWHHA7U6ieKwAK8bV7zpXKSvSDOSgTJn1GnoD49ivjoXoPiLg_P2RuEJfeG2O4PzS_ibR4Xoj0o2bGf-2UHQ7CBdGs3621m8450BdDs0','2019-08-22 18:43:45','2022-07-27 16:15:02','2019-08-22 18:43:45','0b6c3d2c62a1f448f47fb35128f2be8cf6ead27399b778030276425c9175f387'),
                                                                                                                    (114,NULL,'cwYKyetoSsM:APA91bHXXAkBfBBwC0d261-orrRi_h3vbs1cfzd8GUF8uXZTEKsKbDDx4ygq2tlpBTnareXA63yf225EEW-LEDEhVFNt12Uyupov_oxhwhQEBhD6qfEJqE1srjHrChiV7xkoz-cmCP1y','2019-08-22 12:07:35','2021-09-17 10:20:07','2019-09-03 15:55:53','9827475219b70a8f89dfa29abc4dd29f960082cd31991b199e150ae38fedb19d'),
                                                                                                                    (47,NULL,'c32g__9uzFM:APA91bGy-BimZfwlVoKMolN8cBPDYy_gznvxbeb-oiz7sQWcKLtDnqypAJpdj0vE_L8iq9iEywsVue37IOkKkhmK1dNP9aWmAjwFMqu2wAWcRLhoKCiXtwgNZVH5ajr49-S9OfuE0EQX','2019-08-19 11:12:42','2021-09-17 10:20:07','2019-10-12 14:39:56','7d5986f97376e6f680c059a1220e48f18c53abf2d4f40ac875a5633b43ec3257'),
                                                                                                                    (8,NULL,'dTqoSihLFHA:APA91bEoF6eergEyFdcqdPD6dCw1GoopD6LqrjXcLoM18uOhtckYmGGS7I6jPvPQ7AGN5IBSaztJgVUtl3KYsR5vFcQKxQYuPe2hq6-OvM6NN9B_TAAeX0JoIlRfFEOL9Ro7JDiEAb4q','2019-08-22 21:43:41','2024-02-22 20:57:22','2019-08-22 21:43:45','43c4f3f1cb275f3ec8634c7562d340efd45e204b0d3c988056c64e773982715c'),
                                                                                                                    (9,NULL,'cWL1GYSlb6A:APA91bFDJm0KnzHDpCc53JeY4Rj5kQu6oHDIwMkd0iUG80TqsWSzRpmvJZcioh4XmLandiwobAQSWWPx34eG6_6EljifrPlgYhgy9q2FwH8mTcM8AwT1pC7S_q2XH3NaK82g-atgLJez','2019-08-22 22:15:37','2025-02-03 13:56:06','2019-10-02 18:36:20','bc65b44510e29c028326f87b9d1b964ae1110bbb7d9fbac3e03c3ec8046f3771'),
                                                                                                                    (1,NULL,'dRkHyjcBkWU:APA91bFYkDAa5KkNkOATtt09B1YtKx04mPu4FTCh3yikvkDmMfrtdnY3qE0xTNRoH5sttPuMLxmRuCKELACBmbMUMuCiXJXGaLrJKgXHyvKcf8K3fZWMgBs4Z8RosXf2zvAfOBnxeWl6','2019-08-22 16:39:32','2025-02-26 07:11:18','2019-11-20 20:24:56','74c08e7c0dac66b0a973a6db943eefc3a7ade0d6e00f7fab789d36b3554bc5ef'),
                                                                                                                    (17,NULL,'dz-TQ6_0vw0:APA91bHXpfPnU_3hfPJiQODM6-SOq38JL3t408lKKVAeI7LXnzkjmB2YxbExrh8CSlff3RPurn6fGRUHgJOgkbauM_4bPCs_KWd2FteWIrJoO1ZY4ubscl1DOxL8DVGZ2gsCyMcX-jCq','2019-08-23 12:49:55','2020-10-15 18:05:15','2019-10-02 09:05:37','eb39a5e9d4f61c0afed1365a7e0079cad0ed63ab51202890f2e7eec18a079f1d'),
                                                                                                                    (12,NULL,'cyThxERTF-g:APA91bGnxya4vxbaXyZQccgW069Tqqral4__sxojx8Y0Al-i-u7Zn2EVHSiiKmUBy3Lm_tJjOBbup1Oz0upLvUbJx9_FlFRzfr_mfAu9ASl1Xyw51_PeHiDsOd_s2P_cjEhQqR6E51NB','2019-08-23 13:10:06','2020-10-15 18:05:15','2019-08-23 13:10:06','246c447607977c1e6322dbae560231d7e85a91e6fff4d5c8f7a2390e6d6c6511'),
                                                                                                                    (1,NULL,'dbGh5fXYI9E:APA91bGS0SLE2CZ0xBZS5RPtmfTDWv9cRVC80nLRFRgonSOkL6XtVoWjZxhWJveEEBjXYEKO61Jp3aWiuHlhSvLDUqfIbKFHTKNgCMgmP_wjpf9C_XZ7rWu6IhIfWbFyLUEUtHOlje-q','2019-08-23 13:10:56','2025-02-26 07:11:19','2019-08-23 16:32:59','baa6352d084b533d0f626573097ef8856b057a2c117a2f595ee0f2af771e75ac'),
                                                                                                                    (20,NULL,'eFsFhGDoMt4:APA91bHObh6Cvts8BDm4aCYCz1YZc3_YQhU5b39115S1K-xRi7JY8Hq4dPlavsxTTc3Ta_uUnfWQZb051v5X3Rc8jbeeyMzbbI9ioE0_n2aJGCJ7uqw_QtYUU7yPG8dnH9FrtIZpiiEX','2019-08-23 14:39:50','2020-10-15 18:05:15','2019-08-23 14:39:50','4ff59b16042667c75a9416aafef29c67dc4a676757a60dab0a743aabc7bfad95');

CREATE TABLE audit_log (
                           `id` bigint NOT NULL AUTO_INCREMENT,
                           `token` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
                           `reason` TEXT NOT NULL,
                           `created_at` datetime NOT NULL,
                           PRIMARY KEY (`id`)

);

ALTER TABLE `dispositivo` ADD COLUMN count_updated_token INT DEFAULT 0;

SELECT '✅ Script init.sql executado até o fim com sucesso';
