
<h1>TODO app</h1>
<h2>Using</h2>
<p><b>[GET]</b>  /api/todo - get all tasks </p>
<p><b>[GET]</b>  /api/todo/:id - get task by id </p>
<p><b>[POST]</b>  /api/todo - create task </p>
<p><b>[PUT]</b>  /api/todo/:id - update task by id </p>
<p><b>[POST]</b>  /api/todo/:id/execute - mark execute </p>
<p><b>[DELETE]</b>  /api/todo/:id - delete task by id</p>
<hr>
<h2>Dependencies</h2>
<a href="github.com/valyala/fasthttp">fasthttp</a><br>
<a href="github.com/buaazp/fasthttprouter">fasthttprouter</a><br>
<a href="github.com/urfave/cli/v2">cli</a><br>
<a href="github.com/go-pg/pg">go-pg</a><br>
<a href="github.com/go-pg/pg/orm">orm</a><br>
<a href="github.com/joho/godotenv">godotenv</a><br><hr>
<h2>Installation (linux)</h2>
<h3>Clone the project</h3>
<code>git clone https://github.com/gospodinzerkalo/todo_app_golang</code>
<h3>Build and run</h3>
<code>make build</code><br>
<code>make run</code>