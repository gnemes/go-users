package presenter

import (
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
	serializers "github.com/gnemes/go-users/Infrastructure/Serializers"
	usecases "github.com/gnemes/go-users/Domain/UseCases"
)

type JsonApiPresenter struct {
	Logger     logger.Logger
	OutputPort usecases.OutputPort
	Serializer *serializers.Serializer
}

func (p *JsonApiPresenter) Present() ([]byte, error) {
	p.Logger.Debugf("Controller / JsonApiPresenter / Present()")
	defer p.Logger.Debugf("Controller / JsonApiPresenter / Present() ending...")

	data := p.OutputPort.GetData()

	response, err := p.Serializer.Serialize(data, nil)
	if err != nil {
		p.Logger.Debugf("Error serializing Use Case response: %s", err.Error())
		return nil, err
	} 

	return response, nil
}