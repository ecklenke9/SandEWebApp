package todo

import (
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

func GetByID(id string) (Todo, error) {
	_, todo, err := findTodoByLocation(id)
	if err != nil {
		return Todo{}, err
	}
	return todo, err
}

func Add(message string) string {
	t := newTodo(message)
	mtx.Lock()
	list = append(list, t)
	mtx.Unlock()
	return t.ID
}

func Delete(id string) error {
	location, _, err := findTodoByLocation(id)
	if err != nil {
		return err
	}
	removeElementByLocation(location)
	return nil
}

func Complete(id string) (Todo, error) {
	location, todo, err := findTodoByLocation(id)
	if err != nil {
		return Todo{}, err
	}
	setTodoCompleteByLocation(location)
	todo.Complete = true
	return todo, nil
}

func newTodo(msg string) Todo {
	todo := Todo{
		ID:       xid.New().String(),
		Message:  msg,
		Complete: false,
	}
	log.Info().Msgf("New Todo created! [%+v]", todo)
	return todo
}

func findTodoByLocation(id string) (int, Todo, error) {
	mtx.RLock()
	defer mtx.RUnlock()
	for i, t := range list {
		if isMatchingID(t.ID, id) {
			return i, t, nil
		}
	}
	err := fmt.Errorf("could not find todo based on id: %v", id)
	log.Err(err).Send()
	return 0, Todo{}, err
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
