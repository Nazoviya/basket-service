-- name: ListProducts :many
SELECT * FROM products
ORDER BY product_id;

-- name: AddToBasket :one
INSERT INTO userbasket (
    SELECT * FROM products
    WHERE products.product_id = $1
) RETURNING *;

-- name: DeleteFromBasket :exec
DELETE FROM userbasket
WHERE product_id = $1;

-- name: ShowBasket :many
SELECT * FROM userbasket
ORDER BY product_id;

-- name: CalculateBasket :one 
INSERT INTO total_basket (
    price,
    vat,
    total_price,
    discount
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetPrice :many
SELECT product_price
FROM userbasket;

-- name: GetVAT :many
SELECT product_vat
FROM userbasket;

-- name: CompleteOrder :exec
DELETE FROM userbasket;
