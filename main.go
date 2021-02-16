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

const quotepath = "https://raw.githubusercontent.com/Ilya-Meer/marcus/main/data/quotes.json"

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

	formatted.WriteString(AddBreaks(2))

	formatted.WriteString(PadString("From 'Meditations' by Marcus Aurelius \n", 5))

	formatted.WriteString(AddBreaks(2))

	for _, str := range quote.Quote {
		formatted.WriteString(PadString(str, 5))
		formatted.WriteString(AddBreaks(1))
	}

	formatted.WriteString(AddBreaks(2))

	formatted.WriteString(
		PadString(
			fmt.Sprintf(
				"💎 Book %s, Section %s 💎",
				strconv.FormatInt(int64(quote.Book), 10),
				strconv.FormatInt(int64(quote.Section), 10),
			),
			5,
		),
	)

	formatted.WriteString(AddBreaks(2))

	return formatted.String()
}
