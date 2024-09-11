package gomock

type toolsMockReturnValue[U any] struct {
	Value *U
	Error error
}

type ToolMock[T any, U any] struct {
	generalError   error
	inputParams    []T
	alwaysReturn   *U
	alwaysReturnFn func() (*U, error)
	returnValues   []toolsMockReturnValue[U]
}

func GetMock[T any, U any](generalError error) *ToolMock[T, U] {
	return &ToolMock[T, U]{
		generalError: generalError,
		inputParams:  []T{},
	}
}

func (mock *ToolMock[T, U]) AddInput(i T) {
	if mock.inputParams == nil {
		mock.inputParams = []T{}
	}

	mock.inputParams = append(mock.inputParams, i)
}

func (mock ToolMock[T, U]) GetInputs() []T {
	return mock.inputParams
}

func (mock ToolMock[T, U]) HasBeenCalled() int {
	return len(mock.inputParams)
}

func (mock ToolMock[T, U]) GetInput(position int) T {
	return mock.inputParams[position]
}

func (mock ToolMock[T, U]) GetLastInput() T {
	return mock.GetInput(mock.HasBeenCalled() - 1)
}

func (mock *ToolMock[T, U]) AddReturnValue(value *U) {
	if mock.returnValues == nil {
		mock.returnValues = []toolsMockReturnValue[U]{}
	}

	mock.returnValues = append(mock.returnValues, toolsMockReturnValue[U]{
		Value: value,
	})
}

func (mock *ToolMock[T, U]) prepareReturnArray(position int) {
	if mock.returnValues == nil {
		mock.returnValues = []toolsMockReturnValue[U]{}
	}

	if position >= len(mock.returnValues) {
		r := make([]toolsMockReturnValue[U], position+1)
		for i := 0; i < position; i++ {
			if i < len(mock.returnValues) {
				r[i] = mock.returnValues[i]
			}
		}

		mock.returnValues = r
	}
}

func (mock *ToolMock[T, U]) SetAlwaysReturn(value U) {
	mock.alwaysReturn = &value
}

func (mock *ToolMock[T, U]) SetAlwaysReturnFn(fn func() (*U, error)) {
	mock.alwaysReturnFn = fn
}

func (mock *ToolMock[T, U]) SetReturnValue(position int, value *U) {
	mock.prepareReturnArray(position)

	if position < len(mock.returnValues) {
		mock.returnValues[position] = toolsMockReturnValue[U]{
			Value: value,
		}
	} else {
		mock.returnValues[position] = toolsMockReturnValue[U]{
			Value: value,
		}
	}
}

func (mock *ToolMock[T, U]) AddError(err error) {
	if mock.returnValues == nil {
		mock.returnValues = []toolsMockReturnValue[U]{}
	}

	mock.returnValues = append(mock.returnValues, toolsMockReturnValue[U]{
		Error: err,
	})
}

func (mock *ToolMock[T, U]) SetError(position int, err error) {
	mock.prepareReturnArray(position)

	if position < len(mock.returnValues) {
		mock.returnValues[position] = toolsMockReturnValue[U]{
			Error: err,
		}
	} else {
		mock.returnValues[position] = toolsMockReturnValue[U]{
			Error: err,
		}
	}
}

func (mock *ToolMock[T, U]) GetNextResult() (*U, error) {
	if mock.alwaysReturn != nil {
		return mock.alwaysReturn, nil
	}

	if mock.alwaysReturnFn != nil {
		return mock.alwaysReturnFn()
	}

	if mock.returnValues == nil {
		return nil, mock.generalError
	}

	if len(mock.returnValues) < 1 {
		return nil, mock.generalError
	}

	var x toolsMockReturnValue[U]
	x, mock.returnValues = mock.returnValues[0], mock.returnValues[1:]

	if x.Error != nil {
		return nil, x.Error
	}

	if x.Value == nil {
		return nil, mock.generalError
	}

	return x.Value, nil
}

func (mock *ToolMock[T, U]) Reset() {
	mock.returnValues = nil
	mock.inputParams = []T{}
	mock.alwaysReturn = nil
	mock.alwaysReturnFn = nil
}
