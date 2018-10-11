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
 * Type definition for GroupConfirmFile44Enum enum
 */
type GroupConfirmFile44Enum int

/**
 * Value collection for GroupConfirmFile44Enum enum
 */
const (
    GroupConfirmFile44_MP3 GroupConfirmFile44Enum = 1 + iota
    GroupConfirmFile44_WAV
)

func (r GroupConfirmFile44Enum) MarshalJSON() ([]byte, error) { 
    s := GroupConfirmFile44EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *GroupConfirmFile44Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  GroupConfirmFile44EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts GroupConfirmFile44Enum to its string representation
 */
func GroupConfirmFile44EnumToValue(groupConfirmFile44Enum GroupConfirmFile44Enum) string {
    switch groupConfirmFile44Enum {
        case GroupConfirmFile44_MP3:
    		return "mp3"		
        case GroupConfirmFile44_WAV:
    		return "wav"		
        default:
        	return "mp3"
    }
}

/**
 * Converts GroupConfirmFile44Enum Array to its string Array representation
*/
func GroupConfirmFile44EnumArrayToValue(groupConfirmFile44Enum []GroupConfirmFile44Enum) []string {
    convArray := make([]string,len( groupConfirmFile44Enum))
    for i:=0; i<len(groupConfirmFile44Enum);i++ {
        convArray[i] = GroupConfirmFile44EnumToValue(groupConfirmFile44Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func GroupConfirmFile44EnumFromValue(value string) GroupConfirmFile44Enum {
    switch value {
        case "mp3":
            return GroupConfirmFile44_MP3
        case "wav":
            return GroupConfirmFile44_WAV
        default:
            return GroupConfirmFile44_MP3
    }
}
