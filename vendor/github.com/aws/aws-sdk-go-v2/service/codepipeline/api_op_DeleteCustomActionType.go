// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the input of a DeleteCustomActionType operation. The custom action
// will be marked as deleted.
type DeleteCustomActionTypeInput struct {
	_ struct{} `type:"structure"`

	// The category of the custom action that you want to delete, such as source
	// or deploy.
	//
	// Category is a required field
	Category ActionCategory `locationName:"category" type:"string" required:"true" enum:"true"`

	// The provider of the service used in the custom action, such as AWS CodeDeploy.
	//
	// Provider is a required field
	Provider *string `locationName:"provider" min:"1" type:"string" required:"true"`

	// The version of the custom action to delete.
	//
	// Version is a required field
	Version *string `locationName:"version" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCustomActionTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCustomActionTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCustomActionTypeInput"}
	if len(s.Category) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Category"))
	}

	if s.Provider == nil {
		invalidParams.Add(aws.NewErrParamRequired("Provider"))
	}
	if s.Provider != nil && len(*s.Provider) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Provider", 1))
	}

	if s.Version == nil {
		invalidParams.Add(aws.NewErrParamRequired("Version"))
	}
	if s.Version != nil && len(*s.Version) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Version", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteCustomActionTypeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCustomActionTypeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteCustomActionType = "DeleteCustomActionType"

// DeleteCustomActionTypeRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Marks a custom action as deleted. PollForJobs for the custom action fails
// after the action is marked for deletion. Used for custom actions only.
//
// To re-create a custom action after it has been deleted you must use a string
// in the version field that has never been used before. This string can be
// an incremented version number, for example. To restore a deleted custom action,
// use a JSON file that is identical to the deleted action, including the original
// string in the version field.
//
//    // Example sending a request using DeleteCustomActionTypeRequest.
//    req := client.DeleteCustomActionTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/DeleteCustomActionType
func (c *Client) DeleteCustomActionTypeRequest(input *DeleteCustomActionTypeInput) DeleteCustomActionTypeRequest {
	op := &aws.Operation{
		Name:       opDeleteCustomActionType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCustomActionTypeInput{}
	}

	req := c.newRequest(op, input, &DeleteCustomActionTypeOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteCustomActionTypeRequest{Request: req, Input: input, Copy: c.DeleteCustomActionTypeRequest}
}

// DeleteCustomActionTypeRequest is the request type for the
// DeleteCustomActionType API operation.
type DeleteCustomActionTypeRequest struct {
	*aws.Request
	Input *DeleteCustomActionTypeInput
	Copy  func(*DeleteCustomActionTypeInput) DeleteCustomActionTypeRequest
}

// Send marshals and sends the DeleteCustomActionType API request.
func (r DeleteCustomActionTypeRequest) Send(ctx context.Context) (*DeleteCustomActionTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCustomActionTypeResponse{
		DeleteCustomActionTypeOutput: r.Request.Data.(*DeleteCustomActionTypeOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCustomActionTypeResponse is the response type for the
// DeleteCustomActionType API operation.
type DeleteCustomActionTypeResponse struct {
	*DeleteCustomActionTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCustomActionType request.
func (r *DeleteCustomActionTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}