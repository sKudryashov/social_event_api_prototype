package controller

import (
	"io/ioutil"
	"gopkg.in/go-playground/validator.v8"
	"github.com/sKudryashov/social_event_api_prototype/model"
	"github.com/sKudryashov/social_event_api_prototype/router"
	"encoding/json"
	"github.com/go-playground/lars"
	"strconv"
	"net/http"
)

// EventController serves events logic of the application
type EventController struct {
	error EventError
}
// EventError this is the custom error type of event handling
type EventError struct {
	err error
}

// NewEventController initialization of the controller
func NewEventController() *EventController {
	return &EventController{}
}

// PushData adding data to a storage (whatever it is)
func (ec *EventController) PushData (c *router.MyContext) {
	data, _ := ioutil.ReadAll(c.Request().Body)
	request := model.EventData{}
	validate := ec.getValidator()

	if err := json.Unmarshal(data, &request); err != nil {
		c.AppContext.Log.Println("Error with json unmarshalling: " + err.Error())
		return err
	}

	if err := validate.Struct(request); err != nil {
		c.AppContext.Log.Println("Error with validation: " + err.Error())
		return err
	}

	if err := c.AppContext.Storage.AddEvent(request); err == nil {
		c.AppContext.Log.Println("Error adding data: " + err.Error())
		return err
	}

	rsp := ec.getSuccessWriter(c)
	rsp.Write([]byte("Data has been written successfully"))
	c.AppContext.Log.Println("Successfull response")
}

// GetData returns the whole dataset
func (ec *EventController) GetData (c *router.MyContext)  {
	responseModel, err := c.AppContext.Storage.GetAllEvents()
	if err != nil {
		c.AppContext.Log.Println("Error with db fetching: " + err.Error())
	}

	rsp := ec.getSuccessWriter(c)
	dataFoundJson, err := json.Marshal(responseModel)

	if err != nil {
		c.AppContext.Log.Println("Error with unmarshalling: " + err.Error())
	}
	rsp.Write([]byte(dataFoundJson))
}

// GetDataByType Fetching data by event type from storage
func (ec *EventController) GetDataByType(c *router.MyContext) {
	data, _ := ioutil.ReadAll(c.Request().Body)
	request := model.FetchBy{}

	if err := json.Unmarshal(data, &request); err != nil {
		c.AppContext.Log.Println("Error with json unmarshalling: " + err.Error())
	}

	validate := ec.getValidator()

	if err := validate.Struct(request); err != nil {
		c.AppContext.Log.Println("Error with validation: " + err.Error())
	}

	events, err := c.AppContext.Storage.GetEvents(request.Type)

	if err != nil {
		c.AppContext.Log.Println("Data fetching error:" + err.Error())
	}

	rsp := ec.getSuccessWriter(c)
	dataFoundJson, err := json.Marshal(events)

	if err != nil {
		c.AppContext.Log.Println("Unmarshalling error: " + err.Error())
	}

	rsp.Write([]byte(dataFoundJson))
}

// GetDataByRange returns data in a given time range
func (ec *EventController) GetDataByRange(c *router.MyContext)  {
	var start, end int
	var err error

	start, err = strconv.Atoi(c.Ctx.Param("start"))
	if err != nil {
		c.AppContext.Log.Println("Wrong URL start parameter: " + err.Error())
	}

	end, err = strconv.Atoi(c.Ctx.Param("end"))
	if err != nil {
		c.AppContext.Log.Println("Wrong URL end parameter: " + err.Error())
	}

	responseModel, err := c.AppContext.Storage.GetEventsByRange(start, end)

	if err != nil {
		c.AppContext.Log.Println("Error with db fetching: " + err.Error())
	}

	rsp := ec.getSuccessWriter(c)
	dataFoundJson, err := json.Marshal(responseModel)
	rsp.Write([]byte(dataFoundJson))
}

func (ec *EventController) getValidator() *validator.Validate {
	validate := validator.New()
	validate.SetTagName("validate")

	return validate
}

func (ec *EventController) getSuccessWriter(c *router.MyContext) *lars.Response {
	rsp := c.Ctx.Response()
	rsp.WriteHeader(http.StatusOK)
	rsp.Header().Set("Content-Type", "application/json")

	return rsp
}

func (ec *EventController) getErrorNotFoundWriter(c *router.MyContext) *lars.Response {
	rsp := c.Ctx.Response()
	rsp.WriteHeader(http.StatusNotFound)
	rsp.Header().Set("Content-Type", "application/json")

	return rsp
}

// Returns writer for HTTP forbidden
func (ec *EventController) getErrorForbiddenWriter(c *router.MyContext) *lars.Response {
	rsp := c.Ctx.Response()
	rsp.WriteHeader(http.StatusForbidden)
	rsp.Header().Set("Content-Type", "application/json")

	return rsp
}
