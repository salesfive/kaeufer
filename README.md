# kaeufer
A Go client for fetching leads from the kaeuferportal.de API.


### Example
```go
c := kaeufer.Client{
  Id:     "client id",
  Secret: "client secret",
}

leads, err := c.Leads(
  kaeufer.PerPage(10),
  kaeufer.FromTimestamp(1459157308),
)

```
