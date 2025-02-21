package main

import "fmt"

type Product struct{}

func (p Product) getProduct(name string) {
	fmt.Printf("%sを取得しました\n", name)
}

type Payment struct{}

func (p Payment) makePayment(name string) {
	fmt.Printf("%sの支払いが完了しました\n", name)
}

type Invoice struct{}

func (i Invoice) sendInvoice(name string) {
	fmt.Printf("%sの請求書が送信されました\n", name)
}

type Order struct{}

func (o Order) placeOrder(name string) {
	fmt.Println("注文開始")

	product := Product{}
	product.getProduct(name)

	payment := Payment{}
	payment.makePayment(name)

	invoice := Invoice{}
	invoice.sendInvoice(name)

	fmt.Println("注文は正常に完了しました")
}

func main() {
	name := "デザインパターン本"
	order := Order{}
	order.placeOrder(name)
}
