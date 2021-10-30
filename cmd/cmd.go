package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/Toffee-iZt/balaboba"
)

var (
	s    = flag.Int("s", 0, "generation style")
	t    = flag.String("t", "", "text to generate")
	help = flag.Bool("help", false, "print help")
)

func main() {
	flag.Parse()

	if *help {
		fmt.Printf("%s\n\n%s\n%s\n\n", balaboba.About, balaboba.Warn1, balaboba.Warn2)

		flag.PrintDefaults()

		fmt.Println("\nСтили:")
		for s := balaboba.NoStyle; s <= balaboba.XChehov; s++ {
			fmt.Println(uint8(s), "-", s.String(), "-", s.Description())
		}

		return
	}

	text := *t
	if text == "" {
		text = strings.Join(flag.Args(), " ")
	}
	if text == "" {
		fmt.Println("write the text to generate")
		return
	}

	style := balaboba.Style(*s)

	b := balaboba.New()

	fmt.Println("please wait up to 15 seconds")

	r, err := b.Get(text, style)
	if err != nil {
		fmt.Println(err)
		return
	}

	if r.Error != 0 {
		fmt.Println("response code:", r.Error)
		return
	}

	if r.BadQuery != 0 {
		fmt.Println(balaboba.BadQuery)
		return
	}

	fmt.Printf("%s\n%s%s", style.String(), r.Query, r.Text)
}
