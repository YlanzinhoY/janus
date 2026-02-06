package operatorEntity

type Operator struct {
	Name      Name
	Email     Email
	Password  Password
	Phone     Phone
	Address   Address
	VatNumber VatNumber
}

func NewOperator(
	name Name,
	email Email,
	phone Phone,
	password Password,
	address Address,
	vatNumber VatNumber,
) (*Operator, error) {
	operator := &Operator{
		Name:      name,
		Email:     email,
		Phone:     phone,
		Address:   address,
		Password:  password,
		VatNumber: vatNumber,
	}

	return operator, nil
}

func (u *Operator) ChangeEmail(newEmail Email) error {
	u.Email = newEmail
	return nil
}

func (u *Operator) ChangePhone(newPhone Phone) error {
	u.Phone = newPhone
	return nil
}

func (u *Operator) ChangeAddress(newAddress Address) error {
	u.Address = newAddress
	return nil
}

func (u *Operator) ChangePassword(newPassword Password) error {
	u.Password = newPassword
	return nil
}
