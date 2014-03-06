package cf

import (
	. "github.com/onsi/gomega"
	. "github.com/vito/cmdtest/matchers"
)

func AsUser(user User, actions func() error) error {
	defer func() {
		Expect(Cf("logout")).To(ExitWith(0))
	}()

	if Expect(Cf("login", user.Username, user.Password)).To(ExitWith(0)) {
		return actions()
	}

	return nil
}
