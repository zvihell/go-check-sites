package service_test

import (
	"check-domain-api/internal/models"
	"check-domain-api/internal/service"
	"check-domain-api/internal/service/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	mockStore := new(mocks.DomainRepository)
	dom := models.Domain{
		ID:          1,
		Domain:      "google.com",
		Latency:     123,
		Available:   true,
		Last_update: time.Now(),
	}
	mockStore.On("Save", dom).Return(nil).Once()

	service := service.NewDomain(mockStore)

	err := service.Save(dom)

	assert.Nil(t, err)

}

func TestGetName(t *testing.T) {
	mockStore := new(mocks.DomainRepository)
	dom := models.Domain{
		ID:          1,
		Domain:      "google.com",
		Latency:     123,
		Available:   true,
		Last_update: time.Now(),
	}
	mockStore.On("GetByName", "google.com").Return(dom, nil).Once()
	service := service.NewDomain(mockStore)
	result, err := service.GetByName("google.com")

	assert.NoError(t, err)
	assert.Equal(t, dom, result)
	mockStore.AssertExpectations(t)

}

func TestGetMin(t *testing.T) {
	mockStore := new(mocks.DomainRepository)
	dom := models.Domain{
		ID:          1,
		Domain:      "google.com",
		Latency:     123,
		Available:   true,
		Last_update: time.Now(),
	}
	mockStore.On("GetMin").Return(dom, nil).Once()

	service := service.NewDomain(mockStore)
	result, err := service.GetMin()

	assert.NoError(t, err)
	assert.Equal(t, dom, result)
	mockStore.AssertExpectations(t)
}

func TestGetMax(t *testing.T) {
	mockStore := new(mocks.DomainRepository)
	dom := models.Domain{
		ID:          1,
		Domain:      "taobao.com",
		Latency:     1223233,
		Available:   false,
		Last_update: time.Now(),
	}
	mockStore.On("GetMax").Return(dom, nil).Once()

	service := service.NewDomain(mockStore)
	result, err := service.GetMax()

	assert.NoError(t, err)
	assert.Equal(t, dom, result)
	mockStore.AssertExpectations(t)
}
