package serializerentities

import (
	"reflect"

	di "github.com/sarulabs/di/v2"

	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	jsonapi "github.com/gnemes/go-users/Infrastructure/Services/Jsonapi"
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
	serializers "github.com/gnemes/go-users/Domain/Serializers"
	serializersimpl "github.com/gnemes/go-users/Infrastructure/Serializers"
)

type BaseSerializerEntity struct {
	Container           di.Container                              `json:"-"`
	Logger              logger.Logger                             `json:"-"`
	ID                  string                                    `json:"-"`
	ToOneRelationships  []serializersimpl.Relationship            `json:"-"`
	ToManyRelationships map[string][]serializersimpl.Relationship `json:"-"`
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

func (bse *BaseSerializerEntity) AddToOneRelationship(SerializerName string, e entities.Entity, Type string, Name string) {
	var entity serializers.SerializerEntity

	if !reflect.ValueOf(e).IsNil() {
		s := bse.newSerializer(SerializerName)
		if s != nil {
			err := s.Fill(e)
			if err == nil {
				entity = s.GetSerializerEntity()
			}
		}
	}

	relationship := serializersimpl.Relationship{
		Type: Type,
		Name: Name,
		Entity: entity,
	}

	bse.ToOneRelationships = append(bse.ToOneRelationships, relationship)
}

func (bse *BaseSerializerEntity) AddToManyRelationship(SerializerName string, d interface{}, Type string, Name string) {
	if e, ok := d.([]entities.Entity); ok {
		for _, entity := range e {
			s := bse.newSerializer(SerializerName)
			if s != nil {
				err := s.Fill(entity)
				if err == nil {
					relationship := serializersimpl.Relationship{
						Type: Type,
						Name: Name,
						Entity: s.GetSerializerEntity(),
					}

					bse.ToManyRelationships[Name] = append(bse.ToManyRelationships[Name], relationship)
				}
			}
		}
	}

	if bse.ToManyRelationships[Name] == nil || len(bse.ToManyRelationships[Name]) == 0 {
		relationship := serializersimpl.Relationship{
			Type: Type,
			Name: Name,
		}
		bse.ToManyRelationships[Name] = append(bse.ToManyRelationships[Name], relationship)
	}
}

func (bse *BaseSerializerEntity) GetReferences() []jsonapi.Reference {
	var references []jsonapi.Reference

	for _, relationship := range bse.ToOneRelationships {
		reference := jsonapi.Reference{
			Type:        relationship.Type,
			Name:        relationship.Name,
			IsNotLoaded: relationship.IsNotLoaded,
		}

		references = append(references, reference)
	}

	for _, relationship := range bse.ToManyRelationships {
		first := relationship[0]
		reference := jsonapi.Reference{
			Type:        first.Type,
			Name:        first.Name,
			IsNotLoaded: first.IsNotLoaded,
		}

		references = append(references, reference)
	}

	return references
}

func (bse *BaseSerializerEntity) GetReferencedIDs() []jsonapi.ReferenceID {
	var references []jsonapi.ReferenceID

	for _, relationship := range bse.ToOneRelationships {
		if relationship.Entity != nil {
			reference := jsonapi.ReferenceID{
				Type: relationship.Type, 
				Name: relationship.Name,
				ID: relationship.Entity.GetID(),
			}
			references = append(references, reference)
		}
	}

	return references
}

func (bse *BaseSerializerEntity) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	var includes []jsonapi.MarshalIdentifier

	for _, relationship := range bse.ToOneRelationships {
		// if s.Query.Include(jsonapitypes.RecoProgramBudgetsJSONAPIName) {
		if true {
			if relationship.Entity != nil {
				includes = append(includes, relationship.Entity)
			}
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