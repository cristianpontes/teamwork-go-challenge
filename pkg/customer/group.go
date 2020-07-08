package customer

// GroupedCustomers represents a set of customers that have been grouped by a common factor, ie: email
type GroupedCustomers map[string]Customers

// GroupStrategy represents the contract that any grouping strategy must follow
type GroupStrategy interface {
	Execute(customers Customers) map[string]Customers
}

// DomainEmailGroupStrategy represents a groping strategy that groups a set of customers by their email's domain address
type DomainEmailGroupStrategy struct{}

// Execute executes the strategy given a set of customers
func (d *DomainEmailGroupStrategy) Execute(customers Customers) GroupedCustomers {
	group := make(GroupedCustomers)

	for _, customer := range customers {
		domainEmail := customer.Email.Domain()

		group[domainEmail] = append(group[domainEmail], customer)
	}

	return group
}
