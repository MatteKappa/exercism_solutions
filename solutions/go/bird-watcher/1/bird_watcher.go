package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	counter := 0
	for idx := 0; idx < len(birdsPerDay); idx++ {
		counter += birdsPerDay[idx]
	}
	return counter
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	counter := 0
	for idx := 0; idx < 7; idx++ {
		dayIdx := idx + (7 * (week-1))
		counter += birdsPerDay[dayIdx]
	}
	return counter
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for idx := 0; idx < len(birdsPerDay); idx++ {
		if idx == 0 {
			birdsPerDay[idx] += 1
		} else if idx % 2 == 0 {
			birdsPerDay[idx] += 1
		}
	}
	return birdsPerDay
}
