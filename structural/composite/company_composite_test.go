package composite

import (
	"design-patterns/structural/composite/other"
	"design-patterns/structural/composite/other/company_composite"
	"fmt"
	"testing"
)

func TestCompanyComposite(t *testing.T) {

	kroco1 := &company_composite.BabuKroco{}
	kroco2 := &company_composite.BabuKroco{}
	kroco3 := &company_composite.BabuKroco{}
	kroco4 := &company_composite.BabuKroco{}
	kroco5 := &company_composite.BabuKroco{}
	kroco6 := &company_composite.BabuKroco{}
	kroco7 := &company_composite.BabuKroco{}
	kroco8 := &company_composite.BabuKroco{}
	kroco9 := &company_composite.BabuKroco{}

	manager1 := &company_composite.Manager{
		Subordinates: []other.Employee{
			kroco1, kroco2, kroco3, kroco4,
		},
	}

	manager2 := &company_composite.Manager{
		Subordinates: []other.Employee{
			kroco5, kroco6, kroco7,
		},
	}

	manager3 := &company_composite.Manager{
		Subordinates: []other.Employee{
			kroco8, kroco9,
		},
	}

	ceo := &company_composite.CEO{
		Subordinates: []other.Employee{
			manager1, manager2, manager3,
		},
	}

	fmt.Printf("Total salary under Manager 1 is %d\n", manager1.TotalDivisionSalary())
	fmt.Printf("Total salary under Manager 2 is %d\n", manager2.TotalDivisionSalary())
	fmt.Printf("Total salary under Manager 3 is %d\n", manager3.TotalDivisionSalary())
	fmt.Printf("Total salary all employee in company is %d\n", ceo.TotalDivisionSalary())
}
