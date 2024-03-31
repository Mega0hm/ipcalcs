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

/*
IPv4StoH()
Convert IPv4 string to hexvalue 0x0 - 0xFFFFFFFF

Call with input string, or net.IP.String()
*/
func IPv4StoH(ip string) (uint32, error) {
	/* un-string-ify all parts of this mess.
	   Then assemble as a hexval.
	   And who tf thought dots and shit was a good idea? Made ipv4 some mixed-up bullshit.
	   Anyway, undo that. */
	ipparts := strings.Split(ip, ".")
	var result uint32 = 0
	var bs int = 0
	var j = 0 // uh, dude... .sdrawkcab s'tihS LOL - no, really. strings.Split arranges our IP elements
	          // backwards. Great. So I need the j counter to run the list in fucking retrograde. Yahp.
	for i := 3; i >= 0; i-- {
		quad, ok := StringToUint32(ipparts[i])
		if !ok || quad > 255 || len(ipparts) != 4 { // FAIL on: Nan, invalid num, too many/few parts
			err := fmt.Errorf("%s is not a valid IPv4 address", ip)
			return 0x0, err
		}
		bs = j * 8
		j++
		result = result + quad<<bs // bitshift it left and add to result
	}
	return result, nil  // range: 0-4294967295
}

/* Convert IPv4 string to IPv6  (type net.IP, as defined in net.IP) */
func IpStoI(ip string) net.IP {
	return net.ParseIP(ip)
}

/******************************************************************************
CIDRRead()
A small addon to net.ParseCIDR that also returns the mask as an unsigned int
We're relying on net.ParseCIDR to catch anything hinky before we split the input
and convert the mask string to a real number.
******************************************************************************/ 
func CIDRRead(s string) (net.IP, *net.IPNet, uint32) {
	ip, ipn, err := net.ParseCIDR(s)
	if err != nil {
		fmt.Println(err.Error())
		return net.IP{0,0,0,0}, nil, 0
	} 
	ms := strings.Split(s, "/")
	mask, _ := StringToUint32(ms[0])
	return ip, ipn, mask

}

func StringToUint32(s string) (uint32, bool) {
	var isnum = false // Validate: "Is the input string a valid numeric representation?"
	var num uint32 = 0
	slen := len([]rune(s))
	for i := slen - 1; i >= 0; i-- { // Walk the chars of the input string, validate each.
		char := s[i]
		if (int32(char)-48 < 0) || (int32(char)-48) > 9 { // When char->num conversion shows the val is not 0-9 (NaN)...
			isnum = false // ...set failure and...
			break         // ...break outta here.
		} else { // When char is 0-9, go convert it.
			isnum = true
			var mult int32 = 1 
			if char != '0' { // For non-zero numbers represented:
				digit := slen - 1 - i        // ...work through chars L->R... (len-1-i is most signif. digit L->R)
				for j := digit; j > 0; j-- { // ...first note correct pow10 multiplier for the position of the number...
					mult = mult * 10 // [NOTE ON WHY NOT math.Pow10: That uses float64- this is, I think, computationally less expensive.]
				}
				num = num + uint32((int32(char)-48)*mult) // ...then solve for current column and update total.
			}
		}
	}
	if !isnum {
		return 0x00, false // NaN:  Return IPADDR "0" and failure.
	} else {
		return num, true // GOOD: Return IPADDR and success
	}
}