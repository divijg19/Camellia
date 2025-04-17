const { v4: uuidv4 } = require('uuid');

const getAllClasses = async (req, res) => {
  // In production: query DynamoDB here
  const sampleClasses = [
    {
      id: '123',
      title: 'Beginner Painting',
      instructor: 'Jane Doe',
      location: 'Online',
      price: 25
    }
  ];
  res.json(sampleClasses);
};

const createClass = async (req, res) => {
  const newClass = { ...req.body, id: uuidv4() };
  // Save to DynamoDB
  res.status(201).json({ message: 'Class created', class: newClass });
};

module.exports = { getAllClasses, createClass };

