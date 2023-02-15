package entity

import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)
func TestNeegativeName(t *testing.T)  {
	g := NewGomegaWithT(t)
	e := Employee{
		Name: "" , //ผิด
		Email: "supaorn@gmail.com" ,
		EmployeeID: "J12345678",
	}
	ok,err := govalidator.ValidateStruct(e)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Name not to be blank"))
}