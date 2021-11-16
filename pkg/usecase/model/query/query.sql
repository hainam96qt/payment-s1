/* name: CreateWager :exec */
INSERT INTO wager (
    id,
    odds,
    total_wager_value,
    selling_percentage,
    selling_price,
    current_selling_price,
    percentage_sold,
    amount_sold,
    placed_at
) VALUES (
    ?,?,?,?,?,?,?,?,?
);

/* name: UpdateWager :exec */
UPDATE wager set current_selling_price = ?, percentage_sold = ?, amount_sold = ?
WHERE id = ?;

/* name: CreateBuyWagerLog :exec */
INSERT INTO buy_wager_log (
    id,
    wager_id,
    buying_price,
    bought_at
) VALUES (
    ?,?,?,?
);

/* name: ListWagers :many */
SELECT * FROM wager LIMIT ? OFFSET ?;

/* name: GetWager :one */
SELECT * FROM wager where odds = ? and total_wager_value = ? order by id desc LIMIT 1;

/* name: GetBuyWager :one */
SELECT * FROM buy_wager_log where wager_id = ? order by id desc LIMIT 1;

/* name: GetWagerById :one */
SELECT * FROM wager where id = ?;
