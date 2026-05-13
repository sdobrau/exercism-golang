package jedlik

import "fmt"

// TODO: define the 'Drive()' method

func (c *Car) Drive() Car {
	if c.battery >= c.batteryDrain {
		c.battery -= c.batteryDrain
        c.distance += c.speed
        return *c
	} else if c.battery < c.batteryDrain {
		return *c
	}
    return *c
}

// TODO: define the 'DisplayDistance() string' method

func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method

func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method

func (c *Car) CanFinish(trackDistance int) bool {
	var drivesRequired int
	if trackDistance % c.speed != 0 {
		drivesRequired = trackDistance / c.speed + 1
	} else {
		drivesRequired = trackDistance / c.speed
	}
	drivesPossible := c.battery / c.batteryDrain
	if drivesPossible < drivesRequired {
		return false
	} else {
		return true
	}
}
