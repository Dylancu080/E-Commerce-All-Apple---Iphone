//@autor: Dylan Curay
//@Fecha: 06 / 10 / 2024
//@Version: 1.0
//@Funcion: Creamos nuevos pedidos, con funciones de actulizar y eliminar (Esto esta solo disponible para el desarrollador web)

package models

import (
	"fmt"
	"time"
)

type Pedido struct {
	PedidoID    int
	UserID      int
	FechaPedido time.Time
	ValorTotal  int
}

type PedidoInterface interface {
	CrearOrden(pedidoID, userID int, valorTotal int)
	CancelarOrden()
	RevisarPedidos()
}

func (p *Pedido) CrearOrden(pedidoID, userID int, valorTotal int) {
	p.PedidoID = pedidoID
	p.UserID = userID
	p.FechaPedido = time.Now()
	p.ValorTotal = valorTotal

	fmt.Printf("Pedido creado: ID %d, Usuario %d, Fecha %s, Valor Total %d\n",
		p.PedidoID, p.UserID, p.FechaPedido.Format("2006-01-02 15:04:05"), p.ValorTotal)
}

func (p *Pedido) CancelarOrden() {
	fmt.Printf("Pedido a cancelar: ID %d, Usuario %d, Fecha %s, Valor Total %d\n",
		p.PedidoID, p.UserID, p.FechaPedido.Format("2006-01-02 15:04:05"), p.ValorTotal)

	p.PedidoID = 0
	p.UserID = 0
	p.FechaPedido = time.Time{}
	p.ValorTotal = 0
}

func (p *Pedido) RevisarPedidos() {
	fmt.Printf("Revisar pedido: ID %d, Usuario: %d, Fecha: %s, Valor Total: %d\n",
		p.PedidoID, p.UserID, p.FechaPedido.Format("2006-01-02 15:04:05"), p.ValorTotal)
}
