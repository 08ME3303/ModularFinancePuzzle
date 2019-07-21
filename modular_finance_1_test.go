package main

import(
	"testing"
	"fmt"
)

func TestURLGenerator(t *testing.T){
	result := urlGenerator(1, 1)
	if result != "http://fx.modfin.se/2018-01-01?symbols=usd,sek" {
		t.Errorf("Expected: http://fx.modfin.se/2018-01-01?symbols=usd,sek Got: %v", result)
	}

	result = urlGenerator(0, 0)
	if result != "Invalid arguments" {
		t.Errorf("Index out of bound. Ensure the month is from 1 to 12 and day is is from 1 to 31. Got: %v", result)
	}
}

func TestAddDataOdd(t *testing.T){
	structure := Constructor()
	var inp = [7]float32{0.1, 0.13, 0.14, 0.12, 0.11, 0.126, 0.127}

	for i := 0; i < 7; i++ {
		structure.AddData(inp[i])
	}

	fmt.Println("Result from AddData", structure.left, structure.right)

	if structure.left.Len() - structure.right.Len() > 1 {
		t.Errorf("Unbalanced left and right heaps in AddData")
	}
}

func TestAddDataEven(t *testing.T){
	structure := Constructor()
	var inp = [6]float32{0.13, 0.14, 0.12, 0.11, 0.126, 0.127}

	for i := 0; i < 6; i++ {
		structure.AddData(inp[i])
	}

	fmt.Println("Result from AddData", structure.left, structure.right)

	if structure.left.Len() - structure.right.Len() > 1 {
		t.Errorf("Unbalanced left and right heaps in AddData")
	}
}


func TestAddDataID(t *testing.T){
	structure := Constructor()

	for i := 0; i < 10; i++ {
		structure.AddData(0.1)
	}

	fmt.Println("Result from AddData", structure.left, structure.right)

	if structure.left.Len() - structure.right.Len() > 1 {
		t.Errorf("Unbalanced left and right heaps in AddData")
	}

	for i := 0; i < 11; i++ {
		structure.AddData(0.1)
	}

	fmt.Println("Result from AddData", structure.left, structure.right)

	if structure.left.Len() - structure.right.Len() > 1 {
		t.Errorf("Unbalanced left and right heaps in AddData")
	}
}

func TestMedianFinder(t *testing.T) {
	result := GetMedianFromAPI(1, 1, 1, 5)
	fmt.Println(result)
	var expected float32
	expected = 8.154341
	if result != expected {
		t.Errorf("Incorrect result for even data stream")
	}

	result = GetMedianFromAPI(1, 1, 1, 8)
	fmt.Println(result)
	expected = 8.162558
	if result != expected {
		t.Errorf("Incorrect result for odd data stream")
	}
}