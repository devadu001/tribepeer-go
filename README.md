# tribepeer (Go)

Server client for TribePeer — e-campus, organisation training, e-library, and AI-powered learning.

```bash
go get github.com/devadu001/tribepeer-go
```

```go
import tribepeer "github.com/devadu001/tribepeer-go"

client := tribepeer.New(os.Getenv("TP_CLIENT_ID"), os.Getenv("TP_CLIENT_SECRET"))
classes, err := client.Tribes.List(ctx)
```

## Keys

1. Create an account — [tribepeer.com/register](https://www.tribepeer.com/register)
2. Become a Tribe Owner — [tribepeer.com/tribe-owner/apply](https://www.tribepeer.com/tribe-owner/apply)
3. Issue keys — [tribepeer.com/tribe-owner/credentials](https://www.tribepeer.com/tribe-owner/credentials)
4. API guide — [tribepeer.com/institutions/docs](https://www.tribepeer.com/institutions/docs)

You get a **client id** (`tp_id_…`) and a **client secret** (`tp_sec_…`). The secret stays on your server.

```env
TP_CLIENT_ID=tp_id_…
TP_CLIENT_SECRET=tp_sec_…
```

Go 1.22+. MIT.
