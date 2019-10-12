package main

import (
	"fmt"
	"strings"

	"github.com/anaskhan96/soup"
	utils "github.com/aosousa/golang-utils"
)

const (
	baseLeaguepediaURL = "https://lol.gamepedia.com"
	version            = "1.0.0"
)

// Prints the list of accepted commands
func printHelp() {
	fmt.Printf("LoL Pro Player Stats (version %s)\n", version)
	fmt.Println("Available commands:")
	fmt.Println("* -h | --help\t Prints the list of available commands")
	fmt.Println("* CHAMPION CODE SPLIT YEAR Prints the statistics of the requested champion in a given split of a given year (e.g. lol-pro-champion-stats.exe Syndra LEC Summer 2019)")
}

// Prints the current version of the application
func printVersion() {
	fmt.Printf("LoL Pro Champion Stats version %s\n", version)
}

/* Handles a request to lookup information related to a champion
 * Receives:
 * args ([]string) - Arguments passed in the terminal by the user
 */
func handleChampionOptions(args []string) {
	var (
		queryURL, champion, leagueCode, split, year string
	)

	champion, leagueCode, split, year = args[1], args[2], args[3], args[4]
	queryURL = fmt.Sprintf("%s/%s/%s_Season/%s_Season/Champion_Statistics", baseLeaguepediaURL, leagueCode, year, split)
	resp, err := soup.Get(queryURL)
	if err != nil {
		utils.HandleError(err)
	}

	printChampionStats(champion, leagueCode, split, year, resp)
}

/* Print a champion's stats in a given split of a given year.
 * Receives:
 * champion (string) - Name of the champion
 * leagueCode (string) - Code of the league (LCS, LEC, etc.)
 * split (string) - Split of the league season (Spring, Summer, Winter, etc.)
 * year (string) - Year of the season
 * document (string) - Leaguepedia page HTML document
 */
func printChampionStats(champion, leagueCode, split, year, document string) {
	var hasStats bool
	var numPickBan, percentagePickBan, numBans, numPicks, winRatio, numWins, numLosses string

	pickBanString, bannedString, pickedString := "times", "times", "times"

	doc := soup.HTMLParse(document)
	championTable := doc.Find("table", "class", "spstats")
	championTableBody := championTable.Find("tbody")
	championTableBodyRows := championTableBody.Children()

	for _, row := range championTableBodyRows {
		if row.Children()[0].NodeValue == "td" {
			if strings.TrimSpace(row.Children()[0].Children()[1].NodeValue) == champion {
				hasStats = true
				numPickBan = row.Children()[1].Text()
				if numPickBan == "1" {
					pickBanString = "time"
				}

				percentagePickBan = row.Children()[2].Text()

				numBans = row.Children()[3].Text()
				if numBans == "-" {
					numBans = "0"
				}
				if numBans == "1" {
					bannedString = "time"
				}

				numPicks = row.Children()[4].Children()[0].Text()
				if numPicks == "1" {
					pickedString = "time"
				}

				numWins = row.Children()[6].Text()
				numLosses = row.Children()[7].Text()
				winRatio = row.Children()[8].Text()
			}
		}
	}

	if hasStats {
		fmt.Printf("%s %s %s %s Stats\n", champion, leagueCode, split, year)
		fmt.Printf("* Picked/banned %s %s (%s PB)\n", numPickBan, pickBanString, percentagePickBan)
		fmt.Printf("* Picked %s %s\n", numPicks, pickedString)
		fmt.Printf("* Banned %s %s\n", numBans, bannedString)
		fmt.Printf("* %s win ratio (%sW, %sL)\n", winRatio, numWins, numLosses)
	} else {
		fmt.Printf("%s was not played in %s %s %s\n", champion, leagueCode, split, year)
	}
}
