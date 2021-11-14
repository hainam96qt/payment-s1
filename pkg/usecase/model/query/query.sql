/* name: GetPayment :one */
SELECT * FROM payment
WHERE id = ? LIMIT 1;
