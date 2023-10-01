package ipcalcs

/******************************************************************************
* ipcalcs
*
* Functions that help convert wacky ipv4 representations, to enable things like
* doing math on IPs and shit. Helps a lot with stratified random sampling
* method for net discovery we do in NetscanX.
* Key:
* IpStoH(): STRING to HEX, since raw hex would be the easiest to wrangle.
* IpStoI(): STRING to net.IP
* StringToUint32(): STRING to UINT32. Made to feed IpStoH creation of usable
* 	32-bit hexvals from IPv4 dotted-decimal strings. Useful however for
* 	converting any string.
*
* AUTHOR: CT Geigner (chux0r)
* 28SEPT2023
*
* *Viva la resistance!*
******************************************************************************/
import (
	"fmt"
	"net"
	"strings"
)

/* IP string-to-hexvalue */
func IpStoH(ip string) uint32 {
	/* un-string-ify all parts of this mess.
	   Then assemble as a hexval.
	   And who tf thought dots and shit was a good idea? Made ipv4 some mixed-up bullshit.
	   Anyway, undo that. */
	ipparts := strings.Split(ip, ".")
	var result uint32 = 0
	var bs uint32 = 0
	j := 0 //uh, dude... .sdrawkcab s'tihS LOL - no, really. strings.Split arranges our IP elements
	// backwards. Great. So I need the j counter to run the list in fucking retrograde.
	for i := 3; i >= 0; i-- {
		//fmt.Printf("Iteration %d: Value: %s", i, ipparts[i]) //DEBUG
		quad, ok := StringToUint32(ipparts[i])
		if !ok || quad > 255 || len(ipparts) != 4 { // FAIL on: Nan, invalid num, too many/few parts
			fmt.Printf("\nERROR: %s is not a valid IP.\n", ip)
			return 0x0
		}
		bs = uint32(j) * 8
		j++
		result = result + quad<<bs // bitshift it left and add to result
		//fmt.Printf("Bit shift is %d,\tConverterd Number is: %X\t Result is 0x%x\n", bs, quad<<bs, result) //DEBUG
	}
	return result
}

/* Convert IPv4 string to IPv6  (type net.IP, as defined in net.IP) */
func IpStoI(ip string) net.IP {
	return net.ParseIP(ip)
}

func StringToUint32(s string) (uint32, bool) {
	// fmt.Printf("\nDEBUG (stringToUint32) :: INPUT Item: %s", s) // DEBUG
	isnum := false // Validate: "Is the input string a valid numeric representation?"
	var num uint32 = 0
	slen := len([]rune(s))
	for i := slen - 1; i >= 0; i-- { // Walk the chars of the input string, validate each.
		char := s[i]
		if (int32(char)-48 < 0) || (int32(char)-48) > 9 { // When char->num conversion shows the val is not 0-9 (NaN)...
			isnum = false // ...set failure and...
			break         // ...break outta here.
		} else { // When char is 0-9, go convert it.
			isnum = true
			var mlt int32 = 1
			// fmt.Printf("\nDEBUG (stringToUint32) ::\n\tNum char: %c", char) // DEBUG
			if char != '0' { // For non-zero numbers represented:
				digit := slen - 1 - i        // ...work through chars L->R... (len-1-i is most signif. digit L->R)
				for j := digit; j > 0; j-- { // ...first note correct pow10 multiplier for the position of the number...
					mlt = mlt * 10 // [NOTE ON WHY: This is computationally less expensive than using math.Pow10 and float64s, that's why.]
				}
				// fmt.Printf(" multiplied by %d (x10^%d)", mlt, digit) // DEBUG
				num = num + uint32((int32(char)-48)*mlt) // ...then solve for current column and update total.
				// fmt.Printf(", totaling %d\n", num)        // DEBUG
			}
		}
	}
	if isnum == false {
		return 0x00, false // NaN:  Return IPADDR "0" and failure.
	} else {
		return num, true // GOOD: Return IPADDR and success
	}
}
