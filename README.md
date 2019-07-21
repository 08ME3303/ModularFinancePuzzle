# ModularFinancePuzzle

Puzzle:
Write something in any language that uses our FX api to figure out the median exchange rate for USD to SEK during last year,
This should be done by iteration over the dates one at the time, regard it as a stream. ie. do NOT sort and select middle item,
Make it in O(N*log(N))

FX api can be found at http://fx.modfin.se

This repo consists of GoLang code for retrieving and parsing JSON data from the API, and running the data stream through min-max-heap to deduce the median of the stream so far. The repo also contains a unit testing script for the code. 

Run the script as *go run modular_finance_1.go* and running the unit tests as *go test -v*
