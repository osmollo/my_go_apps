package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const (
	letters  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	digits   = "0123456789"
	specials = "~=+%^*/()[]{}/!@#$?|"
)

func randomBase64String(size int, disable_special bool) string {
	all := letters + digits
	rand.Seed(time.Now().UnixNano())
	buf := make([]byte, size)
	buf[0] = digits[rand.Intn(len(digits))]
	if !disable_special {
		buf[1] = specials[rand.Intn(len(specials))]
		all += specials
	}
	for i := 2; i < size; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	for i := len(buf) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		buf[i], buf[j] = buf[j], buf[i]
	}
	str := string(buf)
	return str
}

func main() {
	var (
		size            = flag.Int("size", 24, "Password size")
		disable_special = flag.Bool("disable_special", false, "Disable special characters (default: false)")
	)

	flag.Parse()
	fmt.Println(randomBase64String(*size, *disable_special))
}
