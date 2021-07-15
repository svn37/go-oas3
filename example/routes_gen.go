// This file is generated by github.com/mikekonan/go-oas3. DO NOT EDIT.

package example

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	chi "github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	cast "github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var deleteTransactionsUUIDPathRegexParamRegex = regexp.MustCompile("^[.?\\d]+$")

type Hooks struct {
	RequestSecurityParseFailed    func(*http.Request, string, RequestProcessingResult)
	RequestSecurityParseCompleted func(*http.Request, string)
	RequestSecurityCheckFailed    func(*http.Request, string, string, RequestProcessingResult)
	RequestSecurityCheckCompleted func(*http.Request, string, string)
	RequestBodyUnmarshalFailed    func(*http.Request, string, RequestProcessingResult)
	RequestHeaderParseFailed      func(*http.Request, string, string, RequestProcessingResult)
	RequestPathParseFailed        func(*http.Request, string, string, RequestProcessingResult)
	RequestQueryParseFailed       func(*http.Request, string, string, RequestProcessingResult)
	RequestBodyValidationFailed   func(*http.Request, string, RequestProcessingResult)
	RequestHeaderValidationFailed func(*http.Request, string, RequestProcessingResult)
	RequestPathValidationFailed   func(*http.Request, string, RequestProcessingResult)
	RequestQueryValidationFailed  func(*http.Request, string, RequestProcessingResult)
	RequestBodyUnmarshalCompleted func(*http.Request, string)
	RequestHeaderParseCompleted   func(*http.Request, string)
	RequestPathParseCompleted     func(*http.Request, string)
	RequestQueryParseCompleted    func(*http.Request, string)
	RequestParseCompleted         func(*http.Request, string)
	RequestProcessingCompleted    func(*http.Request, string)
	RequestRedirectStarted        func(*http.Request, string, string)
	ResponseBodyMarshalCompleted  func(*http.Request, string)
	ResponseBodyWriteCompleted    func(*http.Request, string, int)
	ResponseBodyMarshalFailed     func(http.ResponseWriter, *http.Request, string, error)
	ResponseBodyWriteFailed       func(*http.Request, string, int, error)
	ServiceCompleted              func(*http.Request, string)
}

type requestProcessingResultType uint8

const (
	BodyUnmarshalFailed requestProcessingResultType = iota + 1
	BodyValidationFailed
	HeaderParseFailed
	HeaderValidationFailed
	QueryParseFailed
	QueryValidationFailed
	PathParseFailed
	PathValidationFailed
	SecurityParseFailed
	SecurityCheckFailed
	ParseSucceed
)

type RequestProcessingResult struct {
	error error
	typee requestProcessingResultType
}

func NewRequestProcessingResult(t requestProcessingResultType, err error) RequestProcessingResult {
	return RequestProcessingResult{
		error: err,
		typee: t,
	}
}

func (r RequestProcessingResult) Type() requestProcessingResultType {
	return r.typee
}

func (r RequestProcessingResult) Err() error {
	return r.error
}

func TransactionsHandler(impl TransactionsService, r chi.Router, hooks *Hooks, securitySchemas SecuritySchemas) http.Handler {
	router := &transactionsRouter{router: r, service: impl, hooks: hooks}

	router.securityHandlers = map[SecurityScheme]securityProcessor{
		SecuritySchemeBearer: {
			scheme:  SecuritySchemeBearer,
			extract: securityExtractorsFuncs[SecuritySchemeBearer],
			handle:  securitySchemas.SecuritySchemeBearer,
		},
		SecuritySchemeCookie: {
			scheme:  SecuritySchemeCookie,
			extract: securityExtractorsFuncs[SecuritySchemeCookie],
			handle:  securitySchemas.SecuritySchemeCookie,
		},
	}

	router.mount()

	return router.router
}

type transactionsRouter struct {
	router           chi.Router
	service          TransactionsService
	hooks            *Hooks
	securityHandlers map[SecurityScheme]securityProcessor
}

func (router *transactionsRouter) mount() {
	router.router.Post("/transaction", router.PostTransaction)
	router.router.Delete("/transactions/{uuid}", router.DeleteTransactionsUUID)
}

func (router *transactionsRouter) parsePostTransactionRequest(r *http.Request) (request PostTransactionRequest) {
	request.ProcessingResult = RequestProcessingResult{typee: ParseSucceed}

	headerXSignature := r.Header.Get("x-signature")
	request.Header.XSignature = headerXSignature

	if err := request.Header.Validate(); err != nil {
		request.ProcessingResult = RequestProcessingResult{error: err, typee: HeaderValidationFailed}
		if router.hooks.RequestHeaderValidationFailed != nil {
			router.hooks.RequestHeaderValidationFailed(r, "PostTransaction", request.ProcessingResult)
		}

		return
	}

	if router.hooks.RequestHeaderParseCompleted != nil {
		router.hooks.RequestHeaderParseCompleted(r, "PostTransaction")
	}

	var (
		body      CreateTransactionRequest
		decodeErr error
	)
	decodeErr = json.NewDecoder(r.Body).Decode(&body)
	if decodeErr != nil {
		request.ProcessingResult = RequestProcessingResult{error: decodeErr, typee: BodyUnmarshalFailed}
		if router.hooks.RequestBodyUnmarshalFailed != nil {
			router.hooks.RequestBodyUnmarshalFailed(r, "PostTransaction", request.ProcessingResult)

			return
		}

		return
	}

	request.Body = body

	if err := request.Body.Validate(); err != nil {
		request.ProcessingResult = RequestProcessingResult{error: err, typee: BodyValidationFailed}
		if router.hooks.RequestBodyValidationFailed != nil {
			router.hooks.RequestBodyValidationFailed(r, "PostTransaction", request.ProcessingResult)
		}

		return
	}

	if router.hooks.RequestBodyUnmarshalCompleted != nil {
		router.hooks.RequestBodyUnmarshalCompleted(r, "PostTransaction")
	}

	if router.hooks.RequestParseCompleted != nil {
		router.hooks.RequestParseCompleted(r, "PostTransaction")
	}

	return
}

func (router *transactionsRouter) PostTransaction(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	response := router.service.PostTransaction(r.Context(), router.parsePostTransactionRequest(r))

	if response.statusCode() == 302 && response.redirectURL() != "" {
		if router.hooks.RequestRedirectStarted != nil {
			router.hooks.RequestRedirectStarted(r, "PostTransaction", response.redirectURL())
		}

		http.Redirect(w, r, response.redirectURL(), 302)

		if router.hooks.ServiceCompleted != nil {
			router.hooks.ServiceCompleted(r, "PostTransaction")
		}

		return
	}

	for header, value := range response.headers() {
		w.Header().Set(header, value)
	}

	for _, c := range response.cookies() {
		cookie := c
		http.SetCookie(w, &cookie)
	}

	if router.hooks.RequestProcessingCompleted != nil {
		router.hooks.RequestProcessingCompleted(r, "PostTransaction")
	}

	if len(response.contentType()) > 0 {
		w.Header().Set("content-type", response.contentType())
	}

	w.WriteHeader(response.statusCode())

	if response.body() != nil {
		var (
			data []byte
			err  error
		)

		switch response.contentType() {
		case "application/xml":
			data, err = xml.Marshal(response.body())
		case "application/octet-stream":
			var ok bool
			if data, ok = (response.body()).([]byte); !ok {
				err = errors.New("body is not []byte")
			}
		case "text/html":
			data = []byte(fmt.Sprint(response.body()))
		case "application/json":
			fallthrough
		default:
			data, err = json.Marshal(response.body())
		}

		if err != nil {
			if router.hooks.ResponseBodyMarshalFailed != nil {
				router.hooks.ResponseBodyMarshalFailed(w, r, "PostTransaction", err)
			}

			return
		}

		if router.hooks.ResponseBodyMarshalCompleted != nil {
			router.hooks.ResponseBodyMarshalCompleted(r, "PostTransaction")
		}

		count, err := w.Write(data)
		if err != nil {
			if router.hooks.ResponseBodyWriteFailed != nil {
				router.hooks.ResponseBodyWriteFailed(r, "PostTransaction", count, err)
			}

			if router.hooks.ResponseBodyWriteCompleted != nil {
				router.hooks.ResponseBodyWriteCompleted(r, "PostTransaction", count)
			}

			return
		}

		if router.hooks.ResponseBodyWriteCompleted != nil {
			router.hooks.ResponseBodyWriteCompleted(r, "PostTransaction", count)
		}
	}

	if router.hooks.ServiceCompleted != nil {
		router.hooks.ServiceCompleted(r, "PostTransaction")
	}
}

func (router *transactionsRouter) parseDeleteTransactionsUUIDRequest(r *http.Request) (request DeleteTransactionsUUIDRequest) {
	request.ProcessingResult = RequestProcessingResult{typee: ParseSucceed}

	isSecurityCheckPassed := false
	for _, processors := range [][]securityProcessor{{router.securityHandlers[SecuritySchemeBearer]}, {router.securityHandlers[SecuritySchemeCookie]}} {
		isLinkedChecksValid := true

		for _, processor := range processors {
			name, value, isExtracted := processor.extract(r)

			if !isExtracted {
				isLinkedChecksValid = false
				break
			}

			if err := processor.handle(r, processor.scheme, name, value); err != nil {
				if router.hooks.RequestSecurityCheckFailed != nil {
					router.hooks.RequestSecurityCheckFailed(r, "DeleteTransactionsUUID", string(processor.scheme), RequestProcessingResult{error: err, typee: SecurityCheckFailed})
				}

				isLinkedChecksValid = false

				break
			}

			if router.hooks.RequestSecurityCheckCompleted != nil {
				router.hooks.RequestSecurityCheckCompleted(r, "DeleteTransactionsUUID", string(processor.scheme))
			}

			if len(request.SecurityCheckResults) == 0 {
				request.SecurityCheckResults = map[SecurityScheme]string{}
			}

			request.SecurityCheckResults[processor.scheme] = value
		}

		if isLinkedChecksValid {
			isSecurityCheckPassed = true
			break
		}
	}

	if !isSecurityCheckPassed {
		err := fmt.Errorf("failed passing security checks")

		request.ProcessingResult = RequestProcessingResult{error: err, typee: SecurityParseFailed}

		if router.hooks.RequestSecurityParseFailed != nil {
			router.hooks.RequestSecurityParseFailed(r, "DeleteTransactionsUUID", request.ProcessingResult)
		}

		return
	}

	if router.hooks.RequestSecurityParseCompleted != nil {
		router.hooks.RequestSecurityParseCompleted(r, "DeleteTransactionsUUID")
	}

	headerXSignature := r.Header.Get("x-signature")
	request.Header.XSignature = headerXSignature

	if err := request.Header.Validate(); err != nil {
		request.ProcessingResult = RequestProcessingResult{error: err, typee: HeaderValidationFailed}
		if router.hooks.RequestHeaderValidationFailed != nil {
			router.hooks.RequestHeaderValidationFailed(r, "DeleteTransactionsUUID", request.ProcessingResult)
		}

		return
	}

	if router.hooks.RequestHeaderParseCompleted != nil {
		router.hooks.RequestHeaderParseCompleted(r, "DeleteTransactionsUUID")
	}

	pathUUID := chi.URLParam(r, "uuid")
	if pathUUID == "" {
		err := fmt.Errorf("uuid is empty")

		request.ProcessingResult = RequestProcessingResult{error: err, typee: PathParseFailed}
		if router.hooks.RequestPathParseFailed != nil {
			router.hooks.RequestPathParseFailed(r, "DeleteTransactionsUUID", "uuid", request.ProcessingResult)
		}

		return
	}

	request.Path.UUID = pathUUID

	pathRegexParam := chi.URLParam(r, "regexParam")
	if pathRegexParam == "" {
		err := fmt.Errorf("regexParam is empty")

		request.ProcessingResult = RequestProcessingResult{error: err, typee: PathParseFailed}
		if router.hooks.RequestPathParseFailed != nil {
			router.hooks.RequestPathParseFailed(r, "DeleteTransactionsUUID", "regexParam", request.ProcessingResult)
		}

		return
	}

	if !deleteTransactionsUUIDPathRegexParamRegex.MatchString(request.Path.RegexParam) {
		err := fmt.Errorf("regexParam not matched by the '^[.?\\d]+$' regex")

		request.ProcessingResult = RequestProcessingResult{error: err, typee: PathParseFailed}
		if router.hooks.RequestPathParseFailed != nil {
			router.hooks.RequestPathParseFailed(r, "DeleteTransactionsUUID", "regexParam", request.ProcessingResult)
		}

		return
	}

	request.Path.RegexParam = pathRegexParam

	if err := request.Path.Validate(); err != nil {
		request.ProcessingResult = RequestProcessingResult{error: err, typee: PathValidationFailed}
		if router.hooks.RequestPathValidationFailed != nil {
			router.hooks.RequestPathValidationFailed(r, "DeleteTransactionsUUID", request.ProcessingResult)
		}

		return
	}

	if router.hooks.RequestPathParseCompleted != nil {
		router.hooks.RequestPathParseCompleted(r, "DeleteTransactionsUUID")
	}

	queryTimeParamStr := r.URL.Query().Get("timeParam")
	if queryTimeParamStr != "" {
		queryTimeParam, err := cast.ToTimeE(queryTimeParamStr)
		if err != nil {
			request.ProcessingResult = RequestProcessingResult{error: err, typee: QueryParseFailed}
			if router.hooks.RequestQueryParseFailed != nil {
				router.hooks.RequestQueryParseFailed(r, "DeleteTransactionsUUID", "timeParam", request.ProcessingResult)
			}

			return
		}

		request.Query.TimeParam = queryTimeParam
	}

	if err := request.Query.Validate(); err != nil {
		request.ProcessingResult = RequestProcessingResult{error: err, typee: QueryValidationFailed}
		if router.hooks.RequestQueryValidationFailed != nil {
			router.hooks.RequestQueryValidationFailed(r, "DeleteTransactionsUUID", request.ProcessingResult)
		}

		return
	}

	if router.hooks.RequestQueryParseCompleted != nil {
		router.hooks.RequestQueryParseCompleted(r, "DeleteTransactionsUUID")
	}

	if router.hooks.RequestParseCompleted != nil {
		router.hooks.RequestParseCompleted(r, "DeleteTransactionsUUID")
	}

	return
}

func (router *transactionsRouter) DeleteTransactionsUUID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	response := router.service.DeleteTransactionsUUID(r.Context(), router.parseDeleteTransactionsUUIDRequest(r))

	if response.statusCode() == 302 && response.redirectURL() != "" {
		if router.hooks.RequestRedirectStarted != nil {
			router.hooks.RequestRedirectStarted(r, "DeleteTransactionsUUID", response.redirectURL())
		}

		http.Redirect(w, r, response.redirectURL(), 302)

		if router.hooks.ServiceCompleted != nil {
			router.hooks.ServiceCompleted(r, "DeleteTransactionsUUID")
		}

		return
	}

	for header, value := range response.headers() {
		w.Header().Set(header, value)
	}

	for _, c := range response.cookies() {
		cookie := c
		http.SetCookie(w, &cookie)
	}

	if router.hooks.RequestProcessingCompleted != nil {
		router.hooks.RequestProcessingCompleted(r, "DeleteTransactionsUUID")
	}

	if len(response.contentType()) > 0 {
		w.Header().Set("content-type", response.contentType())
	}

	w.WriteHeader(response.statusCode())

	if response.body() != nil {
		var (
			data []byte
			err  error
		)

		switch response.contentType() {
		case "application/xml":
			data, err = xml.Marshal(response.body())
		case "application/octet-stream":
			var ok bool
			if data, ok = (response.body()).([]byte); !ok {
				err = errors.New("body is not []byte")
			}
		case "text/html":
			data = []byte(fmt.Sprint(response.body()))
		case "application/json":
			fallthrough
		default:
			data, err = json.Marshal(response.body())
		}

		if err != nil {
			if router.hooks.ResponseBodyMarshalFailed != nil {
				router.hooks.ResponseBodyMarshalFailed(w, r, "DeleteTransactionsUUID", err)
			}

			return
		}

		if router.hooks.ResponseBodyMarshalCompleted != nil {
			router.hooks.ResponseBodyMarshalCompleted(r, "DeleteTransactionsUUID")
		}

		count, err := w.Write(data)
		if err != nil {
			if router.hooks.ResponseBodyWriteFailed != nil {
				router.hooks.ResponseBodyWriteFailed(r, "DeleteTransactionsUUID", count, err)
			}

			if router.hooks.ResponseBodyWriteCompleted != nil {
				router.hooks.ResponseBodyWriteCompleted(r, "DeleteTransactionsUUID", count)
			}

			return
		}

		if router.hooks.ResponseBodyWriteCompleted != nil {
			router.hooks.ResponseBodyWriteCompleted(r, "DeleteTransactionsUUID", count)
		}
	}

	if router.hooks.ServiceCompleted != nil {
		router.hooks.ServiceCompleted(r, "DeleteTransactionsUUID")
	}
}

func CallbacksHandler(impl CallbacksService, r chi.Router, hooks *Hooks, securitySchemas SecuritySchemas) http.Handler {
	router := &callbacksRouter{router: r, service: impl, hooks: hooks}

	router.securityHandlers = map[SecurityScheme]securityProcessor{
		SecuritySchemeCookie: {
			scheme:  SecuritySchemeCookie,
			extract: securityExtractorsFuncs[SecuritySchemeCookie],
			handle:  securitySchemas.SecuritySchemeCookie,
		},
	}

	router.mount()

	return router.router
}

type callbacksRouter struct {
	router           chi.Router
	service          CallbacksService
	hooks            *Hooks
	securityHandlers map[SecurityScheme]securityProcessor
}

func (router *callbacksRouter) mount() {
	router.router.Post("/callbacks/{callbackType}", router.PostCallbacksCallbackType)
}

func (router *callbacksRouter) parsePostCallbacksCallbackTypeRequest(r *http.Request) (request PostCallbacksCallbackTypeRequest) {
	request.ProcessingResult = RequestProcessingResult{typee: ParseSucceed}

	isSecurityCheckPassed := false
	for _, processors := range [][]securityProcessor{{router.securityHandlers[SecuritySchemeCookie]}} {
		isLinkedChecksValid := true

		for _, processor := range processors {
			name, value, isExtracted := processor.extract(r)

			if !isExtracted {
				isLinkedChecksValid = false
				break
			}

			if err := processor.handle(r, processor.scheme, name, value); err != nil {
				if router.hooks.RequestSecurityCheckFailed != nil {
					router.hooks.RequestSecurityCheckFailed(r, "PostCallbacksCallbackType", string(processor.scheme), RequestProcessingResult{error: err, typee: SecurityCheckFailed})
				}

				isLinkedChecksValid = false

				break
			}

			if router.hooks.RequestSecurityCheckCompleted != nil {
				router.hooks.RequestSecurityCheckCompleted(r, "PostCallbacksCallbackType", string(processor.scheme))
			}

			if len(request.SecurityCheckResults) == 0 {
				request.SecurityCheckResults = map[SecurityScheme]string{}
			}

			request.SecurityCheckResults[processor.scheme] = value
		}

		if isLinkedChecksValid {
			isSecurityCheckPassed = true
			break
		}
	}

	if !isSecurityCheckPassed {
		err := fmt.Errorf("failed passing security checks")

		request.ProcessingResult = RequestProcessingResult{error: err, typee: SecurityParseFailed}

		if router.hooks.RequestSecurityParseFailed != nil {
			router.hooks.RequestSecurityParseFailed(r, "PostCallbacksCallbackType", request.ProcessingResult)
		}

		return
	}

	if router.hooks.RequestSecurityParseCompleted != nil {
		router.hooks.RequestSecurityParseCompleted(r, "PostCallbacksCallbackType")
	}

	pathCallbackType := chi.URLParam(r, "callbackType")
	if pathCallbackType == "" {
		err := fmt.Errorf("callbackType is empty")

		request.ProcessingResult = RequestProcessingResult{error: err, typee: PathParseFailed}
		if router.hooks.RequestPathParseFailed != nil {
			router.hooks.RequestPathParseFailed(r, "PostCallbacksCallbackType", "callbackType", request.ProcessingResult)
		}

		return
	}

	request.Path.CallbackType = pathCallbackType

	if err := request.Path.Validate(); err != nil {
		request.ProcessingResult = RequestProcessingResult{error: err, typee: PathValidationFailed}
		if router.hooks.RequestPathValidationFailed != nil {
			router.hooks.RequestPathValidationFailed(r, "PostCallbacksCallbackType", request.ProcessingResult)
		}

		return
	}

	if router.hooks.RequestPathParseCompleted != nil {
		router.hooks.RequestPathParseCompleted(r, "PostCallbacksCallbackType")
	}

	queryHasSmthStr := r.URL.Query().Get("hasSmth")
	if queryHasSmthStr != "" {
		queryHasSmth, err := strconv.ParseBool(queryHasSmthStr)
		if err != nil {
			request.ProcessingResult = RequestProcessingResult{error: err, typee: QueryParseFailed}
			if router.hooks.RequestQueryParseFailed != nil {
				router.hooks.RequestQueryParseFailed(r, "PostCallbacksCallbackType", "hasSmth", request.ProcessingResult)
			}

			return
		}

		request.Query.HasSmth = queryHasSmth
	}

	if err := request.Query.Validate(); err != nil {
		request.ProcessingResult = RequestProcessingResult{error: err, typee: QueryValidationFailed}
		if router.hooks.RequestQueryValidationFailed != nil {
			router.hooks.RequestQueryValidationFailed(r, "PostCallbacksCallbackType", request.ProcessingResult)
		}

		return
	}

	if router.hooks.RequestQueryParseCompleted != nil {
		router.hooks.RequestQueryParseCompleted(r, "PostCallbacksCallbackType")
	}

	var (
		body      RawPayload
		decodeErr error
	)
	var (
		buf     interface{}
		ok      bool
		readErr error
	)
	if buf, readErr = ioutil.ReadAll(r.Body); readErr == nil {
		if body, ok = buf.(RawPayload); !ok {
			decodeErr = errors.New("body is not []byte")
		}
	}
	if decodeErr != nil {
		request.ProcessingResult = RequestProcessingResult{error: decodeErr, typee: BodyUnmarshalFailed}
		if router.hooks.RequestBodyUnmarshalFailed != nil {
			router.hooks.RequestBodyUnmarshalFailed(r, "PostCallbacksCallbackType", request.ProcessingResult)

			return
		}

		return
	}

	request.Body = body

	if router.hooks.RequestBodyUnmarshalCompleted != nil {
		router.hooks.RequestBodyUnmarshalCompleted(r, "PostCallbacksCallbackType")
	}

	if router.hooks.RequestParseCompleted != nil {
		router.hooks.RequestParseCompleted(r, "PostCallbacksCallbackType")
	}

	return
}

func (router *callbacksRouter) PostCallbacksCallbackType(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	response := router.service.PostCallbacksCallbackType(r.Context(), router.parsePostCallbacksCallbackTypeRequest(r))

	if response.statusCode() == 302 && response.redirectURL() != "" {
		if router.hooks.RequestRedirectStarted != nil {
			router.hooks.RequestRedirectStarted(r, "PostCallbacksCallbackType", response.redirectURL())
		}

		http.Redirect(w, r, response.redirectURL(), 302)

		if router.hooks.ServiceCompleted != nil {
			router.hooks.ServiceCompleted(r, "PostCallbacksCallbackType")
		}

		return
	}

	for header, value := range response.headers() {
		w.Header().Set(header, value)
	}

	for _, c := range response.cookies() {
		cookie := c
		http.SetCookie(w, &cookie)
	}

	if router.hooks.RequestProcessingCompleted != nil {
		router.hooks.RequestProcessingCompleted(r, "PostCallbacksCallbackType")
	}

	if len(response.contentType()) > 0 {
		w.Header().Set("content-type", response.contentType())
	}

	w.WriteHeader(response.statusCode())

	if response.body() != nil {
		var (
			data []byte
			err  error
		)

		switch response.contentType() {
		case "application/xml":
			data, err = xml.Marshal(response.body())
		case "application/octet-stream":
			var ok bool
			if data, ok = (response.body()).([]byte); !ok {
				err = errors.New("body is not []byte")
			}
		case "text/html":
			data = []byte(fmt.Sprint(response.body()))
		case "application/json":
			fallthrough
		default:
			data, err = json.Marshal(response.body())
		}

		if err != nil {
			if router.hooks.ResponseBodyMarshalFailed != nil {
				router.hooks.ResponseBodyMarshalFailed(w, r, "PostCallbacksCallbackType", err)
			}

			return
		}

		if router.hooks.ResponseBodyMarshalCompleted != nil {
			router.hooks.ResponseBodyMarshalCompleted(r, "PostCallbacksCallbackType")
		}

		count, err := w.Write(data)
		if err != nil {
			if router.hooks.ResponseBodyWriteFailed != nil {
				router.hooks.ResponseBodyWriteFailed(r, "PostCallbacksCallbackType", count, err)
			}

			if router.hooks.ResponseBodyWriteCompleted != nil {
				router.hooks.ResponseBodyWriteCompleted(r, "PostCallbacksCallbackType", count)
			}

			return
		}

		if router.hooks.ResponseBodyWriteCompleted != nil {
			router.hooks.ResponseBodyWriteCompleted(r, "PostCallbacksCallbackType", count)
		}
	}

	if router.hooks.ServiceCompleted != nil {
		router.hooks.ServiceCompleted(r, "PostCallbacksCallbackType")
	}
}

type response struct {
	statusCode  int
	body        interface{}
	contentType string
	redirectURL string
	headers     map[string]string
	cookies     []http.Cookie
}

type responseInterface interface {
	statusCode() int
	body() interface{}
	contentType() string
	redirectURL() string
	cookies() []http.Cookie
	headers() map[string]string
}

type PostCallbacksCallbackTypeResponse interface {
	responseInterface
	postCallbacksCallbackTypeResponse()
}

type postCallbacksCallbackTypeResponse struct {
	response
}

func (postCallbacksCallbackTypeResponse) postCallbacksCallbackTypeResponse() {}

func (response postCallbacksCallbackTypeResponse) statusCode() int {
	return response.response.statusCode
}

func (response postCallbacksCallbackTypeResponse) body() interface{} {
	return response.response.body
}

func (response postCallbacksCallbackTypeResponse) contentType() string {
	return response.response.contentType
}

func (response postCallbacksCallbackTypeResponse) redirectURL() string {
	return response.response.redirectURL
}

func (response postCallbacksCallbackTypeResponse) headers() map[string]string {
	return response.response.headers
}

func (response postCallbacksCallbackTypeResponse) cookies() []http.Cookie {
	return response.response.cookies
}

type PostTransactionResponse interface {
	responseInterface
	postTransactionResponse()
}

type postTransactionResponse struct {
	response
}

func (postTransactionResponse) postTransactionResponse() {}

func (response postTransactionResponse) statusCode() int {
	return response.response.statusCode
}

func (response postTransactionResponse) body() interface{} {
	return response.response.body
}

func (response postTransactionResponse) contentType() string {
	return response.response.contentType
}

func (response postTransactionResponse) redirectURL() string {
	return response.response.redirectURL
}

func (response postTransactionResponse) headers() map[string]string {
	return response.response.headers
}

func (response postTransactionResponse) cookies() []http.Cookie {
	return response.response.cookies
}

type DeleteTransactionsUUIDResponse interface {
	responseInterface
	deleteTransactionsUUIDResponse()
}

type deleteTransactionsUUIDResponse struct {
	response
}

func (deleteTransactionsUUIDResponse) deleteTransactionsUUIDResponse() {}

func (response deleteTransactionsUUIDResponse) statusCode() int {
	return response.response.statusCode
}

func (response deleteTransactionsUUIDResponse) body() interface{} {
	return response.response.body
}

func (response deleteTransactionsUUIDResponse) contentType() string {
	return response.response.contentType
}

func (response deleteTransactionsUUIDResponse) redirectURL() string {
	return response.response.redirectURL
}

func (response deleteTransactionsUUIDResponse) headers() map[string]string {
	return response.response.headers
}

func (response deleteTransactionsUUIDResponse) cookies() []http.Cookie {
	return response.response.cookies
}

type postTransactionStatusCodeResponseBuilder struct {
	response
}

func PostTransactionResponseBuilder() *postTransactionStatusCodeResponseBuilder {
	return new(postTransactionStatusCodeResponseBuilder)
}

func (builder *postTransactionStatusCodeResponseBuilder) StatusCode201() *postTransaction201ContentTypeBuilder {
	builder.response.statusCode = 201

	return &postTransaction201ContentTypeBuilder{response: builder.response}
}

type postTransaction201ContentTypeBuilder struct {
	response
}

type PostTransaction201ApplicationJsonResponseBuilder struct {
	response
}

func (builder *PostTransaction201ApplicationJsonResponseBuilder) Build() PostTransactionResponse {
	return postTransactionResponse{response: builder.response}
}

func (builder *postTransaction201ContentTypeBuilder) ApplicationJson() *postTransaction201ApplicationJsonBodyBuilder {
	builder.response.contentType = "application/json"

	return &postTransaction201ApplicationJsonBodyBuilder{response: builder.response}
}

type postTransaction201ApplicationJsonBodyBuilder struct {
	response
}

func (builder *postTransaction201ApplicationJsonBodyBuilder) Body(body GenericResponse) *PostTransaction201ApplicationJsonResponseBuilder {
	builder.response.body = body

	return &PostTransaction201ApplicationJsonResponseBuilder{response: builder.response}
}

func (builder *postTransactionStatusCodeResponseBuilder) StatusCode400() *postTransaction400ContentTypeBuilder {
	builder.response.statusCode = 400

	return &postTransaction400ContentTypeBuilder{response: builder.response}
}

type postTransaction400ContentTypeBuilder struct {
	response
}

type PostTransaction400ApplicationJsonResponseBuilder struct {
	response
}

func (builder *PostTransaction400ApplicationJsonResponseBuilder) Build() PostTransactionResponse {
	return postTransactionResponse{response: builder.response}
}

func (builder *postTransaction400ContentTypeBuilder) ApplicationJson() *postTransaction400ApplicationJsonBodyBuilder {
	builder.response.contentType = "application/json"

	return &postTransaction400ApplicationJsonBodyBuilder{response: builder.response}
}

type postTransaction400ApplicationJsonBodyBuilder struct {
	response
}

func (builder *postTransaction400ApplicationJsonBodyBuilder) Body(body GenericResponse) *PostTransaction400ApplicationJsonResponseBuilder {
	builder.response.body = body

	return &PostTransaction400ApplicationJsonResponseBuilder{response: builder.response}
}

func (builder *postTransactionStatusCodeResponseBuilder) StatusCode500() *postTransaction500ContentTypeBuilder {
	builder.response.statusCode = 500

	return &postTransaction500ContentTypeBuilder{response: builder.response}
}

type postTransaction500ContentTypeBuilder struct {
	response
}

type PostTransaction500ApplicationJsonResponseBuilder struct {
	response
}

func (builder *PostTransaction500ApplicationJsonResponseBuilder) Build() PostTransactionResponse {
	return postTransactionResponse{response: builder.response}
}

func (builder *postTransaction500ContentTypeBuilder) ApplicationJson() *postTransaction500ApplicationJsonBodyBuilder {
	builder.response.contentType = "application/json"

	return &postTransaction500ApplicationJsonBodyBuilder{response: builder.response}
}

type postTransaction500ApplicationJsonBodyBuilder struct {
	response
}

func (builder *postTransaction500ApplicationJsonBodyBuilder) Body(body GenericResponse) *PostTransaction500ApplicationJsonResponseBuilder {
	builder.response.body = body

	return &PostTransaction500ApplicationJsonResponseBuilder{response: builder.response}
}

type deleteTransactionsUUIDStatusCodeResponseBuilder struct {
	response
}

func DeleteTransactionsUUIDResponseBuilder() *deleteTransactionsUUIDStatusCodeResponseBuilder {
	return new(deleteTransactionsUUIDStatusCodeResponseBuilder)
}

func (builder *deleteTransactionsUUIDStatusCodeResponseBuilder) StatusCode200() *deleteTransactionsUUID200ContentTypeBuilder {
	builder.response.statusCode = 200

	return &deleteTransactionsUUID200ContentTypeBuilder{response: builder.response}
}

type deleteTransactionsUUID200ContentTypeBuilder struct {
	response
}

type DeleteTransactionsUUID200ApplicationJsonResponseBuilder struct {
	response
}

func (builder *DeleteTransactionsUUID200ApplicationJsonResponseBuilder) Build() DeleteTransactionsUUIDResponse {
	return deleteTransactionsUUIDResponse{response: builder.response}
}

func (builder *deleteTransactionsUUID200ContentTypeBuilder) ApplicationJson() *deleteTransactionsUUID200ApplicationJsonBodyBuilder {
	builder.response.contentType = "application/json"

	return &deleteTransactionsUUID200ApplicationJsonBodyBuilder{response: builder.response}
}

type deleteTransactionsUUID200ApplicationJsonBodyBuilder struct {
	response
}

func (builder *deleteTransactionsUUID200ApplicationJsonBodyBuilder) Body(body GenericResponse) *DeleteTransactionsUUID200ApplicationJsonResponseBuilder {
	builder.response.body = body

	return &DeleteTransactionsUUID200ApplicationJsonResponseBuilder{response: builder.response}
}

func (builder *deleteTransactionsUUIDStatusCodeResponseBuilder) StatusCode400() *deleteTransactionsUUID400ContentTypeBuilder {
	builder.response.statusCode = 400

	return &deleteTransactionsUUID400ContentTypeBuilder{response: builder.response}
}

type deleteTransactionsUUID400ContentTypeBuilder struct {
	response
}

type DeleteTransactionsUUID400ApplicationJsonResponseBuilder struct {
	response
}

func (builder *DeleteTransactionsUUID400ApplicationJsonResponseBuilder) Build() DeleteTransactionsUUIDResponse {
	return deleteTransactionsUUIDResponse{response: builder.response}
}

func (builder *deleteTransactionsUUID400ContentTypeBuilder) ApplicationJson() *deleteTransactionsUUID400ApplicationJsonBodyBuilder {
	builder.response.contentType = "application/json"

	return &deleteTransactionsUUID400ApplicationJsonBodyBuilder{response: builder.response}
}

type deleteTransactionsUUID400ApplicationJsonBodyBuilder struct {
	response
}

func (builder *deleteTransactionsUUID400ApplicationJsonBodyBuilder) Body(body GenericResponse) *DeleteTransactionsUUID400ApplicationJsonResponseBuilder {
	builder.response.body = body

	return &DeleteTransactionsUUID400ApplicationJsonResponseBuilder{response: builder.response}
}

type postCallbacksCallbackTypeStatusCodeResponseBuilder struct {
	response
}

func PostCallbacksCallbackTypeResponseBuilder() *postCallbacksCallbackTypeStatusCodeResponseBuilder {
	return new(postCallbacksCallbackTypeStatusCodeResponseBuilder)
}

func (builder *postCallbacksCallbackTypeStatusCodeResponseBuilder) StatusCode200() *postCallbacksCallbackType200HeadersBuilder {
	builder.response.statusCode = 200

	return &postCallbacksCallbackType200HeadersBuilder{response: builder.response}
}

type PostCallbacksCallbackType200Headers struct {
	XJwsSignature string
}

func (headers PostCallbacksCallbackType200Headers) toMap() map[string]string {
	return map[string]string{"x-jws-signature": cast.ToString(headers.XJwsSignature)}
}

type postCallbacksCallbackType200HeadersBuilder struct {
	response
}

func (builder *postCallbacksCallbackType200HeadersBuilder) Headers(headers PostCallbacksCallbackType200Headers) *postCallbacksCallbackType200CookiesBuilder {
	builder.headers = headers.toMap()

	return &postCallbacksCallbackType200CookiesBuilder{response: builder.response}
}

type postCallbacksCallbackType200CookiesBuilder struct {
	response
}

func (builder *postCallbacksCallbackType200CookiesBuilder) SetCookie(cookie ...http.Cookie) *postCallbacksCallbackType200ContentTypeBuilder {
	builder.cookies = cookie
	return &postCallbacksCallbackType200ContentTypeBuilder{response: builder.response}
}

type postCallbacksCallbackType200ContentTypeBuilder struct {
	response
}

type PostCallbacksCallbackType200ApplicationOctetStreamResponseBuilder struct {
	response
}

func (builder *PostCallbacksCallbackType200ApplicationOctetStreamResponseBuilder) Build() PostCallbacksCallbackTypeResponse {
	return postCallbacksCallbackTypeResponse{response: builder.response}
}

func (builder *postCallbacksCallbackType200ContentTypeBuilder) ApplicationOctetStream() *postCallbacksCallbackType200ApplicationOctetStreamBodyBuilder {
	builder.response.contentType = "application/octet-stream"

	return &postCallbacksCallbackType200ApplicationOctetStreamBodyBuilder{response: builder.response}
}

type postCallbacksCallbackType200ApplicationOctetStreamBodyBuilder struct {
	response
}

func (builder *postCallbacksCallbackType200ApplicationOctetStreamBodyBuilder) Body(body RawPayload) *PostCallbacksCallbackType200ApplicationOctetStreamResponseBuilder {
	builder.response.body = body

	return &PostCallbacksCallbackType200ApplicationOctetStreamResponseBuilder{response: builder.response}
}

type CallbacksService interface {
	PostCallbacksCallbackType(context.Context, PostCallbacksCallbackTypeRequest) PostCallbacksCallbackTypeResponse
}

type TransactionsService interface {
	PostTransaction(context.Context, PostTransactionRequest) PostTransactionResponse
	DeleteTransactionsUUID(context.Context, DeleteTransactionsUUIDRequest) DeleteTransactionsUUIDResponse
}

type PostCallbacksCallbackTypeRequestQuery struct {
	HasSmth bool
}

func (query PostCallbacksCallbackTypeRequestQuery) GetHasSmth() bool {
	return query.HasSmth
}

func (query PostCallbacksCallbackTypeRequestQuery) Validate() error {
	return nil
}

type PostCallbacksCallbackTypeRequestPath struct {
	CallbackType string
}

func (path PostCallbacksCallbackTypeRequestPath) GetCallbackType() string {
	return path.CallbackType
}

func (path PostCallbacksCallbackTypeRequestPath) Validate() error {
	return nil
}

type PostCallbacksCallbackTypeRequest struct {
	Body                 RawPayload
	Path                 PostCallbacksCallbackTypeRequestPath
	Query                PostCallbacksCallbackTypeRequestQuery
	ProcessingResult     RequestProcessingResult
	SecurityCheckResults map[SecurityScheme]string
}

type PostTransactionRequestHeader struct {
	XSignature string
}

func (header PostTransactionRequestHeader) GetXSignature() string {
	return header.XSignature
}

func (header PostTransactionRequestHeader) Validate() error {
	return validation.ValidateStruct(&header,
		validation.Field(&header.XSignature, validation.RuneLength(0, 5)))
}

type PostTransactionRequest struct {
	Body             CreateTransactionRequest
	Header           PostTransactionRequestHeader
	ProcessingResult RequestProcessingResult
}

type DeleteTransactionsUUIDRequestHeader struct {
	XSignature string
}

func (header DeleteTransactionsUUIDRequestHeader) GetXSignature() string {
	return header.XSignature
}

func (header DeleteTransactionsUUIDRequestHeader) Validate() error {
	return validation.ValidateStruct(&header,
		validation.Field(&header.XSignature, validation.RuneLength(0, 5)))
}

type DeleteTransactionsUUIDRequestPath struct {
	RegexParam string
	UUID       string
}

func (path DeleteTransactionsUUIDRequestPath) GetRegexParam() string {
	return path.RegexParam
}

func (path DeleteTransactionsUUIDRequestPath) GetUUID() string {
	return path.UUID
}

func (path DeleteTransactionsUUIDRequestPath) Validate() error {
	return validation.ValidateStruct(&path,
		validation.Field(&path.RegexParam, validation.Required, validation.RuneLength(5, 0)))
}

type DeleteTransactionsUUIDRequestQuery struct {
	TimeParam time.Time
}

func (query DeleteTransactionsUUIDRequestQuery) GetTimeParam() time.Time {
	return query.TimeParam
}

func (query DeleteTransactionsUUIDRequestQuery) Validate() error {
	return nil
}

type DeleteTransactionsUUIDRequest struct {
	Header               DeleteTransactionsUUIDRequestHeader
	Path                 DeleteTransactionsUUIDRequestPath
	Query                DeleteTransactionsUUIDRequestQuery
	ProcessingResult     RequestProcessingResult
	SecurityCheckResults map[SecurityScheme]string
}

type SecurityScheme string

const (
	SecuritySchemeBasic  SecurityScheme = "Basic"
	SecuritySchemeBearer SecurityScheme = "Bearer"
	SecuritySchemeCookie SecurityScheme = "Cookie"
)

type securityProcessor struct {
	scheme  SecurityScheme
	extract func(r *http.Request) (string, string, bool)
	handle  func(r *http.Request, scheme SecurityScheme, name string, value string) error
}

var securityExtractorsFuncs = map[SecurityScheme]func(r *http.Request) (string, string, bool){
	SecuritySchemeBasic: func(r *http.Request) (string, string, bool) {
		value := r.Header.Get("Authorization")

		if !strings.HasPrefix(value, "Bearer ") {
			return "", "", false
		}

		value = value[7:]

		return "Authorization", value, value != ""
	},
	SecuritySchemeBearer: func(r *http.Request) (string, string, bool) {
		value := r.Header.Get("Authorization")

		return "Authorization", value, value != ""
	},
	SecuritySchemeCookie: func(r *http.Request) (string, string, bool) {
		cookie, err := r.Cookie("JSESSIONID")

		if err != nil {
			return "", "", false
		}

		return cookie.Name, cookie.Value, true
	},
}

type SecuritySchemas interface {
	SecuritySchemeBasic(r *http.Request, scheme SecurityScheme, name string, value string) error
	SecuritySchemeBearer(r *http.Request, scheme SecurityScheme, name string, value string) error
	SecuritySchemeCookie(r *http.Request, scheme SecurityScheme, name string, value string) error
}

type SecurityCheckResult struct {
	Scheme SecurityScheme
	Value  string
}
