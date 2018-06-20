/*
 * ytelapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */

package sharedshortcode_pkg


import(
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/apimatic/unirest-go"
	"ytelapi_lib/apihelper_pkg"
	"ytelapi_lib/configuration_pkg"
)

/*
 * Input structure for the method CreateShortcodeUpdateshortcode
 */
type CreateShortcodeUpdateshortcodeInput struct {
    Shortcode         string          //List of valid shortcode to your Ytel account
    FriendlyName      *string         //User generated name of the shortcode
    CallbackUrl       *string         //URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the call is finished.
    CallbackMethod    *string         //Specifies the HTTP method used to request the required StatusCallBackUrl once call connects.
    FallbackUrl       *string         //URL used if any errors occur during execution of InboundXML or at initial request of the required Url provided with the POST.
    FallbackUrlMethod *string         //Specifies the HTTP method used to request the required FallbackUrl once call connects.
}

/*
 * Input structure for the method CreateShortcodeListshortcode
 */
type CreateShortcodeListshortcodeInput struct {
    Page            *int64          //The page count to retrieve from the total results in the collection. Page indexing starts at 1.
    Pagesize        *int64          //Number of individual resources listed in the response per page
    Shortcode       *string         //Only list keywords of shortcode
}

/*
 * Input structure for the method CreateKeywordLists
 */
type CreateKeywordListsInput struct {
    Page            *int64          //The page count to retrieve from the total results in the collection. Page indexing starts at 1.
    Pagesize        *int64          //Number of individual resources listed in the response per page
    Keyword         *string         //Only list keywords of keyword
    Shortcode       *int64          //Only list keywords of shortcode
}

/*
 * Input structure for the method CreateTemplateLists
 */
type CreateTemplateListsInput struct {
    Type            *string         //The type (category) of template Valid values: marketing, authorization
    Page            *int64          //The page count to retrieve from the total results in the collection. Page indexing starts at 1.
    Pagesize        *int64          //The count of objects to return per page.
    Shortcode       *string         //Only list templates of type
}

/*
 * Input structure for the method CreateShortcodeSendsms
 */
type CreateShortcodeSendsmsInput struct {
    Shortcode             string          //The Short Code number that is the sender of this message
    To                    string          //A valid 10-digit number that should receive the message
    Templateid            uuid.UUID       //The unique identifier for the template used for the message
    Data                  string          //format of your data, example: {companyname}:test,{otpcode}:1234
    Method                *string         //Specifies the HTTP method used to request the required URL once the Short Code message is sent.
    MessageStatusCallback *string         //URL that can be requested to receive notification when Short Code message was sent.
}

/*
 * Input structure for the method CreateShortcodeGetinboundsms
 */
type CreateShortcodeGetinboundsmsInput struct {
    Page            *int64          //The page count to retrieve from the total results in the collection. Page indexing starts at 1.
    Pagesize        *int64          //Number of individual resources listed in the response per page
    From            *string         //From Number to Inbound ShortCode
    Shortcode       *string         //Only list messages sent to this Short Code
    Datecreated     *string         //Only list messages sent with the specified date
}

/*
 * Client structure as interface implementation
 */
type SHAREDSHORTCODE_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * The response returned here contains all resource properties associated with the given Shortcode.
 * @param    string        shortcode     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) CreateShortcodeViewshortcode (
            shortcode string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/shortcode/viewshortcode.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Shortcode" : shortcode,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * View a set of properties for a single keyword.
 * @param    string        keywordid     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) CreateKeywordView (
            keywordid string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/keyword/view.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Keywordid" : keywordid,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Update Assignment
 * @param  CreateShortcodeUpdateshortcodeInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) CreateShortcodeUpdateshortcode (input *CreateShortcodeUpdateshortcodeInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/shortcode/updateshortcode.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Shortcode" : input.Shortcode,
        "FriendlyName" : input.FriendlyName,
        "CallbackUrl" : input.CallbackUrl,
        "CallbackMethod" : input.CallbackMethod,
        "FallbackUrl" : input.FallbackUrl,
        "FallbackUrlMethod" : input.FallbackUrlMethod,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * View a Shared ShortCode Template
 * @param    uuid.UUID        templateId     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) CreateTemplateView (
            templateId uuid.UUID) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/template/view.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "TemplateId" : templateId,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Retrieve a list of shortcode assignment associated with your Ytel account.
 * @param  CreateShortcodeListshortcodeInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) CreateShortcodeListshortcode (input *CreateShortcodeListshortcodeInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/shortcode/listshortcode.json"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "Shortcode" : input.Shortcode,
    })
    if err != nil {
        //error in query param handling
        return "", err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "page" : apihelper_pkg.ToString(*input.Page, "1"),
        "pagesize" : apihelper_pkg.ToString(*input.Pagesize, "10"),

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Retrieve a list of keywords associated with your Ytel account.
 * @param  CreateKeywordListsInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) CreateKeywordLists (input *CreateKeywordListsInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/keyword/lists.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "page" : apihelper_pkg.ToString(*input.Page, "1"),
        "pagesize" : apihelper_pkg.ToString(*input.Pagesize, "10"),
        "Keyword" : input.Keyword,
        "Shortcode" : input.Shortcode,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * List Shortcode Templates by Type
 * @param  CreateTemplateListsInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) CreateTemplateLists (input *CreateTemplateListsInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/template/lists.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "type" : apihelper_pkg.ToString(*input.Type, "authorization"),
        "page" : apihelper_pkg.ToString(*input.Page, "1"),
        "pagesize" : apihelper_pkg.ToString(*input.Pagesize, "10"),
        "Shortcode" : input.Shortcode,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Send an SMS from a Ytel ShortCode
 * @param  CreateShortcodeSendsmsInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) CreateShortcodeSendsms (input *CreateShortcodeSendsmsInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/shortcode/sendsms.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "shortcode" : input.Shortcode,
        "to" : input.To,
        "templateid" : input.Templateid,
        "data" : input.Data,
        "Method" : apihelper_pkg.ToString(*input.Method, "GET"),
        "MessageStatusCallback" : input.MessageStatusCallback,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * List All Inbound ShortCode
 * @param  CreateShortcodeGetinboundsmsInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) CreateShortcodeGetinboundsms (input *CreateShortcodeGetinboundsmsInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/shortcode/getinboundsms.json"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "Datecreated" : input.Datecreated,
    })
    if err != nil {
        //error in query param handling
        return "", err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "page" : apihelper_pkg.ToString(*input.Page, "1"),
        "pagesize" : apihelper_pkg.ToString(*input.Pagesize, "10"),
        "from" : input.From,
        "Shortcode" : input.Shortcode,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}
