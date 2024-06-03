package main

type player struct {
	name     string
	purse    int
	isDealer bool
	isHuman  bool
}

func (p *player) String() string {
	return p.name
}

// returns a new dealer player named "dealer".
func newDealer() *player {
	return &player{name: "dealer", purse: 0, isDealer: true, isHuman: false}
}
