package persistence

//totoService
type EntitySaveDocIn struct {
	entity     string
	connection Connection
}

type EntitySaveDocOut struct {
	success bool
}