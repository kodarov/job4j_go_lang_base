package tracker

type Tracker struct {
	items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) (Item, error) {
	index, exist := t.indexOf(item.ID)
	if exist {
		return t.items[index], ErrAlreadyExists
	}
	t.items = append(t.items, item)
	return item, nil
}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.items))
	copy(res, t.items)
	return res
}

func (t *Tracker) UpdateItem(item Item) error {
	index, ok := t.indexOf(item.ID)
	if !ok {
		return ErrNotFound
	}
	t.items[index] = item
	return nil
}

func (t *Tracker) indexOf(id string) (int, bool) {
	for i, item := range t.items {
		if item.ID == id {
			return i, true
		}
	}
	return -1, false
}
