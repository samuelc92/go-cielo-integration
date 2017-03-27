package cielo

//Enum band of card
type BrandEnum uint8

const (
	ELO BrandEnum = iota
	MASTER
	VISA
)

// EnvironmentEnum of API' environment
type EnvironmentEnum uint8

//Enum of API's environment
const (
	PRODUCTION EnvironmentEnum = iota
	TEST
)
