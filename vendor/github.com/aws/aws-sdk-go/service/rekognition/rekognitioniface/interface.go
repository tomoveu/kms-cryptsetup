// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package rekognitioniface provides an interface to enable mocking the Amazon Rekognition service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package rekognitioniface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

// RekognitionAPI provides an interface to enable mocking the
// rekognition.Rekognition service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Rekognition.
//    func myFunc(svc rekognitioniface.RekognitionAPI) bool {
//        // Make svc.CompareFaces request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := rekognition.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRekognitionClient struct {
//        rekognitioniface.RekognitionAPI
//    }
//    func (m *mockRekognitionClient) CompareFaces(input *rekognition.CompareFacesInput) (*rekognition.CompareFacesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRekognitionClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type RekognitionAPI interface {
	CompareFaces(*rekognition.CompareFacesInput) (*rekognition.CompareFacesOutput, error)
	CompareFacesWithContext(aws.Context, *rekognition.CompareFacesInput, ...request.Option) (*rekognition.CompareFacesOutput, error)
	CompareFacesRequest(*rekognition.CompareFacesInput) (*request.Request, *rekognition.CompareFacesOutput)

	CreateCollection(*rekognition.CreateCollectionInput) (*rekognition.CreateCollectionOutput, error)
	CreateCollectionWithContext(aws.Context, *rekognition.CreateCollectionInput, ...request.Option) (*rekognition.CreateCollectionOutput, error)
	CreateCollectionRequest(*rekognition.CreateCollectionInput) (*request.Request, *rekognition.CreateCollectionOutput)

	DeleteCollection(*rekognition.DeleteCollectionInput) (*rekognition.DeleteCollectionOutput, error)
	DeleteCollectionWithContext(aws.Context, *rekognition.DeleteCollectionInput, ...request.Option) (*rekognition.DeleteCollectionOutput, error)
	DeleteCollectionRequest(*rekognition.DeleteCollectionInput) (*request.Request, *rekognition.DeleteCollectionOutput)

	DeleteFaces(*rekognition.DeleteFacesInput) (*rekognition.DeleteFacesOutput, error)
	DeleteFacesWithContext(aws.Context, *rekognition.DeleteFacesInput, ...request.Option) (*rekognition.DeleteFacesOutput, error)
	DeleteFacesRequest(*rekognition.DeleteFacesInput) (*request.Request, *rekognition.DeleteFacesOutput)

	DetectFaces(*rekognition.DetectFacesInput) (*rekognition.DetectFacesOutput, error)
	DetectFacesWithContext(aws.Context, *rekognition.DetectFacesInput, ...request.Option) (*rekognition.DetectFacesOutput, error)
	DetectFacesRequest(*rekognition.DetectFacesInput) (*request.Request, *rekognition.DetectFacesOutput)

	DetectLabels(*rekognition.DetectLabelsInput) (*rekognition.DetectLabelsOutput, error)
	DetectLabelsWithContext(aws.Context, *rekognition.DetectLabelsInput, ...request.Option) (*rekognition.DetectLabelsOutput, error)
	DetectLabelsRequest(*rekognition.DetectLabelsInput) (*request.Request, *rekognition.DetectLabelsOutput)

	DetectModerationLabels(*rekognition.DetectModerationLabelsInput) (*rekognition.DetectModerationLabelsOutput, error)
	DetectModerationLabelsWithContext(aws.Context, *rekognition.DetectModerationLabelsInput, ...request.Option) (*rekognition.DetectModerationLabelsOutput, error)
	DetectModerationLabelsRequest(*rekognition.DetectModerationLabelsInput) (*request.Request, *rekognition.DetectModerationLabelsOutput)

	DetectText(*rekognition.DetectTextInput) (*rekognition.DetectTextOutput, error)
	DetectTextWithContext(aws.Context, *rekognition.DetectTextInput, ...request.Option) (*rekognition.DetectTextOutput, error)
	DetectTextRequest(*rekognition.DetectTextInput) (*request.Request, *rekognition.DetectTextOutput)

	GetCelebrityInfo(*rekognition.GetCelebrityInfoInput) (*rekognition.GetCelebrityInfoOutput, error)
	GetCelebrityInfoWithContext(aws.Context, *rekognition.GetCelebrityInfoInput, ...request.Option) (*rekognition.GetCelebrityInfoOutput, error)
	GetCelebrityInfoRequest(*rekognition.GetCelebrityInfoInput) (*request.Request, *rekognition.GetCelebrityInfoOutput)

	IndexFaces(*rekognition.IndexFacesInput) (*rekognition.IndexFacesOutput, error)
	IndexFacesWithContext(aws.Context, *rekognition.IndexFacesInput, ...request.Option) (*rekognition.IndexFacesOutput, error)
	IndexFacesRequest(*rekognition.IndexFacesInput) (*request.Request, *rekognition.IndexFacesOutput)

	ListCollections(*rekognition.ListCollectionsInput) (*rekognition.ListCollectionsOutput, error)
	ListCollectionsWithContext(aws.Context, *rekognition.ListCollectionsInput, ...request.Option) (*rekognition.ListCollectionsOutput, error)
	ListCollectionsRequest(*rekognition.ListCollectionsInput) (*request.Request, *rekognition.ListCollectionsOutput)

	ListCollectionsPages(*rekognition.ListCollectionsInput, func(*rekognition.ListCollectionsOutput, bool) bool) error
	ListCollectionsPagesWithContext(aws.Context, *rekognition.ListCollectionsInput, func(*rekognition.ListCollectionsOutput, bool) bool, ...request.Option) error

	ListFaces(*rekognition.ListFacesInput) (*rekognition.ListFacesOutput, error)
	ListFacesWithContext(aws.Context, *rekognition.ListFacesInput, ...request.Option) (*rekognition.ListFacesOutput, error)
	ListFacesRequest(*rekognition.ListFacesInput) (*request.Request, *rekognition.ListFacesOutput)

	ListFacesPages(*rekognition.ListFacesInput, func(*rekognition.ListFacesOutput, bool) bool) error
	ListFacesPagesWithContext(aws.Context, *rekognition.ListFacesInput, func(*rekognition.ListFacesOutput, bool) bool, ...request.Option) error

	RecognizeCelebrities(*rekognition.RecognizeCelebritiesInput) (*rekognition.RecognizeCelebritiesOutput, error)
	RecognizeCelebritiesWithContext(aws.Context, *rekognition.RecognizeCelebritiesInput, ...request.Option) (*rekognition.RecognizeCelebritiesOutput, error)
	RecognizeCelebritiesRequest(*rekognition.RecognizeCelebritiesInput) (*request.Request, *rekognition.RecognizeCelebritiesOutput)

	SearchFaces(*rekognition.SearchFacesInput) (*rekognition.SearchFacesOutput, error)
	SearchFacesWithContext(aws.Context, *rekognition.SearchFacesInput, ...request.Option) (*rekognition.SearchFacesOutput, error)
	SearchFacesRequest(*rekognition.SearchFacesInput) (*request.Request, *rekognition.SearchFacesOutput)

	SearchFacesByImage(*rekognition.SearchFacesByImageInput) (*rekognition.SearchFacesByImageOutput, error)
	SearchFacesByImageWithContext(aws.Context, *rekognition.SearchFacesByImageInput, ...request.Option) (*rekognition.SearchFacesByImageOutput, error)
	SearchFacesByImageRequest(*rekognition.SearchFacesByImageInput) (*request.Request, *rekognition.SearchFacesByImageOutput)
}

var _ RekognitionAPI = (*rekognition.Rekognition)(nil)
