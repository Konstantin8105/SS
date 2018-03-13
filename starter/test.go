package starter

// Test run test of starter name `starterName`
func Test(starterName string) {
	m.Lock()
	defer m.Unlock()
	isStarter(starterName)

	starters[starterName].Test()
}
