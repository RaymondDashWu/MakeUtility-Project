# SVG Graphs Using Go!

[![Go Report Card](https://goreportcard.com/badge/github.com/RaymondDashWu/MakeUtility-Project)](https://goreportcard.com/report/github.com/RaymondDashWu/MakeUtility-Project)

Used to provide basic graphs for my latest NLP project.
https://github.com/RaymondDashWu/DS1.1-Class-Files/blob/master/IMDB%20Data%20Analysis%20%2B%20NLP%20Predictions.ipynb

To run:
go run main.go > demo.svg

Based on Gophercises tutorial
https://gophercises.com/exercises/image


### Tech
This package makes extensive use of the SVGo library from ajstarks 
https://github.com/ajstarks/svgo

### Installation
1. To install this first make sure that your GOPATH is properly set.
2. Git clone the repo and make changes to the values in the data struct.
3. Type the following into the directory's terminal
```sh
$ go run main.go > FILENAME.svg
```
