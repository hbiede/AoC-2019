# AoC-2019
[Advent of Code](adventofcode.com) Solutions for 2019. This was my first year doing Advent of Code, I used it as an opportunity to learn Go(lang).

## Highlights:

#### Favorite problems:

* Day 3's Wire Problem (Specifically Part 2 where wire distance comes into play) was a fun challenge that really tested by new knowledge of
the language. 
* Day 7's multithreading was my first time doing a real application of simultaneous calculations
* A working block breaker with a basic AI? That's so cool!!
* Loved the work put into IntCode by the creators of AoC. It was amazing what it could do.
* A full, working text-based adventure? That's amazing!

#### Interesting approaches:

* For Day 3, I originally attempted to use concurrency to step both wires simultaneously, but I was foiled by Go's built in preventions to 
inhibit concurrent map access, so I landed on a strategy of tracking wire distance along with the wire that interacted with the coordinate
before and comparing the combined steps from that. I am sure there are more efficient solutions and I have a few ideas for how I would do
it in OOP, but this is good enough for now!
* In Day 16 - Part 2, I stumbled upon a pattern that works for the second half of large numbers being manipulated in that way. This was a
truly interesting discovery for me that I did not fully understand until I spent a few hours contemplating it and the math behind it.

## Stats
STATS_TABLE

<img alt="Part 1 Time Stats" src="statsImages/part1time.png" width=400> <img alt="Part 1 Rank" src="statsImages/part1rank.png" width=400>
<img alt="Part 2 Time Stats" src="statsImages/part2time.png" width=400> <img alt="Part 2 Rank" src="statsImages/part2rank.png" width=400>

Note: Times are from time of challenge release, not my start time to completion time

## Scripting initially based on a script from [Ullaakut](https://github.com/Ullaakut/aoc19). Expanded upon by [HBiede](https://github.com/Ullaakut/hbiede)
#### Makefile Automation
* Automatically downloads the challenge and input for the day (e.g.: `make download DAY=03`)
  * In order to use this target, you need to specify your session cookie from adventofcode.com in cookies.txt through the usage of `make cookie SESSION={Insert your session cookie here}`.
  * Parses the challenge into a markdown file (adds Markdown style headers and code blocks).
  * Per Ullaakut, this part still needs a bit of work, as multiline code blocks are not supported yet, and formatting (bold, italics etc.) is lost.
* Setup the new day's source file from a template file while downloading the input and challenge per above (e.g.: `make DAY=03`)
* Create the stats table above by calling `make stats`
  * May require calling `sh scripts/install_stats_dependencies.sh` to ensure you have all the necessary python dependencies
