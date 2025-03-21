package services

type Services struct {
	CourseService *CourseService
}

func NewServices() *Services {
	return &Services{
		CourseService: &CourseService{},
	}
}
