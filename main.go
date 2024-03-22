package main

import (
	"math/rand"
	"sort"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Item struct {
	Value    int
	Change   int
	Weight   int
	EndValue int
	Index    float64
}

func updateEndValue(i *Item, days int) Item {
	endValue := i.Value + (i.Change * days)
	i.Index = float64(endValue) / float64(i.Weight)
	i.EndValue = endValue
	return *i
}

type Ship struct {
	Price    int
	Capacity int
	Days     int
}

func generateItems(numItems int) []Item {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	items := make([]Item, numItems)

	for i := 0; i < numItems; i++ {
		items[i] = Item{
			Value:  rand.Intn(201), // Random value between 0 and 200
			Change: rand.Intn(20) - 10,
			Weight: rand.Intn(46) + 5, // Random weight between 5 and 50
		}
	}

	return items
}
func GetMaximumProfit(items []Item, ship Ship) int {
	indexes := []Item{}
	for _, item := range items {
		indexes = append(indexes, updateEndValue(&item, ship.Days))
	}
	sort.Slice(indexes, func(i, j int) bool {
		if indexes[i].EndValue != indexes[j].EndValue {

			return indexes[i].EndValue > indexes[j].EndValue
		}
		return indexes[i].Weight < indexes[j].Weight
	})
	spew.Dump(indexes)
	profit := 0
	availableWeight := ship.Capacity
	for _, v := range indexes {
		if newW := availableWeight - v.Weight; newW >= 0 {
			availableWeight = newW
			profit += v.EndValue
			if availableWeight <= 0 {
				return profit
			}
		}
	}
	return profit
}

func main() {
	ship := Ship{
		Days:     3,
		Price:    75,
		Capacity: 50,
	}
	items := generateItems(2) // generates a list of 10 items
	profit := GetMaximumProfit(items, ship) - ship.Price

}
