Start

root@kali:/home/kali# systemctl status postgresql
● postgresql.service - PostgreSQL RDBMS
     Loaded: loaded (/lib/systemd/system/postgresql.service; enabled; vendor preset:>
     Active: active (exited) since Sun 2020-06-14 15:13:22 EDT; 1min 0s ago
   Main PID: 1763 (code=exited, status=0/SUCCESS)
      Tasks: 0 (limit: 2317)
     Memory: 0B
     CGroup: /system.slice/postgresql.service

Jun 14 15:13:22 kali systemd[1]: Starting PostgreSQL RDBMS...
Jun 14 15:13:22 kali systemd[1]: Finished PostgreSQL RDBMS.


msfdb init


       =[ metasploit v5.0.92-dev                          ]
+ -- --=[ 2026 exploits - 1102 auxiliary - 343 post       ]
+ -- --=[ 562 payloads - 45 encoders - 10 nops            ]
+ -- --=[ 7 evasion                                       ]


 49  python3 rsf.py
   50  systemctl start postgresql
   51  systemctl status postgresql
   52  systemctl enable  postgresql
   53  systemctl start postgresql
   54  systemctl status postgresql
   55  msfdb init
   56  apt-update install
   57  apt-get install updated
   58  apt-get install update
   59  apt-get  update
   60  msfdb init
   61  apt update; apt install metasploit-framework



gem update --system