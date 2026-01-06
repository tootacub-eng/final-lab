package entity

import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)
func TestCustomerNameNegative(t *testing.T) {
	g := NewGomegaWithT(t)

	customer := Customer{
		Name: "",
		Email: "jd@gmail.com",
		CustomerID: "M1234567",
	}

	ok, err := govalidator.ValidateStruct(customer)

	g.Expect(ok).To(BeFalse())
	g.Expect(err).NotTo(BeNil())
}