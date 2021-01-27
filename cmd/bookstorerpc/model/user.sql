CREATE TABLE `user`
(
    `id` int(10) unsigned  NOT NULL COMMENT 'user id',
    `user` varchar(255) NOT NULL COMMENT 'user name',
    `phone` varchar(255) NOT NULL COMMENT 'phone',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


