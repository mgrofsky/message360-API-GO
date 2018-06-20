/*
 * ytelapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */

package letter_pkg

import "ytelapi_lib/configuration_pkg"

/*
 * Interface for the LETTER_IMPL
 */
type LETTER interface {
    CreateLetterDelete (string) (string, error)

    CreateLetterViewletter (string) (string, error)

    CreateLetterListsletter (*CreateLetterListsletterInput) (string, error)

    CreateLetterCreate (*CreateLetterCreateInput) (string, error)
}

/*
 * Factory for the LETTER interaface returning LETTER_IMPL
 */
func NewLETTER(config configuration_pkg.CONFIGURATION) *LETTER_IMPL {
    client := new(LETTER_IMPL)
    client.config = config
    return client
}