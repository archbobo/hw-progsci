// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Generating random text: a Markov chain algorithm

Based on the program presented in the "Design and Implementation" chapter
of The Practice of Programming (Kernighan and Pike, Addison-Wesley 1999).
See also Computer Recreations, Scientific American 260, 122 - 125 (1989).

A Markov chain algorithm generates text by creating a statistical model of
potential textual suffixes for a given prefix. Consider this text:

	I am not a number! I am a free man!

Our Markov chain algorithm would arrange this text into this set of prefixes
and suffixes, or "chain": (This table assumes a prefix length of two words.)

	Prefix       Suffix

	"" ""        I
	"" I         am
	I am         a
	I am         not
	a free       man!
	am a         free
	am not       a
	a number!    I
	number! I    am
	not a        number!

To generate text using this table we select an initial prefix ("I am", for
example), choose one of the suffixes associated with that prefix at random
with probability determined by the input statistics ("a"),
and then create a new prefix by removing the first word from the prefix
and appending the suffix (making the new prefix is "am a"). Repeat this process
until we can't find any suffixes for the current prefix or we exceed the word
limit. (The word limit is necessary as the chain table may contain cycles.)

Our version of this program reads text from standard input, parsing it into a
Markov chain, and writes generated text to standard output.
The prefix and output lengths can be specified using the -prefix and -words
flags on the command-line.
*/
package main

import (
	"bufio"
	// "flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
	"strconv"
)

// Prefix is a Markov chain prefix of one or more words.
type Prefix []string

// String returns the Prefix as a string (for use as a map key).
func (p Prefix) String() string {
	return strings.Join(p, " ")
}

// Shift removes the first word from the Prefix and appends the given word.
func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

// Chain contains a map ("chain") of prefixes to a list of suffixes.
// A prefix is a string of prefixLen words joined with spaces.
// A suffix is a single word. A prefix can have multiple suffixes.
type Chain struct {
	// user a nested map to store the suffix and its freqency in nested one
	chain map[string]map[string]int
	// the key of the map has already store the two strings that consist of prefix
	prefixLen int
}

// NewChain returns a new Chain with prefixes of prefixLen words.
func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string]map[string]int), prefixLen}
}

// Build reads text from the provided Reader and
// parses it into prefixes and suffixes that are stored in Chain.
func (c *Chain) Build(r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	p := make(Prefix, c.prefixLen)
	for i := 0; i < len(p); i++ {
		p[i] = "\"\""
	}
	for scanner.Scan() {
		var s string
		s = scanner.Text()
		key := p.String()
		if c.chain[key] == nil {
			c.chain[key] = make(map[string]int)
			c.chain[key][s] = 1
		} else {
			c.chain[key][s]++
		}
		p.Shift(s)
	}
}

// Write the frequency table into file
func (c *Chain) WriteFreqTable(freqTable io.Writer) {
	for key, mapEntry := range c.chain {
		fmt.Fprint(freqTable, key, " ")
		for nestKey, nextValue := range mapEntry {
			fmt.Fprint(freqTable, nestKey, " ", nextValue, " ")
		}
		fmt.Fprintln(freqTable)
	}
}

// Generate returns a string of at most n words generated from Chain.
func (c *Chain) Generate(n int) string {
	p := make(Prefix, c.prefixLen)
	var words []string
	var denom int
	var suffixString string
	var next string
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator.
	for i := 0; i < len(p); i++ {
		p[i] = "\"\""
	}

	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		for _, v := range choices {
			denom = denom + v
		}
		remainPos := rand.Intn(denom)
		// Search for the suffix with high freq
		for k, v := range choices {
			if remainPos <= v {
				suffixString = k
				break
			} else {
				remainPos = remainPos - v
			}
		}
		next = suffixString
		words = append(words, next)
		p.Shift(next)
		denom = 0
	}
	return strings.Join(words, " ")
}

// This function will turn the rest of parameter into a file list and return.
func getInfileList(argsList []string) []string {
	infileList := make([]string, 0)
	for i := 4; i < len(os.Args); i++ {
		infileList = append(infileList, argsList[i])
	}
	return infileList
}

//This function will invoke Build and WriteFreqTable to execute the read command.
func (c *Chain)commRead(args []string) {
	var outfilename string
	var infileList []string
	outfilename = args[3]
	out, err := os.Create(outfilename) // make a file to write the freq into
	if err != nil {
		fmt.Println("Sorry: couldn’t create the file!")
	}
	fmt.Fprintln(out, c.prefixLen)
	infileList = getInfileList(args)
	for _, file := range infileList {
		source, err := os.Open(file)
		if err != nil {
			fmt.Println("Sorry: couldn't open the text source file!")
		}
		c.Build(source)
		source.Close()
		c.WriteFreqTable(out)
		source.Close()
	}
	out.Close()
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Error: Please provide at least one parameter for the COMMAND!")
		return
	} else {
		switch os.Args[1] {
			// The part to execute read command.
			case "read" :{
				if len(os.Args) < 5 {
					fmt.Println("Error: Please enough parameters for read command!")
					return
				}
				prefixLen, err := strconv.Atoi(os.Args[2])
				if err != nil {
					fmt.Println("Error: Please provide a valid number for context words")
					return
				}
				c := NewChain(prefixLen)
				c.commRead(os.Args)
			}
			// The part to execute generate command.
			case "generate" :
				if (len(os.Args) != 4) {
					fmt.Println("Error: Please input appropriate number of parameters!")
					return
				}
				modelfile := os.Args[2];
				in, err := os.Open(modelfile) // make a file to write the freq into
				if err != nil {
					fmt.Println("Sorry: couldn’t open the file!")
				}
				scanner := bufio.NewScanner(in)
				scanner.Scan()
				prefixLenStr := scanner.Text()
				prefixLen, _ := strconv.Atoi(prefixLenStr)
				c := NewChain(prefixLen)
				numWords, err := strconv.Atoi(os.Args[3])
				if err != nil {
					fmt.Println("Error: Please provide an actural number for the word to output!")
					return
				}
				for scanner.Scan() {
					chainEntry := scanner.Text()
					mapEntry := strings.Split(chainEntry, " ")
					prefix := mapEntry[0:(c.prefixLen)]
					key := strings.Join(prefix, " ")
					if c.chain[key] == nil {
						c.chain[key] = make(map[string]int)
					}
					for i := c.prefixLen; i < len(mapEntry) - 2; {
						c.chain[key][mapEntry[i]], _ = strconv.Atoi(mapEntry[i + 1])
						i = i + 2
					}
				}
				text := c.Generate(numWords)
				fmt.Println(text)
			default : {
				fmt.Println("Error: No such command!")
				return
			}
		}
	}
}
