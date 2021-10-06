package main

/*Ejercicio 3 - Productos
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar
productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

Requerimientos:
Crear dos estructuras “tiendaUno” y “tiendaDos” (Atributos de la estructura y nombre de la misma a elección).
Crear una interface “Ecommerce” que tenga los métodos “Precio” y “Envio”.
Se requiere una función “nuevaTienda” que reciba el tipo de producto. Luego retorne una interface “Ecommerce”
Interface Ecommerce:
 - El método “Precio” debe retornar el precio total en base al costo del producto y los adicionales si los hubiera.
 - El método “Envio” debe retornar la dirección de entrega especificada por el cliente.
*/

type tiendaUno struct {
	nombre string  `json:"nombre"`
	precio float64 `json:"precio"`
}
type tiendaDos struct {
	nombre string  `json:"nombre"`
	precio float64 `json:"precio"`
}
type Ecommerce interface {
	Precio() float64
	Envio() string
}

const (
	pequenio = "PEQUEÑO"
	mediano  = "MEDIANO"
	grande   = "GRANDE"
)

// func nuevaTienda(tipoProducto string) Ecommerce {

// }

func main() {

}
