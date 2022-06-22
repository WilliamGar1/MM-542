
CREATE TABLE bitacora_precios_productos
(
    IdProducto INT,
    IdPrecio INT,
    PrecioAnterior MONEY,
    PrecioNuevo MONEY,
    FechaCambio DATETIME,
    CONSTRAINT PK_bitacora PRIMARY KEY (IdProducto, IdPrecio)
)

GO

CREATE TRIGGER tg_bitacora_precios
ON Products
AFTER UPDATE
AS
BEGIN
DECLARE @idProducto INT
DECLARE @idPrecio INT
IF UPDATE(UnitPrice)
BEGIN
    SELECT @idProducto = ProductID FROM inserted 
    IF EXISTS (select * from bitacora_precios_productos WHERE IdProducto = @idProducto)
        SELECT @idPrecio = MAX(IdPrecio)+1 FROM bitacora_precios_productos WHERE IdProducto = @idProducto
    ELSE 
        SET @idPrecio = 1
    INSERT INTO bitacora_precios_productos VALUES(
        @idProducto,
        @idPrecio,
        (SELECT UnitPrice FROM deleted),
        (SELECT UnitPrice FROM inserted),
        GETDATE()
    )
    END
END

GO

DROP TRIGGER tg_bitacora_precios