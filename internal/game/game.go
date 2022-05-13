package game

type Game interface {
	StartGameLoop()
	Update()
	HandleInput()
	Draw()
	Kill()
	Save()
}

type tictactoe struct {
}

func NewGame() Game {
	newGame := &tictactoe{}

	return newGame
}

func RunGameLoop(game Game) {

}
