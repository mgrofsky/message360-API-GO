/*
 * ytelapiv3_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package YtelAPIV3Client

import(
	"ytelapiv3_lib/configuration_pkg"
	"ytelapiv3_lib/shortcode_pkg"
	"ytelapiv3_lib/areamail_pkg"
	"ytelapiv3_lib/postcard_pkg"
	"ytelapiv3_lib/letter_pkg"
	"ytelapiv3_lib/call_pkg"
	"ytelapiv3_lib/phonenumber_pkg"
	"ytelapiv3_lib/sms_pkg"
	"ytelapiv3_lib/sharedshortcode_pkg"
	"ytelapiv3_lib/conference_pkg"
	"ytelapiv3_lib/carrier_pkg"
	"ytelapiv3_lib/email_pkg"
	"ytelapiv3_lib/account_pkg"
	"ytelapiv3_lib/subaccount_pkg"
	"ytelapiv3_lib/address_pkg"
	"ytelapiv3_lib/recording_pkg"
	"ytelapiv3_lib/transcription_pkg"
	"ytelapiv3_lib/usage_pkg"
)


/*
 * Interface for the YTELAPIV3_IMPL
 */
type YTELAPIV3 interface {
     ShortCode()             shortcode_pkg.SHORTCODE
     AreaMail()              areamail_pkg.AREAMAIL
     PostCard()              postcard_pkg.POSTCARD
     Letter()                letter_pkg.LETTER
     Call()                  call_pkg.CALL
     PhoneNumber()           phonenumber_pkg.PHONENUMBER
     SMS()                   sms_pkg.SMS
     SharedShortCode()       sharedshortcode_pkg.SHAREDSHORTCODE
     Conference()            conference_pkg.CONFERENCE
     Carrier()               carrier_pkg.CARRIER
     Email()                 email_pkg.EMAIL
     Account()               account_pkg.ACCOUNT
     SubAccount()            subaccount_pkg.SUBACCOUNT
     Address()               address_pkg.ADDRESS
     Recording()             recording_pkg.RECORDING
     Transcription()         transcription_pkg.TRANSCRIPTION
     Usage()                 usage_pkg.USAGE
     Configuration()         configuration_pkg.CONFIGURATION
}

/*
 * Factory for the YTELAPIV3 interaface returning YTELAPIV3_IMPL
 */
func NewYTELAPIV3(basicAuthUserName string, basicAuthPassword string) YTELAPIV3 {
    ytelAPIV3Client := new(YTELAPIV3_IMPL)
    ytelAPIV3Client.config = configuration_pkg.NewCONFIGURATION()
    ytelAPIV3Client.config.SetBasicAuthUserName(basicAuthUserName)
    ytelAPIV3Client.config.SetBasicAuthPassword(basicAuthPassword)
    return ytelAPIV3Client
}
