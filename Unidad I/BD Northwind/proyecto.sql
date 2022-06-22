SELECT TOP (1000) [ProductID]
      ,[ProductName]
      ,[SupplierID]
      ,[CategoryID]
      ,[QuantityPerUnit]
      ,[UnitPrice]
      ,[UnitsInStock]
      ,[UnitsOnOrder]
      ,[ReorderLevel]
      ,[Discontinued]
  FROM [NORTHWND].[dbo].[Products]

GO

CREATE VIEW vista_productos_descontinuados AS
SELECT
  p.ProductID,
  p.ProductName,
  s.CompanyName,
  c.CategoryName,
  p.UnitPrice,
  p.UnitsInStock
FROM Products p
INNER JOIN Suppliers s
ON p.SupplierID = s.SupplierID
INNER JOIN Categories c 
ON p.CategoryID = c.CategoryID
WHERE p.Discontinued = 0

GO

select * from vista_productos_descontinuados WHERE ProductName like '%%'

UPDATE Products SET UnitPrice = 12.5, UnitsInStock = 10 WHERE ProductID = 6