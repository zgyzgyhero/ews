package ewsutil

import (
	"github.com/zgyzgyhero/ews"
)

// FindPeople find persona slice by query string
func GetPersona(c ews.Client, personaID string) (*ews.Persona, error) {

	resp, err := ews.GetPersona(c, &ews.GetPersonaRequest{
		PersonaId: ews.PersonaId{Id: personaID},
	})

	if err != nil {
		return nil, err
	}

	return &resp.Persona, nil
}
