package entity

import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)
func TestPositive(t *testing.T)  {
	g := NewGomegaWithT(t)
	e := Employee{
		Name: "supaporn" ,
		Email: "supaorn@gmail.com" ,
		EmployeeID: "J12345678",
	}
	ok,err := govalidator.ValidateStruct(e)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
	
}
