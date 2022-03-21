package serializerentities

import (
	di "github.com/sarulabs/di/v2"

	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	jsonapi "github.com/gnemes/go-users/Infrastructure/Services/Jsonapi"
	serializers "github.com/gnemes/go-users/Domain/Serializers"
	serializersimpl "github.com/gnemes/go-users/Infrastructure/Serializers"
)

type BaseSerializerEntity struct {
	Container     di.Container                   `json:"-"`
	ID            string                         `json:"-"`
	Relationships []serializersimpl.Relationship `json:"-"`
}

func (bse *BaseSerializerEntity) GetID() string {
	return bse.ID
}

func (bse *BaseSerializerEntity) SetID(id string) error {
	bse.ID = id
	return nil
}

func (bse *BaseSerializerEntity) Serialize() ([]byte, error) {
	return nil, nil
}

func (bse *BaseSerializerEntity) AddRelationship(SerializerName string, e entities.Entity, Type string, Name string) {
	s := bse.newSerializer(SerializerName)
	if s != nil {
		err := s.Fill(e)
		if err == nil {
			relationship := serializersimpl.Relationship{
				Type: Type,
				Name: Name,
				Entity: s.GetSerializerEntity(),
			}
		
			bse.Relationships = append(bse.Relationships, relationship)
		}
	}
}

func (bse *BaseSerializerEntity) GetReferences() []jsonapi.Reference {
	var references []jsonapi.Reference

	for _, relationship := range bse.Relationships {
		reference := jsonapi.Reference{
			Type:        relationship.Type,
			Name:        relationship.Name,
			IsNotLoaded: relationship.IsNotLoaded,
		}

		references = append(references, reference)
	}


	return references
}

func (bse *BaseSerializerEntity) GetReferencedIDs() []jsonapi.ReferenceID {
	var references []jsonapi.ReferenceID

	for _, relationship := range bse.Relationships {
		reference := jsonapi.ReferenceID{
			Type: relationship.Type, 
			Name: relationship.Name,
			ID: relationship.Entity.GetID(),
		}
		references = append(references, reference)
	}

	return references
}

func (bse *BaseSerializerEntity) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	var includes []jsonapi.MarshalIdentifier

	for _, relationship := range bse.Relationships {
		// if s.Query.Include(jsonapitypes.RecoProgramBudgetsJSONAPIName) {
		if true {
			includes = append(includes, relationship.Entity)
		}
	}

	return includes
}

func (bse *BaseSerializerEntity) newSerializer(name string) serializers.Serializer {
	serializer, err := bse.Container.SafeGet(name)
	if err != nil {
		return nil
	}

	return serializer.(serializers.Serializer)
}