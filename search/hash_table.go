package search

type IPhonebook interface {
	Add(name, phone string)
	Get(name string) string
}

type Phonebook struct {
	book map[string]string
}

func NewPhonebook() Phonebook {
	pb := Phonebook{}
	pb.book = make(map[string]string)

	return pb
}

func (p Phonebook) Add(name, phone string) {
	p.book[name] = phone
}

func (p Phonebook) Get(name string) string {
	value, ok := p.book[name]

	if !ok {
		return "not found"
	}

	return value
}
