package guitars

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

// MockService is a mock implementation of the Service interface
type MockService struct {
	ListGuitarsFunc     func(ctx context.Context) ([]Guitar, error)
	FindGuitarByIDFunc  func(ctx context.Context, id int) (Guitar, error)
	CreateGuitarFunc    func(ctx context.Context, guitar Guitar) (Guitar, error)
	UpdateGuitarFunc    func(ctx context.Context, guitar Guitar) (Guitar, error)
	DeleteGuitarFunc    func(ctx context.Context, id int) error
}

func (m *MockService) ListGuitars(ctx context.Context) ([]Guitar, error) {
	if m.ListGuitarsFunc != nil {
		return m.ListGuitarsFunc(ctx)
	}
	return nil, nil
}

func (m *MockService) FindGuitarByID(ctx context.Context, id int) (Guitar, error) {
	if m.FindGuitarByIDFunc != nil {
		return m.FindGuitarByIDFunc(ctx, id)
	}
	return Guitar{}, nil
}

func (m *MockService) CreateGuitar(ctx context.Context, guitar Guitar) (Guitar, error) {
	if m.CreateGuitarFunc != nil {
		return m.CreateGuitarFunc(ctx, guitar)
	}
	return Guitar{}, nil
}

func (m *MockService) UpdateGuitar(ctx context.Context, guitar Guitar) (Guitar, error) {
	if m.UpdateGuitarFunc != nil {
		return m.UpdateGuitarFunc(ctx, guitar)
	}
	return Guitar{}, nil
}

func (m *MockService) DeleteGuitar(ctx context.Context, id int) error {
	if m.DeleteGuitarFunc != nil {
		return m.DeleteGuitarFunc(ctx, id)
	}
	return nil
}

func TestListGuitars(t *testing.T) {
	tests := []struct {
		name           string
		mockService    *MockService
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "success",
			mockService: &MockService{
				ListGuitarsFunc: func(ctx context.Context) ([]Guitar, error) {
					return []Guitar{
						{ID: 1, Brand: "Fender", Model: "Stratocaster"},
						{ID: 2, Brand: "Gibson", Model: "Les Paul"},
					}, nil
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "service error",
			mockService: &MockService{
				ListGuitarsFunc: func(ctx context.Context) ([]Guitar, error) {
					return nil, errors.New("database error")
				},
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(tt.mockService)
			req := httptest.NewRequest(http.MethodGet, "/guitars", nil)
			w := httptest.NewRecorder()

			h.ListGuitars(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestFindGuitarByID(t *testing.T) {
	tests := []struct {
		name           string
		guitarID       string
		mockService    *MockService
		expectedStatus int
	}{
		{
			name:     "success",
			guitarID: "1",
			mockService: &MockService{
				FindGuitarByIDFunc: func(ctx context.Context, id int) (Guitar, error) {
					return Guitar{ID: 1, Brand: "Fender", Model: "Stratocaster"}, nil
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid id",
			guitarID:       "abc",
			mockService:    &MockService{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:     "service error",
			guitarID: "1",
			mockService: &MockService{
				FindGuitarByIDFunc: func(ctx context.Context, id int) (Guitar, error) {
					return Guitar{}, errors.New("not found")
				},
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(tt.mockService)
			req := httptest.NewRequest(http.MethodGet, "/guitars/"+tt.guitarID, nil)
			w := httptest.NewRecorder()

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.guitarID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			h.FindGuitarByID(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestCreateGuitar(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		mockService    *MockService
		expectedStatus int
	}{
		{
			name: "success",
			requestBody: Guitar{
				Brand: "Fender",
				Model: "Telecaster",
			},
			mockService: &MockService{
				CreateGuitarFunc: func(ctx context.Context, guitar Guitar) (Guitar, error) {
					guitar.ID = 1
					return guitar, nil
				},
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "invalid json",
			requestBody:    "invalid json",
			mockService:    &MockService{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "service error",
			requestBody: Guitar{
				Brand: "Fender",
				Model: "Telecaster",
			},
			mockService: &MockService{
				CreateGuitarFunc: func(ctx context.Context, guitar Guitar) (Guitar, error) {
					return Guitar{}, errors.New("database error")
				},
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(tt.mockService)

			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}

			req := httptest.NewRequest(http.MethodPost, "/guitars", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			h.CreateGuitar(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestUpdateGuitar(t *testing.T) {
	tests := []struct {
		name           string
		guitarID       string
		requestBody    interface{}
		mockService    *MockService
		expectedStatus int
	}{
		{
			name:     "success",
			guitarID: "1",
			requestBody: Guitar{
				Brand: "Fender",
				Model: "Stratocaster Updated",
			},
			mockService: &MockService{
				UpdateGuitarFunc: func(ctx context.Context, guitar Guitar) (Guitar, error) {
					return guitar, nil
				},
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "invalid id",
			guitarID:       "abc",
			requestBody:    Guitar{},
			mockService:    &MockService{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalid json",
			guitarID:       "1",
			requestBody:    "invalid json",
			mockService:    &MockService{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:     "service error",
			guitarID: "1",
			requestBody: Guitar{
				Brand: "Fender",
				Model: "Stratocaster",
			},
			mockService: &MockService{
				UpdateGuitarFunc: func(ctx context.Context, guitar Guitar) (Guitar, error) {
					return Guitar{}, errors.New("database error")
				},
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(tt.mockService)

			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}

			req := httptest.NewRequest(http.MethodPut, "/guitars/"+tt.guitarID, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.guitarID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			h.UpdateGuitar(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestDeleteGuitar(t *testing.T) {
	tests := []struct {
		name           string
		guitarID       string
		mockService    *MockService
		expectedStatus int
	}{
		{
			name:     "success",
			guitarID: "1",
			mockService: &MockService{
				DeleteGuitarFunc: func(ctx context.Context, id int) error {
					return nil
				},
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "invalid id",
			guitarID:       "abc",
			mockService:    &MockService{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:     "service error",
			guitarID: "1",
			mockService: &MockService{
				DeleteGuitarFunc: func(ctx context.Context, id int) error {
					return errors.New("database error")
				},
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(tt.mockService)
			req := httptest.NewRequest(http.MethodDelete, "/guitars/"+tt.guitarID, nil)
			w := httptest.NewRecorder()

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.guitarID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			h.DeleteGuitar(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}
