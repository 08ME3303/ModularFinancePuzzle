package main

import (
	"fmt"
	"container/heap"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/* Part 1: Running median calculator
            using a min heap, max heap approach*/

type IntHeap []float32 

type Median struct {
	left *MedianHeap
	right *MedianHeap
}

type MedianHeap struct {
	IntHeap
}

func Constructor() Median {
	left := new(MedianHeap)
	heap.Init(left)
	right := new(MedianHeap)
	heap.Init(right)

	return Median{
		left:  left,
		right: right,
	}
}

func (h MedianHeap) Len() int { return len(h.IntHeap) }
func (h MedianHeap) Less(i, j int) bool { return h.IntHeap[i] < h.IntHeap[j] }
func (h MedianHeap) Swap(i, j int) { h.IntHeap[i], h.IntHeap[j] = h.IntHeap[j], h.IntHeap[i] }
func (h *MedianHeap) Push(value interface{}) { h.IntHeap = append(h.IntHeap, value.(float32)) }
func (h *MedianHeap) Pop() interface{} {
	min := (h.IntHeap)[len(h.IntHeap)-1]
	h.IntHeap = (h.IntHeap)[:len(h.IntHeap)-1]
	return min
}

// Function to add incoming data to the approapriate heap
func (find *Median) AddData(n float32) {
	if find.left.Len() == find.right.Len() {
		heap.Push(find.left, n)
	} else {
		heap.Push(find.right, n)
	}
	if find.right.Len() >  0 && find.left.IntHeap[find.left.Len()-1]>find.right.IntHeap[0] {
		find.left.IntHeap[find.left.Len()-1], find.right.IntHeap[0] = find.right.IntHeap[0], find.left.IntHeap[find.left.Len()-1]
		heap.Fix(find.right, 0)
	}
	
	if find.left.Len()>1 && find.left.IntHeap[find.left.Len()-1] < find.left.IntHeap[find.left.Len()-2] {
		find.left.Swap(find.left.Len()-1, find.left.Len()-2)
	}
}

// Finder() returns the median from the heaps
func (find *Median) Finder() float32 {
	if find.left.Len() == find.right.Len() {
		return float32(find.left.IntHeap[find.left.Len()-1]+find.right.IntHeap[0])/2
	}
	return float32(find.left.IntHeap[find.left.Len()-1])
}


// Part 2: Setup for using the API and parsing into appropriate struct
type Shares struct{
	Date string `json:"date"`
	Rate Rates `json:"rates"`
}

type Rates struct{
	USD float32 `json:"USD"`
	SEK float32 `json:"SEK"`
}

var url_base = "http://fx.modfin.se/"
var url_trail = "?symbols=usd,sek"

func urlDate (mm, dd int) string {
	return (fmt.Sprintf("2018-%02d-%02d", mm, dd))
}

func urlGenerator (mm, dd int) string {
	if mm < 1 || dd < 1 || mm > 12 || dd > 31 {
		return "Invalid arguments"
	}
	var url_date = fmt.Sprintf("2018-%02d-%02d",mm,dd)
	return url_base+url_date+url_trail
}

func GetMedianFromAPI(startMonth, startDay, endMonth, endDay int) float32 {
	MM := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var raw Shares
	var conversionUSDtoSEK, medianConversion float32
	median := Constructor()
	var strt, end int

	for month := startMonth; month <= endMonth; month++ {
		if month == startMonth {
			strt = startDay
		} else {
			strt = 1
		}
		if month == endMonth {
			end = endDay
		} else {
			end = MM[month-1]
		}
		for day := strt; day <= end; day++ {
			url := urlGenerator(month, day)
			response, err := http.Get(url)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
			jsdonData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
			json.Unmarshal(jsdonData, &raw)
			if raw.Date == urlDate(month,day){
				conversionUSDtoSEK = raw.Rate.SEK/raw.Rate.USD
				median.AddData(conversionUSDtoSEK)
				medianConversion = median.Finder();
			}
		}
	}
	fmt.Println("Median: ", median.left, median.right)
	return medianConversion
}

func main(){
	fmt.Println("Main running")
	fmt.Println("Median: ", GetMedianFromAPI(1, 1, 12, 31))
}