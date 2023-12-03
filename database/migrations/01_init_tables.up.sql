CREATE TABLE IF NOT EXISTS `wagers` (
    `id`                    INT       AUTO_INCREMENT,
    `odds`                  INT       NOT NULL, 
    `total_wager_value`     INT       NOT NULL,
    `selling_percentage`    INT       NOT NULL,
    `selling_price`         FLOAT     NOT NULL,
    `current_selling_price` FLOAT     NOT NULL,
    `percentage_sold`       INT       NULL,
    `amount_sold`           INT       NULL,
    `placed_at`             TIMESTAMP NOT NULL,
    `updated_at`            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at`            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `wager_buys` (
    `id`                 INT       AUTO_INCREMENT,
    `wager_id`           INT       NOT NULL,
    `buying_price`       FLOAT     NOT NULL,
    `bought_at`          TIMESTAMP NOT NULL,
    `updated_at`         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at`         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`wager_id`) REFERENCES `wagers`(`id`)
);
