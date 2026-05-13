package tracker

import "github.com/google/uuid"

type Usecase interface {
	Done(in Input, out Output, tracker *Tracker)
}
type AddUsecase struct{}

func (u AddUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name:")
	name := in.Get()
	id := uuid.New().String()
	tracker.AddItem(Item{Name: name, ID: id})
}

type GetUsecase struct{}

func (u GetUsecase) Done(_ Input, out Output, tracker *Tracker) {
	for _, item := range tracker.items {
		out.Out(item.toString())
	}
}

type UpdateUsecase struct{}

func (u UpdateUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter ID:")
	id := in.Get()
	for i := range tracker.items {
		if tracker.items[i].ID == id {
			out.Out("enter new name:")
			tracker.items[i].Name = in.Get()
			out.Out("updated name")
			break
		}
	}
}

type DeleteUsecase struct{}

func (u DeleteUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter ID:")
	id := in.Get()
	for i := range tracker.items {
		if tracker.items[i].ID == id {
			tracker.items = append(tracker.items[:i], tracker.items[i+1:]...)
			out.Out("deleted name")
			break
		}
	}
}

type SearchUsecase struct{}

func (u SearchUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter search name:")
	name := in.Get()
	for _, item := range tracker.items {
		if name == item.Name {
			out.Out("found ID: " + item.ID)
			break
		}
	}
}
