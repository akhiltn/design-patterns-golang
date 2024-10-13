package structural

// Define the IToyBird interface
type IToyBird interface {
	MakeSound() string
}

// Bird is the adaptee, the class we want to adapt
type Bird struct{}

func (b *Bird) Fly() string {
	return "I can fly"
}

func (b *Bird) Chirp() string {
	return "Chirp"
}

// BirdAdapter adapts Bird to the IToyBird interface
type BirdAdapter struct {
	bird *Bird
}

// Constructor function to create a new BirdAdapter
func NewBirdAdapter(bird *Bird) *BirdAdapter {
	return &BirdAdapter{
		bird: bird,
	}
}

// Implement the MakeSound method to call Chirp on the Bird
func (b *BirdAdapter) MakeSound() string {
	return b.bird.Chirp()
}
