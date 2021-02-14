package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const quotepath = "https://gist.githubusercontent.com/Ilya-Meer/953bdc26f7feab9a419f0db8bae7de48/raw/0ff6c5b50c1fcc9633be26e1e329203b7ccdc1d1/marcus.json"

type Quote struct {
	Quote   []string `json:"quote"`
	Book    int      `json:"book"`
	Section int      `json:"section"`
}

func main() {
	var quotes []Quote

	r, err := http.Get(quotepath)
	if err != nil {
		log.Fatal("Encountered an error fetching quotes...")
	}

	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&quotes)

	randomQuote := getRandomQuote(quotes)

	formatted := formatQuote(randomQuote)

	fmt.Println(formatted)
}

func getRandomQuote(quotes []Quote) (quote Quote) {
	rand.Seed(time.Now().UnixNano())
	r1 := rand.Intn(len(quotes) - 1)

	return quotes[r1]
}

func formatQuote(quote Quote) (formattedQuote string) {
	var formatted strings.Builder

	formatted.WriteString("\n\n")

	for _, str := range quote.Quote {
		formatted.WriteString(fmt.Sprintf("     %s     ", str))
		formatted.WriteString("\n")
	}

	formatted.WriteString("\n\n")

	formatted.WriteString(
		fmt.Sprintf(
			"     - Book %s, Section %s",
			strconv.FormatInt(int64(quote.Book), 10),
			strconv.FormatInt(int64(quote.Section), 10),
		),
	)

	formatted.WriteString("\n\n")

	return formatted.String()
}
