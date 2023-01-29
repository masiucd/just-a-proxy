package utils

// 1 = Luke Skywalker
// 2 = C-3PO
// 4 = Darth Vader
// 10 = Obi-Wan Kenobi
// 11 = Anakin Skywalker
// 13 = Chewbacca

// GetUserSlugID returns the user slug id
func GetUserSlugID(userSlug string) string {
	switch userSlug {
	case "luke":
		return "1"
	case "c3po":
		return "2"
	case "darth-vader":
		return "4"
	case "obi-wan-kenobi":
		return "10"
	case "anakin-skywalker":
		return "11"
	case "chewbacca":
		return "13"
	default:
		return "15" // Greedo is default
	}
}
