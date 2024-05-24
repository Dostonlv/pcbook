package sample

import (
	"github.com/Dostonlv/pcbook/pb"
	"math/rand"

	"github.com/google/uuid"
)

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Core i7-7500U",
			"Core i5-6200U",
			"Core i3-5005U",
		)
	}
	return randomStringFromSet(
		"Ryzen 7 Pro 2700U",
		"Ryzen 5 Pro 2500U",
		"Ryzen 3 Pro 2300U",
	)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"GeForce RTX 2080 Ti",
			"GeForce GTX 1660 Ti",
			"GeForce GTX 1070",
		)
	}
	return randomStringFromSet(
		"Radeon RX 5700 XT",
		"Radeon RX 5600 XT",
		"Radeon RX 5500 XT",
	)
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {

		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	resolution := &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
	return resolution
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet("Dell", "HP", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS")
	case "HP":
		return randomStringFromSet("Pavilion", "EliteBook", "Envy")
	default:
		return randomStringFromSet("ThinkPad", "IdeaPad", "Yoga")
	}
}
