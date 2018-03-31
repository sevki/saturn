# Saturn
Saturn is a silly web server experiment, it uses the [`http.Filesystem`](https://golang.org/pkg/net/http/#FileSystem) interface to create middlewares (which probably is not a good idea).

## Moons
Each middleware is named after one os saturn's moons, not because the anology works but because saturn has a shit ton of moons and I'll never run out of things to name the middlewares.

### Pan
[`pan`](https://sevki.org/saturn/pan) is the middleware that converts markdowns to html and pdf.

### Atlas
[`atlas`](https://sevki.org/saturn/atlas) is the middleware that takes paths and guesses additional extentions for them.

### Titan
[`titan`](https://sevki.org/saturn/titan) is a [`http.Filesystem`](https://golang.org/pkg/net/http/#FileSystem) backed by [upspin](https://upspin.io).

## Maturity
Most stuff here is experimental.  	**DO NOT** use in prod.

