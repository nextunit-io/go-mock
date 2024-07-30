package gomock

type awsToolsMockReturnValue[U any] struct {
	Value *U
	Error error
}

type awsToolMock[T any, U any] struct {
	generalError error
	inputParams  []T
	returnValues []awsToolsMockReturnValue[U]
}

func GetMock[T any, U any](generalError error) *awsToolMock[T, U] {
	return &awsToolMock[T, U]{
		generalError: generalError,
		inputParams:  []T{},
	}
}

func (mock *awsToolMock[T, U]) AddInput(i T) {
	if mock.inputParams == nil {
		mock.inputParams = []T{}
	}

	mock.inputParams = append(mock.inputParams, i)
}

func (mock awsToolMock[T, U]) GetInputs() []T {
	return mock.inputParams
}

func (mock awsToolMock[T, U]) HasBeenCalled() int {
	return len(mock.inputParams)
}

func (mock awsToolMock[T, U]) GetInput(position int) T {
	return mock.inputParams[position]
}

func (mock *awsToolMock[T, U]) AddReturnValue(value *U) {
	if mock.returnValues == nil {
		mock.returnValues = []awsToolsMockReturnValue[U]{}
	}

	mock.returnValues = append(mock.returnValues, awsToolsMockReturnValue[U]{
		Value: value,
	})
}

func (mock *awsToolMock[T, U]) prepareReturnArray(position int) {
	if mock.returnValues == nil {
		mock.returnValues = []awsToolsMockReturnValue[U]{}
	}

	if position >= len(mock.returnValues) {
		r := make([]awsToolsMockReturnValue[U], position+1)
		for i := 0; i < position; i++ {
			if i < len(mock.returnValues) {
				r[i] = mock.returnValues[i]
			}
		}

		mock.returnValues = r
	}
}

func (mock *awsToolMock[T, U]) SetReturnValue(position int, value *U) {
	mock.prepareReturnArray(position)

	if position < len(mock.returnValues) {
		mock.returnValues[position] = awsToolsMockReturnValue[U]{
			Value: value,
		}
	} else {
		mock.returnValues[position] = awsToolsMockReturnValue[U]{
			Value: value,
		}
	}
}

func (mock *awsToolMock[T, U]) AddError(err error) {
	if mock.returnValues == nil {
		mock.returnValues = []awsToolsMockReturnValue[U]{}
	}

	mock.returnValues = append(mock.returnValues, awsToolsMockReturnValue[U]{
		Error: err,
	})
}

func (mock *awsToolMock[T, U]) SetError(position int, err error) {
	mock.prepareReturnArray(position)

	if position < len(mock.returnValues) {
		mock.returnValues[position] = awsToolsMockReturnValue[U]{
			Error: err,
		}
	} else {
		mock.returnValues[position] = awsToolsMockReturnValue[U]{
			Error: err,
		}
	}
}

func (mock *awsToolMock[T, U]) GetNextResult() (*U, error) {
	if mock.returnValues == nil {
		return nil, mock.generalError
	}

	if len(mock.returnValues) < 1 {
		return nil, mock.generalError
	}

	var x awsToolsMockReturnValue[U]
	x, mock.returnValues = mock.returnValues[0], mock.returnValues[1:]

	if x.Error != nil {
		return nil, x.Error
	}

	if x.Value == nil {
		return nil, mock.generalError
	}

	return x.Value, nil
}

func (mock *awsToolMock[T, U]) Reset() {
	mock.returnValues = nil
	mock.inputParams = []T{}
}
