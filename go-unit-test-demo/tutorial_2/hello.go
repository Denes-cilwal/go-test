package tutorial2

import "fmt"

func Hello(user string) string {
	if len(user) == 0 {
		return "david"
	} else {
		return fmt.Sprintf("Hello %v", user)
	}
}
