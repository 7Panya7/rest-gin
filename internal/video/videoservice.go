package video

type VideoService interface {
	Save(video Video) Video
	FindAll() []Video
}
