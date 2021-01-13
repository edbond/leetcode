import (
	// "sort"
	"fmt"
)

type Boat struct {
	passengers []int
	left       int
}

func d(s string, obj interface{}) {
	// fmt.Printf("%s: %+v\n", s, obj)
}

func numRescueBoats(people []int, limit int) int {
	boats := []Boat{}

	freq := make(map[int]int)
	for _, p := range people {
		freq[p]++
	}
	d("freq", freq)

	i := limit
	// k := 0
	for {
		// k++
		// if k > 10000 {
		//     panic("too long")
		// }

		if i == 0 {
			break
		}

		if p, found := freq[i]; !found || p == 0 {
			i--
			continue
		}

		boat := Boat{left: limit - i, passengers: []int{i}}
		freq[i]--
		if freq[i] <= 0 {
			delete(freq, i)
		}

		d("New boat", boat)

		if boat.left > 0 {
			moreSpace := boat.left
			for {
				if moreSpace <= 0 {
					break
				}

				if _, found := freq[moreSpace]; !found {
					moreSpace--
					continue
				}

				freq[moreSpace]--
				if freq[moreSpace] <= 0 {
					delete(freq, moreSpace)
				}

				boat.passengers = append(boat.passengers, moreSpace)
				boat.left -= moreSpace
				moreSpace = boat.left

				if len(boat.passengers) >= 2 {
					boat.left = 0
					break
				}
			}
		}
		d("Boat full", boat)
		d("freq", freq)

		boats = append(boats, boat)
	}

	return len(boats)
}