package entity

import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)
func TestCustomerEmail(t *testing.T) {
	g := NewGomegaWithT(t)

	customer := Customer{
		Name: "Jane Doe",
		Email: "invalid-email",
		CustomerID: "L1234567",
	}

	ok, err := govalidator.ValidateStruct(customer)

	g.Expect(ok).To(BeFalse())
	g.Expect(err).NotTo(BeNil())
}