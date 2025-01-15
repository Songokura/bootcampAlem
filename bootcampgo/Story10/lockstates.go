package bootcamp

type Lock struct {
	locked bool
}

func NewLock() Lock {
	return Lock{
		locked: false,
	}
}

func (l *Lock) Lock() bool {
	if l.locked {
		return false
	}
	l.locked = true
	return true
}

func (l *Lock) Unlock() bool {
	if !l.locked {
		return false
	}
	l.locked = false
	return true
}

func (l *Lock) IsLocked() bool {
	return l.locked
}

// func main() {
// 	lock := NewLock()
// 	fmt.Println(lock.IsLocked()) // false
// 	fmt.Println(lock.Unlock())   // false
// 	fmt.Println(lock.Lock())     // true
// 	fmt.Println(lock.Lock())     // false
// 	fmt.Println(lock.IsLocked()) // true
// 	fmt.Println(lock.Unlock())   // true
// }
