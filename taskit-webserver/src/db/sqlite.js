/**
 * Init sqlite db
 */

const sqlite3 = require('sqlite3').verbose();

/**
 * Create database if don't exists
 */
let db = new sqlite3.Database('./src/db/taskit.db', (err) => {
  if (err) {
    return console.log(err.message);
  }
  console.log('Connected to the taskit database.');
});

/**
 * Create table if don't exists
 */
db.run('CREATE TABLE IF NOT EXISTS tasks (title TEXT NOT NULL, description TEXT)', (err) => {
  if (err) {
    return console.log(err.message);
  }
  console.log('Table tasks created successfully!');
});