package domain

type buyer struct {
	name  string
	email string
}

func NewBuyer(name, email string) Buyer {
	return &buyer{name, email}
}

func (p *buyer) Name() string {
	return p.name
}

func (p *buyer) Email() string {
	return p.email
}
