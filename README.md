Translate CLI
-------

Translate between Japanese and English on the CLI.

## Install

1. Create Google App Script below

```
function doGet(e) {
  const p = e.parameter
  const translated = LanguageApp.translate(p.text, p.source, p.target)
  let body = {}
  if (translated) {
    body = {
      status: 200,
      text: translated,
    }
  } else {
    body = {
      status: 400,
      text: "Bad request.",
    }
  }
  const response = ContentService.createTextOutput()
  response.setMimeType(ContentService.MimeType.JSON)
  response.setContent(JSON.stringify(body))

  return response
}
```
2. Set the script URL in the environment variable TRNS_URL.

## Usase

  $ trns **text you want to translate**
