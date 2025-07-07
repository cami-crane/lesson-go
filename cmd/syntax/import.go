package syntax

import (
	"fmt"
	"os/user"
	"time"
)	

func Import() {
	fmt.Println(time.Now())
	fmt.Println(user.Current())
}