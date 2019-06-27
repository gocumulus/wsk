package wsk

import "os"

type envVariables struct {
	APIHost            string `json:"api_host"`
	APIKey             string `json:"api_key"`
	Namespace          string `json:"namespace"`
	NamespaceCRN       string `json:"namespace_crn"`
	ActionName         string `json:"action_name"`
	IAMNamespaceAPIKey string `json:"iam_namespace_api_key"`
	IAMAPIURL          string `json:"iam_api_url"`
	ActivationID       string `json:"activation_id"`
	Deadline           string `json:"deadline"`
}

func (env envVariables) setVariables() error {
	vars := map[string]string{
		"__OW_API_HOST":              env.APIHost,
		"__OW_API_KEY":               env.APIKey,
		"__OW_NAMESPACE":             env.Namespace,
		"__OW_NAMESPACE_CRN":         env.NamespaceCRN,
		"__OW_ACTION_NAME":           env.ActionName,
		"__OW_IAM_NAMESPACE_API_KEY": env.IAMNamespaceAPIKey,
		"__OW_IAM_API_URL":           env.IAMAPIURL,
		"__OW_ACTIVATION_ID":         env.ActivationID,
		"__OW_DEADLINE":              env.Deadline,
	}
	for k, v := range vars {
		err := setVariable(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func setVariable(key, value string) error {
	if value == "" {
		return nil
	}
	return os.Setenv(key, value)
}
