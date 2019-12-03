# Formatting
H=$(shell tput -Txterm setaf 3; tput bold)
B=$(shell tput bold; tput smul)
X=$(shell tput sgr0)

# Removes leading zero from given day
SHORT_DAY := $(shell echo ${DAY} | awk 'sub(/^0+/, "", $$1)')
YEAR ?= 2019

default: setup

setup: download day${DAY}.go

## Downloads the instructions and inputs for a day (e.g. make DAY=02)
download: challenges/day${DAY}.md inputs/day${DAY}.txt

day${DAY}.go:
	@echo "${H}=== Copying template for day ${SHORT_DAY} ===${X}"
	@sed -e "s/!DAY!/${paddedday}/g" -e "s/MAIN/main/" template/template.go > day${DAY}.go

inputs/day${DAY}.txt:
	@[ "${AOC_COOKIE}" ] || ( echo "AOC_COOKIE is not set, please specify your Advent of Code session cookie in order to download challenge and input files"; exit 1 )
	@echo "${H}=== Downloading input for day ${SHORT_DAY} ===${X}"
	@curl -s -b "session=${AOC_COOKIE}" https://adventofcode.com/${YEAR}/day/${SHORT_DAY}/input > inputs/day${DAY}.txt

challenges/day${DAY}.md: challenges/day${DAY}.html
	@echo "${H}=== Parsing input ===${X}"
	@./scripts/parse_challenge.sh ${DAY}

## The AOC_COOKIE environment variable should contain a complete session cookie in order to be able to use the make download target
challenges/day${DAY}.html:
	@[ "${AOC_COOKIE}" ] || ( echo "AOC_COOKIE is not set, please specify your Advent of Code session cookie in order to download challenge and input files"; exit 1 )
	@echo "${H}=== Downloading challenge for day ${SHORT_DAY} ===${X}"
	@curl -s -b "session=${AOC_COOKIE}" https://adventofcode.com/${YEAR}/day/${SHORT_DAY} > challenges/day${DAY}.html

## Print this message
help:
	@./scripts/help.sh $(abspath $(lastword $(MAKEFILE_LIST)))
	
## Set the AOC_COOKIE environment variable (make cookie SESSION=02))
cookie:
	@set -Ux AOC_COOKIE ${SESSION}
