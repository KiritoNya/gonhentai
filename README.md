[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

[![GoDoc](https://godoc.org/github.com/KiritoNya/gonhentai?status.svg)](https://pkg.go.dev/github.com/KiritoNya/gonhentai)
[![Go Report Card](https://goreportcard.com/badge/github.com/KiritoNya/gonhentai)](https://goreportcard.com/report/github.com/KiritoNya/gonhentai)
[![Sourcegraph](https://sourcegraph.com/github.com/KiritoNya/gonhentai/-/badge.svg)](https://sourcegraph.com/github.com/KiritoNya/gonhentai?badge)


<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/KiritoNya/gonhentai">
    <img src="https://files.catbox.moe/dwzbu6.png" alt="Logo" width="500" height="250">
  </a>
  <h3 align="center">GOnhentai</h3>

  <p align="center">
    A simple GoLang library for nHentai
    <br />
    <a href="https://pkg.go.dev/github.com/KiritoNya/gonhentai"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/github_username/repo_name">View Demo</a>
    ·
    <a href="https://github.com/KiritoNya/gonhentai/issues">Report Bug</a>
    ·
    <a href="https://github.com/KiritoNya/gonhentai/issues">Request Feature</a>
  </p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

```GOnhentai``` is a simple generic library that tries to create a base for all programs written in GO that use data from the [nhentai](https://nhentai.net) site.

This library introduces new features, including the management of comments and some user features (WIP).

<!-- GETTING STARTED -->
## Getting Started

To download and start using the library follow these simple:

### Prerequisites

* [go](https://golang.org/) >=1.16

### Installation

   ```sh
   go get github.com/KiritoNya/gonhentai
   ```

---
## Usage

Let's now see how the library is used with some small examples of common use.

### Doujinshi info

```go
// Create Doujinshi object and get some information
doujin, err := gonhentai.NewDoujinshiId(354862)
if err != nil {
	panic(err)
}

fmt.Println("Doujinshi:", doujin)
```

### Doujinshi Page Image

```go
// Create Object
doujin, err := gonhentai.NewDoujinshiId(354862)
if err != nil {
panic(err)
}

page := doujin.Pages[35]

// Generate the url of page
err := page.GetUrl()
if err != nil {
	panic(err)
}

// Get the data of page image
err := page.GetData()
if err != nil {
	panic(err)
}

fmt.Println("DATA:", page.Data)
```

### Comments

```go
// Create Doujinshi object and get some information
doujin, err := gonhentai.NewDoujinshiId(354862)
if err != nil {
	panic(err)
}

// Get comments
err := doujin.GetComments()
if err != nil {
	panic(err)
}

fmt.Println("Comments:", doujin.Comments)
```

### Download Page Image

The library already provides you a method to download the image and save it on your PC. 
Set the ```UseProgressBar``` variable to true if you decide to use the progress bar. Set ```ProgressBarTemplate``` variable If you want to use your progress bar defined by [pb](https://github.com/cheggaaa/pb) template.
```go
// Create Doujinshi object
doujin, err := gonhentai.NewDoujinshiId(354862)
if err != nil {
	panic(err)
}

page := doujin.Pages[36]

// Generate url of page
err := page.GetUrl()
if err != nil {
	panic(err)
}

// Default is false
// Set it to true if you want a progress bar
gonhentai.UseProgressBar = true

// Use this code for set your progress bar template.
// gonhentai.ProgressBarTemplate = `{{ red "With funcs:" }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{speed . | rndcolor }} {{percent .}} {{string . "my_green_string" | green}} {{string . "my_blue_string" | blue}}`

// Download image
err := page.Save("./img.jpg", 0644)
if err != nil {
	panic(err)
}
```


### Download Doujinshi

To download a doujinshi a path is required which must be provided via a template.
To form your path you can use all the fields of the ```Doujinshi``` object and all the fields of the ```page``` object, adding respectively ```.Doujinshi```, ```.Page```.

EX: ````/home/<username>/{{.Doujinshi.Id}} - {{.Doujinshi.Title.Pretty}}/{{.Page.Num}}.{{.Page.Ext}}````

```go
// Create Doujinshi object
doujin, err := gonhentai.NewDoujinshiId(354862)
if err != nil {
panic(err)
}

// You can use also
pathTemplate := gonhentai.DefaultDoujinNameTemplate + "/" + nehentai.DefaultPageNameTemplate

// Save doujinshi
err = d.Save("/home/<username/" + pathTemplate, 0644)
if err != nil {
t.Fatal(err)
}

fmt.Println("Doujinshi saved!")
```

### Search

```go
// Simple search
qr, err := nhentai.Search("Blend s", nhentai.QueryOptions{Page: "1"})
if err != nil {
	panic(err)
}
fmt.Println(qr)

// Tagged search
qr, err = nhentai.SearchTag(29859, nhentai.QueryOptions{Page: "1"})
if err != nil {
    panic(err)
}
fmt.Println(qr)

// Custom search
qr, err := nhentai.SearchCustom("Blend s", nhentai.QueryFilter {
    ToDelete: []nhentai.Filter{
        {
            Id:   0,
            Name: "yaoi",
            Type: nhentai.Tag,
        },
	},
    ToFilter: []nhentai.Filter{
        {
            Id:   0,
            Name: "maika sakuranomiya",
            Type: nhentai.Character,
        },
    },
})

// Check error
if err != nil {
    panic(err)
}
fmt.Println(qr)
```

For more examples, please refer to the [Documentation](https://example.com)

---

<!-- TODO -->
## TODO

- [ ] Random sauce
- [ ] User Info
- [ ] Authenticate
- [ ] Favourites
- [ ] Generate PDF
- [ ] Generate CBR/CBZ
- [ ] Tags list



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

KiritoNya - [@YuriLov95141178](https://twitter.com/YuriLov95141178) - watashiwayuridaisuki@gmail.com

Anilist: [KiritnoNya](https://anilist.co/user/KiritoNya/)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/KiritoNya/gonhentai.svg?style=for-the-badge
[contributors-url]: https://github.com/KiritoNya/gonhentai/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/KiritoNya/gonhentai.svg?style=for-the-badge
[forks-url]: https://github.com/KiritoNya/gonhentai/network/members
[stars-shield]: https://img.shields.io/github/stars/KiritoNya/gonhentai.svg?style=for-the-badge
[stars-url]: https://github.com/KiritoNya/gonhentai/stargazers
[issues-shield]: https://img.shields.io/github/issues/KiritoNya/gonhentai.svg?style=for-the-badge
[issues-url]: https://github.com/KiritoNya/gonhentai/issues
[license-shield]: https://img.shields.io/github/license/KiritoNya/gonhentai.svg?style=for-the-badge
[license-url]: https://github.com/KiritoNya/gonhentai/blob/master/LICENSE.txt
