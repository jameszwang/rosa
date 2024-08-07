/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v2alpha1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v2alpha1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalWifCredentialRequestList writes a list of values of the 'wif_credential_request' type to
// the given writer.
func MarshalWifCredentialRequestList(list []*WifCredentialRequest, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeWifCredentialRequestList(list, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// writeWifCredentialRequestList writes a list of value of the 'wif_credential_request' type to
// the given stream.
func writeWifCredentialRequestList(list []*WifCredentialRequest, stream *jsoniter.Stream) {
	stream.WriteArrayStart()
	for i, value := range list {
		if i > 0 {
			stream.WriteMore()
		}
		writeWifCredentialRequest(value, stream)
	}
	stream.WriteArrayEnd()
}

// UnmarshalWifCredentialRequestList reads a list of values of the 'wif_credential_request' type
// from the given source, which can be a slice of bytes, a string or a reader.
func UnmarshalWifCredentialRequestList(source interface{}) (items []*WifCredentialRequest, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	items = readWifCredentialRequestList(iterator)
	err = iterator.Error
	return
}

// readWifCredentialRequestList reads list of values of the ”wif_credential_request' type from
// the given iterator.
func readWifCredentialRequestList(iterator *jsoniter.Iterator) []*WifCredentialRequest {
	list := []*WifCredentialRequest{}
	for iterator.ReadArray() {
		item := readWifCredentialRequest(iterator)
		list = append(list, item)
	}
	return list
}
