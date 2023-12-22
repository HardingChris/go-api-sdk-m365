// graphbeta_device_management_group_policy_configurations.go
// Graph Beta Api - Intune: Group Policy Configuration Profiles
// Documentation: https://learn.microsoft.com/en-us/mem/intune/configuration/group-policy-analytics
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesWindowsMenu/~/configProfiles
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-grouppolicy-grouppolicyconfiguration?view=graph-rest-beta
// ODATA query options reference: https://learn.microsoft.com/en-us/graph/query-parameters?tabs=http
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"
	"time"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

// Constant for the endpoint URL
const uriBetaDeviceManagementGroupPolicyConfigurations = "/beta/deviceManagement/groupPolicyConfigurations"

/*
// Example of struct hierarchy using embedded anonymous structs
type ResponseDeviceManagementGroupPolicyConfigurationsList struct {
	ODataContext string `json:"@odata.context"`
	Value []struct {
		OdataType                        string    `json:"@odata.type"`
		ID                               string    `json:"id"`
		DisplayName                      string    `json:"displayName"`
		Description                      string    `json:"description"`
		RoleScopeTagIds                  []string  `json:"roleScopeTagIds"`
		PolicyConfigurationIngestionType string    `json:"policyConfigurationIngestionType"`
		CreatedDateTime                  time.Time `json:"createdDateTime"`
		LastModifiedDateTime             time.Time `json:"lastModifiedDateTime"`
		DefinitionValues []struct {
			ID                   string    `json:"id"`
			Enabled              bool      `json:"enabled"`
			ConfigurationType    string    `json:"configurationType"`
			CreatedDateTime      time.Time `json:"createdDateTime"`
			LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
			Definition struct {
				ID          string `json:"id"`
				DisplayName string `json:"displayName"`
				Description string `json:"description"`
			} `json:"definition,omitempty"`
			PresentationValues []struct {
				ID                   string    `json:"id"`
				LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
				CreatedDateTime      time.Time `json:"createdDateTime"`
				Label                string    `json:"label"`
				Description          string    `json:"description"`
				ValueType            string    `json:"valueType"`
				Value                interface{} `json:"value"`
				Presentation struct {
					Label       string `json:"label"`
					ID          string `json:"id"`
					Required    bool   `json:"required"`
					DefaultItem struct {
						DisplayName string `json:"displayName"`
						Value       string `json:"value"`
					} `json:"defaultItem"`
					Items []struct {
						DisplayName string `json:"displayName"`
						Value       string `json:"value"`
					} `json:"items"`
				} `json:"presentation,omitempty"`
			} `json:"presentationValues,omitempty"`
		} `json:"definitionValues,omitempty"`
		Assignments []struct {
			ID                   string    `json:"id"`
			LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
			Target struct {
				ID                                         string `json:"id"`
				Type                                       string `json:"@odata.type"`
				DeviceAndAppManagementAssignmentFilterId   string `json:"deviceAndAppManagementAssignmentFilterId"`
				DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
				CollectionId                               string `json:"collectionId"`
			} `json:"target"`
		} `json:"assignments,omitempty"`
	} `json:"value"`
}
*/

// ResponseDeviceManagementGroupPolicyConfigurationsList is used to parse the list response of Group Policy Configurations from Microsoft Graph API.
type ResponseDeviceManagementGroupPolicyConfigurationsList struct {
	ODataContext string                                             `json:"@odata.context"`
	Value        []ResourceDeviceManagementGroupPolicyConfiguration `json:"value"`
}

// ResourceDeviceManagementGroupPolicyConfiguration represents an individual Group Policy Configuration resource from Microsoft Graph API.
type ResourceDeviceManagementGroupPolicyConfiguration struct {
	OdataType                        string                       `json:"@odata.type"`
	ID                               string                       `json:"id"`
	DisplayName                      string                       `json:"displayName"`
	Description                      string                       `json:"description"`
	RoleScopeTagIds                  []string                     `json:"roleScopeTagIds"`
	PolicyConfigurationIngestionType string                       `json:"policyConfigurationIngestionType"`
	CreatedDateTime                  time.Time                    `json:"createdDateTime"`
	LastModifiedDateTime             time.Time                    `json:"lastModifiedDateTime"`
	DefinitionValues                 []GroupPolicyDefinitionValue `json:"definitionValues,omitempty"`
	Assignments                      []Assignment                 `json:"assignments,omitempty"`
}

// ResponseGroupPolicyDefinitionValuesList is used to parse the list response of Group Policy Definition Values from Microsoft Graph API.
type ResponseGroupPolicyDefinitionValuesList struct {
	Value []GroupPolicyDefinitionValue `json:"value"`
}

// GroupPolicyDefinitionValue represents a single Group Policy Definition Value, including its associated definitions and presentation values.
type GroupPolicyDefinitionValue struct {
	ID                   string                         `json:"id"`
	Enabled              bool                           `json:"enabled"`
	ConfigurationType    string                         `json:"configurationType"`
	CreatedDateTime      time.Time                      `json:"createdDateTime"`
	LastModifiedDateTime time.Time                      `json:"lastModifiedDateTime"`
	Definition           *GroupPolicyDefinition         `json:"definition,omitempty"`
	PresentationValues   []GroupPolicyPresentationValue `json:"presentationValues,omitempty"`
}

// GroupPolicyDefinition represents the basic information of a Group Policy Definition.
type GroupPolicyDefinition struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}

// ResponsePresentationValuesList is used to parse the list response of Group Policy Presentation Values from Microsoft Graph API.
type ResponsePresentationValuesList struct {
	Value []GroupPolicyPresentationValue `json:"value"`
}

// GroupPolicyPresentationValue represents a presentation value for a Group Policy Definition Value, including its type and value.
type GroupPolicyPresentationValue struct {
	ID                   string                  `json:"id"`
	LastModifiedDateTime time.Time               `json:"lastModifiedDateTime"`
	CreatedDateTime      time.Time               `json:"createdDateTime"`
	Label                string                  `json:"label"`
	Description          string                  `json:"description"`
	ValueType            string                  `json:"valueType"`
	Value                DynamicValue            `json:"value"`
	Presentation         GroupPolicyPresentation `json:"presentation,omitempty"`
}

// GroupPolicyPresentation defines the presentation details for a Group Policy Presentation Value, such as labels, items, and default values.
type GroupPolicyPresentation struct {
	Label       string             `json:"label"`
	ID          string             `json:"id"`
	Required    bool               `json:"required"`
	DefaultItem PresentationItem   `json:"defaultItem"`
	Items       []PresentationItem `json:"items"`
}

// PresentationItem represents an individual item in a Group Policy Presentation dropdown or similar collection.
type PresentationItem struct {
	DisplayName string `json:"displayName"`
	Value       string `json:"value"`
}

// DynamicValue is a type that can hold different types of values, allowing for dynamic handling of the 'value' field in Group Policy Presentation Values.
type DynamicValue struct {
	// Use an interface to hold the actual value.
	Value interface{} `json:"Value"`
}

// ResponseAssignmentsList is used to parse the list response of Assignments from Microsoft Graph API.
type ResponseAssignmentsList struct {
	Value []Assignment `json:"value"`
}

// Assignment represents an assignment of a Group Policy Configuration to a target, such as a user or a group.
type Assignment struct {
	ID                   string           `json:"id"`
	LastModifiedDateTime time.Time        `json:"lastModifiedDateTime"`
	Target               AssignmentTarget `json:"target"`
}

// AssignmentTarget represents the target of an Assignment, detailing the type and identifiers for the assignment target.
type AssignmentTarget struct {
	ID                                         string `json:"id"`
	Type                                       string `json:"@odata.type"` // e.g., "#microsoft.graph.groupTarget", etc.
	DeviceAndAppManagementAssignmentFilterId   string `json:"deviceAndAppManagementAssignmentFilterId"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
	CollectionId                               string `json:"collectionId"`
}

// Function to get the list of Group Policy Configurations
func (c *Client) GetDeviceManagementGroupPolicyConfigurations() ([]ResourceDeviceManagementGroupPolicyConfiguration, error) {
	endpoint := uriBetaDeviceManagementGroupPolicyConfigurations

	var responseGroupPolicyConfigurations ResponseDeviceManagementGroupPolicyConfigurationsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseGroupPolicyConfigurations)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device management group policy configurations", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return responseGroupPolicyConfigurations.Value, nil
}

// GetDeviceManagementGroupPolicyConfigurationByID retrieves a specific Group Policy Configuration by its ID with expanded details.
func (c *Client) GetDeviceManagementGroupPolicyConfigurationByID(policyConfigurationId string) (*ResourceDeviceManagementGroupPolicyConfiguration, error) {
	// Retrieve the base Group Policy Configuration
	baseEndpoint := fmt.Sprintf("%s/%s", uriBetaDeviceManagementGroupPolicyConfigurations, policyConfigurationId)
	var baseConfig ResourceDeviceManagementGroupPolicyConfiguration
	_, err := c.HTTP.DoRequest("GET", baseEndpoint, nil, &baseConfig)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "group policy configuration", policyConfigurationId, err)
	}

	// Retrieve Definition Values and expand each definition
	defValuesEndpoint := fmt.Sprintf("%s/definitionValues?$expand=definition", baseEndpoint)
	var defValuesList ResponseGroupPolicyDefinitionValuesList
	_, err = c.HTTP.DoRequest("GET", defValuesEndpoint, nil, &defValuesList)
	if err != nil {
		return nil, fmt.Errorf("failed to get definition values: %v", err)
	}

	// For each Definition Value, retrieve and expand Presentation Values
	for i, defValue := range defValuesList.Value {
		presentationEndpoint := fmt.Sprintf("%s/definitionValues/%s/presentationValues?$expand=presentation", baseEndpoint, defValue.ID)
		var presentationList ResponsePresentationValuesList
		_, err = c.HTTP.DoRequest("GET", presentationEndpoint, nil, &presentationList)
		if err != nil {
			return nil, fmt.Errorf("failed to get presentation values: %v", err)
		}
		defValuesList.Value[i].PresentationValues = presentationList.Value
	}

	// Attach expanded Definition Values to the base configuration
	baseConfig.DefinitionValues = defValuesList.Value

	// Retrieve Assignments
	assignmentsEndpoint := fmt.Sprintf("%s/assignments", baseEndpoint)
	var assignmentsList ResponseAssignmentsList
	_, err = c.HTTP.DoRequest("GET", assignmentsEndpoint, nil, &assignmentsList)
	if err != nil {
		return nil, fmt.Errorf("failed to get assignments: %v", err)
	}

	// Attach Assignments to the base configuration
	baseConfig.Assignments = assignmentsList.Value

	return &baseConfig, nil
}
