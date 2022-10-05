const http = require("http");
const express = require("express");

const app = express();
http.createServer(app).listen(5555, () => {
  console.log("Server is running at port 5555");
});

app.post("/customer/create", (req, res, next) => {
  res.json({ message: "Success" });
});

app.post("/customer/cancel", (req, res, next) => {
  res.json({ message: "Canceled" });
});
