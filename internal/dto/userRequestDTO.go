package dto

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignup struct {
	UserLogin
	Phone string `json:"phone"`
}

type VerificationCodeInput struct {
	Code string `json:"code"`
}

type SellerInput struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Phone             string `json:"phone"`
	BankAccountNumber uint   `json:"bank_account_number"`
	SwiftCode         string `json:"swift_code"`
	PaymentType       string `json:"payment_type"`
}

type AddressInput struct {
	AddressInput1 string `json:"address1"`
	AddressInput2 string `json:"address2"`
	City          string `json:"city"`
	PostCode      uint   `json:"post_code"`
	Country       string `json:"country"`
}

type ProfileInput struct {
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	AddressInput AddressInput `json:"address"`
}
