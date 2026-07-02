package port

type PasswordHasher interface {
	Hash(password string) (string, error)
	Validate(password string, hashedPassword string) (bool, error)
}
