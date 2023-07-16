package main

type Database struct {
	user string
}

type Server struct {
	db *Database // uintptr -> 8 bytes long
}

func (s *Server) GetUserFromDB() string {
	// golang is going to "dereference" the db pointer
	// its going to lookup the memory adress of the pointer
	// it there is no db initialized in the server, it will throw
	// invalid memory address or nil pointer dereference.
	//To avoid this, handle it ourselves
	if s.db == nil {
		panic("database is not initialized")
	}
	return s.db.user
}

func main() {
	s := &Server{
		db: &Database{},
	}
	s.GetUserFromDB()
}
