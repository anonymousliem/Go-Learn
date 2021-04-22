package helper

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDefaultEngine_maxSpeed(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "default engine should have 50 ",
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := DefaultEngine{}
			if got := e.maxSpeed(); got != tt.want {
				t.Errorf("maxSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTurboEngine_maxSpeed(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Turbo Engine have 100",
			want: 100,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := TurboEngine{}
			if got := e.maxSpeed(); got != tt.want {
				t.Errorf("maxSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

type FakeEngine struct {

}

func (e FakeEngine) maxSpeed() int{
	return 5
}

type MockEngine struct {
	mock.Mock
}

func (m MockEngine) maxSpeed() int{
	args := m.Called()
	return args.Get(0).(int)
}

func TestCar_Speed_WithMock(t *testing.T) {
	
	t.Run("just call one time", func(t *testing.T) {
		mock := new(MockEngine)
		car := Car{
			Speeder: mock,
		}
		mock.On("maxSpeed").Return(9).Times(1)
		speed := car.Speed()
		assert.Equal(t, 20, speed)
		mock.AssertExpectations(t)
	})
	
	t.Run("call 3 times", func(t *testing.T) {
		mock := new(MockEngine)
		car := Car{
			Speeder: mock,
		}
		mock.On("maxSpeed").Return(60).Times(3)
		speed := car.Speed()
		assert.Equal(t, 60, speed)
		mock.AssertExpectations(t)
	})
	
}
func TestCar_Speed(t *testing.T) {
	type fields struct {
		Speeder Speeder
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "default engine contain 50",
			fields: fields{Speeder: &DefaultEngine{}},
			want: 50,
		},
		{
			name: "turbo engine contain 80",
			fields: fields{Speeder: &TurboEngine{}},
			want: 80,
		},
		{
			name: "using fake engine should be 20 when maxspeed 10",
			fields: fields{Speeder: &FakeEngine{}},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Car{
				Speeder: tt.fields.Speeder,
			}
			if got := c.Speed(); got != tt.want {
				t.Errorf("Speed() = %v, want %v", got, tt.want)
			}
		})
	}
}