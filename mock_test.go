package gomock_test

import (
	"fmt"
	"testing"

	gomock "github.com/nextunit-io/go-mock"
	"github.com/stretchr/testify/assert"
)

func TestMockInput(t *testing.T) {
	mock := gomock.GetMock[string, string](fmt.Errorf("general error"))

	for i := 0; i < 3; i++ {
		mock.AddInput(fmt.Sprintf("test-input-%d", i))
	}

	assert.Equal(t, 3, mock.HasBeenCalled())
	assert.Equal(t, "test-input-0", mock.GetInput(0))
	assert.Equal(t, "test-input-1", mock.GetInput(1))
	assert.Equal(t, "test-input-2", mock.GetInput(2))

	assert.Equal(t, []string{
		"test-input-0",
		"test-input-1",
		"test-input-2",
	}, mock.GetInputs())
}

func TestMockOutput(t *testing.T) {
	mock := gomock.GetMock[string, string](fmt.Errorf("general error"))

	for i := 0; i < 4; i++ {
		if i%2 == 0 {
			mock.AddError(fmt.Errorf("test-error-%d", i))
		} else {
			v := fmt.Sprintf("test-output-%d", i)
			mock.AddReturnValue(&v)
		}
	}

	for i := 0; i < 4; i++ {
		output, err := mock.GetNextResult()
		if i%2 == 0 {
			assert.Nil(t, output)
			assert.Equal(t, fmt.Sprintf("test-error-%d", i), err.Error())
		} else {
			assert.Nil(t, err)
			assert.Equal(t, fmt.Sprintf("test-output-%d", i), *output)
		}
	}

	output, err := mock.GetNextResult()
	assert.Nil(t, output)
	assert.Equal(t, fmt.Errorf("general error"), err)
}

func TestMockSetOutput(t *testing.T) {
	mock := gomock.GetMock[string, string](fmt.Errorf("general error"))

	for i := 0; i < 4; i++ {
		if i%2 == 0 {
			mock.SetError(i*2, fmt.Errorf("test-error-%d", i))
		} else {
			v := fmt.Sprintf("test-output-%d", i)
			mock.SetReturnValue(i*2, &v)
		}
	}

	for i := 0; i < 8; i++ {
		output, err := mock.GetNextResult()
		if i%4 == 0 {
			assert.Nil(t, output)
			assert.Equal(t, fmt.Sprintf("test-error-%d", i/2), err.Error())
		} else if i%2 == 0 {
			assert.Nil(t, err)
			assert.Equal(t, fmt.Sprintf("test-output-%d", i/2), *output)
		} else {
			assert.Nil(t, output)
			assert.Equal(t, fmt.Errorf("general error"), err)
		}
	}
}

func TestMockSetOutputPrefilled(t *testing.T) {
	mock := gomock.GetMock[string, string](fmt.Errorf("general error"))

	for i := 0; i < 6; i++ {
		v := fmt.Sprintf("test-old-%d", i)
		mock.AddReturnValue(&v)
	}

	for i := 0; i < 4; i++ {
		if i%2 == 0 {
			mock.SetError(i*2, fmt.Errorf("test-error-%d", i))
		} else {
			v := fmt.Sprintf("test-output-%d", i)
			mock.SetReturnValue(i*2, &v)
		}
	}

	for i := 0; i < 8; i++ {
		output, err := mock.GetNextResult()
		if i%4 == 0 {
			assert.Nil(t, output)
			assert.Equal(t, fmt.Sprintf("test-error-%d", i/2), err.Error())
		} else if i%2 == 0 {
			assert.Nil(t, err)
			assert.Equal(t, fmt.Sprintf("test-output-%d", i/2), *output)
		} else if i < 6 {
			assert.Nil(t, err)
			assert.Equal(t, fmt.Sprintf("test-old-%d", i), *output)
		} else {
			assert.Nil(t, output)
			assert.Equal(t, fmt.Errorf("general error"), err)
		}
	}

	output, err := mock.GetNextResult()
	assert.Nil(t, output)
	assert.Equal(t, fmt.Errorf("general error"), err)
}

func TestMockSetAlwaysReturn(t *testing.T) {
	mock := gomock.GetMock[string, string](fmt.Errorf("general error"))
	mock.SetAlwaysReturn("test-alwaysreturn")

	output, err := mock.GetNextResult()

	assert.Nil(t, err)
	assert.Equal(t, "test-alwaysreturn", *output)

	for i := 0; i < 6; i++ {
		v := fmt.Sprintf("test-old-%d", i)
		mock.AddReturnValue(&v)
	}

	for i := 0; i < 4; i++ {
		if i%2 == 0 {
			mock.SetError(i*2, fmt.Errorf("test-error-%d", i))
		} else {
			v := fmt.Sprintf("test-output-%d", i)
			mock.SetReturnValue(i*2, &v)
		}
	}

	for i := 0; i < 8; i++ {
		output, err := mock.GetNextResult()

		assert.Nil(t, err)
		assert.Equal(t, "test-alwaysreturn", *output)
	}

	output, err = mock.GetNextResult()
	assert.Nil(t, err)
	assert.Equal(t, "test-alwaysreturn", *output)
}

func TestMockSetAlwaysReturnFn(t *testing.T) {
	mock := gomock.GetMock[string, string](fmt.Errorf("general error"))
	mock.SetAlwaysReturnFn(func() (*string, error) {
		r := "test-alwaysreturnfn"
		return &r, nil
	})

	output, err := mock.GetNextResult()

	assert.Nil(t, err)
	assert.Equal(t, "test-alwaysreturnfn", *output)

	for i := 0; i < 6; i++ {
		v := fmt.Sprintf("test-old-%d", i)
		mock.AddReturnValue(&v)
	}

	for i := 0; i < 4; i++ {
		if i%2 == 0 {
			mock.SetError(i*2, fmt.Errorf("test-error-%d", i))
		} else {
			v := fmt.Sprintf("test-output-%d", i)
			mock.SetReturnValue(i*2, &v)
		}
	}

	for i := 0; i < 8; i++ {
		output, err := mock.GetNextResult()

		assert.Nil(t, err)
		assert.Equal(t, "test-alwaysreturnfn", *output)
	}

	output, err = mock.GetNextResult()
	assert.Nil(t, err)
	assert.Equal(t, "test-alwaysreturnfn", *output)
}

func TestMockReset(t *testing.T) {
	mock := gomock.GetMock[string, string](fmt.Errorf("general error"))

	for i := 0; i < 4; i++ {
		if i%2 == 0 {
			mock.AddError(fmt.Errorf("test-error-%d", i))
		} else {
			v := fmt.Sprintf("test-output-%d", i)
			mock.AddReturnValue(&v)
		}

		mock.AddInput(fmt.Sprintf("test-input-%d", i))
	}

	mock.SetAlwaysReturn("test-alwaysreturn")
	output, err := mock.GetNextResult()
	assert.Nil(t, err)
	assert.Equal(t, "test-alwaysreturn", *output)

	assert.Equal(t, 4, mock.HasBeenCalled())
	mock.Reset()

	assert.Equal(t, 0, mock.HasBeenCalled())
	output, err = mock.GetNextResult()
	assert.Nil(t, output)
	assert.Equal(t, fmt.Errorf("general error"), err)

	mock.SetAlwaysReturnFn(func() (*string, error) {
		r := "test-alwaysreturnfn"
		return &r, nil
	})
	output, err = mock.GetNextResult()
	assert.Nil(t, err)
	assert.Equal(t, "test-alwaysreturnfn", *output)

	mock.Reset()

	assert.Equal(t, 0, mock.HasBeenCalled())
	output, err = mock.GetNextResult()
	assert.Nil(t, output)
	assert.Equal(t, fmt.Errorf("general error"), err)

}
