package todo

import (
	"errors"
	"fmt"
	"sync"

	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

var (
	list []Todo
	mtx  sync.RWMutex
	once sync.Once
)

func init() {
	once.Do(initialiseList)
}

func initialiseList() {
	list = []Todo{}
}

type Todo struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	Complete bool   `json:"complete"`
}

func Get() []Todo {
	return list
}

func Add(message string) string {
	t := newTodo(message)
	mtx.Lock()
	list = append(list, t)
	mtx.Unlock()
	return t.ID
}

func Delete(id string) error {
	location, err := findTodoLocation(id)
	if err != nil {
		return err
	}
	removeElementByLocation(location)
	return nil
}

func Complete(id string) error {
	location, err := findTodoLocation(id)
	if err != nil {
		return err
	}
	setTodoCompleteByLocation(location)
	return nil
}

func newTodo(msg string) Todo {
	todo := Todo{
		ID: xid.New().String(),
		Message:  msg,
		Complete: false,
	}
	log.Info().Msgf("New Todo created! [%+v]", todo)
	return todo
}

func findTodoLocation(id string) (int, error) {
	mtx.RLock()
	defer mtx.RUnlock()
	for i, t := range list {
		if isMatchingID(t.ID, id) {
			return i, nil
		}
	}
	err := fmt.Errorf("could not find todo based on id: %v", id)
	log.Err(err).Send()
	return 0, err
}

func removeElementByLocation(i int) {
	mtx.Lock()
	log.Info().Msgf("Todo Deleted! [%+v]", list[i])
	list = append(list[:i], list[i+1:]...)
	mtx.Unlock()
}

func setTodoCompleteByLocation(location int) {
	mtx.Lock()
	list[location].Complete = true
	log.Info().Msgf("Todo Completed! [%+v]", list[location])
	mtx.Unlock()
}

func isMatchingID(a string, b string) bool {
	return a == b
}
