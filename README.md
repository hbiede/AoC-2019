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

## Stats
| Day  | Part 1 Completion Time (Rank) | Part 2 Completion Time (Rank) | Part 2 Completion Time (in Minutes) | Note |
|-----:|-------------------------------|-------------------------------|-------------------------------|------|
| 1    | 00:37:34 (1279)               | 01:12:42 (1174)               | 35.13333333                   |      |
| 2    | 00:45:55 (3165)               | 00:59:22 (2722)               | 13.45                         |      |
| 3    | 00:26:33 (1987)               | 00:39:31 (1841)               | 12.96666667                   |      |
| 4    | 24:58:45 (25507)              | 25:32:45 (23574)              | 34                            |      |
| 5    | 01:59:30 (3248)               | 02:37:27 (2997)               | 37.95                         |      |
| 6    | 20:59:10 (15667)              | 21:12:42 (14373)              | 13.53333333                   |      |
| 7    | 02:34:51 (3481)               | 03:18:43 (2031)               | 43.86666667                   | Required Multithreading for part 2 (shy of completely redoing the IntComputer) |
| 8    | 35:14:42 (15590)              | 35:49:13 (15342)              | 34.51666667                   | My first time dealing with generating images     |
| 9    | 17:46:15 (9648)               | 17:49:01 (9648)               | 2.766666667                   |      |
| 10   | 11:46:28 (7072)               | 15:58:12 (5927)               | 251.7333                      |      |
| 11   | 23:56:12 (8402)               | 24:22:53 (8209)               | 26.68333333                   |      |
| 12   | 01:26:42 (2328)               | 02:03:07 (1112)               | 36.41666667                   |      |
| 13   |                               |                               |                               |      |
| 14   |                               |                               |                               |      |
| 15   |                               |                               |                               |      |
| 16   |                               |                               |                               |      |
| 17   |                               |                               |                               |      |
| 18   |                               |                               |                               |      |
| 19   |                               |                               |                               |      |
| 20   |                               |                               |                               |      |
| 21   |                               |                               |                               |      |
| 22   |                               |                               |                               |      |
| 23   |                               |                               |                               |      |
| 24🎅 |                               |                               |                               |      |
| 25🎄 |                               |                               |                               |      |
| Avg  | 11:52:43 (8114)               | 12:37:58 (7456)               | 00:45:15                      |      |

| <img alt="Part 1 Time Stats" src="https://raw.githubusercontent.com/hbiede/AoC-2019/master/stats/Part%201%20Time%20%28minutes%29.png" width=400> | <img alt="Part 1 Rank" src="https://raw.githubusercontent.com/hbiede/AoC-2019/master/stats/Part%201%20Rank.png" width=400> |
| <img alt="Part 2 Time Stats" src="https://raw.githubusercontent.com/hbiede/AoC-2019/master/stats/Part%202%20Time%20%28minutes%29.png" width=400> | <img alt="Part 2 Rank" src="https://raw.githubusercontent.com/hbiede/AoC-2019/master/stats/Part%202%20Rank.png" width=400> |
| <img alt="Part 2 Time Stats" src="https://raw.githubusercontent.com/hbiede/AoC-2019/master/stats/Time%20Difference%20%28Time%20to%20Complete%20Part%202%29.png" width=400> | |

Note: Times are from time of challenge release, not my start time to completion time

## Scripting From [Ullaakut](https://github.com/Ullaakut/aoc19)
#### Makefile Automation
* Automatically downloads the challenge and input for the day (e.g.: make download DAY=03)
  * In order to use this target, you need to specify your session cookie from adventofcode.com in AOC_COOKIE.
  * Parses the challenge into a markdown file (adds Markdown style headers and code blocks).
  * This part still needs a bit of work, as multiline code blocks are not supported yet, and formatting (bold, italics etc.) is lost.
