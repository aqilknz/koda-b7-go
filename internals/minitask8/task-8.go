// payment.go
package minitask8

import "fmt"

type PaymentInterface interface {
	Pay(list []int) (string, error)
}

type Bank struct{}

func (b Bank) Pay(list []int) (string, error) {
	total := sum(list)
	if total <= 0 {
		return "", fmt.Errorf("total pembayaran Bank harus > 0")
	}
	return fmt.Sprintf("Pembayaran Bank: %d", total), nil
}

type Online struct{}

func (o Online) Pay(list []int) (string, error) {
	total := sum(list)
	if total <= 0 {
		return "", fmt.Errorf("total pembayaran online harus > 0")
	}
	return fmt.Sprintf("Pembayaran Online: %d", total), nil
}

type Fiktif struct {
	History []string
}

func (f *Fiktif) Pay(list []int) (string, error) {
	total := sum(list)
	if total <= 0 {
		return "", fmt.Errorf("total pembayaran fiktif harus > 0")
	}
	f.History = append(f.History, fmt.Sprintf("%d", total))

	return "", nil
}

// Dependencies Injections
type Checkout struct {
	payment PaymentInterface
}

// Constructor injection
func NewCheckout(payment PaymentInterface) *Checkout {
	return &Checkout{payment: payment}
}

func (c *Checkout) Process(list []int) {
	result, err := c.payment.Pay(list)

	if err != nil {
		fmt.Println(err)
		return
	}

	if result != "" {
		fmt.Println(result)
	}
}

func sum(list []int) int {
	total := 0

	for _, v := range list {
		if v > 0 {
			total += v
		}
	}

	return total
}
