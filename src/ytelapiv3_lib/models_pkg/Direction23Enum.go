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
 * Type definition for Direction23Enum enum
 */
type Direction23Enum int

/**
 * Value collection for Direction23Enum enum
 */
const (
    Direction23_IN Direction23Enum = 1 + iota
    Direction23_OUT
    Direction23_BOTH
)

func (r Direction23Enum) MarshalJSON() ([]byte, error) { 
    s := Direction23EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Direction23Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Direction23EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Direction23Enum to its string representation
 */
func Direction23EnumToValue(direction23Enum Direction23Enum) string {
    switch direction23Enum {
        case Direction23_IN:
    		return "IN"		
        case Direction23_OUT:
    		return "OUT"		
        case Direction23_BOTH:
    		return "BOTH"		
        default:
        	return "IN"
    }
}

/**
 * Converts Direction23Enum Array to its string Array representation
*/
func Direction23EnumArrayToValue(direction23Enum []Direction23Enum) []string {
    convArray := make([]string,len( direction23Enum))
    for i:=0; i<len(direction23Enum);i++ {
        convArray[i] = Direction23EnumToValue(direction23Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Direction23EnumFromValue(value string) Direction23Enum {
    switch value {
        case "IN":
            return Direction23_IN
        case "OUT":
            return Direction23_OUT
        case "BOTH":
            return Direction23_BOTH
        default:
            return Direction23_IN
    }
}
