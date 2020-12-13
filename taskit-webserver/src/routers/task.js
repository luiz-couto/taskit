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
 * Get a list of all tasks
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
 * Get task by id
 * @param id task id to be fetched
 */
router.get('/tasks/:id', async (req, res) => {

  const id = req.params.id;
  if (!id) {
    res.status(400).send('Need to specify task id!');
  }
  let sql = `SELECT rowid, * FROM tasks WHERE rowid = ${id} `;

  let db = openDatabase();
  
  db.all(sql, [], (err, rows) => {
    if (err) {
      res.status(400).send(err.message);
    }
    res.status(200).send(rows);
  });

  db.close();

});


/**
 * Create a task
 * @param title task title
 * @param description task description
 */
router.post('/tasks', async (req, res) => {

  if (!req.body.title || req.body.title.length == 0) {
    res.status(400).send('Title cant be empty!');
  }
  
  let sql = `INSERT INTO tasks(title, description, status) VALUES(?, ?, ?)`;
  let db = openDatabase();

  db.run(sql, [req.body.title, req.body.description, req.body.status], (err) => {
    if (err) {
      console.log(err.message);
      res.status(500).send('An error has occured');
    }
    res.status(200).send('Row was added to the table');
  });

  db.close();
});


/**
 * Remove task by id
 * @param id task id to be removed
 */
router.delete('/tasks/:id', async (req, res) => {
  const id = req.params.id;
  if (!id) {
    res.status(400).send('Need to specify task id');
  }

  let sql = `DELETE FROM tasks WHERE rowid = ${id}`;
  let db = openDatabase();

  db.run(sql, (err) => {
    if (err) {
      console.log(err.message);
      res.status(500).send('An error has occured');
    }
    res.status(200).send('Deleted row in table');
  });

  db.close();

});

/**
 * Update task
 * @param id task id to be updated
 * @param property property to be updated
 * @param value new value of property
 */
router.patch('/tasks/:id', async (req, res) => {
  const id = req.params.id;
  if (!id) {
    res.status(400).send('Need to specify task id');
  }

  let sql = `UPDATE tasks SET "${req.body.property}" = "${req.body.value}" WHERE rowid = ${id}`;
  let db = openDatabase();

  db.run(sql, (err) => {
    if (err) {
      console.log(err.message);
      res.status(500).send('An error has occured');
    }
    res.status(200).send('Updated row in table');
  });

  db.close();

})

module.exports = router;