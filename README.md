# Goldmark D2

[![Go Reference](https://pkg.go.dev/badge/github.com/antonklava/goldmark-d2.svg)](https://pkg.go.dev/github.com/antonklava/goldmark-d2)

Goldmark D2 is a [Goldmark](https://github.com/yuin/goldmark) extension providing diagram support through [D2](https://d2lang.com/).

## Usage

```go
goldmark.New(
	goldmark.WithExtensions(&Extender{
		// Defaults when omitted
		Layout:  Ptr("dagre"), // "elk" also works
		ThemeID: Ptr(d2themescatalog.CoolClassics.ID),
		Sketch:	 Ptr(false),
	}),
).Convert(src, dst)
```

## Example

<table>
<tr>
<td>

````markdown
The following diagram shows the important link between the letters X and Y:

```d2
x -> y
```
````

</td>
<td>

![](testdata/basic.png)

</td>
</tr>

<tr>
<td>

{Sketch: true}

````markdown
```d2
dogs -> cats -> mice: chase
replica 1 <-> replica 2
a -> b: To err is human, to moo bovine {
  source-arrowhead: 1
  target-arrowhead: * {
    shape: diamond
  }
}
```
````

</td>
<td>

![](testdata/connections.png)

</td>
</tr>
</table>
