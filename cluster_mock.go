package main

import (
	"bytes"
	"context"
)

// MockConsulSource implements ClusterSource for local
// mock of cluster service.
type MockConsulSource struct {
}

// NewMockConsulSource creates new mocked source for tests.
func NewMockConsulSource() ClusterSource {
	return &MockConsulSource{}
}

// Node returns the list of mock nodes for the given datacentre 'dc' and tag.
// Satisfies ClusterSource interface.
func (c *MockConsulSource) IPs(ctx context.Context, tag string) ([]string, error) {
	r := bytes.NewBufferString(mockClusterIPsJSON)
	return ParseConsulResponse(r)
}

const mockClusterIPsJSON = `[{"ID":"edaddefa-f894-703a-69db-6158dd56aa5a","Node":"mail-01.do-ams3.eth.beta","Address":"206.189.243.162","Datacenter":"do-ams3","TaggedAddresses":{"lan":"206.189.243.162","wan":"206.189.243.162"},"NodeMeta":{"consul-network-segment":"","env":"eth","stage":"beta"},"ServiceID":"statusd-mail-rpc-1","ServiceName":"statusd-rpc","ServiceTags":["eth.beta","statusd","mail","rpc"],"ServiceAddress":"10.1.0.13","ServiceMeta":{},"ServicePort":8546,"ServiceEnableTagOverride":true,"CreateIndex":719808,"ModifyIndex":719808},{"ID":"d62fa419-49f7-32e5-52f9-64478b4e104b","Node":"mail-02.do-ams3.eth.beta","Address":"206.189.243.169","Datacenter":"do-ams3","TaggedAddresses":{"lan":"206.189.243.169","wan":"206.189.243.169"},"NodeMeta":{"consul-network-segment":"","env":"eth","stage":"beta"},"ServiceID":"statusd-mail-rpc-1","ServiceName":"statusd-rpc","ServiceTags":["eth.beta","statusd","mail","rpc"],"ServiceAddress":"10.1.0.14","ServiceMeta":{},"ServicePort":8546,"ServiceEnableTagOverride":true,"CreateIndex":718197,"ModifyIndex":718197},{"ID":"ad5cbd92-b28e-d5fa-3e49-17674eb91de9","Node":"mail-03.do-ams3.eth.beta","Address":"206.189.243.168","Datacenter":"do-ams3","TaggedAddresses":{"lan":"206.189.243.168","wan":"206.189.243.168"},"NodeMeta":{"consul-network-segment":"","env":"eth","stage":"beta"},"ServiceID":"statusd-mail-rpc-1","ServiceName":"statusd-rpc","ServiceTags":["eth.beta","statusd","mail","rpc"],"ServiceAddress":"10.1.0.12","ServiceMeta":{},"ServicePort":8546,"ServiceEnableTagOverride":true,"CreateIndex":725592,"ModifyIndex":725592},{"ID":"3b6ae627-d8c2-c39a-0628-0792ad1c46e4","Node":"node-01.do-ams3.eth.beta","Address":"206.189.243.176","Datacenter":"do-ams3","TaggedAddresses":{"lan":"206.189.243.176","wan":"206.189.243.176"},"NodeMeta":{"consul-network-segment":"","env":"eth","stage":"beta"},"ServiceID":"statusd-whisper-rpc-1","ServiceName":"statusd-rpc","ServiceTags":["eth.beta","statusd","whisper","rpc"],"ServiceAddress":"10.1.1.99","ServiceMeta":{},"ServicePort":8546,"ServiceEnableTagOverride":true,"CreateIndex":719848,"ModifyIndex":719848},{"ID":"3b6ae627-d8c2-c39a-0628-0792ad1c46e4","Node":"node-01.do-ams3.eth.beta","Address":"206.189.243.176","Datacenter":"do-ams3","TaggedAddresses":{"lan":"206.189.243.176","wan":"206.189.243.176"},"NodeMeta":{"consul-network-segment":"","env":"eth","stage":"beta"},"ServiceID":"statusd-whisper-rpc-2","ServiceName":"statusd-rpc","ServiceTags":["eth.beta","statusd","whisper","rpc"],"ServiceAddress":"10.1.1.99","ServiceMeta":{},"ServicePort":8547,"ServiceEnableTagOverride":true,"CreateIndex":719849,"ModifyIndex":719849},{"ID":"79eacab7-006c-1a19-0bfd-046f874ec1ec","Node":"node-02.do-ams3.eth.beta","Address":"206.189.243.178","Datacenter":"do-ams3","TaggedAddresses":{"lan":"206.189.243.178","wan":"206.189.243.178"},"NodeMeta":{"consul-network-segment":"","env":"eth","stage":"beta"},"ServiceID":"statusd-whisper-rpc-1","ServiceName":"statusd-rpc","ServiceTags":["eth.beta","statusd","whisper","rpc"],"ServiceAddress":"10.1.0.9","ServiceMeta":{},"ServicePort":8546,"ServiceEnableTagOverride":true,"CreateIndex":718374,"ModifyIndex":718374},{"ID":"79eacab7-006c-1a19-0bfd-046f874ec1ec","Node":"node-02.do-ams3.eth.beta","Address":"206.189.243.178","Datacenter":"do-ams3","TaggedAddresses":{"lan":"206.189.243.178","wan":"206.189.243.178"},"NodeMeta":{"consul-network-segment":"","env":"eth","stage":"beta"},"ServiceID":"statusd-whisper-rpc-2","ServiceName":"statusd-rpc","ServiceTags":["eth.beta","statusd","whisper","rpc"],"ServiceAddress":"10.1.0.9","ServiceMeta":{},"ServicePort":8547,"ServiceEnableTagOverride":true,"CreateIndex":718375,"ModifyIndex":718375},{"ID":"dd349135-a34a-1452-b1ea-b00987d588f2","Node":"node-03.do-ams3.eth.beta","Address":"206.189.243.179","Datacenter":"do-ams3","TaggedAddresses":{"lan":"206.189.243.179","wan":"206.189.243.179"},"NodeMeta":{"consul-network-segment":"","env":"eth","stage":"beta"},"ServiceID":"statusd-whisper-rpc-1","ServiceName":"statusd-rpc","ServiceTags":["eth.beta","statusd","whisper","rpc"],"ServiceAddress":"10.1.0.6","ServiceMeta":{},"ServicePort":8546,"ServiceEnableTagOverride":true,"CreateIndex":719868,"ModifyIndex":719868},{"ID":"dd349135-a34a-1452-b1ea-b00987d588f2","Node":"node-03.do-ams3.eth.beta","Address":"206.189.243.179","Datacenter":"do-ams3","TaggedAddresses":{"lan":"206.189.243.179","wan":"206.189.243.179"},"NodeMeta":{"consul-network-segment":"","env":"eth","stage":"beta"},"ServiceID":"statusd-whisper-rpc-2","ServiceName":"statusd-rpc","ServiceTags":["eth.beta","statusd","whisper","rpc"],"ServiceAddress":"10.1.0.6","ServiceMeta":{},"ServicePort":8547,"ServiceEnableTagOverride":true,"CreateIndex":719869,"ModifyIndex":719869},{"ID":"3e7a4660-ea9f-1d5f-f3e8-94f002f3a4cc","Node":"node-04.do-ams3.eth.beta","Address":"206.189.243.171","Datacenter":"do-ams3","TaggedAddresses":{"lan":"206.189.243.171","wan":"206.189.243.171"},"NodeMeta":{"consul-network-segment":"","env":"eth","stage":"beta"},"ServiceID":"statusd-whisper-rpc-1","ServiceName":"statusd-rpc","ServiceTags":["eth.beta","statusd","whisper","rpc"],"ServiceAddress":"10.1.0.7","ServiceMeta":{},"ServicePort":8546,"ServiceEnableTagOverride":true,"CreateIndex":719858,"ModifyIndex":719858},{"ID":"3e7a4660-ea9f-1d5f-f3e8-94f002f3a4cc","Node":"node-04.do-ams3.eth.beta","Address":"206.189.243.171","Datacenter":"do-ams3","TaggedAddresses":{"lan":"206.189.243.171","wan":"206.189.243.171"},"NodeMeta":{"consul-network-segment":"","env":"eth","stage":"beta"},"ServiceID":"statusd-whisper-rpc-2","ServiceName":"statusd-rpc","ServiceTags":["eth.beta","statusd","whisper","rpc"],"ServiceAddress":"10.1.0.7","ServiceMeta":{},"ServicePort":8547,"ServiceEnableTagOverride":true,"CreateIndex":719859,"ModifyIndex":719859},{"ID":"3d75a106-6bd6-027f-b7bd-371a01739d32","Node":"node-05.do-ams3.eth.beta","Address":"206.189.243.172","Datacenter":"do-ams3","TaggedAddresses":{"lan":"206.189.243.172","wan":"206.189.243.172"},"NodeMeta":{"consul-network-segment":"","env":"eth","stage":"beta"},"ServiceID":"statusd-whisper-rpc-1","ServiceName":"statusd-rpc","ServiceTags":["eth.beta","statusd","whisper","rpc"],"ServiceAddress":"10.1.0.4","ServiceMeta":{},"ServicePort":8546,"ServiceEnableTagOverride":true,"CreateIndex":719817,"ModifyIndex":719817},{"ID":"3d75a106-6bd6-027f-b7bd-371a01739d32","Node":"node-05.do-ams3.eth.beta","Address":"206.189.243.172","Datacenter":"do-ams3","TaggedAddresses":{"lan":"206.189.243.172","wan":"206.189.243.172"},"NodeMeta":{"consul-network-segment":"","env":"eth","stage":"beta"},"ServiceID":"statusd-whisper-rpc-2","ServiceName":"statusd-rpc","ServiceTags":["eth.beta","statusd","whisper","rpc"],"ServiceAddress":"10.1.0.4","ServiceMeta":{},"ServicePort":8547,"ServiceEnableTagOverride":true,"CreateIndex":719818,"ModifyIndex":719818},{"ID":"7c395574-9e32-41d8-05a9-8898c07f65a6","Node":"node-06.do-ams3.eth.beta","Address":"206.189.243.177","Datacenter":"do-ams3","TaggedAddresses":{"lan":"206.189.243.177","wan":"206.189.243.177"},"NodeMeta":{"consul-network-segment":"","env":"eth","stage":"beta"},"ServiceID":"statusd-whisper-rpc-1","ServiceName":"statusd-rpc","ServiceTags":["eth.beta","statusd","whisper","rpc"],"ServiceAddress":"10.1.0.5","ServiceMeta":{},"ServicePort":8546,"ServiceEnableTagOverride":true,"CreateIndex":719842,"ModifyIndex":719842},{"ID":"7c395574-9e32-41d8-05a9-8898c07f65a6","Node":"node-06.do-ams3.eth.beta","Address":"206.189.243.177","Datacenter":"do-ams3","TaggedAddresses":{"lan":"206.189.243.177","wan":"206.189.243.177"},"NodeMeta":{"consul-network-segment":"","env":"eth","stage":"beta"},"ServiceID":"statusd-whisper-rpc-2","ServiceName":"statusd-rpc","ServiceTags":["eth.beta","statusd","whisper","rpc"],"ServiceAddress":"10.1.0.5","ServiceMeta":{},"ServicePort":8547,"ServiceEnableTagOverride":true,"CreateIndex":719843,"ModifyIndex":719843}]`
