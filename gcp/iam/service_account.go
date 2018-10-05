package iam

import (
	"fmt"
)

type ServiceAccount struct {
	client serviceAccountsClient
	name   string
	email  string
}

func NewServiceAccount(client serviceAccountsClient, name, email string) ServiceAccount {
	return ServiceAccount{
		client: client,
		name:   name,
		email:  email,
	}
}

func (s ServiceAccount) Delete() error {
	// err := s.client.DeleteServiceAccount(s.name)
	// if err != nil {
	// 	return fmt.Errorf("Delete: %s", err)
	// }

	//Remove it from the project?
	policy, err := s.client.GetProjectIamPolicy()
	if err != nil {
		panic(err)
	}

	// What roles does this service account have?
	toRemove := []*gcpcrm.Binding{}

	for index, binding := range policy.Bindings {
		for _, member := range binding.Members {
			if member == fmt.Sprintf("serviceAccount:%s", s.email) {
				toRemove = append(toRemove, index)
				fmt.Printf("FOUND Member: %s, Role: %s \n", member, binding.Role)
			}
		}
	}

	for _, b := toRemove {
		p.Bindings = append(policy.Bindings[:toRemove], p.Bindings[toRemove+1:]...)
	}

	fmt.Printf("Final Bindings: %+v \n", p.Bindings)	

	fmt.Printf("%+v \n", bindingDeltas)

	return nil
}

func (s ServiceAccount) Name() string {
	return s.name
}

func (s ServiceAccount) Type() string {
	return "IAM Service Account"
}
