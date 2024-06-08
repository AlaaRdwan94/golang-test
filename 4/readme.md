Demonstrate Database Transaction of at least 10 different tables with proper error handling using
a single api(use GIN).

In this example:

We initialize the database connection in the initDB function.
We define a handleTransaction function to handle the API endpoint for performing a transaction involving multiple tables.
Inside the handleTransaction function, we start a database transaction using db.Begin().
We perform database operations (e.g., insertions) within the transaction. For simplicity, we assume inserting into 10 different tables.
If any error occurs during the transaction or any of the database operations, we rollback the transaction using tx.Rollback().
If all operations are successful, we commit the transaction using tx.Commit().
Proper error handling is implemented to return appropriate HTTP responses in case of errors.
Please adjust the database operations according to your actual database schema and requirements. 

-- more work to do 
ensure handling and logging in place for production use.