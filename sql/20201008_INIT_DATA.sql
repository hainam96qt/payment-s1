create table wager (
    id int primary key AUTO_INCREMENT,
    odds bigint,
    total_wager_value bigint,
    selling_percentage bigint,
    selling_price DOUBLE,
    current_selling_price DOUBLE,
    percentage_sold DOUBLE,
    amount_sold bigint,
    placed_at datetime
);

create table buy_wager_log (
    id int primary key AUTO_INCREMENT,
    wager_id int,
    buying_price DOUBLE,
    bought_at datetime
)
