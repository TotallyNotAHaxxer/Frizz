package Engine

// Generates general templates for tables and columns. Simple static generation

type Values struct {
	TableTitles      int      // For example | 2
	TableTitlesNames []string // For example | useragent, useragentos
	TableData        []string // For example | data
	Tablefirstchild  string   // For example | red
	TableSecondChild string   // For example | blue

}

var ColorList = map[string]string{
	"gorange":    "#e05260", // simple orange but red
	"nblue":      "#1c0f7a", // Navy blue
	"bviolet":    "#9d00ff", // Blue violet
	"highpurple": "540595",  // High purple
}

var Val Values

// broken table generation | d: Fix this
func (V *Values) GenerateTable() string {
	var templ string
	templ += StandardTop
	templ += Standard_LinksList
	templ += Generate_Sec("home-section") + "\n"
	templ += Generate_NavBegin() + "\n"
	templ += `<div class="sidebar-button"><i class='bx bx-menu sidebarBtn'></i><span class="dashboard">Dashboard</span></div>` + "\n"
	templ += Generate_NavEnd()
	templ += Generate_Table("container_Overview") + "\n"
	templ += Generate_HRBR() + "\n"
	templ += Generate_THEAD() + "\n" // generate table head
	for i := 0; i < V.TableTitles; i++ {
		templ += Generate_TH() + "\n"                                                 // start table header
		templ += Generate_H1START() + V.TableTitlesNames[i] + Generate_H1END() + "\n" // load h1 tags
		templ += Generate_THE() + "\n"                                                // End table header
	}
	templ += Generate_THEADEND() + "\n"
	for i := 0; i < len(V.TableData); i++ {
		templ += "<tr>" + "\n"
		for k := 0; k < V.TableTitles; k++ {
			templ += Generate_TD() + V.TableData[i] + Generate_TDE() + "\n"
		}
		templ += "</tr>" + "\n"
	}
	templ += Generate_TBLEEND() + "\n"
	templ += `</section></body></html>` + "\n"
	/*

		why didnt this work




		data := EmbednewStyle(
			"container_Overview th",
			fmt.Sprintf("color: %s;", ColorList[V.TableSecondChild]),
		)
		("Embedding -> ", data)
		templ += data
	*/

	templ += standardJSCSS + "\n"
	return templ
}

var (
	standardJSCSS = `
	<script type="text/javascript"></script>
    <script>
        let sidebar = document.querySelector(".sidebar");
        let sidebarBtn = document.querySelector(".sidebarBtn");
        sidebarBtn.onclick = function() {
            sidebar.classList.toggle("active");
            if (sidebar.classList.contains("active")) {
                sidebarBtn.classList.replace("bx-menu", "bx-menu-alt-right")
            } else {
                sidebarBtn.classList.replace("bx-menu-alt-right", "bx-menu")
            }
        };
        (function(name, factory) {
            if (typeof window === 'object') {
                window[name] = factory()
            }
        })('Ribbons', function() {
            var _w = window,
                _b = document.body,
                _d = document.documentElement;
            var random = function() {
                if (arguments.length === 1) {
                    if (Array.isArray(arguments[0])) {
                        var index = Math.round(random(0, arguments[0].length - 1));
                        return arguments[0][index]
                    }
                    return random(0, arguments[0]);
                } else if (arguments.length === 2) {
                    return Math.random() * (arguments[1] - arguments[0]) + arguments[0]
                } else if (arguments.length === 4) {
                    var array = [arguments[0], arguments[1], arguments[2], arguments[3]];
                    return array[Math.floor(Math.random() * array.length)];
                }
                return 0;
            };
            var screenInfo = function(e) {
                var width = Math.max(0, _w.innerWidth || _d.clientWidth || _b.clientWidth || 0),
                    height = Math.max(0, _w.innerHeight || _d.clientHeight || _b.clientHeight || 0),
                    scrollx = Math.max(0, _w.pageXOffset || _d.scrollLeft || _b.scrollLeft || 0) - (_d.clientLeft || 0),
                    scrolly = Math.max(0, _w.pageYOffset || _d.scrollTop || _b.scrollTop || 0) - (_d.clientTop || 0);
                return {
                    width: width,
                    height: height,
                    ratio: width / height,
                    centerx: width / 2,
                    centery: height / 2,
                    scrollx: scrollx,
                    scrolly: scrolly
                }
            };
            var mouseInfo = function(e) {
                var screen = screenInfo(e),
                    mousex = e ? Math.max(0, e.pageX || e.clientX || 0) : 0,
                    mousey = e ? Math.max(0, e.pageY || e.clientY || 0) : 0;
                return {
                    mousex: mousex,
                    mousey: mousey,
                    centerx: mousex - screen.width / 2,
                    centery: mousey - screen.height / 2
                }
            };
            var Point = function(x, y) {
                this.x = 0;
                this.y = 0;
                this.set(x, y)
            };
            Point.prototype = {
                constructor: Point,
                set: function(x, y) {
                    this.x = x || 0;
                    this.y = y || 0
                },
                copy: function(point) {
                    this.x = point.x || 0;
                    this.y = point.y || 0;
                    return this
                },
                multiply: function(x, y) {
                    this.x *= x || 1;
                    this.y *= y || 1;
                    return this
                },
                divide: function(x, y) {
                    this.x /= x || 1;
                    this.y /= y || 1;
                    return this
                },
                add: function(x, y) {
                    this.x += x || 0;
                    this.y += y || 0;
                    return this
                },
                subtract: function(x, y) {
                    this.x -= x || 0;
                    this.y -= y || 0;
                    return this
                },
                clampX: function(min, max) {
                    this.x = Math.max(min, Math.min(this.x, max));
                    return this
                },
                clampY: function(min, max) {
                    this.y = Math.max(min, Math.min(this.y, max));
                    return this
                },
                flipX: function() {
                    this.x *= -1;
                    return this
                },
                flipY: function() {
                    this.y *= -1;
                    return this
                }
            };
            var Factory = function(options) {
                this._canvas = null;
                this._context = null;
                this._sto = null;
                this._width = 0;
                this._height = 0;
                this._scroll = 0;
                this._ribbons = [];
                this._options = {
                    colorSaturation: '80%',
                    colorBrightness: '60%',
                    colorAlpha: 0.65,
                    colorCycleSpeed: 6,
                    verticalPosition: 'center',
                    horizontalSpeed: 150,
                    ribbonCount: 3,
                    strokeSize: 0,
                    parallaxAmount: -0.5,
                    animateSections: true
                };
                this._onDraw = this._onDraw.bind(this);
                this._onResize = this._onResize.bind(this);
                this._onScroll = this._onScroll.bind(this);
                this.setOptions(options);
                this.init()
            };
            Factory.prototype = {
                constructor: Factory,
                setOptions: function(options) {
                    if (typeof options === 'object') {
                        for (var key in options) {
                            if (options.hasOwnProperty(key)) {
                                this._options[key] = options[key]
                            }
                        }
                    }
                },
                init: function() {
                    try {
                        this._canvas = document.createElement('canvas');
                        this._canvas.style['display'] = 'block';
                        this._canvas.style['position'] = 'fixed';
                        this._canvas.style['margin'] = '0';
                        this._canvas.style['padding'] = '0';
                        this._canvas.style['border'] = '0';
                        this._canvas.style['outline'] = '0';
                        this._canvas.style['left'] = '0';
                        this._canvas.style['top'] = '0';
                        this._canvas.style['width'] = '100%';
                        this._canvas.style['height'] = '100%';
                        this._canvas.style['z-index'] = '-1';
                        this._onResize();
                        this._context = this._canvas.getContext('2d');
                        this._context.clearRect(0, 0, this._width, this._height);
                        this._context.globalAlpha = this._options.colorAlpha;
                        window.addEventListener('resize', this._onResize);
                        window.addEventListener('scroll', this._onScroll);
                        document.body.appendChild(this._canvas)
                    } catch (e) {
                        console.warn('Canvas Context Error: ' + e.toString());
                        return
                    }
                    this._onDraw()
                },
                addRibbon: function() {
                    var dir = Math.round(random(1, 9)) > 5 ? 'right' : 'left',
                        stop = 1000,
                        hide = 200,
                        min = 0 - hide,
                        max = this._width + hide,
                        movex = 0,
                        movey = 0,
                        startx = dir === 'right' ? min : max,
                        starty = Math.round(random(0, this._height));
                    if (/^(top|min)$/i.test(this._options.verticalPosition)) {
                        starty = 0 + hide
                    } else if (/^(middle|center)$/i.test(this._options.verticalPosition)) {
                        starty = this._height / 2
                    } else if (/^(bottom|max)$/i.test(this._options.verticalPosition)) {
                        starty = this._height - hide
                    }
                    var ribbon = [],
                        point1 = new Point(startx, starty),
                        point2 = new Point(startx, starty),
                        point3 = null,
                        color = Math.round(random(900)),
                        delay = 0;
                    while (true) {
                        if (stop <= 0) {
                            break
                        }
                        stop--;
                        movex = Math.round((Math.random() * 1 - 0.2) * this._options.horizontalSpeed);
                        movey = Math.round((Math.random() * 1 - 0.5) * (this._height * 0.25));
                        point3 = new Point();
                        point3.copy(point2);
                        if (dir === 'right') {
                            point3.add(movex, movey);
                            if (point2.x >= max) {
                                break
                            }
                        } else if (dir === 'left') {
                            point3.subtract(movex, movey);
                            if (point2.x <= min) {
                                break
                            }
                        }
                        ribbon.push({
                            point1: new Point(point1.x, point1.y),
                            point2: new Point(point2.x, point2.y),
                            point3: point3,
                            color: 615,
                            delay: delay,
                            dir: dir,
                            alpha: 0,
                            phase: 0
                        });
                        point1.copy(point2);
                        point2.copy(point3);
                        delay += 4
                    }
                    this._ribbons.push(ribbon)
                },
                _drawRibbonSection: function(section) {
                    if (section) {
                        if (section.phase >= 1 && section.alpha <= 0) {
                            return true;
                        }
                        if (section.delay <= 0) {
                            section.phase += 0.02;
                            section.alpha = Math.sin(section.phase) * 1;
                            section.alpha = section.alpha <= 0 ? 0 : section.alpha;
                            section.alpha = section.alpha >= 1 ? 1 : section.alpha;
                            if (this._options.animateSections) {
                                var mod = Math.sin(1 + section.phase * Math.PI / 2) * 0.1;
                                if (section.dir === 'right') {
                                    section.point1.add(mod, 0);
                                    section.point2.add(mod, 0);
                                    section.point3.add(mod, 0)
                                } else {
                                    section.point1.subtract(mod, 0);
                                    section.point2.subtract(mod, 0);
                                    section.point3.subtract(mod, 0)
                                }
                                section.point1.add(0, mod);
                                section.point2.add(0, mod);
                                section.point3.add(0, mod)
                            }
                        } else {
                            section.delay -= 0.5
                        }
                        var s = this._options.colorSaturation,
                            l = this._options.colorBrightness,
                            c = 'hsla(' + section.color + ', ' + s + ', ' + l + ', ' + section.alpha + ' )';
                        this._context.save();
                        if (this._options.parallaxAmount !== 0) {
                            this._context.translate(0, this._scroll * this._options.parallaxAmount)
                        }
                        this._context.beginPath();
                        this._context.moveTo(section.point1.x, section.point1.y);
                        this._context.lineTo(section.point2.x, section.point2.y);
                        this._context.lineTo(section.point3.x, section.point3.y);
                        this._context.fillStyle = c;
                        this._context.fill();
                        if (this._options.strokeSize > 0) {
                            this._context.lineWidth = this._options.strokeSize;
                            this._context.strokeStyle = c;
                            this._context.lineCap = 'round';
                            this._context.stroke()
                        }
                        this._context.restore()
                    }
                    return false;
                },
                _onDraw: function() {
                    for (var i = 0, t = this._ribbons.length; i < t; i += 1) {
                        if (!this._ribbons[i]) {
                            this._ribbons.splice(i, 1)
                        }
                    }
                    this._context.clearRect(0, 0, this._width, this._height);
                    for (var a = 0; a < this._ribbons.length; ++a) {
                        var ribbon = this._ribbons[a],
                            numSections = ribbon.length,
                            numDone = 0;
                        for (var b = 0; b < numSections; ++b) {
                            if (this._drawRibbonSection(ribbon[b])) {
                                numDone++;
                            }
                        }
                        if (numDone >= numSections) {
                            this._ribbons[a] = null
                        }
                    }
                    if (this._ribbons.length < this._options.ribbonCount) {
                        this.addRibbon()
                    }
                    requestAnimationFrame(this._onDraw)
                },
                _onResize: function(e) {
                    var screen = screenInfo(e);
                    this._width = screen.width;
                    this._height = screen.height;
                    if (this._canvas) {
                        this._canvas.width = this._width;
                        this._canvas.height = this._height;
                        if (this._context) {
                            this._context.globalAlpha = this._options.colorAlpha
                        }
                    }
                },
                _onScroll: function(e) {
                    var screen = screenInfo(e);
                    this._scroll = screen.scrolly
                }
            };
            return Factory
        });
        new Ribbons({
            colorSaturation: '60%',
            colorBrightness: '50%',
            colorAlpha: 0.5,
            colorCycleSpeed: 5,
            verticalPosition: 'random',
            horizontalSpeed: 200,
            ribbonCount: 3,
            strokeSize: 0,
            parallaxAmount: -0.2,
            animateSections: true
        });
    </script>
<style>
    @import url('https://fonts.googleapis.com/css2?family=Poppins:wght@200;300;400;500;600;700&display=swap');

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
        margin: 0 0px;
        display: table;
        padding: 0 0 8em
    }

    .container_Overview td,
    .container_Overview th {
        padding-bottom: 2%;
        padding-top: 2%;
        padding-left: 2%;
        max-width: 70px;
		color: #e05260;
    }

    .container_Overview th {
        background-color: #090a0b;
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
        font-family: 'Open Sans', sans-serif;
        font-weight: 300;
        background-color: black;
        background-image: radial-gradient(circle, rgba(0, 0, 0, 0) 0%, rgba(0, 0, 0, 0.8) 100%);
        background-position: center center;
        background-repeat: no-repeat;
        background-attachment: fixed;
        background-size: cover;
    }

    * {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
        font-family: 'Poppins', sans-serif;
    }

    .sidebar {
        position: fixed;
        height: 100%;
        width: 240px;
        transition: all 0.5s ease;
        overflow-y: scroll;
    }

    .sidebar.active {
        width: 60px;
    }

    .sidebar .logo-details {
        height: 80px;
        display: flex;
        align-items: center;
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
        font-weight: 500;
    }

    .sidebar .nav-links {
        margin-top: 10px;
    }

    .sidebar .nav-links li {
        position: relative;
        list-style: none;
        height: 50px;
    }

    .sidebar .nav-links li a {
        height: 100%;
        width: 100%;
        display: flex;
        align-items: center;
        text-decoration: none;
        transition: all 0.4s ease;
    }

    .sidebar .nav-links li a.active {
        background: blueviolet;
    }

    .sidebar .nav-links li a:hover {
        background: red;
    }

    .sidebar .nav-links li i {
        min-width: 60px;
        text-align: center;
        font-size: 18px;
        color: #fff;
    }

    .sidebar .nav-links li a .links_name {
        color: #fff;
        font-size: 15px;
        font-weight: 400;
        white-space: nowrap;
    }

    .home-section {
        position: relative;
        min-height: 100vh;
        width: calc(100% - 240px);
        left: 240px;
        transition: all 0.5s ease;
    }

    .sidebar.active~.home-section {
        width: calc(100% - 60px);
        left: 60px;
    }

    .home-section nav {
        justify-content: space-between;
        height: 80px;
        display: flex;
        align-items: center;
        width: calc(100% - 240px);
        left: 240px;
        z-index: 100;
        padding: 0 20px;
        box-shadow: 0 1px 1px rgba(0, 0, 0, 0.1);
        color: white;
        transition: all 0.5s ease;
    }

    .sidebar.active~.home-section nav {
        left: 60px;
        width: calc(100% - 60px);
    }

    .home-section nav .sidebar-button {
        display: flex;
        align-items: center;
        font-size: 24px;
        font-weight: 500;
        position: fixed;
    }

    nav .sidebar-button i {
        font-size: 35px;
        margin-right: 10px;
    }

    @media (max-width: 1240px) {
        .sidebar {
            width: 60px;
        }

        .sidebar.active {
            width: 220px;
        }

        .home-section {
            width: calc(100% - 60px);
            left: 60px;
        }

        .sidebar.active~.home-section {
            overflow: hidden;
            left: 220px;
        }

        .home-section nav {
            width: calc(100% - 60px);
            left: 60px;
        }

        .sidebar.active~.home-section nav {
            width: calc(100% - 220px);
            left: 220px;
        }
    }

    @media (max-width: 700px) {
        nav .sidebar-button .dashboard {
            height: 50px;
            min-width: 40px;
        }
    }

    @media (max-width: 550px) {
        .sidebar.active~.home-section nav .profile-details {
            display: none;
        }
    }

    @media (max-width: 400px) {
        .sidebar {
            width: 0;
        }

        .sidebar.active {
            width: 60px;
        }

        .home-section {
            width: 100%;
            left: 0;
        }

        .sidebar.active~.home-section {
            left: 60px;
            width: calc(100% - 60px);
        }

        .home-section nav {
            width: 100%;
            left: 0;
        }

        .sidebar.active~.home-section nav {
            left: 60px;
            width: calc(100% - 60px);
        }
    }
</style>
	`
)

/*
EXAMPLE STANDARD HTML DOCUMENT

    <section class="home-section">
        <nav>
        <div class="sidebar-button"><i class='bx bx-menu sidebarBtn'></i><span class="dashboard">Dashboard</span></div>
        <div class="home-content">
            <table class="container_Overview"><br>
                <hr><br>
                <thead>
                    <tr>
                        <th>
                            <h1>Hostname</h1>
                        </th>
                        <th>
                            <h1>UserAgent</h1>
                        </th>
                        <th>
                            <h1>Operating System </h1>
                        </th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>browserspy.dk</td>
                        <td>Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.69 Safari/537.36</td>
                        <td>Mac OS X</td>
                    <tr>
                    <tr>
                        <td>bill.ins.com</td>
                        <td>Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.2; SV1; .NET CLR 1.1.4322)</td>
                        <td>Windows</td>
                    <tr>
                    <tr>
                        <td>192.168.1.8</td>
                        <td>Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.65 Safari/537.36</td>
                        <td>Mac OS X</td>
                    <tr>
                    <tr>
                        <td>packetlife.net</td>
                        <td>Wget/1.12 (linux-gnu)</td>
                        <td>Operating system was not found or supported</td>
                    <tr>
                    <tr>
                        <td>10.0.1.101</td>
                        <td>Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322)</td>
                        <td>Windows</td>
                    <tr>
                    <tr>
                        <td>192.168.0.55</td>
                        <td>Mozilla/5.0 (Windows NT 5.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1</td>
                        <td>Windows</td>
                    <tr>
                    <tr>
                        <td>www.africau.edu</td>
                        <td>Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36</td>
                        <td>Windows</td>
                    <tr>
                    <tr>
                        <td>www.ethereal.com</td>
                        <td>Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.6) Gecko/20040113</td>
                        <td>Windows</td>
                    <tr>
                </tbody>
            </table>
        </div>
    </section>
    <script type="text/javascript"></script>
    <script>
        let sidebar = document.querySelector(".sidebar");
        let sidebarBtn = document.querySelector(".sidebarBtn");
        sidebarBtn.onclick = function() {
            sidebar.classList.toggle("active");
            if (sidebar.classList.contains("active")) {
                sidebarBtn.classList.replace("bx-menu", "bx-menu-alt-right")
            } else {
                sidebarBtn.classList.replace("bx-menu-alt-right", "bx-menu")
            }
        };
        (function(name, factory) {
            if (typeof window === 'object') {
                window[name] = factory()
            }
        })('Ribbons', function() {
            var _w = window,
                _b = document.body,
                _d = document.documentElement;
            var random = function() {
                if (arguments.length === 1) {
                    if (Array.isArray(arguments[0])) {
                        var index = Math.round(random(0, arguments[0].length - 1));
                        return arguments[0][index]
                    }
                    return random(0, arguments[0]);
                } else if (arguments.length === 2) {
                    return Math.random() * (arguments[1] - arguments[0]) + arguments[0]
                } else if (arguments.length === 4) {
                    var array = [arguments[0], arguments[1], arguments[2], arguments[3]];
                    return array[Math.floor(Math.random() * array.length)];
                }
                return 0;
            };
            var screenInfo = function(e) {
                var width = Math.max(0, _w.innerWidth || _d.clientWidth || _b.clientWidth || 0),
                    height = Math.max(0, _w.innerHeight || _d.clientHeight || _b.clientHeight || 0),
                    scrollx = Math.max(0, _w.pageXOffset || _d.scrollLeft || _b.scrollLeft || 0) - (_d.clientLeft || 0),
                    scrolly = Math.max(0, _w.pageYOffset || _d.scrollTop || _b.scrollTop || 0) - (_d.clientTop || 0);
                return {
                    width: width,
                    height: height,
                    ratio: width / height,
                    centerx: width / 2,
                    centery: height / 2,
                    scrollx: scrollx,
                    scrolly: scrolly
                }
            };
            var mouseInfo = function(e) {
                var screen = screenInfo(e),
                    mousex = e ? Math.max(0, e.pageX || e.clientX || 0) : 0,
                    mousey = e ? Math.max(0, e.pageY || e.clientY || 0) : 0;
                return {
                    mousex: mousex,
                    mousey: mousey,
                    centerx: mousex - screen.width / 2,
                    centery: mousey - screen.height / 2
                }
            };
            var Point = function(x, y) {
                this.x = 0;
                this.y = 0;
                this.set(x, y)
            };
            Point.prototype = {
                constructor: Point,
                set: function(x, y) {
                    this.x = x || 0;
                    this.y = y || 0
                },
                copy: function(point) {
                    this.x = point.x || 0;
                    this.y = point.y || 0;
                    return this
                },
                multiply: function(x, y) {
                    this.x *= x || 1;
                    this.y *= y || 1;
                    return this
                },
                divide: function(x, y) {
                    this.x /= x || 1;
                    this.y /= y || 1;
                    return this
                },
                add: function(x, y) {
                    this.x += x || 0;
                    this.y += y || 0;
                    return this
                },
                subtract: function(x, y) {
                    this.x -= x || 0;
                    this.y -= y || 0;
                    return this
                },
                clampX: function(min, max) {
                    this.x = Math.max(min, Math.min(this.x, max));
                    return this
                },
                clampY: function(min, max) {
                    this.y = Math.max(min, Math.min(this.y, max));
                    return this
                },
                flipX: function() {
                    this.x *= -1;
                    return this
                },
                flipY: function() {
                    this.y *= -1;
                    return this
                }
            };
            var Factory = function(options) {
                this._canvas = null;
                this._context = null;
                this._sto = null;
                this._width = 0;
                this._height = 0;
                this._scroll = 0;
                this._ribbons = [];
                this._options = {
                    colorSaturation: '80%',
                    colorBrightness: '60%',
                    colorAlpha: 0.65,
                    colorCycleSpeed: 6,
                    verticalPosition: 'center',
                    horizontalSpeed: 150,
                    ribbonCount: 3,
                    strokeSize: 0,
                    parallaxAmount: -0.5,
                    animateSections: true
                };
                this._onDraw = this._onDraw.bind(this);
                this._onResize = this._onResize.bind(this);
                this._onScroll = this._onScroll.bind(this);
                this.setOptions(options);
                this.init()
            };
            Factory.prototype = {
                constructor: Factory,
                setOptions: function(options) {
                    if (typeof options === 'object') {
                        for (var key in options) {
                            if (options.hasOwnProperty(key)) {
                                this._options[key] = options[key]
                            }
                        }
                    }
                },
                init: function() {
                    try {
                        this._canvas = document.createElement('canvas');
                        this._canvas.style['display'] = 'block';
                        this._canvas.style['position'] = 'fixed';
                        this._canvas.style['margin'] = '0';
                        this._canvas.style['padding'] = '0';
                        this._canvas.style['border'] = '0';
                        this._canvas.style['outline'] = '0';
                        this._canvas.style['left'] = '0';
                        this._canvas.style['top'] = '0';
                        this._canvas.style['width'] = '100%';
                        this._canvas.style['height'] = '100%';
                        this._canvas.style['z-index'] = '-1';
                        this._onResize();
                        this._context = this._canvas.getContext('2d');
                        this._context.clearRect(0, 0, this._width, this._height);
                        this._context.globalAlpha = this._options.colorAlpha;
                        window.addEventListener('resize', this._onResize);
                        window.addEventListener('scroll', this._onScroll);
                        document.body.appendChild(this._canvas)
                    } catch (e) {
                        console.warn('Canvas Context Error: ' + e.toString());
                        return
                    }
                    this._onDraw()
                },
                addRibbon: function() {
                    var dir = Math.round(random(1, 9)) > 5 ? 'right' : 'left',
                        stop = 1000,
                        hide = 200,
                        min = 0 - hide,
                        max = this._width + hide,
                        movex = 0,
                        movey = 0,
                        startx = dir === 'right' ? min : max,
                        starty = Math.round(random(0, this._height));
                    if (/^(top|min)$/i.test(this._options.verticalPosition)) {
                        starty = 0 + hide
                    } else if (/^(middle|center)$/i.test(this._options.verticalPosition)) {
                        starty = this._height / 2
                    } else if (/^(bottom|max)$/i.test(this._options.verticalPosition)) {
                        starty = this._height - hide
                    }
                    var ribbon = [],
                        point1 = new Point(startx, starty),
                        point2 = new Point(startx, starty),
                        point3 = null,
                        color = Math.round(random(900)),
                        delay = 0;
                    while (true) {
                        if (stop <= 0) {
                            break
                        }
                        stop--;
                        movex = Math.round((Math.random() * 1 - 0.2) * this._options.horizontalSpeed);
                        movey = Math.round((Math.random() * 1 - 0.5) * (this._height * 0.25));
                        point3 = new Point();
                        point3.copy(point2);
                        if (dir === 'right') {
                            point3.add(movex, movey);
                            if (point2.x >= max) {
                                break
                            }
                        } else if (dir === 'left') {
                            point3.subtract(movex, movey);
                            if (point2.x <= min) {
                                break
                            }
                        }
                        ribbon.push({
                            point1: new Point(point1.x, point1.y),
                            point2: new Point(point2.x, point2.y),
                            point3: point3,
                            color: 615,
                            delay: delay,
                            dir: dir,
                            alpha: 0,
                            phase: 0
                        });
                        point1.copy(point2);
                        point2.copy(point3);
                        delay += 4
                    }
                    this._ribbons.push(ribbon)
                },
                _drawRibbonSection: function(section) {
                    if (section) {
                        if (section.phase >= 1 && section.alpha <= 0) {
                            return true;
                        }
                        if (section.delay <= 0) {
                            section.phase += 0.02;
                            section.alpha = Math.sin(section.phase) * 1;
                            section.alpha = section.alpha <= 0 ? 0 : section.alpha;
                            section.alpha = section.alpha >= 1 ? 1 : section.alpha;
                            if (this._options.animateSections) {
                                var mod = Math.sin(1 + section.phase * Math.PI / 2) * 0.1;
                                if (section.dir === 'right') {
                                    section.point1.add(mod, 0);
                                    section.point2.add(mod, 0);
                                    section.point3.add(mod, 0)
                                } else {
                                    section.point1.subtract(mod, 0);
                                    section.point2.subtract(mod, 0);
                                    section.point3.subtract(mod, 0)
                                }
                                section.point1.add(0, mod);
                                section.point2.add(0, mod);
                                section.point3.add(0, mod)
                            }
                        } else {
                            section.delay -= 0.5
                        }
                        var s = this._options.colorSaturation,
                            l = this._options.colorBrightness,
                            c = 'hsla(' + section.color + ', ' + s + ', ' + l + ', ' + section.alpha + ' )';
                        this._context.save();
                        if (this._options.parallaxAmount !== 0) {
                            this._context.translate(0, this._scroll * this._options.parallaxAmount)
                        }
                        this._context.beginPath();
                        this._context.moveTo(section.point1.x, section.point1.y);
                        this._context.lineTo(section.point2.x, section.point2.y);
                        this._context.lineTo(section.point3.x, section.point3.y);
                        this._context.fillStyle = c;
                        this._context.fill();
                        if (this._options.strokeSize > 0) {
                            this._context.lineWidth = this._options.strokeSize;
                            this._context.strokeStyle = c;
                            this._context.lineCap = 'round';
                            this._context.stroke()
                        }
                        this._context.restore()
                    }
                    return false;
                },
                _onDraw: function() {
                    for (var i = 0, t = this._ribbons.length; i < t; i += 1) {
                        if (!this._ribbons[i]) {
                            this._ribbons.splice(i, 1)
                        }
                    }
                    this._context.clearRect(0, 0, this._width, this._height);
                    for (var a = 0; a < this._ribbons.length; ++a) {
                        var ribbon = this._ribbons[a],
                            numSections = ribbon.length,
                            numDone = 0;
                        for (var b = 0; b < numSections; ++b) {
                            if (this._drawRibbonSection(ribbon[b])) {
                                numDone++;
                            }
                        }
                        if (numDone >= numSections) {
                            this._ribbons[a] = null
                        }
                    }
                    if (this._ribbons.length < this._options.ribbonCount) {
                        this.addRibbon()
                    }
                    requestAnimationFrame(this._onDraw)
                },
                _onResize: function(e) {
                    var screen = screenInfo(e);
                    this._width = screen.width;
                    this._height = screen.height;
                    if (this._canvas) {
                        this._canvas.width = this._width;
                        this._canvas.height = this._height;
                        if (this._context) {
                            this._context.globalAlpha = this._options.colorAlpha
                        }
                    }
                },
                _onScroll: function(e) {
                    var screen = screenInfo(e);
                    this._scroll = screen.scrolly
                }
            };
            return Factory
        });
        new Ribbons({
            colorSaturation: '60%',
            colorBrightness: '50%',
            colorAlpha: 0.5,
            colorCycleSpeed: 5,
            verticalPosition: 'random',
            horizontalSpeed: 200,
            ribbonCount: 3,
            strokeSize: 0,
            parallaxAmount: -0.2,
            animateSections: true
        });
    </script>
</body>

</html>
<style>
    @import url('https://fonts.googleapis.com/css2?family=Poppins:wght@200;300;400;500;600;700&display=swap');

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
        margin: 0 0px;
        display: table;
        padding: 0 0 8em
    }

    .container_Overview td,
    .container_Overview th {
        padding-bottom: 2%;
        padding-top: 2%;
        padding-left: 2%;
        max-width: 70px;
    }

    .container_Overview th {
        background-color: #090a0b;
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
        font-family: 'Open Sans', sans-serif;
        font-weight: 300;
        background-color: black;
        background-image: radial-gradient(circle, rgba(0, 0, 0, 0) 0%, rgba(0, 0, 0, 0.8) 100%);
        background-position: center center;
        background-repeat: no-repeat;
        background-attachment: fixed;
        background-size: cover;
    }

    * {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
        font-family: 'Poppins', sans-serif;
    }

    .sidebar {
        position: fixed;
        height: 100%;
        width: 240px;
        transition: all 0.5s ease;
        overflow-y: scroll;
    }

    .sidebar.active {
        width: 60px;
    }

    .sidebar .logo-details {
        height: 80px;
        display: flex;
        align-items: center;
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
        font-weight: 500;
    }

    .sidebar .nav-links {
        margin-top: 10px;
    }

    .sidebar .nav-links li {
        position: relative;
        list-style: none;
        height: 50px;
    }

    .sidebar .nav-links li a {
        height: 100%;
        width: 100%;
        display: flex;
        align-items: center;
        text-decoration: none;
        transition: all 0.4s ease;
    }

    .sidebar .nav-links li a.active {
        background: blueviolet;
    }

    .sidebar .nav-links li a:hover {
        background: red;
    }

    .sidebar .nav-links li i {
        min-width: 60px;
        text-align: center;
        font-size: 18px;
        color: #fff;
    }

    .sidebar .nav-links li a .links_name {
        color: #fff;
        font-size: 15px;
        font-weight: 400;
        white-space: nowrap;
    }

    .home-section {
        position: relative;
        min-height: 100vh;
        width: calc(100% - 240px);
        left: 240px;
        transition: all 0.5s ease;
    }

    .sidebar.active~.home-section {
        width: calc(100% - 60px);
        left: 60px;
    }

    .home-section nav {
        justify-content: space-between;
        height: 80px;
        display: flex;
        align-items: center;
        width: calc(100% - 240px);
        left: 240px;
        z-index: 100;
        padding: 0 20px;
        box-shadow: 0 1px 1px rgba(0, 0, 0, 0.1);
        color: white;
        transition: all 0.5s ease;
    }

    .sidebar.active~.home-section nav {
        left: 60px;
        width: calc(100% - 60px);
    }

    .home-section nav .sidebar-button {
        display: flex;
        align-items: center;
        font-size: 24px;
        font-weight: 500;
        position: fixed;
    }

    nav .sidebar-button i {
        font-size: 35px;
        margin-right: 10px;
    }

    @media (max-width: 1240px) {
        .sidebar {
            width: 60px;
        }

        .sidebar.active {
            width: 220px;
        }

        .home-section {
            width: calc(100% - 60px);
            left: 60px;
        }

        .sidebar.active~.home-section {
            overflow: hidden;
            left: 220px;
        }

        .home-section nav {
            width: calc(100% - 60px);
            left: 60px;
        }

        .sidebar.active~.home-section nav {
            width: calc(100% - 220px);
            left: 220px;
        }
    }

    @media (max-width: 700px) {
        nav .sidebar-button .dashboard {
            height: 50px;
            min-width: 40px;
        }
    }

    @media (max-width: 550px) {
        .sidebar.active~.home-section nav .profile-details {
            display: none;
        }
    }

    @media (max-width: 400px) {
        .sidebar {
            width: 0;
        }

        .sidebar.active {
            width: 60px;
        }

        .home-section {
            width: 100%;
            left: 0;
        }

        .sidebar.active~.home-section {
            left: 60px;
            width: calc(100% - 60px);
        }

        .home-section nav {
            width: 100%;
            left: 0;
        }

        .sidebar.active~.home-section nav {
            left: 60px;
            width: calc(100% - 60px);
        }
    }
</style>
*/
