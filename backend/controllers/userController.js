const { v4: uuidv4 } = require('uuid');

const createUser = async (req, res) => {
  const user = { ...req.body, id: uuidv4() };
  // Save to DynamoDB
  res.status(201).json({ message: 'User registered', user });
};

module.exports = { createUser };

