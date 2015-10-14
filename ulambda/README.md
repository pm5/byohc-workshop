
Untyped lambda calculus interpreter
===================================


Data structure
--------------

* Define lambda: `["lam", arg_name, body]`
* Apply lambda: `["app", func, arg_value]`
* Variable: `["var", var_name]`

For example,

```
[ "app", [ "lam", "true", [ "app", [ "lam", "false", [ "app", [ "lam", "and", [ "app", [ "app", [ "var", "and" ], [ "var", "true" ] ], [ "var", "true" ] ] ], [ "lam", "a", [ "lam", "b", [ "app", [ "app", [ "var", "a" ], [ "var", "b" ] ], [ "var", "false" ] ] ] ] ] ], [ "lam", "a", [ "lam", "b", [ "var", "b" ] ] ] ] ], [ "lam", "a", [ "lam", "b", [ "var", "a" ] ] ] ]
```

represents

```
(\true (\false (\and (and true) true)(\a \b (a b) false))(\a \b b))(\a \b a)
```
