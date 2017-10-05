package main

/*
Everyone in front of me will get a maximum of the same number of tickets as me,
everyone behind will get a maximum of n-1 tickets because the count stops as soon
as I have all my tickets.
 */
func getEfficientTicketCount(line []int, myPosition int) int {
	myTickets := line[myPosition]
	totalTickets := 0
	for index, numTickets := range line {
		if numTickets >= myTickets {
			totalTickets += myTickets
			if index > myPosition{
				totalTickets--
			}
		} else {
			totalTickets += numTickets
		}
	}
	return totalTickets
}
