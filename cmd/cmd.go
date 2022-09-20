package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/karalef/balaboba"
)

var (
	style  = flag.Uint("s", 0, "generation style")
	text   = flag.String("t", "", "text to generate")
	styles = flag.Bool("styles", false, "print all available styles")
	help   = flag.Bool("help", false, "print help")
)

var bb = balaboba.ClientRus

func main() {
	flag.Parse()

	if *help {
		fmt.Printf("%s\n\n%s\n%s\n\n", balaboba.AboutRus, balaboba.Warn1Rus, balaboba.Warn2Rus)
		flag.PrintDefaults()
		return
	}

	if *styles {
		printStyles()
		return
	}

	if *text == "" {
		*text = strings.Join(flag.Args(), " ")
	}
	if *text == "" {
		fmt.Println("write the text to generate")
		return
	}

	fmt.Println("please wait up to 20 seconds")

	r, err := bb.Generate(*text, balaboba.Style(*style))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(r.Text)
}

func printStyles() {
	allStyles := [...]balaboba.Style{
		balaboba.Standart, balaboba.UserManual,
		balaboba.Recipes, balaboba.ShortStories,
		balaboba.WikipediaSipmlified, balaboba.MovieSynopses,
		balaboba.FolkWisdom,
	}
	fmt.Println("Styles:")
	for _, s := range allStyles {
		str, desc := s.Description(balaboba.Rus)
		fmt.Println(s, "-", str, "-", desc)
	}
}
