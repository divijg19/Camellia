const { v4: uuidv4 } = require('uuid');

const bookClass = async (req, res) => {
  const { userId, classId } = req.body;
  const booking = {
    id: uuidv4(),
    userId,
    classId,
    status: 'confirmed',
    createdAt: new Date().toISOString()
  };

  // Here: call Lambda to process payment

  res.status(201).json({ message: 'Booking successful', booking });
};

module.exports = { bookClass };

