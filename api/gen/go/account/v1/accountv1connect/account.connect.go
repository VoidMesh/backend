// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: account/v1/account.proto

package accountv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/VoidMesh/backend/api/gen/go/account/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// AccountServiceName is the fully-qualified name of the AccountService service.
	AccountServiceName = "account.v1.AccountService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AccountServiceCreateProcedure is the fully-qualified name of the AccountService's Create RPC.
	AccountServiceCreateProcedure = "/account.v1.AccountService/Create"
	// AccountServiceAuthenticateProcedure is the fully-qualified name of the AccountService's
	// Authenticate RPC.
	AccountServiceAuthenticateProcedure = "/account.v1.AccountService/Authenticate"
	// AccountServiceRefreshAccessTokenProcedure is the fully-qualified name of the AccountService's
	// RefreshAccessToken RPC.
	AccountServiceRefreshAccessTokenProcedure = "/account.v1.AccountService/RefreshAccessToken"
	// AccountServiceDeleteRefreshTokenProcedure is the fully-qualified name of the AccountService's
	// DeleteRefreshToken RPC.
	AccountServiceDeleteRefreshTokenProcedure = "/account.v1.AccountService/DeleteRefreshToken"
	// AccountServiceRequestNewPasswordProcedure is the fully-qualified name of the AccountService's
	// RequestNewPassword RPC.
	AccountServiceRequestNewPasswordProcedure = "/account.v1.AccountService/RequestNewPassword"
	// AccountServiceVerifyEmailProcedure is the fully-qualified name of the AccountService's
	// VerifyEmail RPC.
	AccountServiceVerifyEmailProcedure = "/account.v1.AccountService/VerifyEmail"
	// AccountServiceUpdateProcedure is the fully-qualified name of the AccountService's Update RPC.
	AccountServiceUpdateProcedure = "/account.v1.AccountService/Update"
	// AccountServiceDeleteProcedure is the fully-qualified name of the AccountService's Delete RPC.
	AccountServiceDeleteProcedure = "/account.v1.AccountService/Delete"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	accountServiceServiceDescriptor                  = v1.File_account_v1_account_proto.Services().ByName("AccountService")
	accountServiceCreateMethodDescriptor             = accountServiceServiceDescriptor.Methods().ByName("Create")
	accountServiceAuthenticateMethodDescriptor       = accountServiceServiceDescriptor.Methods().ByName("Authenticate")
	accountServiceRefreshAccessTokenMethodDescriptor = accountServiceServiceDescriptor.Methods().ByName("RefreshAccessToken")
	accountServiceDeleteRefreshTokenMethodDescriptor = accountServiceServiceDescriptor.Methods().ByName("DeleteRefreshToken")
	accountServiceRequestNewPasswordMethodDescriptor = accountServiceServiceDescriptor.Methods().ByName("RequestNewPassword")
	accountServiceVerifyEmailMethodDescriptor        = accountServiceServiceDescriptor.Methods().ByName("VerifyEmail")
	accountServiceUpdateMethodDescriptor             = accountServiceServiceDescriptor.Methods().ByName("Update")
	accountServiceDeleteMethodDescriptor             = accountServiceServiceDescriptor.Methods().ByName("Delete")
)

// AccountServiceClient is a client for the account.v1.AccountService service.
type AccountServiceClient interface {
	// Registers a new user account.
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	// Authenticates a user's login request.
	Authenticate(context.Context, *connect.Request[v1.AuthenticateRequest]) (*connect.Response[v1.AuthenticateResponse], error)
	// Refresh access token for a user account.
	RefreshAccessToken(context.Context, *connect.Request[v1.RefreshAccessTokenRequest]) (*connect.Response[v1.RefreshAccessTokenResponse], error)
	// Delete refresh token.
	DeleteRefreshToken(context.Context, *connect.Request[v1.DeleteRefreshTokenRequest]) (*connect.Response[v1.DeleteRefreshTokenResponse], error)
	// Requests a new password for a user account.
	RequestNewPassword(context.Context, *connect.Request[v1.RequestNewPasswordRequest]) (*connect.Response[v1.RequestNewPasswordResponse], error)
	// Verify account email address.
	VerifyEmail(context.Context, *connect.Request[v1.VerifyEmailRequest]) (*connect.Response[v1.VerifyEmailResponse], error)
	// Updates user account settings (e.g., privacy settings, notification
	// preferences).
	Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error)
	// Removes an account permanently.
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
}

// NewAccountServiceClient constructs a client for the account.v1.AccountService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAccountServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AccountServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &accountServiceClient{
		create: connect.NewClient[v1.CreateRequest, v1.CreateResponse](
			httpClient,
			baseURL+AccountServiceCreateProcedure,
			connect.WithSchema(accountServiceCreateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		authenticate: connect.NewClient[v1.AuthenticateRequest, v1.AuthenticateResponse](
			httpClient,
			baseURL+AccountServiceAuthenticateProcedure,
			connect.WithSchema(accountServiceAuthenticateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		refreshAccessToken: connect.NewClient[v1.RefreshAccessTokenRequest, v1.RefreshAccessTokenResponse](
			httpClient,
			baseURL+AccountServiceRefreshAccessTokenProcedure,
			connect.WithSchema(accountServiceRefreshAccessTokenMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteRefreshToken: connect.NewClient[v1.DeleteRefreshTokenRequest, v1.DeleteRefreshTokenResponse](
			httpClient,
			baseURL+AccountServiceDeleteRefreshTokenProcedure,
			connect.WithSchema(accountServiceDeleteRefreshTokenMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		requestNewPassword: connect.NewClient[v1.RequestNewPasswordRequest, v1.RequestNewPasswordResponse](
			httpClient,
			baseURL+AccountServiceRequestNewPasswordProcedure,
			connect.WithSchema(accountServiceRequestNewPasswordMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		verifyEmail: connect.NewClient[v1.VerifyEmailRequest, v1.VerifyEmailResponse](
			httpClient,
			baseURL+AccountServiceVerifyEmailProcedure,
			connect.WithSchema(accountServiceVerifyEmailMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		update: connect.NewClient[v1.UpdateRequest, v1.UpdateResponse](
			httpClient,
			baseURL+AccountServiceUpdateProcedure,
			connect.WithSchema(accountServiceUpdateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		delete: connect.NewClient[v1.DeleteRequest, v1.DeleteResponse](
			httpClient,
			baseURL+AccountServiceDeleteProcedure,
			connect.WithSchema(accountServiceDeleteMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// accountServiceClient implements AccountServiceClient.
type accountServiceClient struct {
	create             *connect.Client[v1.CreateRequest, v1.CreateResponse]
	authenticate       *connect.Client[v1.AuthenticateRequest, v1.AuthenticateResponse]
	refreshAccessToken *connect.Client[v1.RefreshAccessTokenRequest, v1.RefreshAccessTokenResponse]
	deleteRefreshToken *connect.Client[v1.DeleteRefreshTokenRequest, v1.DeleteRefreshTokenResponse]
	requestNewPassword *connect.Client[v1.RequestNewPasswordRequest, v1.RequestNewPasswordResponse]
	verifyEmail        *connect.Client[v1.VerifyEmailRequest, v1.VerifyEmailResponse]
	update             *connect.Client[v1.UpdateRequest, v1.UpdateResponse]
	delete             *connect.Client[v1.DeleteRequest, v1.DeleteResponse]
}

// Create calls account.v1.AccountService.Create.
func (c *accountServiceClient) Create(ctx context.Context, req *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Authenticate calls account.v1.AccountService.Authenticate.
func (c *accountServiceClient) Authenticate(ctx context.Context, req *connect.Request[v1.AuthenticateRequest]) (*connect.Response[v1.AuthenticateResponse], error) {
	return c.authenticate.CallUnary(ctx, req)
}

// RefreshAccessToken calls account.v1.AccountService.RefreshAccessToken.
func (c *accountServiceClient) RefreshAccessToken(ctx context.Context, req *connect.Request[v1.RefreshAccessTokenRequest]) (*connect.Response[v1.RefreshAccessTokenResponse], error) {
	return c.refreshAccessToken.CallUnary(ctx, req)
}

// DeleteRefreshToken calls account.v1.AccountService.DeleteRefreshToken.
func (c *accountServiceClient) DeleteRefreshToken(ctx context.Context, req *connect.Request[v1.DeleteRefreshTokenRequest]) (*connect.Response[v1.DeleteRefreshTokenResponse], error) {
	return c.deleteRefreshToken.CallUnary(ctx, req)
}

// RequestNewPassword calls account.v1.AccountService.RequestNewPassword.
func (c *accountServiceClient) RequestNewPassword(ctx context.Context, req *connect.Request[v1.RequestNewPasswordRequest]) (*connect.Response[v1.RequestNewPasswordResponse], error) {
	return c.requestNewPassword.CallUnary(ctx, req)
}

// VerifyEmail calls account.v1.AccountService.VerifyEmail.
func (c *accountServiceClient) VerifyEmail(ctx context.Context, req *connect.Request[v1.VerifyEmailRequest]) (*connect.Response[v1.VerifyEmailResponse], error) {
	return c.verifyEmail.CallUnary(ctx, req)
}

// Update calls account.v1.AccountService.Update.
func (c *accountServiceClient) Update(ctx context.Context, req *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// Delete calls account.v1.AccountService.Delete.
func (c *accountServiceClient) Delete(ctx context.Context, req *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// AccountServiceHandler is an implementation of the account.v1.AccountService service.
type AccountServiceHandler interface {
	// Registers a new user account.
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	// Authenticates a user's login request.
	Authenticate(context.Context, *connect.Request[v1.AuthenticateRequest]) (*connect.Response[v1.AuthenticateResponse], error)
	// Refresh access token for a user account.
	RefreshAccessToken(context.Context, *connect.Request[v1.RefreshAccessTokenRequest]) (*connect.Response[v1.RefreshAccessTokenResponse], error)
	// Delete refresh token.
	DeleteRefreshToken(context.Context, *connect.Request[v1.DeleteRefreshTokenRequest]) (*connect.Response[v1.DeleteRefreshTokenResponse], error)
	// Requests a new password for a user account.
	RequestNewPassword(context.Context, *connect.Request[v1.RequestNewPasswordRequest]) (*connect.Response[v1.RequestNewPasswordResponse], error)
	// Verify account email address.
	VerifyEmail(context.Context, *connect.Request[v1.VerifyEmailRequest]) (*connect.Response[v1.VerifyEmailResponse], error)
	// Updates user account settings (e.g., privacy settings, notification
	// preferences).
	Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error)
	// Removes an account permanently.
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
}

// NewAccountServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAccountServiceHandler(svc AccountServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	accountServiceCreateHandler := connect.NewUnaryHandler(
		AccountServiceCreateProcedure,
		svc.Create,
		connect.WithSchema(accountServiceCreateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServiceAuthenticateHandler := connect.NewUnaryHandler(
		AccountServiceAuthenticateProcedure,
		svc.Authenticate,
		connect.WithSchema(accountServiceAuthenticateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServiceRefreshAccessTokenHandler := connect.NewUnaryHandler(
		AccountServiceRefreshAccessTokenProcedure,
		svc.RefreshAccessToken,
		connect.WithSchema(accountServiceRefreshAccessTokenMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServiceDeleteRefreshTokenHandler := connect.NewUnaryHandler(
		AccountServiceDeleteRefreshTokenProcedure,
		svc.DeleteRefreshToken,
		connect.WithSchema(accountServiceDeleteRefreshTokenMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServiceRequestNewPasswordHandler := connect.NewUnaryHandler(
		AccountServiceRequestNewPasswordProcedure,
		svc.RequestNewPassword,
		connect.WithSchema(accountServiceRequestNewPasswordMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServiceVerifyEmailHandler := connect.NewUnaryHandler(
		AccountServiceVerifyEmailProcedure,
		svc.VerifyEmail,
		connect.WithSchema(accountServiceVerifyEmailMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServiceUpdateHandler := connect.NewUnaryHandler(
		AccountServiceUpdateProcedure,
		svc.Update,
		connect.WithSchema(accountServiceUpdateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServiceDeleteHandler := connect.NewUnaryHandler(
		AccountServiceDeleteProcedure,
		svc.Delete,
		connect.WithSchema(accountServiceDeleteMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/account.v1.AccountService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AccountServiceCreateProcedure:
			accountServiceCreateHandler.ServeHTTP(w, r)
		case AccountServiceAuthenticateProcedure:
			accountServiceAuthenticateHandler.ServeHTTP(w, r)
		case AccountServiceRefreshAccessTokenProcedure:
			accountServiceRefreshAccessTokenHandler.ServeHTTP(w, r)
		case AccountServiceDeleteRefreshTokenProcedure:
			accountServiceDeleteRefreshTokenHandler.ServeHTTP(w, r)
		case AccountServiceRequestNewPasswordProcedure:
			accountServiceRequestNewPasswordHandler.ServeHTTP(w, r)
		case AccountServiceVerifyEmailProcedure:
			accountServiceVerifyEmailHandler.ServeHTTP(w, r)
		case AccountServiceUpdateProcedure:
			accountServiceUpdateHandler.ServeHTTP(w, r)
		case AccountServiceDeleteProcedure:
			accountServiceDeleteHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAccountServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAccountServiceHandler struct{}

func (UnimplementedAccountServiceHandler) Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("account.v1.AccountService.Create is not implemented"))
}

func (UnimplementedAccountServiceHandler) Authenticate(context.Context, *connect.Request[v1.AuthenticateRequest]) (*connect.Response[v1.AuthenticateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("account.v1.AccountService.Authenticate is not implemented"))
}

func (UnimplementedAccountServiceHandler) RefreshAccessToken(context.Context, *connect.Request[v1.RefreshAccessTokenRequest]) (*connect.Response[v1.RefreshAccessTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("account.v1.AccountService.RefreshAccessToken is not implemented"))
}

func (UnimplementedAccountServiceHandler) DeleteRefreshToken(context.Context, *connect.Request[v1.DeleteRefreshTokenRequest]) (*connect.Response[v1.DeleteRefreshTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("account.v1.AccountService.DeleteRefreshToken is not implemented"))
}

func (UnimplementedAccountServiceHandler) RequestNewPassword(context.Context, *connect.Request[v1.RequestNewPasswordRequest]) (*connect.Response[v1.RequestNewPasswordResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("account.v1.AccountService.RequestNewPassword is not implemented"))
}

func (UnimplementedAccountServiceHandler) VerifyEmail(context.Context, *connect.Request[v1.VerifyEmailRequest]) (*connect.Response[v1.VerifyEmailResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("account.v1.AccountService.VerifyEmail is not implemented"))
}

func (UnimplementedAccountServiceHandler) Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("account.v1.AccountService.Update is not implemented"))
}

func (UnimplementedAccountServiceHandler) Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("account.v1.AccountService.Delete is not implemented"))
}