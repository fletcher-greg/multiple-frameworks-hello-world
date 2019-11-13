const express = require("express");

const app = express();

app.use((req, res) => {
  res.send("Hello World");
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => console.log(`Express listening on port: ${PORT}`));
