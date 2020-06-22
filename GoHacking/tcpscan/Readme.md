#netmap
Setting up Kali
Installing Tools
 Note: had to re-run tools and choose to overwrite instead of the default. 
 Also move the virtual maching into /home/username/VMware

nmap.org

Nmap done: 1 IP address (1 host up) scanned in 48.31 seconds
kali@kali:/mnt/hgfs/Training/EthicalGo/GoHacking/tcpscan$ nmap 192.168.221.1
Starting Nmap 7.80 ( https://nmap.org ) at 2020-06-21 07:49 EDT
Nmap scan report for 192.168.221.1
Host is up (0.0043s latency).
Not shown: 984 closed ports
PORT      STATE    SERVICE
22/tcp    open     ssh
80/tcp    open     http
88/tcp    open     kerberos-sec
515/tcp   filtered printer
1070/tcp  filtered gmrupdateserv
1107/tcp  filtered isoipsigport-2
1113/tcp  filtered ltp-deepspace
2968/tcp  filtered enpp
3306/tcp  open     mysql
5432/tcp  open     postgresql
5440/tcp  filtered unknown
5900/tcp  open     vnc
9010/tcp  filtered sdr
9102/tcp  filtered jetdirect
14441/tcp filtered unknown
24800/tcp filtered unknown



nmap www.google.com

nmap -p 1-100 192.168.0.11

nmap -A 192.168.221.133

