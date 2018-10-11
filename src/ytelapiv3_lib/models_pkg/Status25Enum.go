/*
 * ytelapiv3_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for Status25Enum enum
 */
type Status25Enum int

/**
 * Value collection for Status25Enum enum
 */
const (
    Status25_INPROGRESS Status25Enum = 1 + iota
    Status25_SUCCESS
    Status25_FAILURE
)

func (r Status25Enum) MarshalJSON() ([]byte, error) { 
    s := Status25EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Status25Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Status25EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Status25Enum to its string representation
 */
func Status25EnumToValue(status25Enum Status25Enum) string {
    switch status25Enum {
        case Status25_INPROGRESS:
    		return "INPROGRESS"		
        case Status25_SUCCESS:
    		return "Success"		
        case Status25_FAILURE:
    		return "Failure"		
        default:
        	return "INPROGRESS"
    }
}

/**
 * Converts Status25Enum Array to its string Array representation
*/
func Status25EnumArrayToValue(status25Enum []Status25Enum) []string {
    convArray := make([]string,len( status25Enum))
    for i:=0; i<len(status25Enum);i++ {
        convArray[i] = Status25EnumToValue(status25Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Status25EnumFromValue(value string) Status25Enum {
    switch value {
        case "INPROGRESS":
            return Status25_INPROGRESS
        case "Success":
            return Status25_SUCCESS
        case "Failure":
            return Status25_FAILURE
        default:
            return Status25_INPROGRESS
    }
}