package storages

type Storage interface {
	Init()
	SaveMessage(msg string)
	FlushStorage()
	Close()
}
