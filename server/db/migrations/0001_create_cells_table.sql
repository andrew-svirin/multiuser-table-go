CREATE TABLE cells
(
    `id`     INT PRIMARY KEY AUTO_INCREMENT,
    `column` VARCHAR(255) NOT NULL,
    `row`    INT          NOT NULL,
    `value`  TEXT
);

CREATE INDEX cells_column_row ON cells (`column`, `row`);