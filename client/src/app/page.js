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
      <h1 className="text-4xl font-bold">Welcome Back!</h1>
    </div>
  );
}