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

	r, err := bb.Generate(nil, *text, balaboba.Style(*style))
	checkErr(err, r.Error, r.BadQuery)

	fmt.Println(r.FullText())
}

func checkErr(err error, rerr int, bad int) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if rerr != 0 {
		fmt.Println("response code:", rerr)
		os.Exit(1)
	}

	if bad != 0 {
		fmt.Println(balaboba.BadQueryRus)
		os.Exit(1)
	}
}
