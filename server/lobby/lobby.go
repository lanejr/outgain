package lobby

import (
	"math/rand"

	"github.com/egnwd/outgain/server/engine"
)

const lobbySize int = 10

var lobbies = make(map[uint64]*Lobby)

// Lobby runs its own instance of an engine, and keeps track of its users
type Lobby struct {
	ID     uint64
	Engine engine.Engineer
	Guests guestList
	size   int
}

// NewLobby creates a new lobby with its own engine and list of guests
func NewLobby() (lobby *Lobby) {
	e := engine.NewEngine()
	id := newID()
	lobby = &Lobby{
		ID:     id,
		Engine: e,
		Guests: generalPopulation(lobbySize),
		size:   lobbySize,
	}

	lobbies[lobby.ID] = lobby

	return
}

// NewTestLobby creates a new lobby with a test engine, a specific
// size and list of guests
func NewTestLobby(e engine.Engineer, size int) (lobby *Lobby) {
	id := newID()
	lobby = &Lobby{
		ID:     id,
		Engine: e,
		Guests: generalPopulation(size),
		size:   size,
	}

	lobbies[lobby.ID] = lobby

	return
}

//This is just for testing until it's fully implemented
var baseID = uint64(rand.Uint32())

func newID() uint64 {
	baseID++
	return baseID
}

func (lobby *Lobby) startEngine() {
	for _, guest := range lobby.Guests.list {
		lobby.Engine.AddEntity(guest.name, engine.RandomCreature)
	}

	go lobby.Engine.Run()
}

// GetLobby returns the Lobby with id: `id` and if it does not exist it returns
// `(nil, false)`
func GetLobby(id uint64) (*Lobby, bool) {
	l, ok := lobbies[id]
	return l, ok
}
