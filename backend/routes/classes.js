const express = require('express');
const router = express.Router();
const { getAllClasses, createClass } = require('../controllers/classController');

router.get('/', getAllClasses);
router.post('/', createClass);

module.exports = router;
