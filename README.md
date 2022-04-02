# Problem 2 - Construct the Thing
How would you construct an event ( message bus ) driven microservices for the given scenario where a user requested to see his/her order, profile, and invoices but the data is distributed separately among Order, Profile, and Invoice Micro Services each containing its own Database. Kindly, please illustrate with a diagram.
 
## Solution:
User requests his/her data Command Query Responsibility Segregation (CQRS), CQRS is one of the important patterns when querying between microservices. Instead of querying data from multiple services, use events to maintain a read-only view database that replicates data from the services. We can use the CQRS design pattern in order to avoid complex queries to get rid of inefficient joins. Basically this pattern separates read and update operations for a database. CQRS is all about the separation of concerns; it is split into a persistent data model and the modules into two sides: command side and the query side. We can define a read-only database that combines data from multiple service syncs and update the data by subscribing to published events from multiple services and updating that database.
 
![CQRS Diagram](cqrs_diagram.png "Command Query Responsibility Segregation (CQRS) Pattern")
 
In the CQRS Diagram we have three microservices Profile, Order and Invoice. In these three services, when any data changes, we just need to publish events from each service in event handlers like profile events, order events and invoice events. CQRS service will consume these events and update the database accordingly. Here we can use kafka as a distributed publish-subscribe messaging system or rabbitmq as a message broker. For the database we can use NoSQL database or any relational database.
When a user requests to see his/her order, profile, and invoices, CQRS service gets data for this user from the database and returns the response to the user.
