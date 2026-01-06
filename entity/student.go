package entity

type Student struct {
	Name        string  `valid:"required~Name is required"`
	Price       float64 `valid:"required~Price must be greater than 0,range(0|999999)~Price must be greater than 0"`
	Stock       int     `valid:"range(0|999999)~Stock cannot be nagative"`
	Description string  `valid:"stringlength(10|999)~Description must be at least 10 characters"`
}
	 