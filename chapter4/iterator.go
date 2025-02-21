package main

import "fmt"

type Patient struct {
	Id   int
	Name string
}

func NewPatient(id int, name string) *Patient {
	return &Patient{
		Id:   id,
		Name: name,
	}
}

type IIterator interface {
	hasNext() bool
	next() *Patient
}

type Aggregate interface {
	getIterator() IIterator
}

type WaitingRoom struct {
	patients []Patient
}

func NewWaitingRoom() *WaitingRoom {
	return &WaitingRoom{
		patients: []Patient{},
	}
}

func (wr *WaitingRoom) getPatients() []Patient {
	return wr.patients
}

func (wr *WaitingRoom) getCount() int {
	return len(wr.patients)
}

func (wr *WaitingRoom) checkIn(patient Patient) {
	wr.patients = append(wr.patients, patient)
}

func (wr *WaitingRoom) getIterator() IIterator {
	return &WaitingRoomIterator{
		aggregate: wr,
		position:  0,
	}
}

type WaitingRoomIterator struct {
	position  int
	aggregate *WaitingRoom
}

func (it *WaitingRoomIterator) hasNext() bool {
	return it.position < it.aggregate.getCount()
}

func (it *WaitingRoomIterator) next() *Patient {
	if !it.hasNext() {
		fmt.Println("患者が居ません")
		return nil
	}

	patient := &it.aggregate.getPatients()[it.position]
	it.position++
	return patient
}

// func main() {
// 	waitingRoom := NewWaitingRoom()
// 	waitingRoom.checkIn(Patient{Id: 1, Name: "Yamada"})
// 	waitingRoom.checkIn(Patient{Id: 2, Name: "Suzuki"})
// 	waitingRoom.checkIn(Patient{Id: 3, Name: "Tanaka"})

// 	iterator := waitingRoom.getIterator()
// 	fmt.Println(iterator.next())
// 	fmt.Println(iterator.next())
// 	fmt.Println(iterator.next())
// 	fmt.Println(iterator.next())
// }
