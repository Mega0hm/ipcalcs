# github.com/Mega0hm/ipcalcs
   
ipcalcs is a network utility package of Go functions that help convert wacky ipv4 representations, to enable things like doing math on IPs and shit. Helps a lot with stratified random sampling method for net discovery we do in NetscanX.

## IpStoH() 
* STRING to HEX, since raw hex would be the easiest to wrangle.   
   
## IpStoI() 
* STRING to net.IP   
  
## StringToUint32() 
* STRING to UINT32. Made to feed IpStoH creation of usable 32-bit hexvals from IPv4 dotted-decimal strings. Useful however for converting any string.
    
**AUTHOR: Chuck Geigner a.k.a. "mongoose", a.k.a. "chux0r"**  
**DATE:   28SEP2023**  
   
*Copyright © 2023 CT Geigner, All rights reserved*
*Free to use under GNU GPL v2, see https://github/Mega0hm/ipcalcs/LICENSE.md*
   
Written for use with Netscanx, but possibly plenty useful for other stuff.
   
*--ctg 26NOV2023*
