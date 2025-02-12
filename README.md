# Go Design Patterns

A collection of commonly used design patterns implemented in Go. This repository serves as a reference for developers looking to understand, implement, and apply design patterns effectively in real-world Go applications.

🚀 *Patterns Covered**

✅ Builder Pattern – Step-by-step object construction with a fluent API.

🏗 **Builder Pattern**

The Builder Pattern helps construct complex objects in a structured manner. Here’s an example implementation using a Notification Builder:

📌 *Example Execution*

```bash
go run builder/main.go builder/notification.go builder/builder.go
```

📜 *Sample Output*

```bash
Notification: &{Title:New Notification SubTitle:This is a subtitle Message:This is a basic notification Image:image.jpg Icon:icon.png Priority:5 Type:alert}
```

🏭 **Factory Pattern**

The Factory Pattern provides a way to create objects without specifying their exact class. It encapsulates object creation, promoting flexibility and scalability. Here’s an example implementation using a Publication Factory:

📌 *Example Execution*

```bash
go run factory/main.go factory/factory.go factory/publication.go factory/newspaper.go factory/magazine.go
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

📌 **Future Plans**

This repository will continue to grow with more design patterns, including:
- Singleton
- Adapter
- Observer
- Strategy
- …and more!

Stay tuned for updates. Contributions and suggestions are welcome!