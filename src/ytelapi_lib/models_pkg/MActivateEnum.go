/*
 * ytelapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */
package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for MActivateEnum enum
 */
type MActivateEnum int

/**
 * Value collection for MActivateEnum enum
 */
const (
    MActivate_ACTIVATE MActivateEnum = 1
    MActivate_DEACTIVATE = 0
)

func (r MActivateEnum) MarshalJSON() ([]byte, error) { 
    s := MActivateEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *MActivateEnum) UnmarshalJSON(data []byte) error { 
    var s int 
    json.Unmarshal(data, &s)
    v :=  MActivateEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts MActivateEnum to its int representation
 */
func MActivateEnumToValue(mActivateEnum MActivateEnum) int {
    switch mActivateEnum {
        case MActivate_ACTIVATE:
    		return 1		
        case MActivate_DEACTIVATE:
    		return 0		
        default:
        	return 1
    }
}

/**
 * Converts MActivateEnum Array to its string Array representation
*/
func MActivateEnumArrayToValue(mActivateEnum []MActivateEnum) []int {
    convArray := make([]int,len( mActivateEnum))
    for i:=0; i<len(mActivateEnum);i++ {
        convArray[i] = MActivateEnumToValue(mActivateEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func MActivateEnumFromValue(value int) MActivateEnum {
    switch value {
        case 1:
            return MActivate_ACTIVATE
        case 0:
            return MActivate_DEACTIVATE
        default:
            return MActivate_ACTIVATE
    }
}