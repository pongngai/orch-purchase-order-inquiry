package _const

type OrderStatus string

const (
	RECEIVED OrderStatus = "RECEIVED"
	REJECTED OrderStatus = "REJECTED"
	ACCEPTED OrderStatus = "ACCEPTED"
)
