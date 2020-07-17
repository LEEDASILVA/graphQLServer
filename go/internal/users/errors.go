package users

type WrongUsernameOrPasswordError struct{}

func (ms *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}
