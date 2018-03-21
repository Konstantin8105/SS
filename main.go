package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	// nginx
	_ "github.com/Konstantin8105/ss/nginx"

	// htop
	_ "github.com/Konstantin8105/ss/htop"

	// vim
	_ "github.com/Konstantin8105/ss/vim"

	// mc
	_ "github.com/Konstantin8105/ss/mc"

	// nano
	_ "github.com/Konstantin8105/ss/nano"

	// ssh
	_ "github.com/Konstantin8105/ss/ssh"

	// backup
	// systemd
	// git server
	// git web
	// localhost
	// router settings
	// system update

	// base `starter` package
	"github.com/Konstantin8105/ss/starter"
)

var (
	helpFlag    = flag.Bool("h", false, "give this help list")
	listFlag    = flag.Bool("l", false, "show list of modules")
	installFlag = flag.Bool("i", false, "install settings")
	prefixFlag  = flag.String("prefix", "", "prefix before each command."+
		" Typically used :\"sudo\" or \"ssh tom@localhost sudo\" or ...")
)

/*
Notes:
* https://blog.golang.org/docker
* https://stackoverflow.com/questions/26411594/executing-docker-command-using-golang-exec-fails
* https://github.com/betweenbrain/ubuntu-web-server-build-script
* https://medium.com/statuscode/golang-docker-for-development-and-production-ce3ad4e69673
*/
// TODO: add logs checking /var/log/
// TODO: add database
/*

# Minimal configuration of commands:

```minimal command
$ iw dev
	Interface wlan0          <-- remember interface
$ ip link show wlan0         <-- checking
$ sudo ip link set wlan0 up  <-- only if interface is not open
$ ip link show wlan0         <-- checking
$ iw wlan0 link              <-- checking connection
$ sudo iw wlan0 scan
$ sudo -s
$ wpa_passphrase WIFI_NAME >> /etc/wpa_supplicant.conf
Enter WIFI_PASSWORD
$ exit
$ sudo vim /etc/network/interfaces
```

Inside `interfaces`:

```
auto wls1
#iface wls1 inet dhcp
iface wls1 inet static
	address 192.168.0.55
	netmask 255.255.255.0
	gateway 192.168.0.1
wpa-conf /etc/wpa_supplicant.conf
```

Install video driver:

```
sudo apt install ubuntu-drivers-common
sudo apt-get install intel-microcode

sudo ubuntu-drivers devices
sudo ubuntu-drivers autoinstall
```

```
$ sudo ufw allow 222
$ sudo ufw enable
```




## How to connect to a WPA/WPA2 WiFi network using Linux command line


1.  Find out the wireless device name.
```
$ /sbin/iw dev
phy#0
	Interface wlan0
		ifindex 3
		type managed
```
The above output showed that the system has 1 physical WiFi card, designated as phy#0. The device name is wlan0. The type specifies the operation mode of the wireless device. managed means the device is a WiFi station or client that connects to an access point.

2. Check that the wireless device is up.
```
$ ip link show wlan0
3: wlan0: (BROADCAST,MULTICAST) mtu 1500 qdisc noop state DOWN mode DEFAULT qlen 1000
    link/ether 74:e5:43:a1:ce:65 brd ff:ff:ff:ff:ff:ff
```
Look for the word **"UP"** inside the brackets in the first line of the output.

In the above example, wlan0 is not UP. Execute the following command to bring it up:
```
$ sudo ip link set wlan0 up
[sudo] password for peter:
Note: you need root privilege for the above operation.
```
If you run the show link command again, you can tell that wlan0 is now UP.
```
$ ip link show wlan0
3: wlan0: (NO-CARRIER,BROADCAST,MULTICAST,UP) mtu 1500 qdisc mq state DOWN mode DEFAULT qlen 1000
    link/ether 74:e5:43:a1:ce:65 brd ff:ff:ff:ff:ff:ff
```
3. Check the connection status.
```
$ /sbin/iw wlan0 link
Not connected.
```
The above output shows that you are not connected to any network.

4. Scan to find out what WiFi network(s) are detected
```
$ sudo /sbin/iw wlan0 scan
BSS 00:14:d1:9c:1f:c8 (on wlan0)
        ... sniped ...
	freq: 2412
	SSID: gorilla
	RSN:	 * Version: 1
		 * Group cipher: CCMP
		 * Pairwise ciphers: CCMP
		 * Authentication suites: PSK
		 * Capabilities: (0x0000)
        ... sniped ...
```
The 2 important pieces of information from the above are the SSID and the
security protocol (WPA/WPA2 vs WEP). The SSID from the above example is
gorilla. The security protocol is RSN, also commonly referred to as WPA2.
The security protocol is important because it determines what tool you use
to connect to the network.

5. Connect to WPA/WPA2 WiFi network.

This is a 2 step process. First, you generate a configuration file for wpa_supplicant that contains the pre-shared key ("passphrase") for the WiFi network.
```
$ sudo -s
[sudo] password for peter:
$ wpa_passphrase gorilla >> /etc/wpa_supplicant.conf
...type in the passphrase and hit enter...
wpa_passphrase takes the SSID as the single argument. You must type in the passphrase for the WiFi network gorilla after you run the command. Using that information, wpa_passphrase will output the necessary configuration statements to the standard output. Those statements are appended to the wpa_supplicant configuration file located at /etc/wpa_supplicant.conf.
```
Note: you need root privilege to write to /etc/wpa_supplicant.conf.
```
$ cat /etc/wpa_supplicant.conf
# reading passphrase from stdin
network={
	ssid="gorilla"
	#psk="testtest"
	psk=4dfe1c985520d26a13e932bf0acb1d4580461dd854ed79ad1a88ec221a802061
}
The second step is to run wpa_supplicant with the new configuration file.

$ sudo wpa_supplicant -B -D wext -i wlan0 -c /etc/wpa_supplicant.conf
```
-B means run wpa_supplicant in the background.
-D specifies the wireless driver. wext is the generic driver.
-c specifies the path for the configuration file.

Use the iw command to verify that you are indeed connected to the SSID.
```
$ /sbin/iw wlan0 link
Connected to 00:14:d1:9c:1f:c8 (on wlan0)
	SSID: gorilla
	freq: 2412
	RX: 63825 bytes (471 packets)
	TX: 1344 bytes (12 packets)
	signal: -27 dBm
	tx bitrate: 6.5 MBit/s MCS 0

	bss flags:	short-slot-time
	dtim period:	0
	beacon int:	100
```

6. Obtain IP address by DHCP
```
$ sudo dhclient wlan0
Use the ip command to verify the IP address assigned by DHCP. The IP address is 192.168.1.113 from below.

$ ip addr show wlan0
3: wlan0:  mtu 1500 qdisc mq state UP qlen 1000
    link/ether 74:e5:43:a1:ce:65 brd ff:ff:ff:ff:ff:ff
    inet 192.168.1.113/24 brd 192.168.1.255 scope global wlan0
    inet6 fe80::76e5:43ff:fea1:ce65/64 scope link
       valid_lft forever preferred_lft forever
```

7. Add default routing rule.
The last configuration step is to make sure that you have the proper routing rules.
```
$ ip route show
192.168.1.0/24 dev wlan0  proto kernel  scope link  src 192.168.1.113
The above routing table contains only 1 rule which redirects all traffic
destined for the local subnet (192.168.1.x) to the wlan0 interface.
You may want to add a default routing rule to pass all other traffic
through wlan0 as well.

$ sudo ip route add default via 192.168.1.254 dev wlan0
$ ip route show
default via 192.168.1.254 dev wlan0
192.168.1.0/24 dev wlan0  proto kernel  scope link  src 192.168.1.113
```

8. Ping external ip address to test connectivity
```
$ ping 8.8.8.8
PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
64 bytes from 8.8.8.8: icmp_req=1 ttl=48 time=135 ms
64 bytes from 8.8.8.8: icmp_req=2 ttl=48 time=135 ms
64 bytes from 8.8.8.8: icmp_req=3 ttl=48 time=134 ms
^C
--- 8.8.8.8 ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 2000ms
rtt min/avg/max/mdev = 134.575/134.972/135.241/0.414 ms
```


## Create internet connection

Add in file `/etc/network/interfaces`:
```
# WiFi connection
auto wls1
iface wls1 inet dhcp
wpa-conf /etc/wpa_supplicant.conf
```

## How to install NVidia driver

К сожалению пробую сначало установить сначало MATE
```
sudo apt install ubuntu-drivers-common
sudo apt-get install intel-microcode

sudo ubuntu-drivers devices
sudo ubuntu-drivers autoinstall
```
Screen Blanks/Monitor Turns Off
Using a laptop with a GeForce Go card, or connecting the sole display via DVI on a dual-head system sometimes results in the screen not receiving a picture. This is caused by the driver outputting video to the VGA port on the graphics card, instead of DVI.
The usual hint that you have this problem is when you hear the startup sound but nothing appears on the screen. If you do not hear any sound, you are more than likely experiencing unrelated problems.

This is a bug about displays on digital outputs being blank when using NVIDIA driver, and can be resolved by editing your /etc/X11/xorg.conf file:

1. Switch to the console by using ctrl+alt+F1, or reboot and select recovery mode from the GRUB menu.
```
mount -o rw,remount /
```
2. Open and edit xorg.conf like this:
```
sudo nano /etc/X11/xorg.conf.
```
3. Find the line that says: `Section "Screen"`
4. Insert a new line that says `Option "UseDisplayDevice" "DFP"`. in intel and nvidia
5. Save the file. If you had to restart into recovery mode, type reboot, otherwise restart your display using `sudo /etc/init.d/gdm restart`.

Change Inactive and Devise 0 on file /etc/X11/xorg.conf

```
$ sudo ufw allow 222
$ sudo ufw enable
```



To disable entering the sleep mode I had to edit the /etc/systemd/logind.conf file and modify the line:
`#HandleLidSwitch=suspend` to `HandleLidSwitch=ignore` . Then do
```
sudo service systemd-logind restart
```


*/
func main() {
	flag.Parse()
	err := run()
	if err != nil {
		fmt.Printf("Error = %v", err)
		os.Exit(1)
	}
}

var output io.Writer = os.Stdout

func run() (err error) {

	if len(*prefixFlag) != 0 {
		starter.SetCommandPrefix(*prefixFlag)
	}

	switch {
	case *listFlag:
		// list of modules
		fmt.Fprintf(output, "List of starters :\n")
		list := starter.List()
		var inx int
		for name := range list {
			inx++
			fmt.Fprintf(output, "%2d%20s\n", inx, name)
		}
		fmt.Fprintf(output, "Amount of starters : %2d\n", len(list))

	case *installFlag:
		// set settings
		list := starter.List()
		var inx int
		for name, s := range list {
			inx++
			fmt.Fprintf(output, "%2d%20s\n", inx, name)
			err = s.Run()
			if err != nil {
				return err
			}
		}

	default:
		// help flag
		flag.Usage()
	}
	return nil
}
