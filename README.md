# lol-pro-champion-stats

**Golang application that displays champion statistics of a LoL tournament split**

### Installation

Download the latest release from the [Releases tab](https://github.com/aosousa/lol-pro-champion-stats/releases)

### Usage

Available league codes and splits

| League Code | Description | Splits |
| --- | --- | --- |
| LCS | LoL Championship Series (NA) | Spring, Summer |
| LEC | LoL European Championship | Spring, Summer |
| LCK | LoL Champions Korea | Spring, Summer | 
| LPL | LoL Pro League (China) | Spring, Summer |
| LMS | League Master Series (Taiwan) | Spring, Summer |
| CBLOL | Campeonato Brasileiro LoL (Brazil) | Winter, Summer |
| LCL | LoL Russia League | Spring, Summer |
| LJL | LoL Japan League | Spring, Summer |
| LLA | Latin America League | Opening, Closing | 
| OPL | Oceanic Pro League | Split_1, Split_2 |
| LST | LoL SEA Tour | Spring, Summer |
| TCL | Turkish Champions League | Winter, Summer |
| VCS | Vietnam Championship Series | Spring, Summer |

```
lol-pro-champion-stats.exe [-h | --help | -v | --version]
```

### Options
```
-h, --help Prints the list of available commands
-v, --version Prints the version of the application
CHAMPION CODE SPLIT YEAR Prints the statistics of a player in a given split of a given year (e.g. lol-pro-champion-stats.exe Syndra LEC Summer 2019)
```

### Examples

#### Show a champion's statistics

`$lol-pro-champion-stats.exe Syndra LEC Summer 2019`

![ScreenShot](/img/champion_stats.png)

### Contribute

Found a bug? Have a feature you'd like to see added or something you'd like to see improved? You can report it to me by [opening a new issue](https://github.com/aosousa/lol-pro-champion-stats/issues)!

### License

MIT © [André Sousa](https://github.com/aosousa)