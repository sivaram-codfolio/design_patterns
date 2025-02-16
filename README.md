# Go Design Patterns

A collection of commonly used design patterns implemented in Go. This repository serves as a reference for developers looking to understand, implement, and apply design patterns effectively in real-world Go applications.

🚀 *Patterns Covered**

✅ Builder Pattern – Step-by-step object construction with a fluent API.

🏗 **Builder Pattern**

The Builder Pattern helps construct complex objects in a structured manner. Here’s an example implementation using a Notification Builder:

📌 *Example Execution*

```bash
go run cmd/builder.go
```

📜 *Sample Output*

```bash
Notification: &{Title:New Notification SubTitle:This is a subtitle Message:This is a basic notification Image:image.jpg Icon:icon.png Priority:5 Type:alert}
```

🏭 **Factory Pattern**

The Factory Pattern provides a way to create objects without specifying their exact class. It encapsulates object creation, promoting flexibility and scalability. Here’s an example implementation using a Publication Factory:

📌 *Example Execution*

```bash
go run cmd/factory.go
```

📜 *Sample Output*

```bash
--------------------
This is a magazine named Tyme
Type: *main.Magazine
Name: Tyme
Pages: 50
Publisher: The Tymes
--------------------
--------------------
This is a magazine named Lyfe
Type: *main.Magazine
Name: Lyfe
Pages: 40
Publisher: Lyfe Inc
--------------------
--------------------
This is a newspaper named The Herald
Type: *main.Newspaper
Name: The Herald
Pages: 60
Publisher: Heralders
--------------------
--------------------
This is a newspaper named The Standard
Type: *main.Newspaper
Name: The Standard
Pages: 30
Publisher: Standarders
--------------------
```

🔄 **Singleton Pattern**
The Singleton Pattern ensures that only one instance of a class exists and provides a global access point to it. It is useful for managing shared resources like logging, database connections, or configurations.

📌 *Example Execution*

```bash
go run cmd/singleton.go
```

📜 *Sample Output*

```bash
Creating Logger Instance
Returning Logger Instance
1 : This is a log message
Returning Logger Instance
2 : This is a log message
Returning Logger Instance
3 : This is a log message
Returning Logger Instance
Returning Logger Instance
Returning Logger Instance
Returning Logger Instance
Returning Logger Instance
Returning Logger Instance
Returning Logger Instance
Returning Logger Instance
Returning Logger Instance
_
```

🏢 **Facade Pattern**

The Facade Pattern provides a simplified interface to a complex subsystem, making it easier to use. It hides the complexities of the system and exposes only the necessary parts.

📌 *Example Execution*

```bash
go run cmd/facade.go
```

📜 *Sample Output*

```bash
Making an Americano
--------------------
Starting coffee order with beans: 5 and grind level 2
Grinding the beans: 5 beans at 2 granularity
Adding hot water: 8
Ending coffee order
Americano is ready!

Making a Latte
--------------------
Starting coffee order with beans: 7.5 and grind level 4
Grinding the beans: 7.5 beans at 4 granularity
Adding hot water: 12
Adding milk: 3 oz
Foam setting: true
Ending coffee order
Latte is ready!
```

🏢 **Adapter Pattern**

The Adapter Pattern allows incompatible interfaces to work together. It acts as a bridge between two interfaces, enabling code reusability and flexibility.

📌 *Example Execution*

```bash
go run cmd/adapter.go
```

📜 *Sample Output*

```bash
SonyTV is now on
Increasing SonyTV volume to 21
Decreasing SonyTV volume to 20
Decreasing SonyTV channel to 10
Decreasing SonyTV channel to 9
Setting SonyTV channel to 68
SonyTV is now off
--------------------
SamsungTV is on
SamsungTV volume is 35
Setting SamsungTV volume to 36
SamsungTV volume is 36
Setting SamsungTV volume to 35
SamsungTV channel is 13
Setting SamsungTV channel to 14
SamsungTV channel is 14
Setting SamsungTV channel to 13
Setting SamsungTV channel to 68
SamsungTV is off
```

🏢 **Observer Pattern**

The Observer Pattern is a behavioral design pattern where an object (the subject) maintains a list of observers that need to be notified of any state changes.

📌 *Example Execution*

```bash
go run cmd/observer.go
```

📜 *Sample Output*

```bash
Listener: Listener 1 got data change: Monday!
Listener: Listener 2 got data change: Monday!
Listener: Listener 1 got data change: Wednesday!
Listener: Listener 2 got data change: Wednesday!
Listener: Listener 1 got data change: Friday!
```

🏢 **Iterator Pattern**

The Iterator Pattern in Go is a behavioral design pattern used to traverse elements of a collection without exposing its underlying structure. Go doesn’t have built-in iterators like some other languages (e.g., Python or Java), but you can implement the iterator pattern using structs and methods.

📌 *Example Execution*

```bash
go run cmd/iterator.go
```

📜 *Sample Output*

```bash
Book title: War and Peace
Book title: Crime and Punishment
Book title: Brave New World
Book title: Catcher in the Rye
Book title: To Kill a Mockingbird
Book title: The Grapes of Wrath
Book title: Wuthering Heights
Book author: Leo Tolstoy
Book author: Leo Tolstoy
Book author: Aldous Huxley
Book author: J.D. Salinger
Book author: Harper Lee
Book author: John Steinbeck
Book author: Emily Bronte
Book &{Name:War and Peace Author:Leo Tolstoy PageCount:864 Type:0}
Book &{Name:Crime and Punishment Author:Leo Tolstoy PageCount:1225 Type:1}
Book &{Name:Brave New World Author:Aldous Huxley PageCount:325 Type:2}
Book &{Name:Catcher in the Rye Author:J.D. Salinger PageCount:206 Type:0}
Book &{Name:To Kill a Mockingbird Author:Harper Lee PageCount:399 Type:2}
Book &{Name:The Grapes of Wrath Author:John Steinbeck PageCount:464 Type:0}
Book &{Name:Wuthering Heights Author:Emily Bronte PageCount:288 Type:3}
```

📌 **Future Plans**

This repository will continue to grow with more design patterns, including:
- Strategy
- …and more!

Stay tuned for updates. Contributions and suggestions are welcome!