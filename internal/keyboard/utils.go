package keyboard

import "strconv"

const (
	BackToMainKey = iota
	BackToGaiKey

	BackToGaishnikKey
)

const (
	BackPattern = "back_"
)

type Back struct {
	Keyboard
}

func BackCallbackData(key int) string {
	return BackPattern + strconv.Itoa(key)

}
