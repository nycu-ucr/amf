package callback

import (
	"context"
	"fmt"

	amf_context "github.com/free5gc/amf/context"
	"github.com/nycu-ucr/openapi/Namf_Communication"
	"github.com/nycu-ucr/openapi/models"
)

func SendN2InfoNotifyN2Handover(ue *amf_context.AmfUe, releaseList []int32) error {
	if ue.HandoverNotifyUri == "" {
		return fmt.Errorf("N2 Info Notify N2Handover failed(uri dose not exist)")
	}
	configuration := Namf_Communication.NewConfiguration()
	client := Namf_Communication.NewAPIClient(configuration)

	n2InformationNotification := models.N2InformationNotification{
		N2NotifySubscriptionId: ue.Supi,
		ToReleaseSessionList:   releaseList,
		NotifyReason:           models.N2InfoNotifyReason_HANDOVER_COMPLETED,
	}

	_, httpResponse, err := client.N2MessageNotifyCallbackDocumentApiServiceCallbackDocumentApi.
		N2InfoNotify(context.Background(), ue.HandoverNotifyUri, n2InformationNotification)

	if err == nil {
		// TODO: handle Msg
	} else {
		if httpResponse == nil {
			HttpLog.Errorln(err.Error())
		} else if err.Error() != httpResponse.Status {
			HttpLog.Errorln(err.Error())
		}
	}
	return nil
}
