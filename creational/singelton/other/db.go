package other

var instance *db

type db struct {
}

// NewDatabase is using singelton
func NewDB() *db {
	if instance == nil {
		instance = &db{}
	}
	return instance
}

func (d db) SayHello(s string) string {
	return "Hallo, " + s
}
