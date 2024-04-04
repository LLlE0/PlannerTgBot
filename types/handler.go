package types

type Handler struct {
	Database *DBInstance
}

func NewHandler() *Handler {
	h := &Handler{
		Database: &DBInstance{},
	}
	h.Database.Instance = NewDBInstance()
	return h
}
