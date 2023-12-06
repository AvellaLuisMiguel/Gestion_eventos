package test

import (
	"Events/api/models"
	"testing"
	"github.com/stretchr/testify/assert"
)

type Test struct{
	model models.EventModel
}

func NewTest(model models.EventModel) *Test{
	return &Test{model:model}
}

func (t *Test) TestModel(a *testing.T){
	event := &models.Event{
		ID:3,
		Name:"evento 6",        
		Type:"Tipo a",        
		Description:"descripcion" ,  
		Date:"12/12/2023",         
		State:1,
		Clasification:"",
	}	
	
	err := t.model.AddEvent(event)
	assert.NoError(a, err)
}