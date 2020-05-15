// count -- subcommand to print the numbers from 1 to 100
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	countCmd.Flags().BoolVarP(&useBeer, "beer", "b", false, "Use beer mode.")
	rootCmd.AddCommand(countCmd)
}

var (
useBeer = false
	countCmd = &cobra.Command{
		Use:   "count",
		Short: "Count from 1 to 100, written out longhand",
		Long:  "Count from 1 to 100. Also can count from 100 to 0 in beer mode.",
		Run:   runCount,
	}
)

var (
	ones = []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten"}

	tensPlace = []string{
		"",
		"ten",
		"twenty",
		"thrirty",
		"fourty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety"}
)

func tensHelper(n int) string {
	if n >= 16 && n <= 19 {
		return ones[n-10] + "teen"
	}
	switch n {
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	case 13:
		return "thirteen"
	case 14:
		return "fourteen"
	case 15:
		return "fifteen"
	default:
		log.Fatal("Error in tensHelper")
		return ""
	}
}

// IntToCount returns the English representation of the number n. Currently,
// this only works for n non-negative and less than 100
// Deal with errors better.
func IntToCount(n int) string {
	if 0 <= n && n <= 10 {
		return ones[n]
	}

	if 11 <= n && n < 20 {
		return tensHelper(n)
	}

	if 20 <= n && n < 100 {
		if n%10 == 0 {
			return tensPlace[n/10]
		} else {
			return tensPlace[n/10] + "-" + ones[n%10]
		}
	}

	return ""
}

func beerCount() {
	b := strings.Builder{}
	for i := 99; i >= 1; i-- {
		number := IntToCount(i)
		next := IntToCount(i - 1)
		numberCap := strings.Title(number)

		s := fmt.Sprintf(`%s bottles of beer on the wall, %s bottles of beer.
Take one down, pass it around, %s bottles of beer on the wall...%s`, numberCap, number, next, "\n\n")
		b.WriteString(s)
	}
	fmt.Println(b.String())
}

func regularCount() {
	b := strings.Builder{}
	for i := 1; i < 100; i++ {
		number := IntToCount(i)
		s := fmt.Sprintf("%s\n", number)
		b.WriteString(s)
	}
	fmt.Println(b.String())
}

func runCount(cmd *cobra.Command, args []string) {
	if useBeer {
		beerCount()
	} else {
		regularCount()
	}
}
