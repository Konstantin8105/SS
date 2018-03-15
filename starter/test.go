package starter

import "fmt"

// Test run test of starter name `starterName`
func Test(starterName string) {
	m.Lock()
	defer m.Unlock()
	isStarter(starterName)

	err := starters[starterName].Test()
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
}
