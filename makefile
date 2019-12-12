# Formatting
H=$(shell tput -Txterm setaf 3; tput bold)
B=$(shell tput bold; tput smul)
X=$(shell tput sgr0)

# Removes leading zero from given day
SHORT_DAY := $(shell echo ${DAY} | awk 'sub(/^0*/, "", $$1)')
YEAR ?= 2019

default: setup

setup: download day${DAY}.go

## Downloads the instructions and inputs for a day (e.g. make DAY=02)
download: challenges/day${DAY}.md inputs/day${DAY}.txt

day${DAY}.go:
	@echo "${H}=== Copying template for day ${SHORT_DAY} ===${X}"
	@sed -e "s/!DAY!/${DAY}/g" -e "s/MAIN/main/" src/template/template.go > src/day${DAY}.go

inputs/day${DAY}.txt:
	@echo "${H}=== Downloading input for day ${SHORT_DAY} ===${X}"
	@curl -s -b cookies.txt https://adventofcode.com/${YEAR}/day/${SHORT_DAY}/input > inputs/day${DAY}.txt

challenges/day${DAY}.md: challenges/html/day${DAY}.html
	@echo "${H}=== Parsing input ===${X}"
	@./scripts/parse_challenge.sh ${DAY}

## The AOC_COOKIE environment variable should contain a complete session cookie in order to be able to use the make download target
challenges/html/day${DAY}.html:
	@echo "${H}=== Downloading challenge for day ${SHORT_DAY} ===${X}"
	@curl -s -b cookies.txt https://adventofcode.com/${YEAR}/day/${SHORT_DAY} > challenges/day${DAY}.html
	
## call `make cookie SESSION=${}`
cookie:
	@curl -c cookies.txt 'http://httpbin.org/cookies/set?session=${SESSION}'
