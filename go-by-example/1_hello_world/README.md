## Package
Go program is created within packages. It means, either we create a package and/or use a package.

Default package that is used as program entry point is `package main`. But, we can create a specific package by defining package name.

## Import
To use a package we could use `import` command. Notice `import "fmt"`. To use more than one package, we can surround `"fmt"` with `(...)`

Example:

```
import (
    "fmt"
    "math/rand"
)
```

Regarding to the naming convention, the package name is the same name as last element of import path. E.g. `math/rand` it means the package name is rand on math directory
