# AoC-2019
[Advent of Code](adventofcode.com) Solutions for 2019. This is my first year doing Advent of Code, used it as an opportunity to learn Go(lang)

## Highlights

Favorite problems:

* Day 3's Wire Problem (Specifically Part 2 where wire distance comes into play) was a fun challenge that really tested by new knowledge of
the language. 

Interesting approaches:

* For Day 3, I originally attempted to use concurrency to step both wires simultaneously, but I was foiled by Go's built in preventions to 
inhibit concurrent map access, so I landed on a strategy of tracking wire distance along with the wire that interacted with the coordinate
before and comparing the combined steps from that. I am sure there are more efficient solutions and I have a few ideas for how I would do
it in OOP, but this is good enough for now!

Leaderboard appearances:

* Not likely given that I'm only just learning Go, but you never know...

## From [Ullaakut](https://github.com/Ullaakut/aoc19)
* Automatically downloads the challenge and input for the day (e.g.: make download DAY=03)
  * In order to use this target, you need to specify your session cookie from adventofcode.com in AOC_COOKIE.
  * Parses the challenge into a markdown file (adds Markdown style headers and code blocks).
  * This part still needs a bit of work, as multiline code blocks are not supported yet, and formatting (bold, italics etc.) is lost.
