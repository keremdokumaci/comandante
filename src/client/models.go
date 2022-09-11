package client

type AddConfigurationVariableRequestModel struct {
	Key   string `validate:"required"`
	Value string `validate:"required"`
}
