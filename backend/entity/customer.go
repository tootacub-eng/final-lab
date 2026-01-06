package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)
type Customer struct{
	gorm.Model
	Name string `valid:"required~Name is required"`
	Email string `valid:"email,stringLength(5|50)~Email must be a valid email address and between 5 to 50 characters"`
	CustomerID string `valid:"matches(^C[LMH][0-9]{7}$)~CustomerID must start with 'C' followed by 'L', 'M', or 'H' and then 7 digits"`
}
func init(){
	govalidator.SetFieldsRequiredByDefault(false)
}