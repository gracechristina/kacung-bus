package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("2JxJ0UclRDLYhslNxX1Rk7Xp")
	user, err := api.GetUserInfo("450073692785.452043493622")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}
