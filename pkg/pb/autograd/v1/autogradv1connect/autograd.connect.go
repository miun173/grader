// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: autograd/v1/autograd.proto

package autogradv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/fahmifan/autograd/pkg/pb/autograd/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// AutogradServiceName is the fully-qualified name of the AutogradService service.
	AutogradServiceName = "autograd.v1.AutogradService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AutogradServicePingProcedure is the fully-qualified name of the AutogradService's Ping RPC.
	AutogradServicePingProcedure = "/autograd.v1.AutogradService/Ping"
	// AutogradServiceCreateManagedUserProcedure is the fully-qualified name of the AutogradService's
	// CreateManagedUser RPC.
	AutogradServiceCreateManagedUserProcedure = "/autograd.v1.AutogradService/CreateManagedUser"
	// AutogradServiceFindAllManagedUsersProcedure is the fully-qualified name of the AutogradService's
	// FindAllManagedUsers RPC.
	AutogradServiceFindAllManagedUsersProcedure = "/autograd.v1.AutogradService/FindAllManagedUsers"
	// AutogradServiceFindAssignmentProcedure is the fully-qualified name of the AutogradService's
	// FindAssignment RPC.
	AutogradServiceFindAssignmentProcedure = "/autograd.v1.AutogradService/FindAssignment"
	// AutogradServiceFindAllAssignmentsProcedure is the fully-qualified name of the AutogradService's
	// FindAllAssignments RPC.
	AutogradServiceFindAllAssignmentsProcedure = "/autograd.v1.AutogradService/FindAllAssignments"
	// AutogradServiceFindSubmissionProcedure is the fully-qualified name of the AutogradService's
	// FindSubmission RPC.
	AutogradServiceFindSubmissionProcedure = "/autograd.v1.AutogradService/FindSubmission"
	// AutogradServiceCreateAssignmentProcedure is the fully-qualified name of the AutogradService's
	// CreateAssignment RPC.
	AutogradServiceCreateAssignmentProcedure = "/autograd.v1.AutogradService/CreateAssignment"
	// AutogradServiceUpdateAssignmentProcedure is the fully-qualified name of the AutogradService's
	// UpdateAssignment RPC.
	AutogradServiceUpdateAssignmentProcedure = "/autograd.v1.AutogradService/UpdateAssignment"
	// AutogradServiceDeleteAssignmentProcedure is the fully-qualified name of the AutogradService's
	// DeleteAssignment RPC.
	AutogradServiceDeleteAssignmentProcedure = "/autograd.v1.AutogradService/DeleteAssignment"
	// AutogradServiceCreateSubmissionProcedure is the fully-qualified name of the AutogradService's
	// CreateSubmission RPC.
	AutogradServiceCreateSubmissionProcedure = "/autograd.v1.AutogradService/CreateSubmission"
	// AutogradServiceUpdateSubmissionProcedure is the fully-qualified name of the AutogradService's
	// UpdateSubmission RPC.
	AutogradServiceUpdateSubmissionProcedure = "/autograd.v1.AutogradService/UpdateSubmission"
	// AutogradServiceDeleteSubmissionProcedure is the fully-qualified name of the AutogradService's
	// DeleteSubmission RPC.
	AutogradServiceDeleteSubmissionProcedure = "/autograd.v1.AutogradService/DeleteSubmission"
	// AutogradServiceFindAllStudentAssignmentsProcedure is the fully-qualified name of the
	// AutogradService's FindAllStudentAssignments RPC.
	AutogradServiceFindAllStudentAssignmentsProcedure = "/autograd.v1.AutogradService/FindAllStudentAssignments"
	// AutogradServiceFindStudentAssignmentProcedure is the fully-qualified name of the
	// AutogradService's FindStudentAssignment RPC.
	AutogradServiceFindStudentAssignmentProcedure = "/autograd.v1.AutogradService/FindStudentAssignment"
	// AutogradServiceCreateStudentSubmissionProcedure is the fully-qualified name of the
	// AutogradService's CreateStudentSubmission RPC.
	AutogradServiceCreateStudentSubmissionProcedure = "/autograd.v1.AutogradService/CreateStudentSubmission"
	// AutogradServiceUpdateStudentSubmissionProcedure is the fully-qualified name of the
	// AutogradService's UpdateStudentSubmission RPC.
	AutogradServiceUpdateStudentSubmissionProcedure = "/autograd.v1.AutogradService/UpdateStudentSubmission"
	// AutogradServiceLoginProcedure is the fully-qualified name of the AutogradService's Login RPC.
	AutogradServiceLoginProcedure = "/autograd.v1.AutogradService/Login"
)

// AutogradServiceClient is a client for the autograd.v1.AutogradService service.
type AutogradServiceClient interface {
	Ping(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.PingResponse], error)
	// User Management
	CreateManagedUser(context.Context, *connect.Request[v1.CreateManagedUserRequest]) (*connect.Response[v1.CreatedResponse], error)
	FindAllManagedUsers(context.Context, *connect.Request[v1.FindAllManagedUsersRequest]) (*connect.Response[v1.FindAllManagedUsersResponse], error)
	// Assignment Submission
	// Assignment Queries
	FindAssignment(context.Context, *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.Assignment], error)
	FindAllAssignments(context.Context, *connect.Request[v1.FindAllAssignmentsRequest]) (*connect.Response[v1.FindAllAssignmentsResponse], error)
	FindSubmission(context.Context, *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.Submission], error)
	// Assignment Command
	CreateAssignment(context.Context, *connect.Request[v1.CreateAssignmentRequest]) (*connect.Response[v1.CreatedResponse], error)
	UpdateAssignment(context.Context, *connect.Request[v1.UpdateAssignmentRequest]) (*connect.Response[v1.Empty], error)
	DeleteAssignment(context.Context, *connect.Request[v1.DeleteByIDRequest]) (*connect.Response[v1.Empty], error)
	CreateSubmission(context.Context, *connect.Request[v1.CreateSubmissionRequest]) (*connect.Response[v1.CreatedResponse], error)
	UpdateSubmission(context.Context, *connect.Request[v1.UpdateSubmissionRequest]) (*connect.Response[v1.Empty], error)
	DeleteSubmission(context.Context, *connect.Request[v1.DeleteByIDRequest]) (*connect.Response[v1.Empty], error)
	// Student Assignment
	// Student Assignment Queries
	FindAllStudentAssignments(context.Context, *connect.Request[v1.FindAllStudentAssignmentsRequest]) (*connect.Response[v1.FindAllStudentAssignmentsResponse], error)
	FindStudentAssignment(context.Context, *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.StudentAssignment], error)
	// Student Assignment Command
	CreateStudentSubmission(context.Context, *connect.Request[v1.CreateStudentSubmissionRequest]) (*connect.Response[v1.CreatedResponse], error)
	UpdateStudentSubmission(context.Context, *connect.Request[v1.UpdateStudentSubmissionRequest]) (*connect.Response[v1.Empty], error)
	// Auth
	// Auth Queries
	Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)
}

// NewAutogradServiceClient constructs a client for the autograd.v1.AutogradService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAutogradServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AutogradServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &autogradServiceClient{
		ping: connect.NewClient[v1.Empty, v1.PingResponse](
			httpClient,
			baseURL+AutogradServicePingProcedure,
			opts...,
		),
		createManagedUser: connect.NewClient[v1.CreateManagedUserRequest, v1.CreatedResponse](
			httpClient,
			baseURL+AutogradServiceCreateManagedUserProcedure,
			opts...,
		),
		findAllManagedUsers: connect.NewClient[v1.FindAllManagedUsersRequest, v1.FindAllManagedUsersResponse](
			httpClient,
			baseURL+AutogradServiceFindAllManagedUsersProcedure,
			opts...,
		),
		findAssignment: connect.NewClient[v1.FindByIDRequest, v1.Assignment](
			httpClient,
			baseURL+AutogradServiceFindAssignmentProcedure,
			opts...,
		),
		findAllAssignments: connect.NewClient[v1.FindAllAssignmentsRequest, v1.FindAllAssignmentsResponse](
			httpClient,
			baseURL+AutogradServiceFindAllAssignmentsProcedure,
			opts...,
		),
		findSubmission: connect.NewClient[v1.FindByIDRequest, v1.Submission](
			httpClient,
			baseURL+AutogradServiceFindSubmissionProcedure,
			opts...,
		),
		createAssignment: connect.NewClient[v1.CreateAssignmentRequest, v1.CreatedResponse](
			httpClient,
			baseURL+AutogradServiceCreateAssignmentProcedure,
			opts...,
		),
		updateAssignment: connect.NewClient[v1.UpdateAssignmentRequest, v1.Empty](
			httpClient,
			baseURL+AutogradServiceUpdateAssignmentProcedure,
			opts...,
		),
		deleteAssignment: connect.NewClient[v1.DeleteByIDRequest, v1.Empty](
			httpClient,
			baseURL+AutogradServiceDeleteAssignmentProcedure,
			opts...,
		),
		createSubmission: connect.NewClient[v1.CreateSubmissionRequest, v1.CreatedResponse](
			httpClient,
			baseURL+AutogradServiceCreateSubmissionProcedure,
			opts...,
		),
		updateSubmission: connect.NewClient[v1.UpdateSubmissionRequest, v1.Empty](
			httpClient,
			baseURL+AutogradServiceUpdateSubmissionProcedure,
			opts...,
		),
		deleteSubmission: connect.NewClient[v1.DeleteByIDRequest, v1.Empty](
			httpClient,
			baseURL+AutogradServiceDeleteSubmissionProcedure,
			opts...,
		),
		findAllStudentAssignments: connect.NewClient[v1.FindAllStudentAssignmentsRequest, v1.FindAllStudentAssignmentsResponse](
			httpClient,
			baseURL+AutogradServiceFindAllStudentAssignmentsProcedure,
			opts...,
		),
		findStudentAssignment: connect.NewClient[v1.FindByIDRequest, v1.StudentAssignment](
			httpClient,
			baseURL+AutogradServiceFindStudentAssignmentProcedure,
			opts...,
		),
		createStudentSubmission: connect.NewClient[v1.CreateStudentSubmissionRequest, v1.CreatedResponse](
			httpClient,
			baseURL+AutogradServiceCreateStudentSubmissionProcedure,
			opts...,
		),
		updateStudentSubmission: connect.NewClient[v1.UpdateStudentSubmissionRequest, v1.Empty](
			httpClient,
			baseURL+AutogradServiceUpdateStudentSubmissionProcedure,
			opts...,
		),
		login: connect.NewClient[v1.LoginRequest, v1.LoginResponse](
			httpClient,
			baseURL+AutogradServiceLoginProcedure,
			opts...,
		),
	}
}

// autogradServiceClient implements AutogradServiceClient.
type autogradServiceClient struct {
	ping                      *connect.Client[v1.Empty, v1.PingResponse]
	createManagedUser         *connect.Client[v1.CreateManagedUserRequest, v1.CreatedResponse]
	findAllManagedUsers       *connect.Client[v1.FindAllManagedUsersRequest, v1.FindAllManagedUsersResponse]
	findAssignment            *connect.Client[v1.FindByIDRequest, v1.Assignment]
	findAllAssignments        *connect.Client[v1.FindAllAssignmentsRequest, v1.FindAllAssignmentsResponse]
	findSubmission            *connect.Client[v1.FindByIDRequest, v1.Submission]
	createAssignment          *connect.Client[v1.CreateAssignmentRequest, v1.CreatedResponse]
	updateAssignment          *connect.Client[v1.UpdateAssignmentRequest, v1.Empty]
	deleteAssignment          *connect.Client[v1.DeleteByIDRequest, v1.Empty]
	createSubmission          *connect.Client[v1.CreateSubmissionRequest, v1.CreatedResponse]
	updateSubmission          *connect.Client[v1.UpdateSubmissionRequest, v1.Empty]
	deleteSubmission          *connect.Client[v1.DeleteByIDRequest, v1.Empty]
	findAllStudentAssignments *connect.Client[v1.FindAllStudentAssignmentsRequest, v1.FindAllStudentAssignmentsResponse]
	findStudentAssignment     *connect.Client[v1.FindByIDRequest, v1.StudentAssignment]
	createStudentSubmission   *connect.Client[v1.CreateStudentSubmissionRequest, v1.CreatedResponse]
	updateStudentSubmission   *connect.Client[v1.UpdateStudentSubmissionRequest, v1.Empty]
	login                     *connect.Client[v1.LoginRequest, v1.LoginResponse]
}

// Ping calls autograd.v1.AutogradService.Ping.
func (c *autogradServiceClient) Ping(ctx context.Context, req *connect.Request[v1.Empty]) (*connect.Response[v1.PingResponse], error) {
	return c.ping.CallUnary(ctx, req)
}

// CreateManagedUser calls autograd.v1.AutogradService.CreateManagedUser.
func (c *autogradServiceClient) CreateManagedUser(ctx context.Context, req *connect.Request[v1.CreateManagedUserRequest]) (*connect.Response[v1.CreatedResponse], error) {
	return c.createManagedUser.CallUnary(ctx, req)
}

// FindAllManagedUsers calls autograd.v1.AutogradService.FindAllManagedUsers.
func (c *autogradServiceClient) FindAllManagedUsers(ctx context.Context, req *connect.Request[v1.FindAllManagedUsersRequest]) (*connect.Response[v1.FindAllManagedUsersResponse], error) {
	return c.findAllManagedUsers.CallUnary(ctx, req)
}

// FindAssignment calls autograd.v1.AutogradService.FindAssignment.
func (c *autogradServiceClient) FindAssignment(ctx context.Context, req *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.Assignment], error) {
	return c.findAssignment.CallUnary(ctx, req)
}

// FindAllAssignments calls autograd.v1.AutogradService.FindAllAssignments.
func (c *autogradServiceClient) FindAllAssignments(ctx context.Context, req *connect.Request[v1.FindAllAssignmentsRequest]) (*connect.Response[v1.FindAllAssignmentsResponse], error) {
	return c.findAllAssignments.CallUnary(ctx, req)
}

// FindSubmission calls autograd.v1.AutogradService.FindSubmission.
func (c *autogradServiceClient) FindSubmission(ctx context.Context, req *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.Submission], error) {
	return c.findSubmission.CallUnary(ctx, req)
}

// CreateAssignment calls autograd.v1.AutogradService.CreateAssignment.
func (c *autogradServiceClient) CreateAssignment(ctx context.Context, req *connect.Request[v1.CreateAssignmentRequest]) (*connect.Response[v1.CreatedResponse], error) {
	return c.createAssignment.CallUnary(ctx, req)
}

// UpdateAssignment calls autograd.v1.AutogradService.UpdateAssignment.
func (c *autogradServiceClient) UpdateAssignment(ctx context.Context, req *connect.Request[v1.UpdateAssignmentRequest]) (*connect.Response[v1.Empty], error) {
	return c.updateAssignment.CallUnary(ctx, req)
}

// DeleteAssignment calls autograd.v1.AutogradService.DeleteAssignment.
func (c *autogradServiceClient) DeleteAssignment(ctx context.Context, req *connect.Request[v1.DeleteByIDRequest]) (*connect.Response[v1.Empty], error) {
	return c.deleteAssignment.CallUnary(ctx, req)
}

// CreateSubmission calls autograd.v1.AutogradService.CreateSubmission.
func (c *autogradServiceClient) CreateSubmission(ctx context.Context, req *connect.Request[v1.CreateSubmissionRequest]) (*connect.Response[v1.CreatedResponse], error) {
	return c.createSubmission.CallUnary(ctx, req)
}

// UpdateSubmission calls autograd.v1.AutogradService.UpdateSubmission.
func (c *autogradServiceClient) UpdateSubmission(ctx context.Context, req *connect.Request[v1.UpdateSubmissionRequest]) (*connect.Response[v1.Empty], error) {
	return c.updateSubmission.CallUnary(ctx, req)
}

// DeleteSubmission calls autograd.v1.AutogradService.DeleteSubmission.
func (c *autogradServiceClient) DeleteSubmission(ctx context.Context, req *connect.Request[v1.DeleteByIDRequest]) (*connect.Response[v1.Empty], error) {
	return c.deleteSubmission.CallUnary(ctx, req)
}

// FindAllStudentAssignments calls autograd.v1.AutogradService.FindAllStudentAssignments.
func (c *autogradServiceClient) FindAllStudentAssignments(ctx context.Context, req *connect.Request[v1.FindAllStudentAssignmentsRequest]) (*connect.Response[v1.FindAllStudentAssignmentsResponse], error) {
	return c.findAllStudentAssignments.CallUnary(ctx, req)
}

// FindStudentAssignment calls autograd.v1.AutogradService.FindStudentAssignment.
func (c *autogradServiceClient) FindStudentAssignment(ctx context.Context, req *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.StudentAssignment], error) {
	return c.findStudentAssignment.CallUnary(ctx, req)
}

// CreateStudentSubmission calls autograd.v1.AutogradService.CreateStudentSubmission.
func (c *autogradServiceClient) CreateStudentSubmission(ctx context.Context, req *connect.Request[v1.CreateStudentSubmissionRequest]) (*connect.Response[v1.CreatedResponse], error) {
	return c.createStudentSubmission.CallUnary(ctx, req)
}

// UpdateStudentSubmission calls autograd.v1.AutogradService.UpdateStudentSubmission.
func (c *autogradServiceClient) UpdateStudentSubmission(ctx context.Context, req *connect.Request[v1.UpdateStudentSubmissionRequest]) (*connect.Response[v1.Empty], error) {
	return c.updateStudentSubmission.CallUnary(ctx, req)
}

// Login calls autograd.v1.AutogradService.Login.
func (c *autogradServiceClient) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	return c.login.CallUnary(ctx, req)
}

// AutogradServiceHandler is an implementation of the autograd.v1.AutogradService service.
type AutogradServiceHandler interface {
	Ping(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.PingResponse], error)
	// User Management
	CreateManagedUser(context.Context, *connect.Request[v1.CreateManagedUserRequest]) (*connect.Response[v1.CreatedResponse], error)
	FindAllManagedUsers(context.Context, *connect.Request[v1.FindAllManagedUsersRequest]) (*connect.Response[v1.FindAllManagedUsersResponse], error)
	// Assignment Submission
	// Assignment Queries
	FindAssignment(context.Context, *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.Assignment], error)
	FindAllAssignments(context.Context, *connect.Request[v1.FindAllAssignmentsRequest]) (*connect.Response[v1.FindAllAssignmentsResponse], error)
	FindSubmission(context.Context, *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.Submission], error)
	// Assignment Command
	CreateAssignment(context.Context, *connect.Request[v1.CreateAssignmentRequest]) (*connect.Response[v1.CreatedResponse], error)
	UpdateAssignment(context.Context, *connect.Request[v1.UpdateAssignmentRequest]) (*connect.Response[v1.Empty], error)
	DeleteAssignment(context.Context, *connect.Request[v1.DeleteByIDRequest]) (*connect.Response[v1.Empty], error)
	CreateSubmission(context.Context, *connect.Request[v1.CreateSubmissionRequest]) (*connect.Response[v1.CreatedResponse], error)
	UpdateSubmission(context.Context, *connect.Request[v1.UpdateSubmissionRequest]) (*connect.Response[v1.Empty], error)
	DeleteSubmission(context.Context, *connect.Request[v1.DeleteByIDRequest]) (*connect.Response[v1.Empty], error)
	// Student Assignment
	// Student Assignment Queries
	FindAllStudentAssignments(context.Context, *connect.Request[v1.FindAllStudentAssignmentsRequest]) (*connect.Response[v1.FindAllStudentAssignmentsResponse], error)
	FindStudentAssignment(context.Context, *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.StudentAssignment], error)
	// Student Assignment Command
	CreateStudentSubmission(context.Context, *connect.Request[v1.CreateStudentSubmissionRequest]) (*connect.Response[v1.CreatedResponse], error)
	UpdateStudentSubmission(context.Context, *connect.Request[v1.UpdateStudentSubmissionRequest]) (*connect.Response[v1.Empty], error)
	// Auth
	// Auth Queries
	Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)
}

// NewAutogradServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAutogradServiceHandler(svc AutogradServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	autogradServicePingHandler := connect.NewUnaryHandler(
		AutogradServicePingProcedure,
		svc.Ping,
		opts...,
	)
	autogradServiceCreateManagedUserHandler := connect.NewUnaryHandler(
		AutogradServiceCreateManagedUserProcedure,
		svc.CreateManagedUser,
		opts...,
	)
	autogradServiceFindAllManagedUsersHandler := connect.NewUnaryHandler(
		AutogradServiceFindAllManagedUsersProcedure,
		svc.FindAllManagedUsers,
		opts...,
	)
	autogradServiceFindAssignmentHandler := connect.NewUnaryHandler(
		AutogradServiceFindAssignmentProcedure,
		svc.FindAssignment,
		opts...,
	)
	autogradServiceFindAllAssignmentsHandler := connect.NewUnaryHandler(
		AutogradServiceFindAllAssignmentsProcedure,
		svc.FindAllAssignments,
		opts...,
	)
	autogradServiceFindSubmissionHandler := connect.NewUnaryHandler(
		AutogradServiceFindSubmissionProcedure,
		svc.FindSubmission,
		opts...,
	)
	autogradServiceCreateAssignmentHandler := connect.NewUnaryHandler(
		AutogradServiceCreateAssignmentProcedure,
		svc.CreateAssignment,
		opts...,
	)
	autogradServiceUpdateAssignmentHandler := connect.NewUnaryHandler(
		AutogradServiceUpdateAssignmentProcedure,
		svc.UpdateAssignment,
		opts...,
	)
	autogradServiceDeleteAssignmentHandler := connect.NewUnaryHandler(
		AutogradServiceDeleteAssignmentProcedure,
		svc.DeleteAssignment,
		opts...,
	)
	autogradServiceCreateSubmissionHandler := connect.NewUnaryHandler(
		AutogradServiceCreateSubmissionProcedure,
		svc.CreateSubmission,
		opts...,
	)
	autogradServiceUpdateSubmissionHandler := connect.NewUnaryHandler(
		AutogradServiceUpdateSubmissionProcedure,
		svc.UpdateSubmission,
		opts...,
	)
	autogradServiceDeleteSubmissionHandler := connect.NewUnaryHandler(
		AutogradServiceDeleteSubmissionProcedure,
		svc.DeleteSubmission,
		opts...,
	)
	autogradServiceFindAllStudentAssignmentsHandler := connect.NewUnaryHandler(
		AutogradServiceFindAllStudentAssignmentsProcedure,
		svc.FindAllStudentAssignments,
		opts...,
	)
	autogradServiceFindStudentAssignmentHandler := connect.NewUnaryHandler(
		AutogradServiceFindStudentAssignmentProcedure,
		svc.FindStudentAssignment,
		opts...,
	)
	autogradServiceCreateStudentSubmissionHandler := connect.NewUnaryHandler(
		AutogradServiceCreateStudentSubmissionProcedure,
		svc.CreateStudentSubmission,
		opts...,
	)
	autogradServiceUpdateStudentSubmissionHandler := connect.NewUnaryHandler(
		AutogradServiceUpdateStudentSubmissionProcedure,
		svc.UpdateStudentSubmission,
		opts...,
	)
	autogradServiceLoginHandler := connect.NewUnaryHandler(
		AutogradServiceLoginProcedure,
		svc.Login,
		opts...,
	)
	return "/autograd.v1.AutogradService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AutogradServicePingProcedure:
			autogradServicePingHandler.ServeHTTP(w, r)
		case AutogradServiceCreateManagedUserProcedure:
			autogradServiceCreateManagedUserHandler.ServeHTTP(w, r)
		case AutogradServiceFindAllManagedUsersProcedure:
			autogradServiceFindAllManagedUsersHandler.ServeHTTP(w, r)
		case AutogradServiceFindAssignmentProcedure:
			autogradServiceFindAssignmentHandler.ServeHTTP(w, r)
		case AutogradServiceFindAllAssignmentsProcedure:
			autogradServiceFindAllAssignmentsHandler.ServeHTTP(w, r)
		case AutogradServiceFindSubmissionProcedure:
			autogradServiceFindSubmissionHandler.ServeHTTP(w, r)
		case AutogradServiceCreateAssignmentProcedure:
			autogradServiceCreateAssignmentHandler.ServeHTTP(w, r)
		case AutogradServiceUpdateAssignmentProcedure:
			autogradServiceUpdateAssignmentHandler.ServeHTTP(w, r)
		case AutogradServiceDeleteAssignmentProcedure:
			autogradServiceDeleteAssignmentHandler.ServeHTTP(w, r)
		case AutogradServiceCreateSubmissionProcedure:
			autogradServiceCreateSubmissionHandler.ServeHTTP(w, r)
		case AutogradServiceUpdateSubmissionProcedure:
			autogradServiceUpdateSubmissionHandler.ServeHTTP(w, r)
		case AutogradServiceDeleteSubmissionProcedure:
			autogradServiceDeleteSubmissionHandler.ServeHTTP(w, r)
		case AutogradServiceFindAllStudentAssignmentsProcedure:
			autogradServiceFindAllStudentAssignmentsHandler.ServeHTTP(w, r)
		case AutogradServiceFindStudentAssignmentProcedure:
			autogradServiceFindStudentAssignmentHandler.ServeHTTP(w, r)
		case AutogradServiceCreateStudentSubmissionProcedure:
			autogradServiceCreateStudentSubmissionHandler.ServeHTTP(w, r)
		case AutogradServiceUpdateStudentSubmissionProcedure:
			autogradServiceUpdateStudentSubmissionHandler.ServeHTTP(w, r)
		case AutogradServiceLoginProcedure:
			autogradServiceLoginHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAutogradServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAutogradServiceHandler struct{}

func (UnimplementedAutogradServiceHandler) Ping(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.PingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autograd.v1.AutogradService.Ping is not implemented"))
}

func (UnimplementedAutogradServiceHandler) CreateManagedUser(context.Context, *connect.Request[v1.CreateManagedUserRequest]) (*connect.Response[v1.CreatedResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autograd.v1.AutogradService.CreateManagedUser is not implemented"))
}

func (UnimplementedAutogradServiceHandler) FindAllManagedUsers(context.Context, *connect.Request[v1.FindAllManagedUsersRequest]) (*connect.Response[v1.FindAllManagedUsersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autograd.v1.AutogradService.FindAllManagedUsers is not implemented"))
}

func (UnimplementedAutogradServiceHandler) FindAssignment(context.Context, *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.Assignment], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autograd.v1.AutogradService.FindAssignment is not implemented"))
}

func (UnimplementedAutogradServiceHandler) FindAllAssignments(context.Context, *connect.Request[v1.FindAllAssignmentsRequest]) (*connect.Response[v1.FindAllAssignmentsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autograd.v1.AutogradService.FindAllAssignments is not implemented"))
}

func (UnimplementedAutogradServiceHandler) FindSubmission(context.Context, *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.Submission], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autograd.v1.AutogradService.FindSubmission is not implemented"))
}

func (UnimplementedAutogradServiceHandler) CreateAssignment(context.Context, *connect.Request[v1.CreateAssignmentRequest]) (*connect.Response[v1.CreatedResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autograd.v1.AutogradService.CreateAssignment is not implemented"))
}

func (UnimplementedAutogradServiceHandler) UpdateAssignment(context.Context, *connect.Request[v1.UpdateAssignmentRequest]) (*connect.Response[v1.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autograd.v1.AutogradService.UpdateAssignment is not implemented"))
}

func (UnimplementedAutogradServiceHandler) DeleteAssignment(context.Context, *connect.Request[v1.DeleteByIDRequest]) (*connect.Response[v1.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autograd.v1.AutogradService.DeleteAssignment is not implemented"))
}

func (UnimplementedAutogradServiceHandler) CreateSubmission(context.Context, *connect.Request[v1.CreateSubmissionRequest]) (*connect.Response[v1.CreatedResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autograd.v1.AutogradService.CreateSubmission is not implemented"))
}

func (UnimplementedAutogradServiceHandler) UpdateSubmission(context.Context, *connect.Request[v1.UpdateSubmissionRequest]) (*connect.Response[v1.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autograd.v1.AutogradService.UpdateSubmission is not implemented"))
}

func (UnimplementedAutogradServiceHandler) DeleteSubmission(context.Context, *connect.Request[v1.DeleteByIDRequest]) (*connect.Response[v1.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autograd.v1.AutogradService.DeleteSubmission is not implemented"))
}

func (UnimplementedAutogradServiceHandler) FindAllStudentAssignments(context.Context, *connect.Request[v1.FindAllStudentAssignmentsRequest]) (*connect.Response[v1.FindAllStudentAssignmentsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autograd.v1.AutogradService.FindAllStudentAssignments is not implemented"))
}

func (UnimplementedAutogradServiceHandler) FindStudentAssignment(context.Context, *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.StudentAssignment], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autograd.v1.AutogradService.FindStudentAssignment is not implemented"))
}

func (UnimplementedAutogradServiceHandler) CreateStudentSubmission(context.Context, *connect.Request[v1.CreateStudentSubmissionRequest]) (*connect.Response[v1.CreatedResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autograd.v1.AutogradService.CreateStudentSubmission is not implemented"))
}

func (UnimplementedAutogradServiceHandler) UpdateStudentSubmission(context.Context, *connect.Request[v1.UpdateStudentSubmissionRequest]) (*connect.Response[v1.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autograd.v1.AutogradService.UpdateStudentSubmission is not implemented"))
}

func (UnimplementedAutogradServiceHandler) Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autograd.v1.AutogradService.Login is not implemented"))
}