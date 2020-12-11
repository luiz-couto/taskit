/**
 * Configuring server, using the created routes
 */

const express = require('express');
require('./db/sqlite');

const taskRouter = require('./routers/task');

const app = express();

app.use(express.json());
app.use(taskRouter);

/**
 * Just a test route for now
 */
app.get('/', (req, res) => {
  res.send('Hello World!');
});


module.exports = app;