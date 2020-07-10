// Inverse Time Overcurrent calculations
package main

import "fmt"

func main() {

	// example to calculate the desired timedial for a backup relay given the desired margin
	primary := itoc{
		pickup:   1,
		timedial: 1,
		curve:    U3,
	}
	backup := itoc{
		pickup: 1,
		curve:  U3,
	}

	ampsSecondary := 10.0
	coordinationMargin := 0.33

	primaryTime := primary.opTime(ampsSecondary)

	backupTime := primaryTime + coordinationMargin
	backup.timedial = backup.tdCalc(backupTime, 10)

	fmt.Printf("Primary opTime of %fs, backup %fs, cti=%fs", primaryTime, backupTime, backupTime-primaryTime)
}
