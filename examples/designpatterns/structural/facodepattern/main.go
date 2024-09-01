package main
 
import ("fmt")


type AudioPlayer struct{}

func (a *AudioPlayer)PlayAudio(){
	fmt.Println("Playing Audio")
}

type VideoPlayer struct{}

func (v *VideoPlayer)PlayVideo(){
	fmt.Println("Playing Video")
}


type ScreenManager struct{}

func (s *ScreenManager)ShowScreen(){
	fmt.Println("Showing Screen...")
}

type MultimediaFacade struct{
	audioPlayer *AudioPlayer
	videoPlayer *VideoPlayer
	screenManager *ScreenManager
}


func NewMultimediaFacade()*MultimediaFacade{
	return &MultimediaFacade{
		audioPlayer: &AudioPlayer{},
		videoPlayer: &VideoPlayer{},
		screenManager: &ScreenManager{},
	}
}


func (m *MultimediaFacade)PlayMovie(){
	m.audioPlayer.PlayAudio()
	m.videoPlayer.PlayVideo()
	m.screenManager.ShowScreen()
}

func main(){
	fmt.Println("Facade Pattern")
	fmt.Println("--------------")

	multimediaSystem:=NewMultimediaFacade()
	multimediaSystem.PlayMovie()
}