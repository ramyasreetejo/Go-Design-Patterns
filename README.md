# Go Design Patterns

A collection of Go implementations of classic design patterns, written with clarity and idiomatic Go practices in mind.

This repository focuses on demonstrating how common object-oriented design patterns translate into Go’s interface-driven, composition-first design philosophy.

---

## Contents

### Creational Patterns

#### Abstract Factory

* **Creates related products as a family**: Each factory creates a matching `Lipstick` and `Lipliner` of the same brand (Maybelline or Revlon).
* **Hides concrete implementations**: The client works only with interfaces, not concrete product types.
* **Keeps variants consistent**: A factory guarantees that all products it creates belong to the same brand family.

---

## Project Structure

```
Go-Design-Patterns/
├── go.mod
├── Creational/
│   └── AbstractFactory.go
├── README.md
└── LICENSE
```

---

## Getting Started

Clone the repository:

```bash
git clone https://github.com/ramyasreetejo/Go-Design-Patterns.git
cd Go-Design-Patterns
```

Run the Abstract Factory example:

```bash
go run Creational/AbstractFactory.go
```

---

## About

I’m Ramya Sree Tejomurtula, a backend-focused Software Engineer and a graduate from NIT Andhra Pradesh.
I love building projects that make sense to me and spark my curiosity.

---

## License

This project is licensed under the MIT License. See the LICENSE file for details.
