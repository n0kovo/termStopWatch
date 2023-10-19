<a name="readme-top"></a>
<div align="center">

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![GPL License][license-shield]][license-url]

<br />
  <a href="https://github.com/n0kovo/termstopwatch">
<img src="https://readme-typing-svg.demolab.com?font=MesloLGS+NF&size=40&pause=1000&color=D7D7D7&background=FFFFFF00&center=true&vCenter=true&random=false&width=500&height=80&lines=.%2Ftermstopwatch" alt="Logo" />
  </a>

  <h3 align="center">An incredible, elegant multi-font ASCII art CLI stopwatch</h3>

  <p align="center">
    State-of-the-art terminal chrono-measurement boii. It starts. It stops. It's got all sorts of tricks.
    <br />
    <br />
    ·
    <a href="https://github.com/n0kovo/termstopwatch/issues">Report Bug</a>
    ·
    <a href="https://github.com/n0kovo/termstopwatch/issues">Request Feature</a>
  </p>

## Demo

<img src="https://github.com/n0kovo/termStopWatch/assets/16690056/18ed8792-cf90-4de5-b059-c700dcbd310c" alt="Logo" align="center" width=70% />
</div>
<br />

_"Why settle for an average, humdrum, run-of-the-mill stopwatch when you can dazzle your retinas with a text-based spectacle like no other?"_
   – Steve Jobs
<br />

`termstopwatch` is not just a stopwatch; it's an experience. Written in Go, this terminal app manifests the future of time-keeping through the cutting-edge technology of ASCII art. Yes, we went there.

### Features So Magical, They'll Turn You Into a Newt:

* **Stopwatch Reimagined**
  * Behold the sheer elegance of elapsed time ticking away — in beautifully hand-crafted ASCII art.<br /><br />
* **Fonts Galore**
  * Tired of regular old LCD displays? Choose from an overwhelming array of fonts stored in our luxury `fonts/` directory, in which you can also put your own `figlet` font files. It's basically like changing the drapes in a mansion.<br /><br />
* **Keyboard Shortcuts**
  * With hotkeys so intuitive, you'd think they were telepathic. Pause, reset, toggle milliseconds like you're typing Shakespeare.<br /><br />
* **Help Text**
  * For the unlikely event that a modern-day Einstein like yourself needs it. Toggle on and off for those "just in case I forget I'm a genius" moments.<br /><br />
  
_Welcome to the future, where even your stopwatch is cooler than you._

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

It almost couldn't be easier, my friend.

### Prerequisites

You'll ned `go` to compile `termstopwatch`, so
* Install [Go](https://go.dev/doc/install)

### Installation
When you have `go`, you can install `termstopwatch` like so:
```console
go install https://github.com/n0kovo/termstopwatch@latest
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- USAGE EXAMPLES -->
## Usage

When `termstopwatch` is running, you can use these carefully selected keys to change the behavior:
* **(SPACE)** - Start/Stop the timer
* **(M)** - Toggle displaying of milliseconds
* **(R)** - Reset to 00:00:00.000
* **(H)** - Hide help text
* **(ESC)** or **Ctrl-c** - Exit :(

You can also run it with these command line flags, if you're an _advanced_ user:
```console
Usage of ./termstopwatch:
  -d, --debug         Show debug text.
  -f, --font string   Font to use. Must be in fonts/ directory. (default "Georgia11")
  -q, --hide-help     Hide help text.
  -n, --noms          Hide milliseconds.
  -p, --pause         Start paused.
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

I know it's hard to imagine, but if you have a suggestion that would improve `termstopwatch`, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- LICENSE -->
## License

Distributed under the GNU GPL 3 License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

n0kovo - [Mastodon](https://infosec.exchange/@n0kovo)

Project Link: [https://github.com/n0kovo/termstopwatch](https://github.com/n0kovo/termstopwatch)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [Figlet](http://www.figlet.org/)
* [figlet4go](https://github.com/mbndr/figlet4go)
* Font authors <3

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/n0kovo/termstopwatch.svg?style=for-the-badge
[contributors-url]: https://github.com/n0kovo/termstopwatch/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/n0kovo/termstopwatch.svg?style=for-the-badge
[forks-url]: https://github.com/n0kovo/termstopwatch/network/members
[stars-shield]: https://img.shields.io/github/stars/n0kovo/termstopwatch.svg?style=for-the-badge
[stars-url]: https://github.com/n0kovo/termstopwatch/stargazers
[issues-shield]: https://img.shields.io/github/issues/n0kovo/termstopwatch.svg?style=for-the-badge
[issues-url]: https://github.com/n0kovo/termstopwatch/issues
[license-shield]: https://img.shields.io/github/license/n0kovo/termstopwatch.svg?style=for-the-badge
[license-url]: https://github.com/n0kovo/termstopwatch/blob/master/LICENSE.txt
