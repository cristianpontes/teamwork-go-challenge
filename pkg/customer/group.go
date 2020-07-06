package customer

type GroupedCustomers map[string]Customers

type GroupStrategy interface {
	Execute(customers Customers) map[string]Customers
}

type DomainEmailGroupStrategy struct{}

func (d *DomainEmailGroupStrategy) Execute(customers Customers) GroupedCustomers {
	group := make(GroupedCustomers)

	for _, customer := range customers {
		domainEmail := customer.Email.Domain()

		group[domainEmail] = append(group[domainEmail], customer)
	}

	return group
}
