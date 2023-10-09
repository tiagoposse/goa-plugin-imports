# goa-plugin-imoprts

Provides a way to specify additional imports for a file in goa.design. You can add a meta attribute "import" to a type, so the file in which the type is generated gets such import added.

The motivation for this plugin is to be able to use custom types for attributes, such as custom structs.

## Example

Consider the following example to assign the user type a Group attribute of the Group type, which is declared in another module.

```
## utils/utils.go

type Group struct {
  Name string
}

## design/user.go

var User = ResultType("Group", func() {
	Meta("import", "github.com/tiagoposse/app/utils")
	Attribute("name", String, "Name of the user")
	Attribute("group", String, "Group the user belongs to", func() {
    Meta("struct:field:type", "utils.Group")
	})
})
```

It will work even you pass a `Meta("struct:pkg:path", val)` is passed to the type.
