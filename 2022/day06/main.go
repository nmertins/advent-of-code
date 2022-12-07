package main

const (
	StartOfPacketMarkerLength = 4
)

func FindMarker(input string) string {
	for i := 0; i < len(input)-4; i++ {
		marker := input[i : i+StartOfPacketMarkerLength]
		if AllCharactersUnique(marker) {
			return marker
		}
	}

	return ""
}

func AllCharactersUnique(input string) bool {
	return false
}
