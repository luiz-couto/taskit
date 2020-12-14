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
db.run('CREATE TABLE IF NOT EXISTS tasks (title TEXT NOT NULL, description TEXT, status TEXT NOT NULL, priority INTEGER NOT NULL DEFAULT 0, blocked INTEGER NOT NULL DEFAULT -1, deadline TEXT NOT NULL DEFAULT "", workingEnter TEXT NOT NULL DEFAULT "0", workingElapsed TEXT NOT NULL DEFAULT "0")', (err) => {
  if (err) {
    return console.log(err.message);
  }
  console.log('Table tasks created successfully!');
});