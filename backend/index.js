// backend/index.js
const express = require("express");
const cors = require("cors");
const bodyParser = require("body-parser");
const classRoutes = require('./routes/classes');
const bookingRoutes = require('./routes/bookings');
const userRoutes = require('./routes/users');

const app = express();
const PORT = process.env.PORT || 3000;

app.use(cors());
app.use(bodyParser.json());

app.use('/api/classes', classRoutes);
app.use('/api/bookings', bookingRoutes);
app.use('/api/users', userRoutes);

// Dummy endpoints
app.get("/api/classes", (req, res) => {
  res.json([
    { id: 1, name: "Painting Basics", instructor: "Alice", price: 20 },
    { id: 2, name: "Python 101", instructor: "Bob", price: 30 }
  ]);
});

app.post("/api/bookings", (req, res) => {
  const { user, classId } = req.body;
  res.json({ success: true, message: `User ${user} booked class ${classId}` });
});

app.listen(PORT, () => console.log(`Server running on port ${PORT}`));
