package controller

import (
	"testing"
	"github.com/sKudryashov/social_event_api_prototype/router"
	"github.com/sKudryashov/go-playground/lars"
	"github.com/sKudryashov/social_event_api_prototype/model"
	"net/http/httptest"
	//"net/http"
	"fmt"
	"github.com/karlseguin/gofake"
)

func TestMain(m *testing.M) {
	m.Run()
}

type ApplicationGlobals struct {
	Storage *model.Storage
}

type StubReader struct{
}


type Storage struct {}

// GetAllEvents mock
func (s* Storage) GetAllEvents() (ed *[]model.EventData, err error) {
	responseModel := []model.EventData{
	}
	return &responseModel, nil
}

// AddEvent mock
func (s* Storage) AddEvent (ed *model.EventData) error {
	return nil
}

// GetEvents mock
func (s* Storage) GetEvents(eventType string)  (rm *[]model.EventData, err error) {
	responseModel := []model.EventData{
	}
	return &responseModel, nil
}

// GetEventsByRange mock
func (s* Storage) GetEventsByRange (start, end int) (ed *[]model.EventData, err error) {
	responseModel := []model.EventData{
	}
	return &responseModel, nil
}

// initTestModel initializes a fake model
func initTestModel () *Storage {
	return &Storage{}
}

// initContext initializes context mock
func initContext() *router.MyContext {
	ctx := lars.New()
	return &router.MyContext {
		Ctx:        lars.NewContext(ctx),
		AppContext: newGlobals(),
	}
}

// newGlobals initializes globals for our controller
func newGlobals() *ApplicationGlobals {
	return &ApplicationGlobals{
		Storage: initTestModel(),
	}
}

// Read is a mock for reader
func (r *StubReader) Read(p []byte) (n int, err error){
	fmt.Println("A reader has been called")
	return 22, nil
}

/**
 * ==[ Tests ]==
 *
 * go test -v -run=EventController_PushData
 */
func TestEventController_PushData(t *testing.T) {
	recorder := httptest.NewRecorder()
	fake := gofake.New()
	fake.Stub("getRequestBody").Returning()
	reader := new(StubReader)
	request := httptest.NewRequest("POST", "/add", reader)
	request.Body
	recorder.Body
	ec := new(EventController)
	context := initContext()
	err := ec.PushData(context)
	if err != nil {
		t.Fatal("Push data controller error")
	}
}
