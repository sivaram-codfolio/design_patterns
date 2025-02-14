package facade

import "fmt"

type CoffeeMachine struct {
	BeanAmount   float32
	GrinderLevel int
	WaterTemp    int
	WaterAmt     float32
	MilkAmount   float32
	AddFoam      bool
}

func (c *CoffeeMachine) startCoffee(beanAmount float32, grind int) {
	c.BeanAmount = beanAmount
	c.GrinderLevel = grind
	fmt.Println("Starting coffee order with beans:", beanAmount, "and grind level", grind)
}

func (c *CoffeeMachine) endCoffee() {
	fmt.Println("Ending coffee order")
}

func (c *CoffeeMachine) grindBeans() bool {
	fmt.Println("Grinding the beans:", c.BeanAmount, "beans at", c.GrinderLevel, "granularity")
	return true
}

func (c *CoffeeMachine) useMilk(amount float32) float32 {
	fmt.Println("Adding milk:", amount, "oz")
	c.MilkAmount = amount
	return amount
}

func (c *CoffeeMachine) useHotWater(amount float32) float32 {
	fmt.Println("Adding hot water:", amount)
	c.WaterAmt = amount
	return amount
}

func (c *CoffeeMachine) doFoam(useFoam bool) {
	fmt.Println("Foam setting:", useFoam)
	c.AddFoam = useFoam
}

// func (c *CoffeeMachine) useWaterTemp(temp int) int {
// 	fmt.Println("Setting water temp:", temp)
// 	c.WaterTemp = temp
// 	return temp
// }
