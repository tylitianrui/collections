package set

// set
type Set struct {
	data   map[interface{}]bool
	capity int
}

func NewSet(capity int) *Set {
	initCapity := 16
	if initCapity > capity {
		capity = initCapity
	}
	return &Set{
		data:   make(map[interface{}]bool, capity),
		capity: capity,
	}
}

// the amount of  elements  in  the set
func (set *Set) Len() int {
	return len(set.data)
}

// return true  if  item  included in  this  set
func (set *Set) Eixst(item interface{}) bool {
	data := set.data[item]
	return data
}

// Add an element to a set.
// This has no effect if the element is already present
func (set *Set) Add(item interface{}) {
	if exist := set.Eixst(item); exist == false {
		set.data[item] = true
	}
}

// Return the difference of two  sets as a new set
func (set *Set) Diff(other *Set) *Set {
	var capity int

	if set.capity >= other.capity {
		capity = set.capity
	} else {
		capity = other.capity
	}
	newSet := NewSet(capity)

	for k := range set.data {
		if other.Eixst(k) != true {
			newSet.Add(k)
		}
	}
	return newSet
}

// Return the intersection of two sets as a new set
func (set *Set) Intersect(other *Set) *Set {
	var capity int
	if set.capity >= other.capity {
		capity = set.capity
	} else {
		capity = other.capity
	}
	newSet := NewSet(capity)
	for k := range set.data {
		if other.Eixst(k) == true {
			newSet.Add(k)
		}
	}
	return newSet
}

// Return the union of sets as a new set
func (set *Set) Union(other *Set) *Set {
	capity := set.capity*3/2 + other.capity
	newSet := NewSet(capity)

	for k := range set.data {
		newSet.Add(k)
	}
	for k := range other.data {
		newSet.Add(k)
	}

	return newSet
}

// Remove all elements from this set
func (set *Set) Clear() {
	for item := range set.data {
		delete(set.data, item)
	}
}

//Remove given element from this set
func (set *Set) Delete(item interface{}) {
	delete(set.data, item)
}
