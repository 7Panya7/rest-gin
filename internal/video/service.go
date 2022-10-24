package video

type videoService struct {
	videos []Video
}

func NewService() VideoService {
	return &videoService{}
}

func (v *videoService) Save(video Video) Video {
	v.videos = append(v.videos, video)
	return video
}
func (v *videoService) FindAll() []Video {
	return v.videos
}
