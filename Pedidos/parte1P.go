package pedidos

import (
	"fmt"
	"time"
)

type Pedidos struct {
	pedidoID    int
	userID      int
	fechaPedido time.Time
	valorTotal  int
}

func (u *Pedidos) CrearOrden(pedidoID, userID int, valorTotal int) {
	u.pedidoID = pedidoID
	u.userID = userID
	u.fechaPedido = time.Now()
	u.valorTotal = valorTotal

	fmt.Printf("Pedido creado: ID %d, Usuario %d, Fecha %s, Valor Total %d",
		u.pedidoID, u.userID, u.fechaPedido, u.valorTotal)
}

func (a *Pedidos) CancelarOrden() {
	fmt.Printf("Pedido a cancelar: ID %d, Usuario %d, Fecha %s, valor total %d",
		a.pedidoID, a.userID, a.fechaPedido, a.valorTotal)

	a.pedidoID = 0
	a.userID = 0
	a.fechaPedido = time.Time{}
	a.valorTotal = 0
}

func (d Pedidos) RevisarPedidos() {
	fmt.Printf("Revisar pedido: %d, Usuario: %d, fecha: %s, valorTotal: %d",
		d.pedidoID, d.userID, d.fechaPedido, d.valorTotal)
}
