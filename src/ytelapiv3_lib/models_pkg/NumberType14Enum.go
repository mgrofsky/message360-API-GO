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
 * Type definition for NumberType14Enum enum
 */
type NumberType14Enum int

/**
 * Value collection for NumberType14Enum enum
 */
const (
    NumberType14_ALL NumberType14Enum = 1 + iota
    NumberType14_VOICE
    NumberType14_SMS
)

func (r NumberType14Enum) MarshalJSON() ([]byte, error) { 
    s := NumberType14EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *NumberType14Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  NumberType14EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts NumberType14Enum to its string representation
 */
func NumberType14EnumToValue(numberType14Enum NumberType14Enum) string {
    switch numberType14Enum {
        case NumberType14_ALL:
    		return "all"		
        case NumberType14_VOICE:
    		return "voice"		
        case NumberType14_SMS:
    		return "sms"		
        default:
        	return "all"
    }
}

/**
 * Converts NumberType14Enum Array to its string Array representation
*/
func NumberType14EnumArrayToValue(numberType14Enum []NumberType14Enum) []string {
    convArray := make([]string,len( numberType14Enum))
    for i:=0; i<len(numberType14Enum);i++ {
        convArray[i] = NumberType14EnumToValue(numberType14Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func NumberType14EnumFromValue(value string) NumberType14Enum {
    switch value {
        case "all":
            return NumberType14_ALL
        case "voice":
            return NumberType14_VOICE
        case "sms":
            return NumberType14_SMS
        default:
            return NumberType14_ALL
    }
}
