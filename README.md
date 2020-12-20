# ImageDropShadow

Image drop shadow using Go.

# Usage

Download the executable and run from the command prompt

## Options

* _-i_: Path to input image
* _-b_: Toggle image border. No border by default
* _-t_: Transparent background

e.g.

```
$ ./imageDropShadow-mac -i cnh.png
```

```
$ ./imageDropShadow-mac -i joey.png -t
```

# Samples

## Input 1

<img src="images/cnh.png" alt="Calvin And Hobbes" height="200px">

Output

<img src="images/cnh-out.png" alt="Calvin And Hobbes" height="208px">

---

## Input 2

<img src="images/joey.png" alt="Joey" height="200px">

Output

<img src="images/joey-out.png" alt="Joey" height="210px">

# Releases

| Release | Description                                                                                      |
| ------- | ------------------------------------------------------------------------------------------------ |
| 1.0     | <ul><li>Basic working version where a shadow effect is added to a given png/jpg image.</li></ul> |
| 1.1     | <ul><li>Added stroke</li><li>Fixed the shadow effect</li></ul>                                   |
| 1.2     | <ul><li>Optional stroke</li></ul>                                   |