# nocors
```
fetch('http://localhost:5555?url=https://my-example', {
    method: "POST",
    headers: { "Content-type": "application/json" },
    body: JSON.stringify({ market: "SE" })
})
.then(v => v.json())
.then(console.log)
.catch(console.log);
```