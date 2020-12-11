/**
 * Task Router
 */
const sqlite3 = require('sqlite3').verbose();
const express = require('express');
const router = new express.Router();

/**
 * Open database and return it
 */
const openDatabase = () => {
  let db = new sqlite3.Database('./src/db/taskit.db');
  return db;
}

/**
 * Get a list of all tasks (used for debug only)
 */
router.get('/tasks', async (req, res) => {
  let sql = `SELECT rowid, * FROM tasks`;

  let db = openDatabase();
  
  db.all(sql, [], (err, rows) => {
    if (err) {
      throw err;
    }
    res.status(200).send(rows);
  });

  db.close();

});

/**
 * Create a task
 */
router.post('/tasks', async (req, res) => {

  if (!req.body.title || req.body.title.length == 0) {
    res.status(400).send('Title cant be empty!');
  }
  
  let sql = `INSERT INTO tasks(title, description) VALUES(?, ?)`;
  let db = openDatabase();

  db.run(sql, [req.body.title, req.body.description], (err) => {
    if (err) {
      console.log(err.message);
      res.status(500).send('An error has occured');
    }
    res.status(200).send('Row was added to the table');
  });

  db.close();
});

module.exports = router;