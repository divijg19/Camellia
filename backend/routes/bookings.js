const express = require('express');
const router = express.Router();
const { bookClass } = require('../controllers/bookingController');

router.post('/', bookClass);

module.exports = router;

