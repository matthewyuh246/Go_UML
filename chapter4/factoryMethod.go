package main

// import "fmt"

type CreditCard interface {
	GetOwner() string
	GetCardType() string
	GetAnnualCharge() int
}

type BaseCreditCard struct {
	Owner string
}

func (b *BaseCreditCard) GetOwner() string {
	return b.Owner
}

type Platinum struct {
	BaseCreditCard
}

func NewPlatinum(owner string) *Platinum {
	return &Platinum{BaseCreditCard{Owner: owner}}
}

func (p *Platinum) GetCardType() string {
	return "Platinum"
}

func (p *Platinum) GetAnnualCharge() int {
	return 30000
}

type Gold struct {
	BaseCreditCard
}

func NewGold(owner string) *Gold {
	return &Gold{BaseCreditCard{Owner: owner}}
}

func (g *Gold) GetCardType() string {
	return "Gold"
}

func (g *Gold) GetAnnualCharge() int {
	return 10000
}

var creditCardDatabase []CreditCard

type PlatinumCreditCardFactory struct{}

func (f *PlatinumCreditCardFactory) CreateCreditCard(owner string) CreditCard {
	return NewPlatinum(owner)
}

func (f *PlatinumCreditCardFactory) RegisterCreditCard(cc CreditCard) {
	creditCardDatabase = append(creditCardDatabase, cc)
}

func (f *PlatinumCreditCardFactory) Create(owner string) CreditCard {
	cc := f.CreateCreditCard(owner)
	f.RegisterCreditCard(cc)
	return cc
}

type GoldCreditCardFactory struct{}

func (f *GoldCreditCardFactory) CreateCreditCard(owner string) CreditCard {
	return NewGold(owner)
}

func (f *GoldCreditCardFactory) RegisterCreditCard(cc CreditCard) {
	creditCardDatabase = append(creditCardDatabase, cc)
}

func (f *GoldCreditCardFactory) Create(owner string) CreditCard {
	cc := f.CreateCreditCard(owner)
	f.RegisterCreditCard(cc)
	return cc
}

// func main() {
// 	platinumFactory := &PlatinumCreditCardFactory{}
// 	platinumCard := platinumFactory.Create("Tanaka")
// 	fmt.Printf("Platinum Card: Owner=%s, Type=%s, AnnualCharge=%d\n",
// 		platinumCard.GetOwner(), platinumCard.GetCardType(), platinumCard.GetAnnualCharge())

// 	GoldFactory := &GoldCreditCardFactory{}
// 	GoldCard := GoldFactory.Create("Yamada")
// 	fmt.Printf("Gold Card: Owner=%s, Type=%s, AnnualCharge=%d\n",
// 		GoldCard.GetOwner(), GoldCard.GetCardType(), GoldCard.GetAnnualCharge())

// 	for _, creditcard := range creditCardDatabase {
// 		fmt.Println(creditcard)
// 	}
// }
