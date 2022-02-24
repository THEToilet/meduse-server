package model

type ControllerName string

const (
	CON0 = ControllerName("Controller0")
	CON1 = ControllerName("Controller1")
	CON2 = ControllerName("Controller2")
	CON3 = ControllerName("Controller3")
	CON4 = ControllerName("Controller4")
)

func GetControllerName(num int) ControllerName {
	switch num {
	case 1:
		return CON1
	case 2:
		return CON2
	case 3:
		return CON3
	case 4:
		return CON4
	default:
	}
	return CON0
}
