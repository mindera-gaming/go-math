# Polygon Triangulation - Ear Clipping Method
This package offers a polygon triangulation solution using the Ear Clipping method

- [Installation](#installation)
- [Usage](#usage)
    - [Requirements](#mequirements)
    - [Mentioned structures](#mentioned-structures)


## Installation

Simply run the following command to install this package to your GOPATH:
```shell
go get github.com/mindera-gaming/go-ear-clipping
```


## Usage

The `triangulation` package API exposes one function:

```go
// Triangulate decomposes a simple polygon into a set of triangles
func Triangulate(vertices []vector.Vector2, options TriangulationOptions) (triangles []int, area float64, err error)
```

`Triangulate` takes the vertices of your polygon `[]Vector2` and some triangulation options `TriangulationOptions`.  
At the end, it returns the set of indices `[]int` of the vertices of the calculated triangles and the area `float64` of the polygon.


### Requirements

In order for your polygon to be successfully triangulated, it needs to satisfy a few requirements:

- `[Mandatory]` It must be a simple polygon
- `[Mandatory]` The polygon must not contain collinear edges
- `[Optional]` The vertices of the polygon need to be sent clockwise

If you think your polygon satisfies the indicated requirements, you can use the `TriangulationOptions` to skip some validations.


### Mentioned structures

```go
// TriangulationOptions is the structure that defines the triangulation options
type TriangulationOptions struct {
	SkipSimplePolygonValidation bool // set to true to skip the simple polygon verification
	SkipColinearEdgesValidation bool // set to true to skip the collinear edge verification
	SkipWindingOrderValidation  bool // set to true to skip the winding order verification
}

// Vector2 represents a 2D vector, point or position
type Vector2 struct {
	X float64
	Y float64
}
```
