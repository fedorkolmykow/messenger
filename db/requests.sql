`SELECT 
	Orders.order_id, 
	Seller.phone, 
	Seller.name,
	Buyer.phone,
	Buyer.name,
	Courier.phone,
	Courier.name,
	City.city,
	Start_add.address,
	End_add.address,
	Notice.title,
	Notice.price,
	Orders.delivery_price
FROM Orders
LEFT JOIN Client AS Buyer ON Buyer.client_id = Orders.buyer_id
LEFT JOIN Courier ON Courier.courier_id = Orders.courier_id
LEFT JOIN Notice ON Notice.notice_id = Orders.notice_id
LEFT JOIN Client AS Seller ON Notice.seller_id = Seller.client_id
LEFT JOIN Address AS Start_add ON Start_add.address_id = Notice.start_address_id
LEFT JOIN Address AS End_add ON End_add.address_id = Orders.end_address_id
LEFT JOIN City ON End_add.city_id = City.city_id WHERE Orders.order_id = $1;`

`SELECT 
	Orders.order_id,  
	Seller.name,
	Buyer.name,
	Courier.name,
	Notice.title,
	Notice.price,
	Orders.delivery_price
FROM Orders
LEFT JOIN Client AS Buyer ON Buyer.client_id = Orders.buyer_id
LEFT JOIN Courier ON Courier.courier_id = Orders.courier_id
LEFT JOIN Notice ON Notice.notice_id = Orders.notice_id
LEFT JOIN Client AS Seller ON Notice.seller_id = Seller.client_id
WHERE Seller.client_id = $1;`

`INSERT INTO Orders (courier_id, buyer_id, end_address_id, notice_id, delivery_price)
	VALUES	
	($1, $2, $3, $4, $5);`
	
`SELECT Courier.courier_id, 0 as co FROM Courier
WHERE courier_id not in (SELECT courier_id FROM Orders)
UNION
SELECT courier_id, count(*) as co
FROM Orders
group by courier_id
order by co asc
limit 1;`

`SELECT seller_id From Notice
WHERE notice_id = $1;`

`SELECT `
