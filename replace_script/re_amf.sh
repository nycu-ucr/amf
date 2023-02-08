#!/bin/bash
s="util/convert.go
util/json.go
util/init_context.go
oam/api_registered_ue_context.go
oam/routers.go
mt/api_ue_reach_ind_document.go
mt/api_ue_context_document.go
mt/routers.go
eventexposure/api_individual_subscription_document.go
eventexposure/api_subscriptions_collection_document.go
eventexposure/routers.go
service/init.go
ngap/handler.go
ngap/service/service.go
ngap/message/send.go
ngap/message/build.go
ngap/message/send_test.go.test
ngap/message/forward_ie.go
ngap/dispatcher.go
factory/factory.go
producer/n1n2message.go
producer/subscription.go
producer/oam.go
producer/location_info.go
producer/mt.go
producer/callback/n1n2message.go
producer/callback/subscription.go
producer/callback/ue_context.go
producer/callback.go
producer/ue_context.go
producer/event_exposure.go
amf.go
consumer/am_policy.go
consumer/nf_mangement.go
consumer/ue_context_management.go
consumer/nsselection.go
consumer/subscriber_data_management.go
consumer/nf_discovery.go
consumer/ue_authentication.go
consumer/sm_context.go
consumer/communication.go
gmm/sm.go
gmm/init.go
gmm/handler.go
gmm/init_test.go
gmm/message/send.go
gmm/message/build.go
httpcallback/router.go
httpcallback/api_n1_message_notify.go
httpcallback/api_am_policy_control_update_notify.go
httpcallback/api_sm_context_status_notify.go
context/common_function.go
context/context.go
context/amf_ue.go
context/amf_ran.go
context/ran_ue.go
communication/api_individual_subscription_document.go
communication/api_n1_n2_message_collection_document.go
communication/api_non_uen2_message_notification_individual_subscription_document.go
communication/api_n1_n2_individual_subscription_document.go
communication/api_non_uen2_messages_subscriptions_collection_document.go
communication/api_individual_ue_context_document.go
communication/api_n1_n2_subscriptions_collection_for_individual_ue_contexts_document.go
communication/api_subscriptions_collection_document.go
communication/routers.go
communication/api_non_uen2_messages_collection_document.go
nas/handler.go
nas/nas_security/security.go
nas/dispatch.go
location/api_individual_ue_context_document.go
location/routers.go"
for f in $s
do
    sed -i 's#github.com/free5gc/amf#github.com/nycu-ucr/amf#g' $f
done