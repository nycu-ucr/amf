#!/bin/bash
s="service/init.go
ngap/handler.go
ngap/service/service.go
ngap/message/send.go
ngap/message/build.go
ngap/message/send_test.go.test
ngap/message/forward_ie.go
ngap/dispatcher.go
producer/n1n2message.go
producer/callback.go
gmm/handler.go
gmm/message/send.go
context/amf_ran.go
context/ran_ue.go"
for f in $s
do
    sed -i 's#github.com/free5gc/ngap#github.com/nycu-ucr/ngap#g' $f
done
