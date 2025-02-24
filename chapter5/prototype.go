package main

import "fmt"

type Detail struct {
	Comment []string
}

type Item interface {
	AddComment(comment string)
	CreateCopy() Item
	GetName() string
	GetDetail() *Detail
}

type ItemPrototype struct {
	Name   string
	Detail *Detail
}

func (ip *ItemPrototype) AddComment(comment string) {
	ip.Detail.Comment = append(ip.Detail.Comment, comment)
}

func (ip *ItemPrototype) GetName() string {
	return ip.Name
}

func (ip *ItemPrototype) GetDetail() *Detail {
	return ip.Detail
}

type DeepCopyItem struct {
	*ItemPrototype
}

func NewDeepCopyItem(name string) *DeepCopyItem {
	return &DeepCopyItem{
		ItemPrototype: &ItemPrototype{
			Name:   name,
			Detail: &Detail{Comment: []string{}},
		},
	}
}

func (d *DeepCopyItem) CreateCopy() Item {
	newComments := make([]string, len(d.Detail.Comment))
	copy(newComments, d.Detail.Comment)
	newDetail := &Detail{Comment: newComments}

	return &DeepCopyItem{
		ItemPrototype: &ItemPrototype{
			Name:   d.Name,
			Detail: newDetail,
		},
	}
}

type ShallowCopyItem struct {
	*ItemPrototype
}

func NewShallowCopyItem(name string) *ShallowCopyItem {
	return &ShallowCopyItem{
		ItemPrototype: &ItemPrototype{
			Name:   name,
			Detail: &Detail{Comment: []string{}},
		},
	}
}

func (s *ShallowCopyItem) CreateCopy() Item {
	return &ShallowCopyItem{
		ItemPrototype: &ItemPrototype{
			Name:   s.Name,
			Detail: s.Detail,
		},
	}
}

type ItemManager struct {
	items map[string]Item
}

func NewItemManager() *ItemManager {
	return &ItemManager{
		items: make(map[string]Item),
	}
}

func (im *ItemManager) RegisterItem(key string, item Item) {
	im.items[key] = item
}

func (im *ItemManager) Create(key string) (Item, error) {
	item, ok := im.items[key]
	if !ok {
		return nil, fmt.Errorf("指定されたキーは存在しません")
	}
	return item.CreateCopy(), nil
}

// func main() {
// 	mouse := NewDeepCopyItem("マウス")
// 	mouse.AddComment("original")

// 	keyboard := NewShallowCopyItem("キーボード")
// 	keyboard.AddComment("original")

// 	manager := NewItemManager()
// 	manager.RegisterItem("mouse", mouse)
// 	manager.RegisterItem("keyboard", keyboard)

// 	clonedMouse, err := manager.Create("mouse")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	clonedKeyboard, err := manager.Create("keyboard")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Printf("マウス（オリジナル）: Name=%s, Comments=%v\n", mouse.Name, mouse.Detail.Comment)
// 	fmt.Printf("マウス（コピー）: Name=%s, Comments=%v\n", clonedMouse.GetName(), clonedMouse.GetDetail().Comment)
// 	fmt.Printf("キーボード（オリジナル）: Name=%s, Comments=%v\n", keyboard.Name, keyboard.Detail.Comment)
// 	fmt.Printf("キーボード（コピー）: Name=%s, Comments=%v\n", clonedKeyboard.GetName(), clonedKeyboard.GetDetail().Comment)

// 	clonedMouse.AddComment("Good")
// 	clonedKeyboard.AddComment("SoSo")

// 	fmt.Println("")
// 	fmt.Printf("マウス（オリジナル）: Name=%s, Comments=%v\n", mouse.Name, mouse.Detail.Comment)
// 	fmt.Printf("マウス（コピー）: Name=%s, Comments=%v\n", clonedMouse.GetName(), clonedMouse.GetDetail().Comment)
// 	fmt.Printf("キーボード（オリジナル）: Name=%s, Comments=%v\n", keyboard.Name, keyboard.Detail.Comment)
// 	fmt.Printf("キーボード（コピー）: Name=%s, Comments=%v\n", clonedKeyboard.GetName(), clonedKeyboard.GetDetail().Comment)
// }
