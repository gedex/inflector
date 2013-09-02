Inflector
=========

Inflector pluralizes and singularizes English nouns.

**TODO** Add godoc and build links

## Basic Usage

There are only two methods: `Pluralize` and `Singularize`.

~~~go
s := "People"
fmt.Println(inflector.Singularize(s)) // will print "Person"

s2 := "octopus"
fmt.Println(inflector.Pluralize(s2)) // will print "octopuses"
~~~

Please see [example/example.go](./example/example.go) for a complete example.

## Credits

* [CakePHP's Inflector](https://github.com/cakephp/cakephp/blob/master/lib/Cake/Utility/Inflector.php)

## License

This library is distributed under the BSD-style license found in the LICENSE.md file.
