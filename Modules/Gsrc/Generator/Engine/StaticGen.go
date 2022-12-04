package Engine

// Static data for example of credential HTML documents
var StandardTop = `<!DOCTYPE html><html lang="en" dir="ltr"><head><title>Frizz | Credentials</title><meta charset="UTF-8"><link rel="stylesheet" href="style.css"><link href='https://unpkg.com/boxicons@2.1.2/css/boxicons.min.css' rel='stylesheet'><meta name="viewport" content="width=device-width, initial-scale=1.0"></head><source src="../HTML/LobbyMisc/Lobby_Music" type="audio/mpeg"><source src="../HTML/Future/Future_Lobby" type="audio/mpeg">`
var StaticJS = `<script>let sidebar=document.querySelector(".sidebar");let sidebarBtn=document.querySelector(".sidebarBtn");sidebarBtn.onclick=function(){sidebar.classList.toggle("active");if(sidebar.classList.contains("active")){sidebarBtn.classList.replace("bx-menu","bx-menu-alt-right")}else{sidebarBtn.classList.replace("bx-menu-alt-right","bx-menu")}};(function(name,factory){if(typeof window==='object'){window[name]=factory()}})('Ribbons',function(){var _w=window,_b=document.body,_d=document.documentElement;var random=function(){if(arguments.length===1){if(Array.isArray(arguments[0])){var index=Math.round(random(0,arguments[0].length-1));return arguments[0][index]}return random(0,arguments[0]);}else if(arguments.length===2){return Math.random()*(arguments[1]-arguments[0])+arguments[0]}else if(arguments.length===4){var array=[arguments[0],arguments[1],arguments[2],arguments[3]];return array[Math.floor(Math.random()*array.length)];}return 0;};var screenInfo=function(e){var width=Math.max(0,_w.innerWidth||_d.clientWidth||_b.clientWidth||0),height=Math.max(0,_w.innerHeight||_d.clientHeight||_b.clientHeight||0),scrollx=Math.max(0,_w.pageXOffset||_d.scrollLeft||_b.scrollLeft||0)-(_d.clientLeft||0),scrolly=Math.max(0,_w.pageYOffset||_d.scrollTop||_b.scrollTop||0)-(_d.clientTop||0);return{width:width,height:height,ratio:width/height,centerx:width/2,centery:height/2,scrollx:scrollx,scrolly:scrolly}};var mouseInfo=function(e){var screen=screenInfo(e),mousex=e?Math.max(0,e.pageX||e.clientX||0):0,mousey=e?Math.max(0,e.pageY||e.clientY||0):0;return{mousex:mousex,mousey:mousey,centerx:mousex-screen.width/2,centery:mousey-screen.height/2}};var Point=function(x,y){this.x=0;this.y=0;this.set(x,y)};Point.prototype={constructor:Point,set:function(x,y){this.x=x||0;this.y=y||0},copy:function(point){this.x=point.x||0;this.y=point.y||0;return this},multiply:function(x,y){this.x*=x||1;this.y*=y||1;return this},divide:function(x,y){this.x/=x||1;this.y/=y||1;return this},add:function(x,y){this.x+=x||0;this.y+=y||0;return this},subtract:function(x,y){this.x-=x||0;this.y-=y||0;return this},clampX:function(min,max){this.x=Math.max(min,Math.min(this.x,max));return this},clampY:function(min,max){this.y=Math.max(min,Math.min(this.y,max));return this},flipX:function(){this.x*=-1;return this},flipY:function(){this.y*=-1;return this}};var Factory=function(options){this._canvas=null;this._context=null;this._sto=null;this._width=0;this._height=0;this._scroll=0;this._ribbons=[];this._options={colorSaturation:'80%',colorBrightness:'60%',colorAlpha:0.65,colorCycleSpeed:6,verticalPosition:'center',horizontalSpeed:150,ribbonCount:3,strokeSize:0,parallaxAmount:-0.5,animateSections:true};this._onDraw=this._onDraw.bind(this);this._onResize=this._onResize.bind(this);this._onScroll=this._onScroll.bind(this);this.setOptions(options);this.init()};Factory.prototype={constructor:Factory,setOptions:function(options){if(typeof options==='object'){for(var key in options){if(options.hasOwnProperty(key)){this._options[key]=options[key]}}}},init:function(){try{this._canvas=document.createElement('canvas');this._canvas.style['display']='block';this._canvas.style['position']='fixed';this._canvas.style['margin']='0';this._canvas.style['padding']='0';this._canvas.style['border']='0';this._canvas.style['outline']='0';this._canvas.style['left']='0';this._canvas.style['top']='0';this._canvas.style['width']='100%';this._canvas.style['height']='100%';this._canvas.style['z-index']='-1';this._onResize();this._context=this._canvas.getContext('2d');this._context.clearRect(0,0,this._width,this._height);this._context.globalAlpha=this._options.colorAlpha;window.addEventListener('resize',this._onResize);window.addEventListener('scroll',this._onScroll);document.body.appendChild(this._canvas)}catch(e){console.warn('Canvas Context Error: '+e.toString());return}this._onDraw()},addRibbon:function(){var dir=Math.round(random(1,9))>5?'right':'left',stop=1000,hide=200,min=0-hide,max=this._width+hide,movex=0,movey=0,startx=dir==='right'?min:max,starty=Math.round(random(0,this._height));if(/^(top|min)$/i.test(this._options.verticalPosition)){starty=0+hide}else if(/^(middle|center)$/i.test(this._options.verticalPosition)){starty=this._height/2}else if(/^(bottom|max)$/i.test(this._options.verticalPosition)){starty=this._height-hide}var ribbon=[],point1=new Point(startx,starty),point2=new Point(startx,starty),point3=null,color=Math.round(random(900)),delay=0;while(true){if(stop<=0){break}stop--;movex=Math.round((Math.random()*1-0.2)*this._options.horizontalSpeed);movey=Math.round((Math.random()*1-0.5)*(this._height*0.25));point3=new Point();point3.copy(point2);if(dir==='right'){point3.add(movex,movey);if(point2.x>=max){break}}else if(dir==='left'){point3.subtract(movex,movey);if(point2.x<=min){break}}ribbon.push({point1:new Point(point1.x,point1.y),point2:new Point(point2.x,point2.y),point3:point3,color:615,delay:delay,dir:dir,alpha:0,phase:0});point1.copy(point2);point2.copy(point3);delay+=4}this._ribbons.push(ribbon)},_drawRibbonSection:function(section){if(section){if(section.phase>=1&&section.alpha<=0){return true;}if(section.delay<=0){section.phase+=0.02;section.alpha=Math.sin(section.phase)*1;section.alpha=section.alpha<=0?0:section.alpha;section.alpha=section.alpha>=1?1:section.alpha;if(this._options.animateSections){var mod=Math.sin(1+section.phase*Math.PI/2)*0.1;if(section.dir==='right'){section.point1.add(mod,0);section.point2.add(mod,0);section.point3.add(mod,0)}else{section.point1.subtract(mod,0);section.point2.subtract(mod,0);section.point3.subtract(mod,0)}section.point1.add(0,mod);section.point2.add(0,mod);section.point3.add(0,mod)}}else{section.delay-=0.5}var s=this._options.colorSaturation,l=this._options.colorBrightness,c='hsla('+section.color+', '+s+', '+l+', '+section.alpha+' )';this._context.save();if(this._options.parallaxAmount!==0){this._context.translate(0,this._scroll*this._options.parallaxAmount)}this._context.beginPath();this._context.moveTo(section.point1.x,section.point1.y);this._context.lineTo(section.point2.x,section.point2.y);this._context.lineTo(section.point3.x,section.point3.y);this._context.fillStyle=c;this._context.fill();if(this._options.strokeSize>0){this._context.lineWidth=this._options.strokeSize;this._context.strokeStyle=c;this._context.lineCap='round';this._context.stroke()}this._context.restore()}return false;},_onDraw:function(){for(var i=0,t=this._ribbons.length;i<t;i+=1){if(!this._ribbons[i]){this._ribbons.splice(i,1)}}this._context.clearRect(0,0,this._width,this._height);for(var a=0;a<this._ribbons.length;++a ){var ribbon=this._ribbons[a],numSections=ribbon.length,numDone=0;for(var b=0;b<numSections;++b ){if(this._drawRibbonSection(ribbon[b])){numDone++;}}if(numDone>=numSections){this._ribbons[a]=null}}if(this._ribbons.length<this._options.ribbonCount){this.addRibbon()}requestAnimationFrame(this._onDraw)},_onResize:function(e){var screen=screenInfo(e);this._width=screen.width;this._height=screen.height;if(this._canvas){this._canvas.width=this._width;this._canvas.height=this._height;if(this._context){this._context.globalAlpha=this._options.colorAlpha}}},_onScroll:function(e){var screen=screenInfo(e);this._scroll=screen.scrolly}};return Factory});new Ribbons({colorSaturation:'60%',colorBrightness:'50%',colorAlpha:0.5,colorCycleSpeed:5,verticalPosition:'random',horizontalSpeed:200,ribbonCount:3,strokeSize:0,parallaxAmount:-0.2,animateSections:true})</script> ;`
var Standard_LinksList = `
<body>
    <div class="sidebar">
        <div class="logo-details"><i class='bx bxs-injection'></i><span class="logo_name">Frizz</span></div>
        <ul class=nav-links>
            <li><a href=/ class=active><i class="bx bx-grid-alt"></i><span class=links_name>Analytics</span></a>
                <li><a href=/ParseNew.html><i class="bx bx-rocket"></i><span class=links_name>Parse New</span></a>
                    <li><a href=/Useragents.html><i class="bx bx-user-circle"></i><span class=links_name>HTTP Useragents</span></a>
                        <li><a href=/Hostnames.html><i class="bx bxs-ghost"></i><span class=links_name>HTTP Hostnames</span></a>
                            <li><a href=/URLs.html><i class="bx bxl-sketch"></i><span class=links_name>HTTP URLs</span></a>
                                <li><a href=/HTTPSESSION.html><i class="bx bxs-business"></i><span class=links_name>HTTP General</span></a>
                                    
                                        <li><a href=/OpenPorts.html><i class="bx bx-fingerprint"></i><span class=links_name>Open Ports</span></a>
                                            <li><a href=/ARP.html><i class="bx bx-broadcast"></i><span class=links_name>ARP</span></a>
                                                <li><a href=/Ethernet.html><i class="bx bx-wifi-1"></i><span class=links_name>Ethernet</span></a>
                                                    <li><a href=/Servers.html><i class="bx bx-server"></i><span class=links_name>Servers</span></a>
                                                        <li><a href=/Wifi.html><i class="bx bx-wifi"></i><span class=links_name>Wifi</span></a>
                                                            <li><a href=/WifiOspf.html><i class="bx bx-wifi"></i><span class=links_name>Wifi Warnings</span></a>
                                                                <li><a href=/FTP.html><i class="bx bx-folder"></i><span class=links_name>FTP</span></a>
                                                                    <li><a href=/SSH.html><i class="bx bx-terminal"></i><span class=links_name>SSH</span></a>
                                                                        <li><a href=/SMTP.html><i class="bx bx-envelope"></i><span class=links_name>SMTP</span></a>
                                                                            <li><a href=/Telnet.html><i class="bx bx-desktop"></i><span class=links_name>Telnet</span></a>
                                                                                <li><a href=/SIP.html><i class="bx bx-phone-incoming"></i><span class=links_name>SIP Invites</span></a>
                                                                                    <li><a href=/AuthFTPCreds.html><i class="bx bx-dialpad"></i><span class=links_name>FTP Credentials</span></a>
                                                                                        <li><a href=/AuthSSHCreds.html><i class="bx bxs-key"></i><span class=links_name>SSH Credentials</span></a>
                                                                                            <li><a href=/AuthIMAP.html><i class="bx bxs-lock"></i><span class=links_name>IMAP Credentials</span></a>
                                                                                                <li><a href=/AuthTelnet.html><i class="bx bx-laptop"></i><span class=links_name>Telnet Credentials</span></a>
                                                                                                    <li><a href=/AuthDigest.html><i class="bx bxs-user-pin"></i><span class=links_name>HTTP Digest</span></a>
                                                                                                        <li><a href=/AuthNTLM.html><i class="bx bx-coffee"></i><span class=links_name>HTTP NTLM</span></a>
                                                                                                            <li><a href=/AuthBASIC.html><i class="bx bxs-contact"></i><span class=links_name>HTTP BASIC</span></a>
                                                                                                                <li><a href=/AuthNegotiation.html><i class="bx bx-share-alt"></i><span class=links_name>HTTP Negotiate</span></a>
                                                                                                                    <li><a href=/AuthSMTP.html><i class="bx bx-envelope"></i><span class=links_name>SMTP Credentials</span></a>
                                                                                                                        <li><a href=/Emails.html><i class="bx bx-shape-triangle"></i><span class=links_name>Found Emails</span></a>
                                                                                                                            <li><a href=/Cc.html><i class="bx bxs-chat"></i><span class=links_name>POP3 Cc payload</span></a>
                                                                                                                                <li><a href=/From.html><i class="bx bx-comment-dots"></i><span class=links_name>POP3 From payload</span></a>
                                                                                                                                    <li><a href=/Recv.html><i class="bx bx-mail-send"></i><span class=links_name>POP3 Recv payload</span></a>
                                                                                                                                        <li><a href=/POP3><i class="bx bx-conversation"></i><span class=links_name>[Beta] Conversation</span></a>
                                                                                                                                            <li><a href=/Masher.html><i class="bx bx-meteor"></i><span class=links_name>Packet masher</span></a>
                                                                                                                                                <li><a href=/Raw.html><i class="bx bx-meteor"></i><span class=links_name>Packets RAW</span></a>
                                                                                                                                                    <li><a href=/Extractor.html><i class="bx bxl-google-cloud"></i><span class=links_name>Packet Extractor</span></a>
                                                                                                                                                        <li><a href=/ServerRequirements.html><i class="bx bx-cctv"></i><span class=links_name>Info this server needs</span></a>
                                                                                                                                                            <li><a href=/JSONDB.html><i class="bx bxs-file-json"></i><span class=links_name>JSON Server file</span></a>
                                                                                                                                                                <li><a href=/AppInfo.html><i class="bx bx-landscape"></i><span class=links_name>Application information</span></a>
                                                                                                                                                                    <li><a href=/ServerInfo.html><i class="bx bxs-component"></i><span class=links_name>Server information</span></a>
                                                                                                                                                                        <li><a href=/Documentation.html><i class="bx bxs-book-content"></i><span class=links_name>Documentation</span></a>
                                                                                                                                                                            <li><a href=https://discord.gg/5WfgbMdfWp><i class="bx bxl-discord-alt"></i><span class=links_name>Discord</span></a>
                                                                                                                                                                                <li><a href=https://account.venmo.com/u/Scare-Security-Development><i class="bx bxl-venmo"></i><span class=links_name>Donate [Venmo]</span></a>
                                                                                                                                                                                    <li><a href=https://cash.app/$TotallyNotAHaxxer><i class="bx bx-money"></i><span class=links_name>Donate [CashApp]</span></a>
                                                                                                                                                                                        <li><a href=https://www.medium.com/@Totally_Not_A_Haxxer><i class="bx bxl-medium-square"></i><span class=links_name>Medium Articles</span></a>
                                                                                                                                                                                            <li><a href=https://www.github.com/ArkAngeL43><i class="bx bxl-git"></i><span class=links_name>Github</span></a>
                                                                                                                                                                                                <li><a href=https://www.github.com/orgs/Scare-Security><i class="bx bxl-github"></i><span class=links_name>Github Organization</span></a>
                                                                                                                                                                                                    <li><a href=https://www.instagram.com/Totally_Not_A_Haxxer><i class="bx bxl-instagram"></i><span class=links_name>Instagram</span></a>
                                                                                                                                                                                                        <li><a href=https://twitter.com/NotAHaxxor><i class="bx bxl-twitter"></i><span class=links_name>Twitter</span></a></ul>
    </div>

`
var StaticSection = `<section class="home-section"><nav><div class="sidebar-button"><i class='bx bx-menu sidebarBtn'></i><span class="dashboard">Dashboard</span></div></nav>`
var StaticStyle = `<style>@import url(https://fonts.googleapis.com/css?family=Roboto:500); .container_Overview th h1{font-weight: 700; font-size: 1em; text-align: left; color: #9d00ff;}.container_Overview td{font-weight: 400; font-size: 1em; -webkit-box-shadow: 0 2px 2px -2px #0e1119; -moz-box-shadow: 0 2px 2px -2px #0e1119; box-shadow: 0 2px 2px -2px #0e1119;}.container_Overview{text-align: left; overflow: scroll; width: 100%; margin: 0 0px; display: table; padding: 0 0 8em;}.container_Overview td,.container_Overview th{padding-bottom: 2%; padding-top: 2%; padding-left: 2%; max-width: 70px; color: #e05260;word-wrap: break-word;}.container_Overview th{background-color: #090a0b;}.container_Overview td:first-child{color: #9d00ff;}.container_Overview tr:hover{background-color: #222222; -webkit-box-shadow: 0 6px 6px -6px #0e1119; -moz-box-shadow: 0 6px 6px -6px #0e1119; box-shadow: 0 6px 6px -6px #0e1119;}</style><!--JAVA SXCRIPT BAR CHART CDN AND SRC --></html><style>/* Googlefont Poppins CDN Link */ @import url('https://fonts.googleapis.com/css2?family=Poppins:wght@200;300;400;500;600;700&display=swap'); body{background-color: black; background-image: radial-gradient(circle, rgba(0, 0, 0, 0) 0%, rgba(0, 0, 0, 0.8) 100%); background-position: center center; background-repeat: no-repeat; background-attachment: fixed; background-size: cover;}*{margin: 0; padding: 0; box-sizing: border-box; font-family: 'Poppins', sans-serif;}.sidebar{position: fixed; height: 100%; width: 240px; transition: all 0.5s ease; overflow-y: scroll;}.sidebar.active{width: 60px;}.sidebar .logo-details{height: 80px; display: flex; align-items: center;}.sidebar .logo-details i{font-size: 28px; font-weight: 500; color: #fff; min-width: 60px; text-align: center}.sidebar .logo-details .logo_name{color: #fff; font-size: 24px; font-weight: 500;}.sidebar .nav-links{margin-top: 10px;}.sidebar .nav-links li{position: relative; list-style: none; height: 50px;}.sidebar .nav-links li a{height: 100%; width: 100%; display: flex; align-items: center; text-decoration: none; transition: all 0.4s ease;}.sidebar .nav-links li a.active{background: blueviolet;}.sidebar .nav-links li a:hover{background: red;}.sidebar .nav-links li i{min-width: 60px; text-align: center; font-size: 18px; color: #fff;}.sidebar .nav-links li a .links_name{color: #fff; font-size: 15px; font-weight: 400; white-space: nowrap;}.sidebar .nav-links .log_out{position: absolute; bottom: 0; width: 100%;}.home-section{position: relative; min-height: 100vh; width: calc(100% - 240px); left: 240px; transition: all 0.5s ease;}.sidebar.active~.home-section{width: calc(100% - 60px); left: 60px;}.home-section nav{display: flex; justify-content: space-between; height: 80px; display: flex; align-items: center; position: fixed; width: calc(100% - 240px); left: 240px; z-index: 100; padding: 0 20px; box-shadow: 0 1px 1px rgba(0, 0, 0, 0.1); color: white; transition: all 0.5s ease;}.sidebar.active~.home-section nav{left: 60px; width: calc(100% - 60px);}.home-section nav .sidebar-button{display: flex; align-items: center; font-size: 24px; font-weight: 500;}.ChartCanvasBar_HTTP{float: right; left: 200px; margin-top: 100px; max-width: 600px;}nav .sidebar-button i{font-size: 35px; margin-right: 10px;}.home-section .home-content{position: relative; padding-top: 104px;}.home-content .overview-boxes{display: flex; align-items: center; justify-content: space-between; flex-wrap: wrap; padding: 0 20px; margin-bottom: 26px;}/* END OF SECTION HOME BOXES */ /* Responsive Media Query */ @media (max-width: 1240px){.sidebar{width: 60px;}.sidebar.active{width: 220px;}.home-section{width: calc(100% - 60px); left: 60px;}.sidebar.active~.home-section{/* width: calc(100% - 220px); */ overflow: hidden; left: 220px;}.home-section nav{width: calc(100% - 60px); left: 60px;}.sidebar.active~.home-section nav{width: calc(100% - 220px); left: 220px;}}@media (max-width: 1150px){.home-content .table-info{flex-direction: column;}.home-content .table-info .box{width: 100%; overflow-x: scroll; margin-bottom: 30px;}.home-content .table-info .top-sales{margin: 0;}}@media (max-width: 1000px){.overview-boxes .box{width: calc(100% / 2 - 15px); margin-bottom: 15px;}}@media (max-width: 700px){nav .sidebar-button .dashboard, nav .profile-details .admin_name, nav .profile-details i{display: none;}.home-section nav .profile-details{height: 50px; min-width: 40px;}.home-content .table-info .sales-details{width: 560px;}}@media (max-width: 550px){.overview-boxes .box{width: 100%; margin-bottom: 15px;}.sidebar.active~.home-section nav .profile-details{display: none;}}@media (max-width: 400px){.sidebar{width: 0;}.sidebar.active{width: 60px;}.home-section{width: 100%; left: 0;}.sidebar.active~.home-section{left: 60px; width: calc(100% - 60px);}.home-section nav{width: 100%; left: 0;}.sidebar.active~.home-section nav{left: 60px; width: calc(100% - 60px);}}.overview-boxes .box{display: flex; align-items: center; justify-content: center; width: calc(100% / 4 - 15px); background: white; padding: 15px 14px; border-radius: 12px; box-shadow: 0 5px 10px rgba(0, 0, 0, 0.1);}.overview-boxes .box-topic a{color: blueviolet; text-decoration: none; font-size: 20px; font-weight: 500;}.overview-boxes .box-topic2 a{color: #e05260; text-decoration: none; font-size: 20px; font-weight: 500;}</style>`

// Static data for example of session HTML documents | We can use the same StaticJS and Standard_LinksList for this file as well
var StandardTopSessions = `
<!DOCTYPE html>
<html lang="en" dir="ltr">

<head>
    <title>Frizz | Session Information</title>
    <meta charset="UTF-8">
    <link rel="stylesheet" href="style.css">
    <link href='https://unpkg.com/boxicons@2.1.2/css/boxicons.min.css' rel='stylesheet'>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<source src="../HTML/LobbyMisc/Lobby_Music" type="audio/mpeg">
<source src="../HTML/Future/Future_Lobby" type="audio/mpeg">
`
var StandardStyleSheet = `<style>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@200;300;400;500;600;700&display=swap');

#template {
	display: none
}

.num {
	font-weight: 600;
	color: #d0782a
}

[class*=var] {
	font-weight: 600
}

.clr {
	border-radius: 1px;
	color: white;
	font-weight: 400
}

.str,
.str i {
	color: #ddca7e !important;
	font-weight: 400 !important
}

.reg,
.reg i {
	color: #ddca7e !important;
	font-weight: 500 !important
}

.cmnt,
.cmnt i {
	color: #555 !important;
	font-weight: 400 !important
}

@media (max-width: 400px) {
	.syntax {
		font-size: 12px
	}
}

.container_Overview th h1 {
	font-weight: 700;
	font-size: 1em;
	text-align: left;
	color: #9d00ff
}

.container_Overview td {
	font-weight: 400;
	font-size: 1em;
	-webkit-box-shadow: 0 2px 2px -2px #0e1119;
	-moz-box-shadow: 0 2px 2px -2px #0e1119;
	box-shadow: 0 2px 2px -2px #0e1119
}

.container_Overview {
	text-align: left;
	overflow: scroll;
	width: 100%;
	margin: 0;
	display: table;
	padding: 0 0 8em
}

.container_Overview td,
.container_Overview th {
	padding-bottom: 2%;
	padding-top: 2%;
	padding-left: 2%;
	max-width: 70px;
	color: #e05260
}

.container_Overview th {
	background-color: #090a0b
}

.container_Overview td:first-child {
	color: #9d00ff
}

.container_Overview tr:hover {
	background-color: #222222;
	-webkit-box-shadow: 0 6px 6px -6px #0e1119;
	-moz-box-shadow: 0 6px 6px -6px #0e1119;
	box-shadow: 0 6px 6px -6px #0e1119
}

body {
	background-color: black;
	background-image: radial-gradient(circle, rgba(0, 0, 0, 0) 0%, rgba(0, 0, 0, 0.8) 100%);
	background-position: center center;
	background-repeat: no-repeat;
	background-attachment: fixed;
	background-size: cover
}

* {
	margin: 0;
	padding: 0;
	box-sizing: border-box;
	font-family: 'Poppins', sans-serif
}

.sidebar {
	position: fixed;
	height: 100%;
	width: 240px;
	transition: all 0.5s ease;
	overflow-y: scroll
}

.sidebar.active {
	width: 60px
}

.sidebar .logo-details {
	height: 80px;
	display: flex;
	align-items: center
}

.sidebar .logo-details i {
	font-size: 28px;
	font-weight: 500;
	color: #fff;
	min-width: 60px;
	text-align: center
}

.sidebar .logo-details .logo_name {
	color: #fff;
	font-size: 24px;
	font-weight: 500
}

.sidebar .nav-links {
	margin-top: 10px
}

.sidebar .nav-links li {
	position: relative;
	list-style: none;
	height: 50px
}

.sidebar .nav-links li a {
	height: 100%;
	width: 100%;
	display: flex;
	align-items: center;
	text-decoration: none;
	transition: all 0.4s ease
}

.sidebar .nav-links li a.active {
	background: blueviolet
}

.sidebar .nav-links li a:hover {
	background: red
}

.sidebar .nav-links li i {
	min-width: 60px;
	text-align: center;
	font-size: 18px;
	color: #fff
}

.sidebar .nav-links li a .links_name {
	color: #fff;
	font-size: 15px;
	font-weight: 400;
	white-space: nowrap
}

.sidebar .nav-links .log_out {
	position: absolute;
	bottom: 0;
	width: 100%
}

.home-section {
	position: relative;
	min-height: 100vh;
	width: calc(100% - 240px);
	left: 240px;
	transition: all 0.5s ease
}

.sidebar.active~.home-section {
	width: calc(100% - 60px);
	left: 60px
}

.home-section nav {
	display: flex;
	justify-content: space-between;
	height: 80px;
	display: flex;
	align-items: center;
	position: fixed;
	width: calc(100% - 240px);
	left: 240px;
	z-index: 100;
	padding: 0 20px;
	box-shadow: 0 1px 1px rgba(0, 0, 0, 0.1);
	color: white;
	transition: all 0.5s ease
}

.sidebar.active~.home-section nav {
	left: 60px;
	width: calc(100% - 60px)
}

.home-section nav .sidebar-button {
	display: flex;
	align-items: center;
	font-size: 24px;
	font-weight: 500
}

nav .sidebar-button i {
	font-size: 35px;
	margin-right: 10px
}

.home-section .home-content {
	position: relative;
	padding-top: 104px
}

.home-content .overview-boxes {
	display: flex;
	align-items: center;
	justify-content: space-between;
	flex-wrap: wrap;
	padding: 0 20px;
	margin-bottom: 26px
}

.overview-boxes .box {
	display: flex;
	align-items: center;
	justify-content: center;
	width: calc(100% / 4 - 15px);
	background: white;
	padding: 15px 14px;
	border-radius: 12px;
	box-shadow: 0 5px 10px rgba(0, 0, 0, 0.1)
}

.overview-boxes .box2 {
	margin-top: 50px;
	display: flex;
	align-items: center;
	justify-content: center;
	width: calc(100% / 4 - 15px);
	padding: 15px 14px;
	border-radius: 12px;
	box-shadow: 0 5px 10px rgba(0, 0, 0, 0.1)
}

.overview-boxes .box-topic a {
	color: blueviolet;
	text-decoration: none;
	font-size: 20px;
	font-weight: 500
}

.overview-boxes .box-topic2 a {
	color: #e05260;
	text-decoration: none;
	font-size: 20px;
	font-weight: 500
}

.home-content .box .number {
	display: inline-block;
	font-size: 35px;
	margin-top: -6px;
	font-weight: 500
}

.home-content .box2 .number2 {
	display: inline-block;
	font-size: 35px;
	margin-top: -6px;
	font-weight: 500
}

.home-content .overview-boxes {
	display: flex;
	align-items: center;
	justify-content: space-between;
	flex-wrap: wrap;
	padding: 0 20px;
	margin-bottom: 26px
}

.overview-boxes .box-topic {
	font-size: 20px;
	font-weight: 500
}

.home-content .box .number {
	display: inline-block;
	font-size: 35px;
	margin-top: -6px;
	font-weight: 500
}

.home-content .sales-boxes {
	display: flex;
	justify-content: space-between
}

@media (max-width: 1240px) {
	.sidebar {
		width: 60px
	}

	.sidebar.active {
		width: 220px
	}

	.home-section {
		width: calc(100% - 60px);
		left: 60px
	}

	.sidebar.active~.home-section {
		overflow: hidden;
		left: 220px
	}

	.home-section nav {
		width: calc(100% - 60px);
		left: 60px
	}

	.sidebar.active~.home-section nav {
		width: calc(100% - 220px);
		left: 220px
	}
}

@media (max-width: 1150px) {
	.home-content .sales-boxes {
		flex-direction: column
	}

	.home-content .sales-boxes .box {
		width: 100%;
		overflow-x: scroll;
		margin-bottom: 30px
	}

	.home-content .sales-boxes .top-sales {
		margin: 0
	}
}

@media (max-width: 1000px) {
	.overview-boxes .box {
		width: calc(100% / 2 - 15px);
		margin-bottom: 15px
	}
}

@media (max-width: 700px) {

	nav .profile-details .admin_name,
	nav .profile-details i,
	nav .sidebar-button .dashboard {
		display: none
	}

	.home-section nav .profile-details {
		height: 50px;
		min-width: 40px
	}

	.home-content .sales-boxes .sales-details {
		width: 560px
	}
}

@media (max-width: 550px) {
	.overview-boxes .box {
		width: 100%;
		margin-bottom: 15px
	}

	.sidebar.active~.home-section nav .profile-details {
		display: none
	}
}

@media (max-width: 400px) {
	.sidebar {
		width: 0
	}

	.sidebar.active {
		width: 60px
	}

	.home-section {
		width: 100%;
		left: 0
	}

	.sidebar.active~.home-section {
		left: 60px;
		width: calc(100% - 60px)
	}

	.home-section nav {
		width: 100%;
		left: 0
	}

	.sidebar.active~.home-section nav {
		left: 60px;
		width: calc(100% - 60px)
	}
}
</style>`

var WIFIHTMLTOP = `
<!DOCTYPE html>
<html lang="en" dir="ltr">

<head>
    <title>Frizz | Session Information</title>
    <meta charset="UTF-8">
    <link rel="stylesheet" href="style.css">
    <link href='https://unpkg.com/boxicons@2.1.2/css/boxicons.min.css' rel='stylesheet'>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<source src="../HTML/LobbyMisc/Lobby_Music" type="audio/mpeg">
<source src="../HTML/Future/Future_Lobby" type="audio/mpeg">

<body>
    <div class="sidebar">
        <div class="logo-details"><i class='bx bxs-injection'></i><span class="logo_name">Frizz</span></div>
        <ul class=nav-links>
            <li><a href=/ class=active><i class="bx bx-grid-alt"></i><span class=links_name>Analytics</span></a>
            <li><a href=/ParseNew.html><i class="bx bx-rocket"></i><span class=links_name>Parse New</span></a>
            <li><a href=/Useragents.html><i class="bx bx-user-circle"></i><span class=links_name>HTTP
                        Useragents</span></a>
            <li><a href=/Hostnames.html><i class="bx bxs-ghost"></i><span class=links_name>HTTP Hostnames</span></a>
            <li><a href=/URLs.html><i class="bx bxl-sketch"></i><span class=links_name>HTTP URLs</span></a>
            <li><a href=/HTTPSESSION.html><i class="bx bxs-business"></i><span class=links_name>HTTP General</span></a>
            
            <li><a href=/OpenPorts.html><i class="bx bx-fingerprint"></i><span class=links_name>Open Ports</span></a>
            <li><a href=/ARP.html><i class="bx bx-broadcast"></i><span class=links_name>ARP</span></a>
            <li><a href=/Ethernet.html><i class="bx bx-wifi-1"></i><span class=links_name>Ethernet</span></a>
            <li><a href=/Servers.html><i class="bx bx-server"></i><span class=links_name>Servers</span></a>
            <li><a href=/Wifi.html><i class="bx bx-wifi"></i><span class=links_name>Wifi</span></a>
            <li><a href=/WifiOspf.html><i class="bx bx-wifi"></i><span class=links_name>Wifi Warnings</span></a>
            <li><a href=/FTP.html><i class="bx bx-folder"></i><span class=links_name>FTP</span></a>
            <li><a href=/SSH.html><i class="bx bx-terminal"></i><span class=links_name>SSH</span></a>
            <li><a href=/SMTP.html><i class="bx bx-envelope"></i><span class=links_name>SMTP</span></a>
            <li><a href=/Telnet.html><i class="bx bx-desktop"></i><span class=links_name>Telnet</span></a>
            <li><a href=/SIP.html><i class="bx bx-phone-incoming"></i><span class=links_name>SIP Invites</span></a>
            <li><a href=/AuthFTPCreds.html><i class="bx bx-dialpad"></i><span class=links_name>FTP
                        Credentials</span></a>
            <li><a href=/AuthSSHCreds.html><i class="bx bxs-key"></i><span class=links_name>SSH Credentials</span></a>
            <li><a href=/AuthIMAP.html><i class="bx bxs-lock"></i><span class=links_name>IMAP Credentials</span></a>
            <li><a href=/AuthTelnet.html><i class="bx bx-laptop"></i><span class=links_name>Telnet
                        Credentials</span></a>
            <li><a href=/AuthDigest.html><i class="bx bxs-user-pin"></i><span class=links_name>HTTP Digest</span></a>
            <li><a href=/AuthNTLM.html><i class="bx bx-coffee"></i><span class=links_name>HTTP NTLM</span></a>
            <li><a href=/AuthBASIC.html><i class="bx bxs-contact"></i><span class=links_name>HTTP BASIC</span></a>
            <li><a href=/AuthNegotiation.html><i class="bx bx-share-alt"></i><span class=links_name>HTTP
                        Negotiate</span></a>
            <li><a href=/AuthSMTP.html><i class="bx bx-envelope"></i><span class=links_name>SMTP Credentials</span></a>
            <li><a href=/Emails.html><i class="bx bx-shape-triangle"></i><span class=links_name>Found Emails</span></a>
            <li><a href=/Cc.html><i class="bx bxs-chat"></i><span class=links_name>POP3 Cc payload</span></a>
            <li><a href=/From.html><i class="bx bx-comment-dots"></i><span class=links_name>POP3 From payload</span></a>
            <li><a href=/Recv.html><i class="bx bx-mail-send"></i><span class=links_name>POP3 Recv payload</span></a>
            <li><a href=/POP3><i class="bx bx-conversation"></i><span class=links_name>[Beta] Conversation</span></a>
            <li><a href=/Masher.html><i class="bx bx-meteor"></i><span class=links_name>Packet masher</span></a>
            <li><a href=/Raw.html><i class="bx bx-meteor"></i><span class=links_name>Packets RAW</span></a>
            <li><a href=/Extractor.html><i class="bx bxl-google-cloud"></i><span class=links_name>Packet
                        Extractor</span></a>
            <li><a href=/ServerRequirements.html><i class="bx bx-cctv"></i><span class=links_name>Info this server
                        needs</span></a>
            <li><a href=/JSONDB.html><i class="bx bxs-file-json"></i><span class=links_name>JSON Server file</span></a>
            <li><a href=/AppInfo.html><i class="bx bx-landscape"></i><span class=links_name>Application
                        information</span></a>
            <li><a href=/ServerInfo.html><i class="bx bxs-component"></i><span class=links_name>Server
                        information</span></a>
            <li><a href=/Documentation.html><i class="bx bxs-book-content"></i><span
                        class=links_name>Documentation</span></a>
            <li><a href=https://discord.gg/5WfgbMdfWp><i class="bx bxl-discord-alt"></i><span
                        class=links_name>Discord</span></a>
            <li><a href=https://account.venmo.com/u/Scare-Security-Development><i class="bx bxl-venmo"></i><span
                        class=links_name>Donate [Venmo]</span></a>
            <li><a href=https://cash.app/$TotallyNotAHaxxer><i class="bx bx-money"></i><span class=links_name>Donate
                        [CashApp]</span></a>
            <li><a href=https://www.medium.com/@Totally_Not_A_Haxxer><i class="bx bxl-medium-square"></i><span
                        class=links_name>Medium Articles</span></a>
            <li><a href=https://www.github.com/ArkAngeL43><i class="bx bxl-git"></i><span
                        class=links_name>Github</span></a>
            <li><a href=https://www.github.com/orgs/Scare-Security><i class="bx bxl-github"></i><span
                        class=links_name>Github Organization</span></a>
            <li><a href=https://www.instagram.com/Totally_Not_A_Haxxer><i class="bx bxl-instagram"></i><span
                        class=links_name>Instagram</span></a>
            <li><a href=https://twitter.com/NotAHaxxor><i class="bx bxl-twitter"></i><span
                        class=links_name>Twitter</span></a>
        </ul>
    </div>

    <section class="home-section">
        <nav>
            <div class="sidebar-button"><i class='bx bx-menu sidebarBtn'></i><span class="dashboard">Dashboard</span>
            </div>
        </nav>
`

var WIFIHTMLSTYLE = `
<style>
                @import url('https://fonts.googleapis.com/css2?family=Poppins:wght@200;300;400;500;600;700&display=swap');


                .codeheader {
                    padding: 5px 5px 5px 10px;
                    font-family: 'Roboto', sans-serif;
                    font-size: 1.1em;
                    color: #fff;
                    -webkit-border-radius: 6px 6px 0 0;
                    -moz-border-radius: 6px 6px 0 0;
                    border-radius: 6px 6px 0 0;
                    user-select: none;
                    -ms-user-select: none;
                    -moz-user-select: none;
                    -webkit-user-select: none;
                    margin-bottom: -10px;
                    margin-left: 10px;
                    width: 99%;
                }


                #Topic {
                    color: black;
                    background-color: #e05260
                }

                @media (max-width: 400px) {
                    .syntax {
                        font-size: 12px
                    }
                }

                .container_Overview th h1 {
                    font-weight: 700;
                    font-size: 1em;
                    text-align: left;
                    color: #9d00ff
                }

                .container_Overview td {
                    font-weight: 400;
                    font-size: 1em;
                    -webkit-box-shadow: 0 2px 2px -2px #0e1119;
                    -moz-box-shadow: 0 2px 2px -2px #0e1119;
                    box-shadow: 0 2px 2px -2px #0e1119
                }

                .container_Overview {
                    text-align: left;
                    overflow: scroll;
                    width: 100%;
                    margin: 0;
                    display: table;
                    padding: 0 0 8em;
                    margin-top: -30px;
                }

                .container_Overview td,
                .container_Overview th {
                    padding-bottom: 2%;
                    padding-top: 2%;
                    padding-left: 2%;
                    max-width: 70px;
                    color: #e05260
                }

                .container_Overview th {
                    background-color: #090a0b
                }

                .container_Overview td:first-child {
                    color: #9d00ff
                }

                .container_Overview tr:hover {
                    background-color: #222222;
                    -webkit-box-shadow: 0 6px 6px -6px #0e1119;
                    -moz-box-shadow: 0 6px 6px -6px #0e1119;
                    box-shadow: 0 6px 6px -6px #0e1119
                }

                body {
                    background-color: black;
                    background-image: radial-gradient(circle, rgba(0, 0, 0, 0) 0%, rgba(0, 0, 0, 0.8) 100%);
                    background-position: center center;
                    background-repeat: no-repeat;
                    background-attachment: fixed;
                    background-size: cover
                }

                * {
                    margin: 0;
                    padding: 0;
                    box-sizing: border-box;
                    font-family: 'Poppins', sans-serif
                }

                .sidebar {
                    position: fixed;
                    height: 100%;
                    width: 240px;
                    transition: all 0.5s ease;
                    overflow-y: scroll
                }

                .sidebar.active {
                    width: 60px
                }

                .sidebar .logo-details {
                    height: 80px;
                    display: flex;
                    align-items: center
                }

                .sidebar .logo-details i {
                    font-size: 28px;
                    font-weight: 500;
                    color: #fff;
                    min-width: 60px;
                    text-align: center
                }

                .sidebar .logo-details .logo_name {
                    color: #fff;
                    font-size: 24px;
                    font-weight: 500
                }

                .sidebar .nav-links {
                    margin-top: 10px
                }

                .sidebar .nav-links li {
                    position: relative;
                    list-style: none;
                    height: 50px
                }

                .sidebar .nav-links li a {
                    height: 100%;
                    width: 100%;
                    display: flex;
                    align-items: center;
                    text-decoration: none;
                    transition: all 0.4s ease
                }

                .sidebar .nav-links li a.active {
                    background: blueviolet
                }

                .sidebar .nav-links li a:hover {
                    background: red
                }

                .sidebar .nav-links li i {
                    min-width: 60px;
                    text-align: center;
                    font-size: 18px;
                    color: #fff
                }

                .sidebar .nav-links li a .links_name {
                    color: #fff;
                    font-size: 15px;
                    font-weight: 400;
                    white-space: nowrap
                }

                .sidebar .nav-links .log_out {
                    position: absolute;
                    bottom: 0;
                    width: 100%
                }

                .home-section {
                    position: relative;
                    min-height: 100vh;
                    width: calc(100% - 240px);
                    left: 240px;
                    transition: all 0.5s ease
                }

                .sidebar.active~.home-section {
                    width: calc(100% - 60px);
                    left: 60px
                }

                .home-section nav {
                    display: flex;
                    justify-content: space-between;
                    height: 80px;
                    display: flex;
                    align-items: center;
                    position: fixed;
                    width: calc(100% - 240px);
                    left: 240px;
                    z-index: 100;
                    padding: 0 20px;
                    box-shadow: 0 1px 1px rgba(0, 0, 0, 0.1);
                    color: white;
                    transition: all 0.5s ease
                }

                .sidebar.active~.home-section nav {
                    left: 60px;
                    width: calc(100% - 60px)
                }

                .home-section nav .sidebar-button {
                    display: flex;
                    align-items: center;
                    font-size: 24px;
                    font-weight: 500
                }

                nav .sidebar-button i {
                    font-size: 35px;
                    margin-right: 10px
                }

                .home-section .home-content {
                    position: relative;
                    padding-top: 104px
                }

                .home-content .overview-boxes {
                    display: flex;
                    align-items: center;
                    justify-content: space-between;
                    flex-wrap: wrap;
                    padding: 0 20px;
                    margin-bottom: 26px
                }

                .overview-boxes .box {
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    width: calc(100% / 4 - 15px);
                    background: white;
                    padding: 15px 14px;
                    border-radius: 12px;
                    box-shadow: 0 5px 10px rgba(0, 0, 0, 0.1)
                }

                .overview-boxes .box2 {
                    margin-top: 50px;
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    width: calc(100% / 4 - 15px);
                    padding: 15px 14px;
                    border-radius: 12px;
                    box-shadow: 0 5px 10px rgba(0, 0, 0, 0.1)
                }

                .overview-boxes .box-topic a {
                    color: blueviolet;
                    text-decoration: none;
                    font-size: 20px;
                    font-weight: 500
                }

                .overview-boxes .box-topic2 a {
                    color: #e05260;
                    text-decoration: none;
                    font-size: 20px;
                    font-weight: 500
                }

                .home-content .box .number {
                    display: inline-block;
                    font-size: 35px;
                    margin-top: -6px;
                    font-weight: 500
                }

                .home-content .box2 .number2 {
                    display: inline-block;
                    font-size: 35px;
                    margin-top: -6px;
                    font-weight: 500
                }

                .home-content .overview-boxes {
                    display: flex;
                    align-items: center;
                    justify-content: space-between;
                    flex-wrap: wrap;
                    padding: 0 20px;
                    margin-bottom: 26px
                }

                .overview-boxes .box-topic {
                    font-size: 20px;
                    font-weight: 500
                }

                .home-content .box .number {
                    display: inline-block;
                    font-size: 35px;
                    margin-top: -6px;
                    font-weight: 500
                }

                .home-content .sales-boxes {
                    display: flex;
                    justify-content: space-between
                }

                @media (max-width: 1240px) {
                    .sidebar {
                        width: 60px
                    }

                    .sidebar.active {
                        width: 220px
                    }

                    .home-section {
                        width: calc(100% - 60px);
                        left: 60px
                    }

                    .sidebar.active~.home-section {
                        overflow: hidden;
                        left: 220px
                    }

                    .home-section nav {
                        width: calc(100% - 60px);
                        left: 60px
                    }

                    .sidebar.active~.home-section nav {
                        width: calc(100% - 220px);
                        left: 220px
                    }
                }

                @media (max-width: 1150px) {
                    .home-content .sales-boxes {
                        flex-direction: column
                    }

                    .home-content .sales-boxes .box {
                        width: 100%;
                        overflow-x: scroll;
                        margin-bottom: 30px
                    }

                    .home-content .sales-boxes .top-sales {
                        margin: 0
                    }
                }

                @media (max-width: 1000px) {
                    .overview-boxes .box {
                        width: calc(100% / 2 - 15px);
                        margin-bottom: 15px
                    }
                }

                @media (max-width: 700px) {

                    nav .profile-details .admin_name,
                    nav .profile-details i,
                    nav .sidebar-button .dashboard {
                        display: none
                    }

                    .home-section nav .profile-details {
                        height: 50px;
                        min-width: 40px
                    }

                    .home-content .sales-boxes .sales-details {
                        width: 560px
                    }
                }

                @media (max-width: 550px) {
                    .overview-boxes .box {
                        width: 100%;
                        margin-bottom: 15px
                    }

                    .sidebar.active~.home-section nav .profile-details {
                        display: none
                    }
                }

                @media (max-width: 400px) {
                    .sidebar {
                        width: 0
                    }

                    .sidebar.active {
                        width: 60px
                    }

                    .home-section {
                        width: 100%;
                        left: 0
                    }

                    .sidebar.active~.home-section {
                        left: 60px;
                        width: calc(100% - 60px)
                    }

                    .home-section nav {
                        width: 100%;
                        left: 0
                    }

                    .sidebar.active~.home-section nav {
                        left: 60px;
                        width: calc(100% - 60px)
                    }
                }
            </style>
`

// LAST HTML DOCUMENT YAY!

var DOCTOP = `

<!DOCTYPE html>
<html lang="en" dir="ltr">

<head>
<title>Frizz | Server Information</title>
    <meta charset="UTF-8">
    <link rel="stylesheet" href="style.css">
    <link href='https://unpkg.com/boxicons@2.1.2/css/boxicons.min.css' rel='stylesheet'>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<source src="../HTML/LobbyMisc/Lobby_Music" type="audio/mpeg">
<source src="../HTML/Future/Future_Lobby" type="audio/mpeg">

<body>
    <div class="sidebar">
        <div class="logo-details">
            <i class='bx bxs-injection'></i>
            <span class="logo_name">Frizzed</span>
        </div>
        <ul class="nav-links">
        <li>
                    <a href="/" class="active"><i class="bx bx-grid-alt"></i><span class="links_name">Analytics</span></a>
                </li>
                <li><a href="/ParseNew.html"><i class="bx bx-rocket"></i><span class="links_name">Parse New</span></a></li>
                <li>
                    <a href="/Useragents.html"><i class="bx bx-user-circle"></i><span class="links_name">HTTP Useragents</span></a>
                </li>
                <li>
                    <a href="/Hostnames.html"><i class="bx bxs-ghost"></i><span class="links_name">HTTP Hostnames</span></a>
                </li>
                <li>
                    <a href="/URLs.html"><i class="bx bxl-sketch"></i><span class="links_name">HTTP URLs</span></a>
                </li>
                <li>
                    <a href="/HTTPSESSION.html"><i class="bx bxs-business"></i><span class="links_name">HTTP General</span></a>
                </li>
                <li>
                    <a href="/DNS.html"><i class="bx bx-cabinet"></i><span class="links_name">DNS</span></a>
                </li>
                <li>
                    <a href="/OpenPorts.html"><i class="bx bx-fingerprint"></i><span class="links_name">Open Ports</span></a>
                </li>
                <li>
                    <a href="/ARP.html"><i class="bx bx-broadcast"></i><span class="links_name">ARP</span></a>
                </li>
                <li>
                    <a href="/Ethernet.html"><i class="bx bx-wifi-1"></i><span class="links_name">Ethernet</span></a>
                </li>
                <li>
                    <a href="/Servers.html"><i class="bx bx-server"></i><span class="links_name">Servers</span></a>
                </li>
                <li>
                    <a href="/Wifi.html"><i class="bx bx-wifi"></i><span class="links_name">Wifi</span></a>
                </li>
                <li>
                    <a href="/WifiOspf.html"><i class="bx bx-wifi"></i><span class="links_name">Wifi Warnings</span></a>
                </li>
                <li>
                    <a href="/FTP.html"><i class="bx bx-folder"></i><span class="links_name">FTP</span></a>
                </li>
                <li>
                    <a href="/SSH.html"><i class="bx bx-terminal"></i><span class="links_name">SSH</span></a>
                </li>
                <li>
                    <a href="/SMTP.html"><i class="bx bx-envelope"></i><span class="links_name">SMTP</span></a>
                </li>
                <li>
                    <a href="/Telnet.html"><i class="bx bx-desktop"></i><span class="links_name">Telnet</span></a>
                </li>
                <li>
                    <a href="/SIP.html"><i class="bx bx-phone-incoming"></i><span class="links_name">SIP Invites</span></a>
                </li>
                <li>
                    <a href="/AuthFTPCreds.html"><i class="bx bx-dialpad"></i><span class="links_name">FTP Credentials</span></a>
                </li>
                <li>
                    <a href="/AuthSSHCreds.html"><i class="bx bxs-key"></i><span class="links_name">SSH Credentials</span></a>
                </li>
                <li>
                    <a href="/AuthIMAP.html"><i class="bx bxs-lock"></i><span class="links_name">IMAP Credentials</span></a>
                </li>
                <li><a href=/AuthTelnet.html><i class="bx bx-laptop"></i><span class=links_name>Telnet Credentials</span></a>

                <li>
                    <a href="/AuthDigest.html"><i class="bx bxs-user-pin"></i><span class="links_name">HTTP Digest</span></a>
                </li>
                <li>
                    <a href="/AuthNTLM.html"><i class="bx bx-coffee"></i><span class="links_name">HTTP NTLM</span></a>
                </li>
                <li>
                    <a href="/AuthBASIC.html"><i class="bx bxs-contact"></i><span class="links_name">HTTP BASIC</span></a>
                </li>
                <li>
                    <a href="/AuthNegotiation.html"><i class="bx bx-share-alt"></i><span class="links_name">HTTP Negotiate</span></a>
                </li>
                <li>
                    <a href="/AuthSMTP.html"><i class="bx bx-envelope"></i><span class="links_name">SMTP Credentials</span></a>
                </li>
                <li>
                    <a href="/Emails.html"><i class="bx bx-shape-triangle"></i><span class="links_name">Found Emails</span></a>
                </li>
                <li>
                    <a href="/Cc.html"><i class="bx bxs-chat"></i><span class="links_name">POP3 Cc payload</span></a>
                </li>
                <li>
                    <a href="/From.html"><i class="bx bx-comment-dots"></i><span class="links_name">POP3 From payload</span></a>
                </li>
                <li>
                    <a href="/Recv.html"><i class="bx bx-mail-send"></i><span class="links_name">POP3 Recv payload</span></a>
                </li>
                <li>
                    <a href="/POP3"><i class="bx bx-conversation"></i><span class="links_name">[Beta] Conversation</span></a>
                </li>
                <li>
                    <a href="/Masher.html"><i class="bx bx-meteor"></i><span class="links_name">Packet masher</span></a>
                </li>
                <li>
                    <a href="/Raw.html"><i class="bx bx-meteor"></i><span class="links_name">Packets RAW</span></a>
                </li>
                <li>
                    <a href="/Extractor.html"><i class="bx bxl-google-cloud"></i><span class="links_name">Packet Extractor</span></a>
                </li>
                <li>
                    <a href="/ServerRequirements.html"><i class="bx bx-cctv"></i><span class="links_name">Info this server needs</span></a>
                </li>
                <li>
                    <a href="/JSONDB.html"><i class="bx bxs-file-json"></i><span class="links_name">JSON Server file</span></a>
                </li>
                <li>
                    <a href="/AppInfo.html"><i class="bx bx-landscape"></i><span class="links_name">Application information</span></a>
                </li>
                <li>
                    <a href="/ServerInfo.html"><i class="bx bxs-component"></i><span class="links_name">Server information</span></a>
                </li>
                <li>
                    <a href="/Documentation.html"><i class="bx bxs-book-content"></i><span class="links_name">Documentation</span></a>
                </li>
                <li>
                    <a href="https://discord.gg/5WfgbMdfWp"><i class="bx bxl-discord-alt"></i><span class="links_name">Discord</span></a>
                </li>
                <li>
                    <a href="https://account.venmo.com/u/Scare-Security-Development"><i class="bx bxl-venmo"></i><span class="links_name">Donate [Venmo]</span></a>
                </li>
                <li>
                    <a href="https://cash.app/$TotallyNotAHaxxer"><i class="bx bx-money"></i><span class="links_name">Donate [CashApp]</span></a>
                </li>
                <li>
                    <a href="https://www.medium.com/@Totally_Not_A_Haxxer"><i class="bx bxl-medium-square"></i><span class="links_name">Medium Articles</span></a>
                </li>
                <li>
                    <a href="https://www.github.com/ArkAngeL43"><i class="bx bxl-git"></i><span class="links_name">Github</span></a>
                </li>
                <li>
                    <a href="https://www.github.com/orgs/Scare-Security"><i class="bx bxl-github"></i><span class="links_name">Github Organization</span></a>
                </li>
                <li>
                    <a href="https://www.instagram.com/Totally_Not_A_Haxxer"><i class="bx bxl-instagram"></i><span class="links_name">Instagram</span></a>
                </li>
                <li>
                    <a href="https://twitter.com/NotAHaxxor"><i class="bx bxl-twitter"></i><span class="links_name">Twitter</span></a>
                </li>
        </ul>
    </div>
    <section class="home-section">
        <nav>
            <div class="sidebar-button">
                <i class='bx bx-menu sidebarBtn'></i>
                <span class="dashboard">Dashboard</span>
            </div>
        </nav>
        <div class="home-content">
`

var DOCSTYLE = `
<script type="text/javascript"><script type="text/javascript"></script><script>let sidebar=document.querySelector(".sidebar");let sidebarBtn=document.querySelector(".sidebarBtn");sidebarBtn.onclick=function(){sidebar.classList.toggle("active");if(sidebar.classList.contains("active")){sidebarBtn.classList.replace("bx-menu","bx-menu-alt-right")}else{sidebarBtn.classList.replace("bx-menu-alt-right","bx-menu")}};(function(name,factory){if(typeof window==='object'){window[name]=factory()}})('Ribbons',function(){var _w=window,_b=document.body,_d=document.documentElement;var random=function(){if(arguments.length===1){if(Array.isArray(arguments[0])){var index=Math.round(random(0,arguments[0].length-1));return arguments[0][index]}return random(0,arguments[0]);}else if(arguments.length===2){return Math.random()*(arguments[1]-arguments[0])+arguments[0]}else if(arguments.length===4){var array=[arguments[0],arguments[1],arguments[2],arguments[3]];return array[Math.floor(Math.random()*array.length)];}return 0;};var screenInfo=function(e){var width=Math.max(0,_w.innerWidth||_d.clientWidth||_b.clientWidth||0),height=Math.max(0,_w.innerHeight||_d.clientHeight||_b.clientHeight||0),scrollx=Math.max(0,_w.pageXOffset||_d.scrollLeft||_b.scrollLeft||0)-(_d.clientLeft||0),scrolly=Math.max(0,_w.pageYOffset||_d.scrollTop||_b.scrollTop||0)-(_d.clientTop||0);return{width:width,height:height,ratio:width/height,centerx:width/2,centery:height/2,scrollx:scrollx,scrolly:scrolly}};var mouseInfo=function(e){var screen=screenInfo(e),mousex=e?Math.max(0,e.pageX||e.clientX||0):0,mousey=e?Math.max(0,e.pageY||e.clientY||0):0;return{mousex:mousex,mousey:mousey,centerx:mousex-screen.width/2,centery:mousey-screen.height/2}};var Point=function(x,y){this.x=0;this.y=0;this.set(x,y)};Point.prototype={constructor:Point,set:function(x,y){this.x=x||0;this.y=y||0},copy:function(point){this.x=point.x||0;this.y=point.y||0;return this},multiply:function(x,y){this.x*=x||1;this.y*=y||1;return this},divide:function(x,y){this.x/=x||1;this.y/=y||1;return this},add:function(x,y){this.x+=x||0;this.y+=y||0;return this},subtract:function(x,y){this.x-=x||0;this.y-=y||0;return this},clampX:function(min,max){this.x=Math.max(min,Math.min(this.x,max));return this},clampY:function(min,max){this.y=Math.max(min,Math.min(this.y,max));return this},flipX:function(){this.x*=-1;return this},flipY:function(){this.y*=-1;return this}};var Factory=function(options){this._canvas=null;this._context=null;this._sto=null;this._width=0;this._height=0;this._scroll=0;this._ribbons=[];this._options={colorSaturation:'80%',colorBrightness:'60%',colorAlpha:0.65,colorCycleSpeed:6,verticalPosition:'center',horizontalSpeed:150,ribbonCount:3,strokeSize:0,parallaxAmount:-0.5,animateSections:true};this._onDraw=this._onDraw.bind(this);this._onResize=this._onResize.bind(this);this._onScroll=this._onScroll.bind(this);this.setOptions(options);this.init()};Factory.prototype={constructor:Factory,setOptions:function(options){if(typeof options==='object'){for(var key in options){if(options.hasOwnProperty(key)){this._options[key]=options[key]}}}},init:function(){try{this._canvas=document.createElement('canvas');this._canvas.style['display']='block';this._canvas.style['position']='fixed';this._canvas.style['margin']='0';this._canvas.style['padding']='0';this._canvas.style['border']='0';this._canvas.style['outline']='0';this._canvas.style['left']='0';this._canvas.style['top']='0';this._canvas.style['width']='100%';this._canvas.style['height']='100%';this._canvas.style['z-index']='-1';this._onResize();this._context=this._canvas.getContext('2d');this._context.clearRect(0,0,this._width,this._height);this._context.globalAlpha=this._options.colorAlpha;window.addEventListener('resize',this._onResize);window.addEventListener('scroll',this._onScroll);document.body.appendChild(this._canvas)}catch(e){console.warn('Canvas Context Error: '+e.toString());return}this._onDraw()},addRibbon:function(){var dir=Math.round(random(1,9))>5?'right':'left',stop=1000,hide=200,min=0-hide,max=this._width+hide,movex=0,movey=0,startx=dir==='right'?min:max,starty=Math.round(random(0,this._height));if(/^(top|min)$/i.test(this._options.verticalPosition)){starty=0+hide}else if(/^(middle|center)$/i.test(this._options.verticalPosition)){starty=this._height/2}else if(/^(bottom|max)$/i.test(this._options.verticalPosition)){starty=this._height-hide}var ribbon=[],point1=new Point(startx,starty),point2=new Point(startx,starty),point3=null,color=Math.round(random(900)),delay=0;while(true){if(stop<=0){break}stop--;movex=Math.round((Math.random()*1-0.2)*this._options.horizontalSpeed);movey=Math.round((Math.random()*1-0.5)*(this._height*0.25));point3=new Point();point3.copy(point2);if(dir==='right'){point3.add(movex,movey);if(point2.x>=max){break}}else if(dir==='left'){point3.subtract(movex,movey);if(point2.x<=min){break}}ribbon.push({point1:new Point(point1.x,point1.y),point2:new Point(point2.x,point2.y),point3:point3,color:615,delay:delay,dir:dir,alpha:0,phase:0});point1.copy(point2);point2.copy(point3);delay+=4}this._ribbons.push(ribbon)},_drawRibbonSection:function(section){if(section){if(section.phase>=1&&section.alpha<=0){return true;}if(section.delay<=0){section.phase+=0.02;section.alpha=Math.sin(section.phase)*1;section.alpha=section.alpha<=0?0:section.alpha;section.alpha=section.alpha>=1?1:section.alpha;if(this._options.animateSections){var mod=Math.sin(1+section.phase*Math.PI/2)*0.1;if(section.dir==='right'){section.point1.add(mod,0);section.point2.add(mod,0);section.point3.add(mod,0)}else{section.point1.subtract(mod,0);section.point2.subtract(mod,0);section.point3.subtract(mod,0)}section.point1.add(0,mod);section.point2.add(0,mod);section.point3.add(0,mod)}}else{section.delay-=0.5}var s=this._options.colorSaturation,l=this._options.colorBrightness,c='hsla('+section.color+', '+s+', '+l+', '+section.alpha+' )';this._context.save();if(this._options.parallaxAmount!==0){this._context.translate(0,this._scroll*this._options.parallaxAmount)}this._context.beginPath();this._context.moveTo(section.point1.x,section.point1.y);this._context.lineTo(section.point2.x,section.point2.y);this._context.lineTo(section.point3.x,section.point3.y);this._context.fillStyle=c;this._context.fill();if(this._options.strokeSize>0){this._context.lineWidth=this._options.strokeSize;this._context.strokeStyle=c;this._context.lineCap='round';this._context.stroke()}this._context.restore()}return false;},_onDraw:function(){for(var i=0,t=this._ribbons.length;i<t;i+=1){if(!this._ribbons[i]){this._ribbons.splice(i,1)}}this._context.clearRect(0,0,this._width,this._height);for(var a=0;a<this._ribbons.length;++a ){var ribbon=this._ribbons[a],numSections=ribbon.length,numDone=0;for(var b=0;b<numSections;++b ){if(this._drawRibbonSection(ribbon[b])){numDone++;}}if(numDone>=numSections){this._ribbons[a]=null}}if(this._ribbons.length<this._options.ribbonCount){this.addRibbon()}requestAnimationFrame(this._onDraw)},_onResize:function(e){var screen=screenInfo(e);this._width=screen.width;this._height=screen.height;if(this._canvas){this._canvas.width=this._width;this._canvas.height=this._height;if(this._context){this._context.globalAlpha=this._options.colorAlpha}}},_onScroll:function(e){var screen=screenInfo(e);this._scroll=screen.scrolly}};return Factory});new Ribbons({colorSaturation:'60%',colorBrightness:'50%',colorAlpha:0.5,colorCycleSpeed:5,verticalPosition:'random',horizontalSpeed:200,ribbonCount:3,strokeSize:0,parallaxAmount:-0.2,animateSections:true});</script><script>let sidebar=document.querySelector(".sidebar");let sidebarBtn=document.querySelector(".sidebarBtn");sidebarBtn.onclick=function(){sidebar.classList.toggle("active");if(sidebar.classList.contains("active")){sidebarBtn.classList.replace("bx-menu","bx-menu-alt-right")}else{sidebarBtn.classList.replace("bx-menu-alt-right","bx-menu")}};(function(name,factory){if(typeof window==='object'){window[name]=factory()}})('Ribbons',function(){var _w=window,_b=document.body,_d=document.documentElement;var random=function(){if(arguments.length===1){if(Array.isArray(arguments[0])){var index=Math.round(random(0,arguments[0].length-1));return arguments[0][index]}return random(0,arguments[0]);}else if(arguments.length===2){return Math.random()*(arguments[1]-arguments[0])+arguments[0]}else if(arguments.length===4){var array=[arguments[0],arguments[1],arguments[2],arguments[3]];return array[Math.floor(Math.random()*array.length)];}return 0;};var screenInfo=function(e){var width=Math.max(0,_w.innerWidth||_d.clientWidth||_b.clientWidth||0),height=Math.max(0,_w.innerHeight||_d.clientHeight||_b.clientHeight||0),scrollx=Math.max(0,_w.pageXOffset||_d.scrollLeft||_b.scrollLeft||0)-(_d.clientLeft||0),scrolly=Math.max(0,_w.pageYOffset||_d.scrollTop||_b.scrollTop||0)-(_d.clientTop||0);return{width:width,height:height,ratio:width/height,centerx:width/2,centery:height/2,scrollx:scrollx,scrolly:scrolly}};var mouseInfo=function(e){var screen=screenInfo(e),mousex=e?Math.max(0,e.pageX||e.clientX||0):0,mousey=e?Math.max(0,e.pageY||e.clientY||0):0;return{mousex:mousex,mousey:mousey,centerx:mousex-screen.width/2,centery:mousey-screen.height/2}};var Point=function(x,y){this.x=0;this.y=0;this.set(x,y)};Point.prototype={constructor:Point,set:function(x,y){this.x=x||0;this.y=y||0},copy:function(point){this.x=point.x||0;this.y=point.y||0;return this},multiply:function(x,y){this.x*=x||1;this.y*=y||1;return this},divide:function(x,y){this.x/=x||1;this.y/=y||1;return this},add:function(x,y){this.x+=x||0;this.y+=y||0;return this},subtract:function(x,y){this.x-=x||0;this.y-=y||0;return this},clampX:function(min,max){this.x=Math.max(min,Math.min(this.x,max));return this},clampY:function(min,max){this.y=Math.max(min,Math.min(this.y,max));return this},flipX:function(){this.x*=-1;return this},flipY:function(){this.y*=-1;return this}};var Factory=function(options){this._canvas=null;this._context=null;this._sto=null;this._width=0;this._height=0;this._scroll=0;this._ribbons=[];this._options={colorSaturation:'80%',colorBrightness:'60%',colorAlpha:0.65,colorCycleSpeed:6,verticalPosition:'center',horizontalSpeed:150,ribbonCount:3,strokeSize:0,parallaxAmount:-0.5,animateSections:true};this._onDraw=this._onDraw.bind(this);this._onResize=this._onResize.bind(this);this._onScroll=this._onScroll.bind(this);this.setOptions(options);this.init()};Factory.prototype={constructor:Factory,setOptions:function(options){if(typeof options==='object'){for(var key in options){if(options.hasOwnProperty(key)){this._options[key]=options[key]}}}},init:function(){try{this._canvas=document.createElement('canvas');this._canvas.style['display']='block';this._canvas.style['position']='fixed';this._canvas.style['margin']='0';this._canvas.style['padding']='0';this._canvas.style['border']='0';this._canvas.style['outline']='0';this._canvas.style['left']='0';this._canvas.style['top']='0';this._canvas.style['width']='100%';this._canvas.style['height']='100%';this._canvas.style['z-index']='-1';this._onResize();this._context=this._canvas.getContext('2d');this._context.clearRect(0,0,this._width,this._height);this._context.globalAlpha=this._options.colorAlpha;window.addEventListener('resize',this._onResize);window.addEventListener('scroll',this._onScroll);document.body.appendChild(this._canvas)}catch(e){console.warn('Canvas Context Error: '+e.toString());return}this._onDraw()},addRibbon:function(){var dir=Math.round(random(1,9))>5?'right':'left',stop=1000,hide=200,min=0-hide,max=this._width+hide,movex=0,movey=0,startx=dir==='right'?min:max,starty=Math.round(random(0,this._height));if(/^(top|min)$/i.test(this._options.verticalPosition)){starty=0+hide}else if(/^(middle|center)$/i.test(this._options.verticalPosition)){starty=this._height/2}else if(/^(bottom|max)$/i.test(this._options.verticalPosition)){starty=this._height-hide}var ribbon=[],point1=new Point(startx,starty),point2=new Point(startx,starty),point3=null,color=Math.round(random(900)),delay=0;while(true){if(stop<=0){break}stop--;movex=Math.round((Math.random()*1-0.2)*this._options.horizontalSpeed);movey=Math.round((Math.random()*1-0.5)*(this._height*0.25));point3=new Point();point3.copy(point2);if(dir==='right'){point3.add(movex,movey);if(point2.x>=max){break}}else if(dir==='left'){point3.subtract(movex,movey);if(point2.x<=min){break}}ribbon.push({point1:new Point(point1.x,point1.y),point2:new Point(point2.x,point2.y),point3:point3,color:615,delay:delay,dir:dir,alpha:0,phase:0});point1.copy(point2);point2.copy(point3);delay+=4}this._ribbons.push(ribbon)},_drawRibbonSection:function(section){if(section){if(section.phase>=1&&section.alpha<=0){return true;}if(section.delay<=0){section.phase+=0.02;section.alpha=Math.sin(section.phase)*1;section.alpha=section.alpha<=0?0:section.alpha;section.alpha=section.alpha>=1?1:section.alpha;if(this._options.animateSections){var mod=Math.sin(1+section.phase*Math.PI/2)*0.1;if(section.dir==='right'){section.point1.add(mod,0);section.point2.add(mod,0);section.point3.add(mod,0)}else{section.point1.subtract(mod,0);section.point2.subtract(mod,0);section.point3.subtract(mod,0)}section.point1.add(0,mod);section.point2.add(0,mod);section.point3.add(0,mod)}}else{section.delay-=0.5}var s=this._options.colorSaturation,l=this._options.colorBrightness,c='hsla('+section.color+', '+s+', '+l+', '+section.alpha+' )';this._context.save();if(this._options.parallaxAmount!==0){this._context.translate(0,this._scroll*this._options.parallaxAmount)}this._context.beginPath();this._context.moveTo(section.point1.x,section.point1.y);this._context.lineTo(section.point2.x,section.point2.y);this._context.lineTo(section.point3.x,section.point3.y);this._context.fillStyle=c;this._context.fill();if(this._options.strokeSize>0){this._context.lineWidth=this._options.strokeSize;this._context.strokeStyle=c;this._context.lineCap='round';this._context.stroke()}this._context.restore()}return false;},_onDraw:function(){for(var i=0,t=this._ribbons.length;i<t;i+=1){if(!this._ribbons[i]){this._ribbons.splice(i,1)}}this._context.clearRect(0,0,this._width,this._height);for(var a=0;a<this._ribbons.length;++a ){var ribbon=this._ribbons[a],numSections=ribbon.length,numDone=0;for(var b=0;b<numSections;++b ){if(this._drawRibbonSection(ribbon[b])){numDone++;}}if(numDone>=numSections){this._ribbons[a]=null}}if(this._ribbons.length<this._options.ribbonCount){this.addRibbon()}requestAnimationFrame(this._onDraw)},_onResize:function(e){var screen=screenInfo(e);this._width=screen.width;this._height=screen.height;if(this._canvas){this._canvas.width=this._width;this._canvas.height=this._height;if(this._context){this._context.globalAlpha=this._options.colorAlpha}}},_onScroll:function(e){var screen=screenInfo(e);this._scroll=screen.scrolly}};return Factory});new Ribbons({colorSaturation:'60%',colorBrightness:'50%',colorAlpha:0.5,colorCycleSpeed:5,verticalPosition:'random',horizontalSpeed:200,ribbonCount:3,strokeSize:0,parallaxAmount:-0.2,animateSections:true})</script> ;</body></html><style>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@200;300;400;500;600;700&display=swap');.syntax{background:#222;color:#eee;margin:10px;border:solid thin #333;max-width:500px;width:100%;display:inline-block}.syntax span{counter-increment:linecounter}.codeheader{margin-left:10px;margin-bottom:-10px;padding:5px 5px 5px 15px;font-family:'Roboto', sans-serif;font-size:1.1em;color:#fff;-webkit-border-radius:6px 6px 0 0;-moz-border-radius:6px 6px 0 0;border-radius:6px 6px 0 0;user-select:none;-ms-user-select:none;-moz-user-select:none;-webkit-user-select:none;width:100%;max-width:500px}#Extra{background-color:#e05260}.syntax span:before{content:counter(linecounter);width:1.2em;text-align:center;display:inline-block;border-right:1px solid #444;margin-right:10px;font-style:normal !important;color:#444 !important}#template{display:none}.num{font-weight:600;color:#d0782a}[class*=var]{font-weight:600}.clr{border-radius:1px;color:white;font-weight:400}.str,.str i{color:#ddca7e !important;font-weight:400 !important}.reg,.reg i{color:#ddca7e !important;font-weight:500 !important}.cmnt,.cmnt i{color:#555 !important;font-weight:400 !important}@media (max-width: 400px){.syntax{font-size:12px}}.container_Overview th h1{font-weight:700;font-size:1em;text-align:left;color:#9d00ff}.container_Overview td{font-weight:400;font-size:1em;-webkit-box-shadow:0 2px 2px -2px #0e1119;-moz-box-shadow:0 2px 2px -2px #0e1119;box-shadow:0 2px 2px -2px #0e1119}.container_Overview{text-align:left;overflow:scroll;width:100%;margin:0;display:table;padding:0 0 8em}.container_Overview td,.container_Overview th{padding-bottom:2%;padding-top:2%;padding-left:2%;max-width:70px;color:#e05260}.container_Overview th{background-color:#090a0b}.container_Overview td:first-child{color:#9d00ff}.container_Overview tr:hover{background-color:#222222;-webkit-box-shadow:0 6px 6px -6px #0e1119;-moz-box-shadow:0 6px 6px -6px #0e1119;box-shadow:0 6px 6px -6px #0e1119}body{background-color:black;background-image:radial-gradient(circle, rgba(0, 0, 0, 0) 0%, rgba(0, 0, 0, 0.8) 100%);background-position:center center;background-repeat:no-repeat;background-attachment:fixed;background-size:cover}*{margin:0;padding:0;box-sizing:border-box;font-family:'Poppins', sans-serif}.sidebar{position:fixed;height:100%;width:240px;transition:all 0.5s ease;overflow-y:scroll}.sidebar.active{width:60px}.sidebar .logo-details{height:80px;display:flex;align-items:center}.sidebar .logo-details i{font-size:28px;font-weight:500;color:#fff;min-width:60px;text-align:center}.sidebar .logo-details .logo_name{color:#fff;font-size:24px;font-weight:500}.sidebar .nav-links{margin-top:10px}.sidebar .nav-links li{position:relative;list-style:none;height:50px}.sidebar .nav-links li a{height:100%;width:100%;display:flex;align-items:center;text-decoration:none;transition:all 0.4s ease}.sidebar .nav-links li a.active{background:blueviolet}.sidebar .nav-links li a:hover{background:red}.sidebar .nav-links li i{min-width:60px;text-align:center;font-size:18px;color:#fff}.sidebar .nav-links li a .links_name{color:#fff;font-size:15px;font-weight:400;white-space:nowrap}.sidebar .nav-links .log_out{position:absolute;bottom:0;width:100%}.home-section{position:relative;min-height:100vh;width:calc(100% - 240px);left:240px;transition:all 0.5s ease}.sidebar.active~.home-section{width:calc(100% - 60px);left:60px}.home-section nav{display:flex;justify-content:space-between;height:80px;display:flex;align-items:center;position:fixed;width:calc(100% - 240px);left:240px;z-index:100;padding:0 20px;box-shadow:0 1px 1px rgba(0, 0, 0, 0.1);color:white;transition:all 0.5s ease}.sidebar.active~.home-section nav{left:60px;width:calc(100% - 60px)}.home-section nav .sidebar-button{display:flex;align-items:center;font-size:24px;font-weight:500}nav .sidebar-button i{font-size:35px;margin-right:10px}.home-section .home-content{position:relative;padding-top:104px}.home-content .overview-boxes{display:flex;align-items:center;justify-content:space-between;flex-wrap:wrap;padding:0 20px;margin-bottom:26px}.overview-boxes .box{display:flex;align-items:center;justify-content:center;width:calc(100% / 4 - 15px);background:white;padding:15px 14px;border-radius:12px;box-shadow:0 5px 10px rgba(0, 0, 0, 0.1)}.overview-boxes .box2{margin-top:50px;display:flex;align-items:center;justify-content:center;width:calc(100% / 4 - 15px);padding:15px 14px;border-radius:12px;box-shadow:0 5px 10px rgba(0, 0, 0, 0.1)}.overview-boxes .box-topic a{color:blueviolet;text-decoration:none;font-size:20px;font-weight:500}.overview-boxes .box-topic2 a{color:#e05260;text-decoration:none;font-size:20px;font-weight:500}.home-content .box .number{display:inline-block;font-size:35px;margin-top:-6px;font-weight:500}.home-content .box2 .number2{display:inline-block;font-size:35px;margin-top:-6px;font-weight:500}.home-content .overview-boxes{display:flex;align-items:center;justify-content:space-between;flex-wrap:wrap;padding:0 20px;margin-bottom:26px}.overview-boxes .box-topic{font-size:20px;font-weight:500}.home-content .box .number{display:inline-block;font-size:35px;margin-top:-6px;font-weight:500}.home-content .sales-boxes{display:flex;justify-content:space-between}@media (max-width: 1240px){.sidebar{width:60px}.sidebar.active{width:220px}.home-section{width:calc(100% - 60px);left:60px}.sidebar.active~.home-section{overflow:hidden;left:220px}.home-section nav{width:calc(100% - 60px);left:60px}.sidebar.active~.home-section nav{width:calc(100% - 220px);left:220px}}@media (max-width: 1150px){.home-content .sales-boxes{flex-direction:column}.home-content .sales-boxes .box{width:100%;overflow-x:scroll;margin-bottom:30px}.home-content .sales-boxes .top-sales{margin:0}}@media (max-width: 1000px){.overview-boxes .box{width:calc(100% / 2 - 15px);margin-bottom:15px}}@media (max-width: 700px){nav .profile-details .admin_name,nav .profile-details i,nav .sidebar-button .dashboard{display:none}.home-section nav .profile-details{height:50px;min-width:40px}.home-content .sales-boxes .sales-details{width:560px}}@media (max-width: 550px){.overview-boxes .box{width:100%;margin-bottom:15px}.sidebar.active~.home-section nav .profile-details{display:none}}@media (max-width: 400px){.sidebar{width:0}.sidebar.active{width:60px}.home-section{width:100%;left:0}.sidebar.active~.home-section{left:60px;width:calc(100% - 60px)}.home-section nav{width:100%;left:0}.sidebar.active~.home-section nav{left:60px;width:calc(100% - 60px)}}</style>
`
