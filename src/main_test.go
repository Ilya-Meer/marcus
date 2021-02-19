package main

import "testing"

func TestFormatQuote(t *testing.T) {
	q1 := Quote{
		Text: []string{
			"The best revenge is not to be like that.",
		},
		Book:    6,
		Section: 6,
	}

	q2 := Quote{
		Text: []string{
			"How to act:",
			"Never under compulsion, out of selfishness, without forethought, with misgivings.",
			"Don't gussy up your thoughts.",
			"No surplus words or unnecessary actions.",
			"Let the spirit in you represent a man, an adult, a citizen, a Roman, a ruler. Taking up his post like a soldier and patiently awaiting his recall from life. Needing no oath or witness.",
			"Cheerfulness. Without requiring other people's help. Or serenity supplied by others.",
			"To stand up straight - not straightened.",
		},
		Book:    3,
		Section: 5,
	}

	expected := `


     From 'Meditations' by Marcus Aurelius


     The best revenge is not to be like that.


     💎 Book 6, Section 6 💎

`

	actual := formatQuote(q1)

	if len([]byte(expected)) != len([]byte(actual)) {
		t.Errorf("Q1 was not formatted correctly, got : %s, want: %s.", actual, expected)
	}

	expected = `

     From 'Meditations' by Marcus Aurelius


     How to act:
     Never under compulsion, out of selfishness, without forethought, with misgivings.
     Don't gussy up your thoughts.
     No surplus words or unnecessary actions.
     Let the spirit in you represent a man, an adult, a citizen, a Roman, a ruler. Taking up his post like a soldier and patiently awaiting his recall from life. Needing no oath or witness.
     Cheerfulness. Without requiring other people's help. Or serenity supplied by others.
     To stand up straight - not straightened.


     💎 Book 3, Section 5 💎

	`

	actual = formatQuote(q2)

	if len([]byte(expected)) != len([]byte(actual)) {
		t.Errorf("Q1 was not formatted correctly, got : %s, want: %s.", actual, expected)
	}
}
