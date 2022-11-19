package domain

type Company struct {
	id       uint
	name     string
	postcode uint
	state    string
	city     string
	street   string
	house    string
	office   string
}

func (c *Company) Id() uint {
	return c.id
}

func (c *Company) Name() string {
	return c.name
}

func (c *Company) Postcode() uint {
	return c.postcode
}

func (c *Company) State() string {
	return c.state
}

func (c *Company) City() string {
	return c.city
}

func (c *Company) Street() string {
	return c.street
}

func (c *Company) House() string {
	return c.house
}

func (c *Company) Office() string {
	return c.office
}
