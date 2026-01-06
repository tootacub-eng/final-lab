	package entity

	import(
		"testing"
		"github.com/asaskevich/govalidator"
		. "github.com/onsi/gomega"
	)
	func TestCustomerNegative(t *testing.T) {
		g := NewGomegaWithT(t)

		customer := Customer{
			Name: "Alice",
			Email: "a@gmail.com",
			CustomerID: "X123456",
		}
		ok, err := govalidator.ValidateStruct(customer)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())

	}