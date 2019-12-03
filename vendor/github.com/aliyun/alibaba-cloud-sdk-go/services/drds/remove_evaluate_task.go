package drds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// RemoveEvaluateTask invokes the drds.RemoveEvaluateTask API synchronously
// api document: https://help.aliyun.com/api/drds/removeevaluatetask.html
func (client *Client) RemoveEvaluateTask(request *RemoveEvaluateTaskRequest) (response *RemoveEvaluateTaskResponse, err error) {
	response = CreateRemoveEvaluateTaskResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveEvaluateTaskWithChan invokes the drds.RemoveEvaluateTask API asynchronously
// api document: https://help.aliyun.com/api/drds/removeevaluatetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveEvaluateTaskWithChan(request *RemoveEvaluateTaskRequest) (<-chan *RemoveEvaluateTaskResponse, <-chan error) {
	responseChan := make(chan *RemoveEvaluateTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveEvaluateTask(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// RemoveEvaluateTaskWithCallback invokes the drds.RemoveEvaluateTask API asynchronously
// api document: https://help.aliyun.com/api/drds/removeevaluatetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveEvaluateTaskWithCallback(request *RemoveEvaluateTaskRequest, callback func(response *RemoveEvaluateTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveEvaluateTaskResponse
		var err error
		defer close(result)
		response, err = client.RemoveEvaluateTask(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// RemoveEvaluateTaskRequest is the request struct for api RemoveEvaluateTask
type RemoveEvaluateTaskRequest struct {
	*requests.RpcRequest
	TaskId requests.Integer `position:"Query" name:"TaskId"`
}

// RemoveEvaluateTaskResponse is the response struct for api RemoveEvaluateTask
type RemoveEvaluateTaskResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	Success          bool             `json:"Success" xml:"Success"`
	TaskManageResult TaskManageResult `json:"TaskManageResult" xml:"TaskManageResult"`
}

// CreateRemoveEvaluateTaskRequest creates a request to invoke RemoveEvaluateTask API
func CreateRemoveEvaluateTaskRequest() (request *RemoveEvaluateTaskRequest) {
	request = &RemoveEvaluateTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "RemoveEvaluateTask", "Drds", "openAPI")
	return
}

// CreateRemoveEvaluateTaskResponse creates a response to parse from RemoveEvaluateTask response
func CreateRemoveEvaluateTaskResponse() (response *RemoveEvaluateTaskResponse) {
	response = &RemoveEvaluateTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
