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
 * Type definition for NumberType15Enum enum
 */
type NumberType15Enum int

/**
 * Value collection for NumberType15Enum enum
 */
const (
    NumberType15_ALL NumberType15Enum = 1 + iota
    NumberType15_VOICE
    NumberType15_SMS
)

func (r NumberType15Enum) MarshalJSON() ([]byte, error) { 
    s := NumberType15EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *NumberType15Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  NumberType15EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts NumberType15Enum to its string representation
 */
func NumberType15EnumToValue(numberType15Enum NumberType15Enum) string {
    switch numberType15Enum {
        case NumberType15_ALL:
    		return "all"		
        case NumberType15_VOICE:
    		return "voice"		
        case NumberType15_SMS:
    		return "sms"		
        default:
        	return "all"
    }
}

/**
 * Converts NumberType15Enum Array to its string Array representation
*/
func NumberType15EnumArrayToValue(numberType15Enum []NumberType15Enum) []string {
    convArray := make([]string,len( numberType15Enum))
    for i:=0; i<len(numberType15Enum);i++ {
        convArray[i] = NumberType15EnumToValue(numberType15Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func NumberType15EnumFromValue(value string) NumberType15Enum {
    switch value {
        case "all":
            return NumberType15_ALL
        case "voice":
            return NumberType15_VOICE
        case "sms":
            return NumberType15_SMS
        default:
            return NumberType15_ALL
    }
}
