package cf

import (
	. "github.com/vito/cmdtest/matchers"
	. "github.com/onsi/gomega"
)

func CfAsUser(user User, actions func() error) error {
	defer func() {
		Expect(Cf("logout")).To(ExitWith(0))
	}()

	if Expect(Cf("login", user.Username, user.Password)).To(ExitWith(0)) {
		return actions()
	}

	return nil
}
