# github.com/Mega0hm/ipcalcs
   
ipcalcs is a utility package of Go functions that help convert wacky IPv4/v6 representations, to enable things like doing math on IPs and shit. Helps a lot with stratified random sampling method for net discovery in netBang.

## IPv4StoH() 
* STRING to HEX, since raw hex is sometimes the easiest to wrangle.   
   
## IpStoI() 
* STRING to net.IP   

## CIDRRead()
* Calls net.ParseCIDR to return net.IP, net.IPNet, but also also returns the netmask bitcount as an unsigned int.
  
## StringToUint32() 
* STRING to UINT32. Made to feed IpStoH creation of usable 32-bit hexvals from IPv4 dotted-decimal strings. Useful however for converting any string.
    
**AUTHOR: Chuck Geigner a.k.a. "mongoose", a.k.a. "chux0r"**  
**DATE:   28SEP2023**  
   
*Copyright © 2023 CT Geigner, All rights reserved*
*Free to use under GNU GPL v2, see https://github/Mega0hm/ipcalcs/LICENSE.md*
