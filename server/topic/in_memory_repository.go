package topic

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/satori/go.uuid"
)

type topic struct {
	name string
}

type sprintTopics map[uuid.UUID]topic

type inMemoryRepository struct {
	topics map[uuid.UUID]sprintTopics
}

func (m *inMemoryRepository) List(command *listTopicQuery) ([]topicModel, error) {
	sprintTopics, ok := m.topics[command.Sprint]
	if !ok {
		return nil, errors.New("sprint not found")
	}

	result := make([]topicModel, 0, len(sprintTopics))

	for topicID, value := range sprintTopics {
		model := topicModel{
			topicID,
			command.Sprint,
			value.name,
		}
		result = append(result, model)
	}

	return result, nil
}

func (m *inMemoryRepository) Save(command *addTopicCommand) error {

	if _, ok := m.topics[command.Sprint]; !ok {
		m.topics[command.Sprint] = sprintTopics{}
	}

	sprintTopics := m.topics[command.Sprint]

	sprintTopics[command.ID] = topic{command.Text}

	spew.Dump(m.topics)

	return nil
}
