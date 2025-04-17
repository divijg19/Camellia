// backend/index.js
const express = require("express");
const cors = require("cors");
const bodyParser = require("body-parser");

const app = express();
const PORT = process.env.PORT || 5000;

app.use(cors());
app.use(bodyParser.json());

// Dummy endpoints
app.get("/api/classes", (req, res) => {
  res.json([
    { id: 1, name: "Painting Basics", instructor: "Alice", price: 20 },
    { id: 2, name: "Python 101", instructor: "Bob", price: 30 }
  ]);
});

app.post("/api/book", (req, res) => {
  const { user, classId } = req.body;
  res.json({ success: true, message: `User ${user} booked class ${classId}` });
});

app.listen(PORT, () => console.log(`Server running on port ${PORT}`));
