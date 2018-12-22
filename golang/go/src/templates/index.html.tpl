<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="UTF-8">
<title>=(.^.)=</title>
</head>
<body>
<p>
  <div id="app">
     message
  </div>
</p>
</body>
<script type="text/javascript">
  const httpGet = (url, responseType = 'json') => {
    return new Promise(function (resolve, reject) {
      const request = new XMLHttpRequest()
      request.onload = function () {
        if (this.status === 200) {
          resolve(this.response)
        } else {
          reject(new Error(this.statusText))
        }
      }
      request.onerror = function () {
        reject(new Error(`XMLHttpRequest Error: ${this.statusText}`))
      }

      request.open('GET', url)
      request.responseType = responseType
      request.send()
    })
  }
  httpGet('/api').then(data => {
    console.log(data)
  }).catch(status => {
    console.error(`Error: status code ${status}`)
  })

</script>

<script src="https://cdn.jsdelivr.net/npm/vue"></script>

</html>

