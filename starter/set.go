package starter

// Set change setting of starter name `starterName`
func Set(starterName string) {
	m.Lock()
	defer m.Unlock()
	isStarter(starterName)

	starters[starterName].Set()
}
