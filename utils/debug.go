package utils

const Debug = true

type DebugActionType func()

func DebugAction(action DebugActionType, activate ...bool) {
	if len(activate) == 0 {
		activate = append(activate, true)
	}
	if Debug && activate[0] {
		action()
	}
}
