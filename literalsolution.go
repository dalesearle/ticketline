package main

func getLiteralTicketCount(line []int, myPosition int) int {
	totalTickets := 0
	for {
		totalTickets++
		guy := line[0] - 1
		line = line[1:]
		myPosition--
		if guy > 0 {
			line = append(line, guy)
		}
		if myPosition < 0 {
			if guy > 0 {
				myPosition = len(line) -1
			} else {
				return totalTickets
			}
		}
	}
}
