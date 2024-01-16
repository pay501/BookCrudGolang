"use client"
import axios from "axios";
import Link from "next/link";
import { useEffect, useState } from "react";

export default function HomePage() {
  const [items, setItems] = useState([]);

  async function getBooks() {
    try {
      const res = await axios.get("http://localhost:8080/books",{ withCredentials: true });
      setItems(res.data);
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
            <Link className="bg-blue-500 rounded-md" href={`/books/${book.ID}`}>
              Read More
            </Link>
          </li>
        ))}
      </ul>
    </div>
  );
}