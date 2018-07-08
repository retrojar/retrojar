package topic

type useCase struct {
	repository Repository
}

type Repository interface {
	Save(topic *addTopicCommand) error
	List(topic *listTopicQuery) ([]topicModel, error)
}

func (u *useCase) add(command *addTopicCommand) error {
	return u.repository.Save(command)
}

func (u *useCase) list(command *listTopicQuery) ([]topicModel, error) {
	return u.repository.List(command)
}
