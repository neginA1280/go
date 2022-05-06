package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
    result := fmt.Sprintf("Welcome to my party, %s!", name)
	return result
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
    result := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
	return result
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
    
    location := fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", table, direction, distance, neighbor)
    result := Welcome(name) + "\n" + location

    return result
}
