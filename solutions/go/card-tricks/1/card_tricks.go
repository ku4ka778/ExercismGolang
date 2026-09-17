package cards
// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
    favCards := []int{2,6,9}
	return favCards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    if 0 <= index && index < len(slice){
        return slice[index]
    } else {
        return -1 
    }
    
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    if 0 <= index && index < len(slice){
    	slice[index] = value
        return slice
    } else {
        slice = append(slice,value)
        return slice
    }
	
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	res := append(values,slice...)
    return res
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    if 0 <= index && index < len(slice){
    	res :=append(slice[:index],slice[index+1:]...)
        return res
    } else {
        return slice
    }
}
