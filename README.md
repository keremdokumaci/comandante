![Version](https://img.shields.io/badge/version-0.1.17-orange.svg)
[![GolangCI Lint](https://github.com/keremdokumaci/comandante/actions/workflows/go-lint.yml/badge.svg)](https://github.com/keremdokumaci/comandante/actions/workflows/go-lint.yml)

# Comandante

Comandante helps to manage server configurations for your application with a basic user interface.

#### Installation

Run command `go get github.com/keremdokumaci/comandante`

#### Usage

Use **Configure** function in the entry point of your service.

**StoreIn** values

- redis
- psql (coming soon)
- mysql (coming soon)
- mongo (coming soon)

![comandante configuration](https://iili.io/ildOzJ.png)

After your app runs, you can access the comandante UI via using configured endpoint. For the above image, this is **/comandante**.

![comandante ui](https://iili.io/ilf2yP.png)

You can easily add, update or delete configuration variables from the page.

In your application, you can use **Get** method to get configuration variables. This method supports generics.

**Example usage**

`val, err := comandante.Get[SomeType]("some_key")`
