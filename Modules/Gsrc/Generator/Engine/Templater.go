/*
┌──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│    ___o .--.               ____                                          ____                                      __                      ___o .--.    │
│   /___| |--|              /\  _`\                                       /\  _`\                                 __/\ \__                  /___| |OO|    │
│  /'   |_|  |_             \ \,\L\_\    ___     __     _ __    __        \ \,\L\_\     __    ___   __  __  _ __ /\_\ \ ,_\  __  __        /'   |_|  |_   │
│       (_    _)             \/_\__ \   /'___\ /'__`\  /\`'__\/'__`\  __o__\/_\__ \   /'__`\ /'___\/\ \/\ \/\`'__\/\ \ \ \/ /\ \/\ \           (_     _)  │
│       | |   \                /\ \L\ \/\ \__//\ \L\.\_\ \ \//\  __/    |    /\ \L\ \/\  __//\ \__/\ \ \_\ \ \ \/ \ \ \ \ \_\ \ \_\ \           | |   \   │
│       | |___/                \ `\____\ \____\ \__/.\_\\ \_\\ \____\  / \   \ `\____\ \____\ \____\\ \____/\ \_\  \ \_\ \__\\/`____ \          | |___/   │
│                               \/_____/\/____/\/__/\/_/ \/_/ \/____/  _______\/_____/\/____/\/____/ \/___/  \/_/   \/_/\/__/ `/___/> \                   │
│                                                                     /\______\                                                  /\___/                   │
│                                                                     \/______/                                                  \/__/                    │
│                                                                                                                                                         │
│           Professional Digital forensics, Network hacking, Stegonography, Recon, OSINT, Bluetooth, CAN and Web Exploitation Secruity Team               │
│																																						  │
│																																						  │
│																																						  │
│																																						  │
│━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━│
│																Package: Engine															             	  │
│																																						  │
│	This file is the main templating engine, the templater that holds diagrams, for loops itterating over the frizz J structure and the loader for the    │
│   HTTP server. This means this file is completely dedicated to parsing and filling the pre generated HTML templates in TemplaterVars.go. The reason that│
│	we generate the data and fill it into currently existing files is just for templating reasons, everytime this code is ran the channels will be forced │
│	to delete the current HTML files and re generate new ones thus this file and module code will be called to re fill those templates up. Of course this │
│	will not happen everytime you make a GET request but everytime the files or data inside of the structure or JSON file is changed via POST.            │
│   Frizz has a built in packet masher and re uploader to re upload and parse new files, once the POST request is made or found with a specific header    │
│	the go code will be ran by the C++ inline which is run by the server. This makes it a bit more performant rather than keeping it with RPCT or re mod- │
│	ing the files with new templates that are constantly generated one by one by the server. Using threads is a much better way to do this same with      │
│	loading the data from the channels and instead of using Golangs templating language to just manually create the bodies with the engine                │
│																																						  │
│━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━│
│																																						  │
│Package status		     -> OK | Working																												  │
│Security status         -> OK | Secure																													  │
│Performance Status      -> OK | Performant 																											  │
│Bug Status              -> OK | NONE																													  │
│Error status            -> OK | NONE																													  │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
*/
package Engine

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Itterator(table bool, tabletype string, Template string, Element ...[]string) string {
	sizeof := len(Element[0]) // First array
	// Range and detect
	for i := 0; i < sizeof; i++ {
		if i > len(Element[1]) || i == len(Element[1]) {
			break
		}
		switch table {
		case true:
			if tabletype == "uagent" {
				Template += "<tr>"
				Template += fmt.Sprintf("<td>%s</td>", Element[0][i])
				Template += fmt.Sprintf("<td>%s</td>", Element[1][i])
				udata := AgentLoader(Element[1][i]) // Load host and agent information
				if udata.OS == "" {
					Template += fmt.Sprintf("<td>%s</td>", "Operating system was not found or supported")
				} else {
					Template += fmt.Sprintf("<td>%s</td>", udata.OS)
				}
				Template += "<tr>"
			}
		}
	}
	Template += Template_EndTags
	Template += Template_CSS
	Template += fmt.Sprintf(TemplatebarvarJS, OSC.Linux, OSC.Windows, OSC.IOS, BC.Apple, BC.Google, BC.XiaoMi, BC.Microsoft)
	Template += TemplatebarJS
	// generate bar values
	if Template != "" {
		return Template
	} else {
		return ""
	}
}

func Read() { // Read go simple
	LoadOntoType()
	var HTMLTemplate string
	HTMLTemplate += Template_Top
	Write("Modules/Server/HTML/Useragents.html", Itterator(true, "uagent", HTMLTemplate, StructureFrizzPointer.Httpd.UagentHostHost, StructureFrizzPointer.Httpd.UagentHostUagent))
	Write("Modules/Server/HTML/JSONDB.html", CalltoStoreServerFile())
	Write("Modules/Server/HTML/ServerRequirements.html", Call_To_Store_PreProc())
	Write("Modules/Server/HTML/AppInfo.html", Call_To_Store_ServerInfo())
	Write("Modules/Server/HTML/ServerInfo.html", StoreServerInformationFileFromDB())
	LoadCredentials("Modules/Server/HTML/AuthIMAP.html", "imap")
	LoadCredentials("Modules/Server/HTML/AuthFTPCreds.html", "ftp")
	LoadCredentials("Modules/Server/HTML/AuthSMTP.html", "smtp")
	LoadCredentials("Modules/Server/HTML/AuthSSHCreds.html", "ssh")
	LoadCredentials("Modules/Server/HTML/AuthDigest.html", "httpdigest")
	LoadCredentials("Modules/Server/HTML/AuthBASIC.html", "httpbasic")
	LoadCredentials("Modules/Server/HTML/AuthNTLM.html", "httpntlm")
	LoadCredentials("Modules/Server/HTML/AuthNegotiation.html", "httpnegotiate")
	LoadCredentials("Modules/Server/HTML/AuthTelnet.html", "telnet")
	LoadSessionTemplates("Modules/Server/HTML/FTP.html", "ftp")
	LoadSessionTemplates("Modules/Server/HTML/SSH.html", "ssh")
	LoadSessionTemplates("Modules/Server/HTML/SMTP.html", "smtp")
	LoadSessionTemplates("Modules/Server/HTML/Telnet.html", "telnet")
	DrawDocumentPOP("Modules/Server/HTML/Cc.html", "cc")
	DrawDocumentPOP("Modules/Server/HTML/From.html", "from")
	DrawDocumentPOP("Modules/Server/HTML/Recv.html", "recv")
	DrawDocumentPOP("Modules/Server/HTML/Convos.html", "conversation")
	DrawDocumentPOP("Modules/Server/HTML/Emails.html", "*em")
	DrawDocumentHTTP("Modules/Server/HTML/URLs.html", "urls")
	DrawDocumentHTTP("Modules/Server/HTML/Hostnames.html", "hosts")
	DrawDocHTTPGENERAL("Modules/Server/HTML/HTTPSESSION.html")
	DrawHome("Modules/Server/HTML/Home.html")
	DrawWifi("Modules/Server/HTML/Wifi.html")
	GenerateARP("Modules/Server/HTML/ARP.html")
	GenSourcePorts("Modules/Server/HTML/OpenPorts.html")
	GenerateEthernet("Modules/Server/HTML/Ethernet.html")
	GenerateDangerouWifi("Modules/Server/HTML/WifiOspf.html") // too lazy to change the name from OSPF to dangerous.html dont ask
	var StaticTmplMasherHTMKL = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Form</title></head><body></body></html><!DOCTYPE html><html lang="en" dir="ltr"><head><meta charset="UTF-8"><link rel="stylesheet" href="style.css"><link href="https://unpkg.com/boxicons@2.1.2/css/boxicons.min.css" rel="stylesheet"><meta name="viewport" content="width=device-width,initial-scale=1"></head><source src="../HTML/LobbyMisc/Lobby_Music" type="audio/mpeg"><source src="../HTML/Future/Future_Lobby" type="audio/mpeg"><body><div class="sidebar"><div class="logo-details"><i class="bx bxs-injection"></i><span class="logo_name">Frizzed</span></div><ul class="nav-links"><li><a href="/" class="active"><i class="bx bx-grid-alt"></i><span class="links_name">Analytics</span></a></li><li><a href="/ParseNew.html"><i class="bx bx-rocket"></i><span class="links_name">Parse New</span></a></li><li><a href="/Useragents.html"><i class="bx bx-user-circle"></i><span class="links_name">HTTP Useragents</span></a></li><li><a href="/Hostnames.html"><i class="bx bxs-ghost"></i><span class="links_name">HTTP Hostnames</span></a></li><li><a href="/URLs.html"><i class="bx bxl-sketch"></i><span class="links_name">HTTP URLs</span></a></li><li><a href="/HTTPSESSION.html"><i class="bx bxs-business"></i><span class="links_name">HTTP General</span></a></li><li><a href="/DNS.html"><i class="bx bx-cabinet"></i><span class="links_name">DNS</span></a></li><li><a href="/OpenPorts.html"><i class="bx bx-fingerprint"></i><span class="links_name">Open Ports</span></a></li><li><a href="/ARP.html"><i class="bx bx-broadcast"></i><span class="links_name">ARP</span></a></li><li><a href="/Ethernet.html"><i class="bx bx-wifi-1"></i><span class="links_name">Ethernet</span></a></li><li><a href="/Servers.html"><i class="bx bx-server"></i><span class="links_name">Servers</span></a></li><li><a href="/Wifi.html"><i class="bx bx-wifi"></i><span class="links_name">Wifi</span></a></li><li><a href="/WifiOspf.html"><i class="bx bx-wifi"></i><span class="links_name">Wifi Warnings</span></a></li><li><a href="/FTP.html"><i class="bx bx-folder"></i><span class="links_name">FTP</span></a></li><li><a href="/SSH.html"><i class="bx bx-terminal"></i><span class="links_name">SSH</span></a></li><li><a href="/SMTP.html"><i class="bx bx-envelope"></i><span class="links_name">SMTP</span></a></li><li><a href="/Telnet.html"><i class="bx bx-desktop"></i><span class="links_name">Telnet</span></a></li><li><a href="/SIP.html"><i class="bx bx-phone-incoming"></i><span class="links_name">SIP Invites</span></a></li><li><a href="/AuthFTPCreds.html"><i class="bx bx-dialpad"></i><span class="links_name">FTP Credentials</span></a></li><li><a href="/AuthSSHCreds.html"><i class="bx bxs-key"></i><span class="links_name">SSH Credentials</span></a></li><li><a href="/AuthIMAP.html"><i class="bx bxs-lock"></i><span class="links_name">IMAP Credentials</span></a></li><li><a href="/AuthDigest.html"><i class="bx bxs-user-pin"></i><span class="links_name">HTTP Digest</span></a></li><li><a href="/AuthNTLM.html"><i class="bx bx-coffee"></i><span class="links_name">HTTP NTLM</span></a></li><li><a href="/AuthBASIC.html"><i class="bx bxs-contact"></i><span class="links_name">HTTP BASIC</span></a></li><li><a href="/AuthNegotiation.html"><i class="bx bx-share-alt"></i><span class="links_name">HTTP Negotiate</span></a></li><li><a href="/AuthSMTP.html"><i class="bx bx-envelope"></i><span class="links_name">SMTP Credentials</span></a></li><li><a href="/Emails.html"><i class="bx bx-shape-triangle"></i><span class="links_name">Found Emails</span></a></li><li><a href="/Cc.html"><i class="bx bxs-chat"></i><span class="links_name">POP3 Cc payload</span></a></li><li><a href="/From.html"><i class="bx bx-comment-dots"></i><span class="links_name">POP3 From payload</span></a></li><li><a href="/Recv.html"><i class="bx bx-mail-send"></i><span class="links_name">POP3 Recv payload</span></a></li><li><a href="/POP3"><i class="bx bx-conversation"></i><span class="links_name">[Beta] Conversation</span></a></li><li><a href="/Masher.html"><i class="bx bx-meteor"></i><span class="links_name">Packet masher</span></a></li><li><a href="/Raw.html"><i class="bx bx-meteor"></i><span class="links_name">Packets RAW</span></a></li><li><a href="/Extractor.html"><i class="bx bxl-google-cloud"></i><span class="links_name">Packet Extractor</span></a></li><li><a href="/ServerRequirements.html"><i class="bx bx-cctv"></i><span class="links_name">Info this server needs</span></a></li><li><a href="/JSONDB.html"><i class="bx bxs-file-json"></i><span class="links_name">JSON Server file</span></a></li><li><a href="/AppInfo.html"><i class="bx bx-landscape"></i><span class="links_name">Application information</span></a></li><li><a href="/ServerInfo.html"><i class="bx bxs-component"></i><span class="links_name">Server information</span></a></li><li><a href="/Documentation.html"><i class="bx bxs-book-content"></i><span class="links_name">Documentation</span></a></li><li><a href="https://discord.gg/5WfgbMdfWp"><i class="bx bxl-discord-alt"></i><span class="links_name">Discord</span></a></li><li><a href="https://account.venmo.com/u/Scare-Security-Development"><i class="bx bxl-venmo"></i><span class="links_name">Donate [Venmo]</span></a></li><li><a href="https://cash.app/$TotallyNotAHaxxer"><i class="bx bx-money"></i><span class="links_name">Donate [CashApp]</span></a></li><li><a href="https://www.medium.com/@Totally_Not_A_Haxxer"><i class="bx bxl-medium-square"></i><span class="links_name">Medium Articles</span></a></li><li><a href="https://www.github.com/ArkAngeL43"><i class="bx bxl-git"></i><span class="links_name">Github</span></a></li><li><a href="https://www.github.com/orgs/Scare-Security"><i class="bx bxl-github"></i><span class="links_name">Github Organization</span></a></li><li><a href="https://www.instagram.com/Totally_Not_A_Haxxer"><i class="bx bxl-instagram"></i><span class="links_name">Instagram</span></a></li><li><a href="https://twitter.com/NotAHaxxor"><i class="bx bxl-twitter"></i><span class="links_name">Twitter</span></a></li></ul></div><section class="home-section"><nav><div class="sidebar-button"><i class="bx bx-menu sidebarBtn"></i><span class="dashboard">Dashboard</span></div></nav><div class="home-content"><form method="POST" enctype="multipart/form-data"><label for="list">List of directories to .pcap files</label><br><input type="text" id="list" name="list"><label for="pcapout">Output PcapFile</label><input type="text" id="pcapout" name="pcapout"><br><button type="submit" value="submit">Submit</button></form></div></section><script>let sidebar=document.querySelector(".sidebar");let sidebarBtn=document.querySelector(".sidebarBtn");sidebarBtn.onclick=function(){sidebar.classList.toggle("active");if(sidebar.classList.contains("active")){sidebarBtn.classList.replace("bx-menu","bx-menu-alt-right")}else{sidebarBtn.classList.replace("bx-menu-alt-right","bx-menu")}};(function(name,factory){if(typeof window==='object'){window[name]=factory()}})('Ribbons',function(){var _w=window,_b=document.body,_d=document.documentElement;var random=function(){if(arguments.length===1){if(Array.isArray(arguments[0])){var index=Math.round(random(0,arguments[0].length-1));return arguments[0][index]}return random(0,arguments[0]);}else if(arguments.length===2){return Math.random()*(arguments[1]-arguments[0])+arguments[0]}else if(arguments.length===4){var array=[arguments[0],arguments[1],arguments[2],arguments[3]];return array[Math.floor(Math.random()*array.length)];}return 0;};var screenInfo=function(e){var width=Math.max(0,_w.innerWidth||_d.clientWidth||_b.clientWidth||0),height=Math.max(0,_w.innerHeight||_d.clientHeight||_b.clientHeight||0),scrollx=Math.max(0,_w.pageXOffset||_d.scrollLeft||_b.scrollLeft||0)-(_d.clientLeft||0),scrolly=Math.max(0,_w.pageYOffset||_d.scrollTop||_b.scrollTop||0)-(_d.clientTop||0);return{width:width,height:height,ratio:width/height,centerx:width/2,centery:height/2,scrollx:scrollx,scrolly:scrolly}};var mouseInfo=function(e){var screen=screenInfo(e),mousex=e?Math.max(0,e.pageX||e.clientX||0):0,mousey=e?Math.max(0,e.pageY||e.clientY||0):0;return{mousex:mousex,mousey:mousey,centerx:mousex-screen.width/2,centery:mousey-screen.height/2}};var Point=function(x,y){this.x=0;this.y=0;this.set(x,y)};Point.prototype={constructor:Point,set:function(x,y){this.x=x||0;this.y=y||0},copy:function(point){this.x=point.x||0;this.y=point.y||0;return this},multiply:function(x,y){this.x*=x||1;this.y*=y||1;return this},divide:function(x,y){this.x/=x||1;this.y/=y||1;return this},add:function(x,y){this.x+=x||0;this.y+=y||0;return this},subtract:function(x,y){this.x-=x||0;this.y-=y||0;return this},clampX:function(min,max){this.x=Math.max(min,Math.min(this.x,max));return this},clampY:function(min,max){this.y=Math.max(min,Math.min(this.y,max));return this},flipX:function(){this.x*=-1;return this},flipY:function(){this.y*=-1;return this}};var Factory=function(options){this._canvas=null;this._context=null;this._sto=null;this._width=0;this._height=0;this._scroll=0;this._ribbons=[];this._options={colorSaturation:'80%',colorBrightness:'60%',colorAlpha:0.65,colorCycleSpeed:6,verticalPosition:'center',horizontalSpeed:150,ribbonCount:3,strokeSize:0,parallaxAmount:-0.5,animateSections:true};this._onDraw=this._onDraw.bind(this);this._onResize=this._onResize.bind(this);this._onScroll=this._onScroll.bind(this);this.setOptions(options);this.init()};Factory.prototype={constructor:Factory,setOptions:function(options){if(typeof options==='object'){for(var key in options){if(options.hasOwnProperty(key)){this._options[key]=options[key]}}}},init:function(){try{this._canvas=document.createElement('canvas');this._canvas.style['display']='block';this._canvas.style['position']='fixed';this._canvas.style['margin']='0';this._canvas.style['padding']='0';this._canvas.style['border']='0';this._canvas.style['outline']='0';this._canvas.style['left']='0';this._canvas.style['top']='0';this._canvas.style['width']='100%';this._canvas.style['height']='100%';this._canvas.style['z-index']='-1';this._onResize();this._context=this._canvas.getContext('2d');this._context.clearRect(0,0,this._width,this._height);this._context.globalAlpha=this._options.colorAlpha;window.addEventListener('resize',this._onResize);window.addEventListener('scroll',this._onScroll);document.body.appendChild(this._canvas)}catch(e){console.warn('Canvas Context Error: '+e.toString());return}this._onDraw()},addRibbon:function(){var dir=Math.round(random(1,9))>5?'right':'left',stop=1000,hide=200,min=0-hide,max=this._width+hide,movex=0,movey=0,startx=dir==='right'?min:max,starty=Math.round(random(0,this._height));if(/^(top|min)$/i.test(this._options.verticalPosition)){starty=0+hide}else if(/^(middle|center)$/i.test(this._options.verticalPosition)){starty=this._height/2}else if(/^(bottom|max)$/i.test(this._options.verticalPosition)){starty=this._height-hide}var ribbon=[],point1=new Point(startx,starty),point2=new Point(startx,starty),point3=null,color=Math.round(random(900)),delay=0;while(true){if(stop<=0){break}stop--;movex=Math.round((Math.random()*1-0.2)*this._options.horizontalSpeed);movey=Math.round((Math.random()*1-0.5)*(this._height*0.25));point3=new Point();point3.copy(point2);if(dir==='right'){point3.add(movex,movey);if(point2.x>=max){break}}else if(dir==='left'){point3.subtract(movex,movey);if(point2.x<=min){break}}ribbon.push({point1:new Point(point1.x,point1.y),point2:new Point(point2.x,point2.y),point3:point3,color:color,delay:delay,dir:dir,alpha:0,phase:0});point1.copy(point2);point2.copy(point3);delay+=4}this._ribbons.push(ribbon)},_drawRibbonSection:function(section){if(section){if(section.phase>=1&&section.alpha<=0){return true;}if(section.delay<=0){section.phase+=0.02;section.alpha=Math.sin(section.phase)*1;section.alpha=section.alpha<=0?0:section.alpha;section.alpha=section.alpha>=1?1:section.alpha;if(this._options.animateSections){var mod=Math.sin(1+section.phase*Math.PI/2)*0.1;if(section.dir==='right'){section.point1.add(mod,0);section.point2.add(mod,0);section.point3.add(mod,0)}else{section.point1.subtract(mod,0);section.point2.subtract(mod,0);section.point3.subtract(mod,0)}section.point1.add(0,mod);section.point2.add(0,mod);section.point3.add(0,mod)}}else{section.delay-=0.5}var s=this._options.colorSaturation,l=this._options.colorBrightness,c='hsla('+section.color+', '+s+', '+l+', '+section.alpha+' )';this._context.save();if(this._options.parallaxAmount!==0){this._context.translate(0,this._scroll*this._options.parallaxAmount)}this._context.beginPath();this._context.moveTo(section.point1.x,section.point1.y);this._context.lineTo(section.point2.x,section.point2.y);this._context.lineTo(section.point3.x,section.point3.y);this._context.fillStyle=c;this._context.fill();if(this._options.strokeSize>0){this._context.lineWidth=this._options.strokeSize;this._context.strokeStyle=c;this._context.lineCap='round';this._context.stroke()}this._context.restore()}return false;},_onDraw:function(){for(var i=0,t=this._ribbons.length;i<t;i+=1){if(!this._ribbons[i]){this._ribbons.splice(i,1)}}this._context.clearRect(0,0,this._width,this._height);for(var a=0;a<this._ribbons.length;++a ){var ribbon=this._ribbons[a],numSections=ribbon.length,numDone=0;for(var b=0;b<numSections;++b ){if(this._drawRibbonSection(ribbon[b])){numDone++;}}if(numDone>=numSections){this._ribbons[a]=null}}if(this._ribbons.length<this._options.ribbonCount){this.addRibbon()}requestAnimationFrame(this._onDraw)},_onResize:function(e){var screen=screenInfo(e);this._width=screen.width;this._height=screen.height;if(this._canvas){this._canvas.width=this._width;this._canvas.height=this._height;if(this._context){this._context.globalAlpha=this._options.colorAlpha}}},_onScroll:function(e){var screen=screenInfo(e);this._scroll=screen.scrolly}};return Factory});new Ribbons({colorSaturation:'60%',colorBrightness:'50%',colorAlpha:0.5,colorCycleSpeed:5,verticalPosition:'random',horizontalSpeed:200,ribbonCount:3,strokeSize:0,parallaxAmount:-0.2,animateSections:true});</script><script>let sidebar=document.querySelector(".sidebar");let sidebarBtn=document.querySelector(".sidebarBtn");sidebarBtn.onclick=function(){sidebar.classList.toggle("active");if(sidebar.classList.contains("active")){sidebarBtn.classList.replace("bx-menu","bx-menu-alt-right")}else{sidebarBtn.classList.replace("bx-menu-alt-right","bx-menu")}};(function(name,factory){if(typeof window==='object'){window[name]=factory()}})('Ribbons',function(){var _w=window,_b=document.body,_d=document.documentElement;var random=function(){if(arguments.length===1){if(Array.isArray(arguments[0])){var index=Math.round(random(0,arguments[0].length-1));return arguments[0][index]}return random(0,arguments[0]);}else if(arguments.length===2){return Math.random()*(arguments[1]-arguments[0])+arguments[0]}else if(arguments.length===4){var array=[arguments[0],arguments[1],arguments[2],arguments[3]];return array[Math.floor(Math.random()*array.length)];}return 0;};var screenInfo=function(e){var width=Math.max(0,_w.innerWidth||_d.clientWidth||_b.clientWidth||0),height=Math.max(0,_w.innerHeight||_d.clientHeight||_b.clientHeight||0),scrollx=Math.max(0,_w.pageXOffset||_d.scrollLeft||_b.scrollLeft||0)-(_d.clientLeft||0),scrolly=Math.max(0,_w.pageYOffset||_d.scrollTop||_b.scrollTop||0)-(_d.clientTop||0);return{width:width,height:height,ratio:width/height,centerx:width/2,centery:height/2,scrollx:scrollx,scrolly:scrolly}};var mouseInfo=function(e){var screen=screenInfo(e),mousex=e?Math.max(0,e.pageX||e.clientX||0):0,mousey=e?Math.max(0,e.pageY||e.clientY||0):0;return{mousex:mousex,mousey:mousey,centerx:mousex-screen.width/2,centery:mousey-screen.height/2}};var Point=function(x,y){this.x=0;this.y=0;this.set(x,y)};Point.prototype={constructor:Point,set:function(x,y){this.x=x||0;this.y=y||0},copy:function(point){this.x=point.x||0;this.y=point.y||0;return this},multiply:function(x,y){this.x*=x||1;this.y*=y||1;return this},divide:function(x,y){this.x/=x||1;this.y/=y||1;return this},add:function(x,y){this.x+=x||0;this.y+=y||0;return this},subtract:function(x,y){this.x-=x||0;this.y-=y||0;return this},clampX:function(min,max){this.x=Math.max(min,Math.min(this.x,max));return this},clampY:function(min,max){this.y=Math.max(min,Math.min(this.y,max));return this},flipX:function(){this.x*=-1;return this},flipY:function(){this.y*=-1;return this}};var Factory=function(options){this._canvas=null;this._context=null;this._sto=null;this._width=0;this._height=0;this._scroll=0;this._ribbons=[];this._options={colorSaturation:'80%',colorBrightness:'60%',colorAlpha:0.65,colorCycleSpeed:6,verticalPosition:'center',horizontalSpeed:150,ribbonCount:3,strokeSize:0,parallaxAmount:-0.5,animateSections:true};this._onDraw=this._onDraw.bind(this);this._onResize=this._onResize.bind(this);this._onScroll=this._onScroll.bind(this);this.setOptions(options);this.init()};Factory.prototype={constructor:Factory,setOptions:function(options){if(typeof options==='object'){for(var key in options){if(options.hasOwnProperty(key)){this._options[key]=options[key]}}}},init:function(){try{this._canvas=document.createElement('canvas');this._canvas.style['display']='block';this._canvas.style['position']='fixed';this._canvas.style['margin']='0';this._canvas.style['padding']='0';this._canvas.style['border']='0';this._canvas.style['outline']='0';this._canvas.style['left']='0';this._canvas.style['top']='0';this._canvas.style['width']='100%';this._canvas.style['height']='100%';this._canvas.style['z-index']='-1';this._onResize();this._context=this._canvas.getContext('2d');this._context.clearRect(0,0,this._width,this._height);this._context.globalAlpha=this._options.colorAlpha;window.addEventListener('resize',this._onResize);window.addEventListener('scroll',this._onScroll);document.body.appendChild(this._canvas)}catch(e){console.warn('Canvas Context Error: '+e.toString());return}this._onDraw()},addRibbon:function(){var dir=Math.round(random(1,9))>5?'right':'left',stop=1000,hide=200,min=0-hide,max=this._width+hide,movex=0,movey=0,startx=dir==='right'?min:max,starty=Math.round(random(0,this._height));if(/^(top|min)$/i.test(this._options.verticalPosition)){starty=0+hide}else if(/^(middle|center)$/i.test(this._options.verticalPosition)){starty=this._height/2}else if(/^(bottom|max)$/i.test(this._options.verticalPosition)){starty=this._height-hide}var ribbon=[],point1=new Point(startx,starty),point2=new Point(startx,starty),point3=null,color=Math.round(random(900)),delay=0;while(true){if(stop<=0){break}stop--;movex=Math.round((Math.random()*1-0.2)*this._options.horizontalSpeed);movey=Math.round((Math.random()*1-0.5)*(this._height*0.25));point3=new Point();point3.copy(point2);if(dir==='right'){point3.add(movex,movey);if(point2.x>=max){break}}else if(dir==='left'){point3.subtract(movex,movey);if(point2.x<=min){break}}ribbon.push({point1:new Point(point1.x,point1.y),point2:new Point(point2.x,point2.y),point3:point3,color:color,delay:delay,dir:dir,alpha:0,phase:0});point1.copy(point2);point2.copy(point3);delay+=4}this._ribbons.push(ribbon)},_drawRibbonSection:function(section){if(section){if(section.phase>=1&&section.alpha<=0){return true;}if(section.delay<=0){section.phase+=0.02;section.alpha=Math.sin(section.phase)*1;section.alpha=section.alpha<=0?0:section.alpha;section.alpha=section.alpha>=1?1:section.alpha;if(this._options.animateSections){var mod=Math.sin(1+section.phase*Math.PI/2)*0.1;if(section.dir==='right'){section.point1.add(mod,0);section.point2.add(mod,0);section.point3.add(mod,0)}else{section.point1.subtract(mod,0);section.point2.subtract(mod,0);section.point3.subtract(mod,0)}section.point1.add(0,mod);section.point2.add(0,mod);section.point3.add(0,mod)}}else{section.delay-=0.5}var s=this._options.colorSaturation,l=this._options.colorBrightness,c='hsla('+section.color+', '+s+', '+l+', '+section.alpha+' )';this._context.save();if(this._options.parallaxAmount!==0){this._context.translate(0,this._scroll*this._options.parallaxAmount)}this._context.beginPath();this._context.moveTo(section.point1.x,section.point1.y);this._context.lineTo(section.point2.x,section.point2.y);this._context.lineTo(section.point3.x,section.point3.y);this._context.fillStyle=c;this._context.fill();if(this._options.strokeSize>0){this._context.lineWidth=this._options.strokeSize;this._context.strokeStyle=c;this._context.lineCap='round';this._context.stroke()}this._context.restore()}return false;},_onDraw:function(){for(var i=0,t=this._ribbons.length;i<t;i+=1){if(!this._ribbons[i]){this._ribbons.splice(i,1)}}this._context.clearRect(0,0,this._width,this._height);for(var a=0;a<this._ribbons.length;++a ){var ribbon=this._ribbons[a],numSections=ribbon.length,numDone=0;for(var b=0;b<numSections;++b ){if(this._drawRibbonSection(ribbon[b])){numDone++;}}if(numDone>=numSections){this._ribbons[a]=null}}if(this._ribbons.length<this._options.ribbonCount){this.addRibbon()}requestAnimationFrame(this._onDraw)},_onResize:function(e){var screen=screenInfo(e);this._width=screen.width;this._height=screen.height;if(this._canvas){this._canvas.width=this._width;this._canvas.height=this._height;if(this._context){this._context.globalAlpha=this._options.colorAlpha}}},_onScroll:function(e){var screen=screenInfo(e);this._scroll=screen.scrolly}};return Factory});new Ribbons({colorSaturation:'60%',colorBrightness:'50%',colorAlpha:0.5,colorCycleSpeed:5,verticalPosition:'random',horizontalSpeed:200,ribbonCount:3,strokeSize:0,parallaxAmount:-0.2,animateSections:true})</script> ;</body></html><style>@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@200;300;400;500;600;700&display=swap');body{background-color:black;background-image:radial-gradient(circle, rgba(0, 0, 0, 0) 0%, rgba(0, 0, 0, 0.8) 100%);background-position:center center;background-repeat:no-repeat;background-attachment:fixed;background-size:cover}*{margin:0;padding:0;box-sizing:border-box;font-family:'Poppins', sans-serif}.sidebar{position:fixed;height:100%;width:240px;transition:all 0.5s ease;overflow-y:scroll}.sidebar.active{width:60px}.sidebar .logo-details{height:80px;display:flex;align-items:center}.sidebar .logo-details i{font-size:28px;font-weight:500;color:#fff;min-width:60px;text-align:center}.sidebar .logo-details .logo_name{color:#fff;font-size:24px;font-weight:500}.sidebar .nav-links{margin-top:10px}.sidebar .nav-links li{position:relative;list-style:none;height:50px}.sidebar .nav-links li a{height:100%;width:100%;display:flex;align-items:center;text-decoration:none;transition:all 0.4s ease}.sidebar .nav-links li a.active{background:blueviolet}.sidebar .nav-links li a:hover{background:red}.sidebar .nav-links li i{min-width:60px;text-align:center;font-size:18px;color:#fff}.sidebar .nav-links li a .links_name{color:#fff;font-size:15px;font-weight:400;white-space:nowrap}.sidebar .nav-links .log_out{position:absolute;bottom:0;width:100%}.home-section{position:relative;min-height:100vh;width:calc(100% - 240px);left:240px;transition:all 0.5s ease}.sidebar.active~.home-section{width:calc(100% - 60px);left:60px}.home-section nav{display:flex;justify-content:space-between;height:80px;display:flex;align-items:center;position:fixed;width:calc(100% - 240px);left:240px;z-index:100;padding:0 20px;box-shadow:0 1px 1px rgba(0, 0, 0, 0.1);color:white;transition:all 0.5s ease}.sidebar.active~.home-section nav{left:60px;width:calc(100% - 60px)}.home-section nav .sidebar-button{display:flex;align-items:center;font-size:24px;font-weight:500}nav .sidebar-button i{font-size:35px;margin-right:10px}.home-section .home-content{position:relative;padding-top:104px}.home-content .overview-boxes{display:flex;align-items:center;justify-content:space-between;flex-wrap:wrap;padding:0 20px;margin-bottom:26px}.overview-boxes .box{display:flex;align-items:center;justify-content:center;width:calc(100% / 4 - 15px);padding:15px 14px;border-radius:12px;box-shadow:0 5px 10px rgba(0, 0, 0, 0.1)}.overview-boxes .box2{margin-top:50px;display:flex;align-items:center;justify-content:center;width:calc(100% / 4 - 15px);padding:15px 14px;border-radius:12px;box-shadow:0 5px 10px rgba(0, 0, 0, 0.1)}.overview-boxes .box-topic a{color:#e05260;text-decoration:none;font-size:20px;font-weight:500}.overview-boxes .box-topic2 a{color:#e05260;text-decoration:none;font-size:20px;font-weight:500}.home-content .box .number{display:inline-block;font-size:35px;margin-top:-6px;font-weight:500}.home-content .box2 .number2{display:inline-block;font-size:35px;margin-top:-6px;font-weight:500}.home-content .sales-boxes{display:flex;justify-content:space-between}.home-content .sales-boxes .recent-sales{width:65%;background:#fff;padding:20px 30px;margin:0 20px;border-radius:12px;box-shadow:0 5px 10px rgba(0, 0, 0, 0.1)}.home-content .sales-boxes .sales-details{display:flex;align-items:center;justify-content:space-between}.sales-boxes .box .title{font-size:24px;font-weight:500}.sales-boxes .sales-details li.topic{font-size:20px;font-weight:500}.sales-boxes .sales-details li{list-style:none;margin:8px 0}.sales-boxes .sales-details li a{font-size:18px;color:#333;font-size:400;text-decoration:none}.sales-boxes .box .button{width:100%;display:flex;justify-content:flex-end}.sales-boxes .box .button a{color:#fff;background:#0A2558;padding:4px 12px;font-size:15px;font-weight:400;border-radius:4px;text-decoration:none;transition:all 0.3s ease}.sales-boxes .box .button a:hover{background:#0d3073}.home-content .sales-boxes .top-sales{width:35%;background:#fff;padding:20px 30px;margin:0 20px 0 0;border-radius:12px;box-shadow:0 5px 10px rgba(0, 0, 0, 0.1)}.sales-boxes .top-sales li{display:flex;align-items:center;justify-content:space-between;margin:10px 0}.sales-boxes .top-sales li a img{height:40px;width:40px;object-fit:cover;border-radius:12px;margin-right:10px;background:#333}.sales-boxes .top-sales li a{display:flex;align-items:center;text-decoration:none}.price,.sales-boxes .top-sales li .product{font-size:17px;font-weight:400;color:#333}@media (max-width: 1240px){.sidebar{width:60px}.sidebar.active{width:220px}.home-section{width:calc(100% - 60px);left:60px}.sidebar.active~.home-section{overflow:hidden;left:220px}.home-section nav{width:calc(100% - 60px);left:60px}.sidebar.active~.home-section nav{width:calc(100% - 220px);left:220px}}@media (max-width: 1150px){.home-content .sales-boxes{flex-direction:column}.home-content .sales-boxes .box{width:100%;overflow-x:scroll;margin-bottom:30px}.home-content .sales-boxes .top-sales{margin:0}}@media (max-width: 1000px){.overview-boxes .box{width:calc(100% / 2 - 15px);margin-bottom:15px}}@media (max-width: 700px){nav .profile-details .admin_name,nav .profile-details i,nav .sidebar-button .dashboard{display:none}.home-section nav .profile-details{height:50px;min-width:40px}.home-content .sales-boxes .sales-details{width:560px}}@media (max-width: 550px){.overview-boxes .box{width:100%;margin-bottom:15px}.sidebar.active~.home-section nav .profile-details{display:none}}@media (max-width: 400px){.sidebar{width:0}.sidebar.active{width:60px}.home-section{width:100%;left:0}.sidebar.active~.home-section{left:60px;width:calc(100% - 60px)}.home-section nav{width:100%;left:0}.sidebar.active~.home-section nav{left:60px;width:calc(100% - 60px)}}*,*:after,*:before{padding:0;margin:0;box-sizing:border-box}.background{width:430px;height:520px;position:absolute;transform:translate(-50%, -50%);left:50%;top:50%}form{height:340px;width:400px;background-color:rgba(255, 255, 255, 0.13);position:absolute;transform:translate(-50%, -50%);top:400%;left:50%;border-radius:10px;backdrop-filter: blur(10px);border:2px solid rgba(255, 255, 255, 0.1);box-shadow:0 0 40px rgba(8, 7, 16, 0.6);padding:50px 35px}form *{font-family:'Poppins', sans-serif;color:#ffffff;letter-spacing:0.5px;outline:none;border:none}label{display:block;margin-top:30px;font-size:16px;font-weight:500}input{display:block;height:50px;width:100%;background-color:rgba(255, 255, 255, 0.07);border-radius:3px;padding:0 10px;margin-top:8px;font-size:14px;font-weight:300}::placeholder{color:#e5e5e5}button{margin-top:50px;width:100%;background-color:#ffffff;color:#080710;padding:15px 0;font-size:18px;font-weight:600;border-radius:5px;cursor:pointer}</style>`
	Write("Modules/Server/HTML/Masher.html", StaticTmplMasherHTMKL)

}

func UnmarshalServerFiles() []string {
	var Files []string
	f, x := os.Open("Modules/Gsrc/Database/ServerInfo.json")
	if x != nil {
		log.Fatal(x)
	} else {
		defer f.Close()
		bv, _ := ioutil.ReadAll(f)
		var exp JSONSERVERINFO
		json.Unmarshal([]byte(bv), &exp)
		for i := 0; i < len(exp.ServerDatabase); i++ {
			Files = append(Files, exp.ServerDatabase[i])
		}
	}
	if Files != nil {
		return Files
	} else {
		fmt.Println("Error: Returning empty value, for some reason the files array was empty, this should not happen which must mean the server info file located at | Modules/Gsrc/Database/ServerInfo.json is corrupted or missing the section.")
		return []string{"ERROR"} // Return error based array | Will crash the applications
	}
}

func Opener(filename string) []string {
	var text []string
	f, x := os.Open(filename)
	if x != nil {
		log.Fatal(x)
	} else {
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			text = append(text, scanner.Text())
		}
	}
	return text
}

func CalltoStoreServerFile() string {
	var Templated2 string
	Templated2 += JSONTEMPLATETOP
	File := UnmarshalServerFiles()
	for o := range File {
		TextData := Opener(File[o])
		Templated2 += fmt.Sprintf(`<div class="codeheader" id="JSON">JSON FILE (DB) | %s </div>`, File[o])
		Templated2 += `<pre class="syntax">`
		for i := range TextData {
			Templated2 += TextData[i] + "<br>"
		}
		Templated2 += `</pre></div>`

	}
	Templated2 += JSONTEMPLATEBOTTOM
	if Templated2 != "" {
		return Templated2
	} else {
		return "ERROR:"
	}

}

func Call_To_Store_ServerInfo() string {
	var Template4 string
	Template4 += APPINFOSERVERTEMPLATETOP
	Template4 = fmt.Sprintf(Template4, StructureAppInfo.Version, StructureAppInfo.ApplicationRunBinary, StructureAppInfo.Name)
	Template4 += fmt.Sprintf("<tr><td>File Formats</td><td>%s</td></tr>", StructureAppInfo.ApplicationFileFormats)
	Template4 += fmt.Sprintf("<tr><td>App Languages</td><td>%s</td></tr>", StructureAppInfo.ApplicationLanguages)
	Template4 += fmt.Sprintf("<tr><td>App Security</td><td>%s</td></tr>", StructureAppInfo.ApplicationSecurity)
	Template4 += fmt.Sprintf("<tr><td>App Dash Tabs</td><td>%s</td></tr>", StructureAppInfo.ApplicationTabs)
	Template4 += fmt.Sprintf("<tr><td>App Indexes</td><td>%s</td></tr>", StructureAppInfo.ApplicationIndexes)
	Template4 += fmt.Sprintf("<tr><td>App Suported Protos</td><td>%s</td></tr>", StructureAppInfo.ApplicationProtocols)
	Template4 += fmt.Sprintf("<tr><td>App Front support</td><td>%s</td></tr>", StructureAppInfo.ApplicationSupport)
	Template4 += fmt.Sprintf("<tr><td>App Binaries</td><td>%s</td></tr>", StructureAppInfo.ApplicationBinaries)
	Template4 += fmt.Sprintf("<tr><td>App Run Binary</td><td>%s</td></tr>", StructureAppInfo.ApplicationRunBinary)
	Template4 += fmt.Sprintf("<tr><td>App Aprox Runtime</td><td>%s</td></tr>", StructureAppInfo.ApplicationRunTime)
	Template4 += fmt.Sprintf("<tr><td>App Support File Formats</td><td>%s</td></tr>", StructureAppInfo.ApplicationSupportedFileFormats)
	Template4 += APPINFOSERVERTEMPLATEBOTTOm
	return Template4
}

func Call_To_Store_PreProc() string {
	var Template3 string
	Template3 += INFOINFOSERVERNEEDSTEMPLATETOP
	Template3 += fmt.Sprintf("<tr><td>OS</td><td>%s</td></tr>", StructurePreProcessor.ServerOperatingSystem)
	Template3 += fmt.Sprintf("<tr><td>OS FS (FileSystem)</td><td>%s</td></tr>", StructurePreProcessor.ServerOperatingSystemFileSystem)
	Template3 += fmt.Sprintf("<tr><td>OS Version</td><td>%s</td></tr>", StructurePreProcessor.ServerOperatingSystemVersion)
	Template3 += fmt.Sprintf("<tr><td>CPU Architecture</td><td>%s</td></tr>", StructurePreProcessor.ServerOperatingArchitecture)
	Template3 += fmt.Sprintf("<tr><td>CPU Vendor ID</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoCPUVendorID)
	Template3 += fmt.Sprintf("<tr><td>CPU Index Number</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoCPUIndexNum)
	Template3 += fmt.Sprintf("<tr><td>CPU Family Number</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoCPUFamily)
	Template3 += fmt.Sprintf("<tr><td>CPU NUMBER OF CORES</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoCPUNumberOfCores)
	Template3 += fmt.Sprintf("<tr><td>CPU Model Name</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoCPUModelName)
	Template3 += fmt.Sprintf("<tr><td>CPU Speed Mhz</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoCPUSpeed)
	Template3 += fmt.Sprintf("<tr><td>CPU Cache Size</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoCPUCacheSize)
	Template3 += fmt.Sprintf("<tr><td>CPU Micronode</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoCPUMicronode)
	Template3 += fmt.Sprintf("<tr><td>CPU Model</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoCPUModelName)
	Template3 += fmt.Sprintf("<tr><td>CPU Physical ID</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoCPUPhysID)
	Template3 += fmt.Sprintf("<tr><td>CPU Step</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoCPUStep)
	Template3 += fmt.Sprintf("<tr><td>Free Memory</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoMEMFree)
	Template3 += fmt.Sprintf("<tr><td>Total Memory</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoMEMTotal)
	Template3 += fmt.Sprintf("<tr><td>Used Memory</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoMEMUsed)
	Template3 += fmt.Sprintf("<tr><td>Hostname (OS)</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoOSPHostname)
	Template3 += fmt.Sprintf("<tr><td>Uptime</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoOSPUptime)
	Template3 += fmt.Sprintf("<tr><td>Processes Running</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoOSPProcRunning)
	Template3 += fmt.Sprintf("<tr><td>Host ID</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoOSPHOSTID)
	Template3 += fmt.Sprintf("<tr><td>Host OS</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoOSPHOSTOS)
	Template3 += fmt.Sprintf("<tr><td>Host OS Family</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoOSPHOSTPLAT)
	Template3 += fmt.Sprintf("<tr><td>OS Kernel Version</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoOSPHOSTKERNELVERSION)
	Template3 += fmt.Sprintf("<tr><td>OS Kernel Architecture</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoOSPHOSTKERNELARCHITECTURE)
	Template3 += fmt.Sprintf("<tr><td>Platform Name</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoOSPHOSTPLAT)
	Template3 += fmt.Sprintf("<tr><td>Platform Family</td><td>%s</td></tr>", StructurePreProcessor.ServerHardwareInfoOSPHOSTPLATFORMVERSION)
	Template3 += INFOSERVERNEEDSTEMPLATEBOTTOM
	return Template3
}
func StoreServerInformationFileFromDB() string {
	var Template5 string
	Template5 += ServerINFORMATIONTOP
	// Init data
	Template5 += Generate_Box("Server Version", StructureServerInfo.ServerVersion)
	Template5 += Generate_Box("Server Port", fmt.Sprint(StructureServerInfo.ServerMainPort))
	Template5 += Generate_Box("Server URL", StructureServerInfo.ServerMainURL)
	Template5 += "</div><hr><br>"
	Template5 += Generate_CH("Extra", "All Server URLS")
	Template5 += `<pre class="syntax">`
	for i := 0; i < len(StructureServerInfo.ServerFiles); i++ {
		Template5 += StructureServerInfo.ServerFiles[i] + " \n "
	}
	Template5 += `</pre>`
	Template5 += `<table class="container_Overview"><br><br><thead><tr><th><h1>Tag Value</h1></th><th><h1>Information</h1></th></tr></thead><tbody>`
	Template5 += `<tr><td>All Server URL's</td><td>`
	for k := 0; k < len(StructureServerInfo.ServerUrls); k++ {
		Template5 += StructureServerInfo.ServerUrls[k] + " \n "
	}
	Template5 += `</td></tr>`
	Template5 += `<tr><td>Database</td><td>`
	for k := 0; k < len(StructureServerInfo.ServerDatabase); k++ {
		Template5 += StructureServerInfo.ServerDatabase[k] + " \n "
	}
	Template5 += `</td></tr>`
	Template5 += `<tr><td>Browser Support</td><td>`
	for k := 0; k < len(StructureServerInfo.ServerSupport); k++ {
		Template5 += StructureServerInfo.ServerSupport[k] + " \n "
	}
	Template5 += `</td></tr>`
	Template5 += `<tr><td>Server Import List</td><td>`
	for k := 0; k < len(StructureServerInfo.ServerImports); k++ {
		Template5 += StructureServerInfo.ServerImports[k] + " \n "
	}
	Template5 += `</td></tr>`
	Template5 += `<tr><td>Suggested Browsers</td><td>`
	for k := 0; k < len(StructureServerInfo.ServerSuggests); k++ {
		Template5 += StructureServerInfo.ServerSuggests[k] + " \n "
	}
	Template5 += `</td></tr>`
	Template5 += `<tr><td>Server Ports</td><td>`
	for k := 0; k < len(StructureServerInfo.ServerPorts); k++ {
		Template5 += fmt.Sprint(StructureServerInfo.ServerPorts[k]) + " \n "
	}
	Template5 += `</td></tr>`
	Template5 += ServerINFORMATIONBOTTOM

	return Template5

}
