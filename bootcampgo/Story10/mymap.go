package bootcamp

// KeyValuePair represents a key-value pair.
type KeyValuePair struct {
	Key   string
	Value interface{}
}

// MyMap represents a custom map structure.
type MyMap struct {
	data []KeyValuePair
}

// Set sets the value v associated with key k.
func (m *MyMap) Set(k string, v interface{}) {
	// Check if key already exists, update its value
	for i := range m.data {
		if m.data[i].Key == k {
			m.data[i].Value = v
			return
		}
	}
	// If key doesn't exist, append new key-value pair
	m.data = append(m.data, KeyValuePair{Key: k, Value: v})
}

// Get retrieves the value associated with key k.
func (m *MyMap) Get(k string) interface{} {
	for _, kv := range m.data {
		if kv.Key == k {
			return kv.Value
		}
	}
	return nil // Return nil if key is not found
}

// Has returns true if the map contains key k, otherwise false.
func (m *MyMap) Has(k string) bool {
	for _, kv := range m.data {
		if kv.Key == k {
			return true
		}
	}
	return false
}

// Delete deletes the key k and its associated value from the map.
func (m *MyMap) Delete(k string) {
	for i := range m.data {
		if m.data[i].Key == k {
			// Delete key-value pair by slicing
			m.data = append(m.data[:i], m.data[i+1:]...)
			return
		}
	}
}

// Items returns a slice of structures containing all key-value pairs in the map.
func (m *MyMap) Items() []KeyValuePair {
	return m.data
}

// func main() {
// 	myMap := MyMap{}

// 	myMap.Set("key1", 42)
// 	myMap.Set("key2", "value2")

// 	fmt.Println(myMap.Get("key1")) // Output: 42
// 	fmt.Println(myMap.Get("key2")) // Output: value2
// 	fmt.Println(myMap.Has("key1")) // Output: true
// 	fmt.Println(myMap.Has("key3")) // Output: false

// 	myMap.Delete("key2")
// 	fmt.Println(myMap.Has("key2")) // Output: false

// 	items := myMap.Items()
// 	for _, item := range items {
// 		fmt.Printf("Key: %s, Value: %v\n", item.Key, item.Value)
// 		// Output: Key: key1, Value: 42
// 	}
// }
