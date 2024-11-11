package contextkey

// Definir tipos personalizados para las claves en el contexto
type key string

const (
	UserIDKey key = "userID"
	EmailKey  key = "email"
)
