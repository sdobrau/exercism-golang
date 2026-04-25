package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var totalBirds int
	for i := 0; i < len(birdsPerDay); i++ {
        totalBirds += birdsPerDay[i]
    }
    return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    var birdsCounted int
    var counts int
	for i := week * 7 - 7; counts < 7; i++ {
        birdsCounted += birdsPerDay[i]
        counts += 1
    }
    return birdsCounted
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i += 2 {
        birdsPerDay[i] += 1
    }
    return birdsPerDay
}
