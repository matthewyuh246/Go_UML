package main

import "fmt"

type Observer interface {
	update(name string)
}

type StoreObserver struct{}

func (s StoreObserver) update(name string) {
	fmt.Printf("%sが入荷されました。仕入れが可能です。 \n", name)
}

type PersonalObserver struct{}

func (p PersonalObserver) update(name string) {
	fmt.Printf("%sが入荷されました。仕入れが可能です。 \n", name)
}

type ItemSubject struct {
	name      string
	observers []Observer
}

func NewItemSubject(name string) ItemSubject {
	return ItemSubject{
		name:      name,
		observers: make([]Observer, 0),
	}
}

func (s *ItemSubject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ItemSubject) Detach(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ItemSubject) Notify() {
	for _, obs := range s.observers {
		obs.update(s.name)
	}
}

type TvGameSubject struct {
	ItemSubject
	inStock bool
}

func NewTvGameSubject(name string) *TvGameSubject {
	return &TvGameSubject{
		ItemSubject: NewItemSubject(name),
		inStock:     false,
	}
}

func (t *TvGameSubject) Restock() {
	fmt.Println("TVゲームの入荷")
	t.inStock = true
	t.Notify()
}

// func main() {
// 	store := StoreObserver{}
// 	personal := PersonalObserver{}
// 	tvGame := NewTvGameSubject("New RPG Game")

// 	tvGame.Attach(store)
// 	tvGame.Attach(personal)

// 	tvGame.Restock()
// }
