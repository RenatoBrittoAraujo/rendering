package ebiten_helper

type EbitenHelper interface {
	Start()
	FullScreen()
	SetScale()
	SetLayout()
}

type EbitenConfig struct {
}

type ebitenHelper struct {
}

func NewEbitenConfig() *EbitenConfig {
	return &EbitenConfig{}
}

func NewEbitenHelper(ebitenConfig *EbitenConfig) EbitenHelper {
	return &ebitenHelper{}
}

func (e *ebitenHelper) Start() {

}

func (e *ebitenHelper) FullScreen() {

}

func (e *ebitenHelper) SetScale() {

}

func (e *ebitenHelper) SetLayout() {

}
