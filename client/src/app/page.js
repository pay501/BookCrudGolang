"use client"
import { useEffect, useState } from "react";

export default function HomePage() {
  const [items, setItems] = useState([]);

  async function getBooks() {
    try {
      const res = await fetch("http://localhost:8080/books");
      const data = await res.json();
      setItems(data);
    } catch (error) {
      console.error("Error fetching books:", error);
    }
  }

  useEffect(() => {
    getBooks();
  }, []);
  
  return (
    <div>
      <h1>Book List</h1>
      <ul>
        {items.map((book) => (
          <li key={book.ID}>
            <strong>{book.name}</strong> by {book.author} - ${book.price}
          </li>
        ))}
      </ul>
    </div>
  );
}