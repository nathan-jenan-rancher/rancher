package clusteregistrationtokens

import (
	"fmt"

	"encoding/json"

	"github.com/rancher/norman/httperror"
	"github.com/rancher/norman/types"
	"github.com/rancher/types/apis/management.cattle.io/v3"
	"github.com/rancher/types/client/management/v3"
)

func Validator(_ *types.APIContext, _ *types.Schema, data map[string]interface{}) error {
	if data[client.ClusterFieldRancherKubernetesEngineConfig] != nil {
		jsonBlob, err := json.Marshal(data)
		if err != nil {
			return httperror.NewAPIError(httperror.InvalidBodyContent, fmt.Sprintf("unparsable rke config: %v",
				err))
		}

		rkeConfig := v3.RancherKubernetesEngineConfig{}
		err = json.Unmarshal(jsonBlob, rkeConfig)
		if err != nil {
			return httperror.NewAPIError(httperror.InvalidBodyContent, fmt.Sprintf(
				"error marshalling config: %v", err))
		}

		if rkeConfig.Services.KubeAPI.PodSecurityPolicy &&
			data[client.ClusterFieldDefaultPodSecurityPolicyTemplateId] == nil {
			return httperror.NewAPIError(httperror.InvalidBodyContent, fmt.Sprintf(
				"PSP controller cannot be enabled without a default PSPT set"))
		}
	}

	return nil
}
