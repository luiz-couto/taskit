/**
 * Raise server and listen on the configured port
 */

const app = require('./app');
const PORT = 8080;

app.listen(PORT, () => {
  console.log('Server is up on port ' + PORT);
});


