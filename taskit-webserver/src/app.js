/**
 * Configuring server, using the created routes
 */

const express = require('express');
require('./db/sqlite');

const app = express();

/**
 * Just a test route for now
 */
app.get('/', (req, res) => {
  res.send('Hello World!');
});

module.exports = app;