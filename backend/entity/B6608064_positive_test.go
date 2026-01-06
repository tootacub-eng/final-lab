package entity

import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)
func TestCustomerPositive(t *testing.T) {
	g:= NewGomegaWithT(t)

	customer := Customer{
		Name: "John Doe",
		Email: "john.doe@example.com",
		CustomerID: "CM1234567",
	}

	ok, err := govalidator.ValidateStruct(customer)

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}