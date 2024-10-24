"use client"

import axios from "axios"
import { useRouter } from "next/navigation"
import { useEffect, useState } from "react"
import Swal from "sweetalert2"
import style from "./manageBook.module.css"
import Link from "next/link"

export default function ManageBookPage() {

    const router = useRouter()
    const [books, setBooks] = useState([])

    async function getBooks() {
        try {
            await axios.get("http://localhost:8080/books", { withCredentials: true })
                .then(res => {
                    if(res.data.code == 200){
                        setBooks(res.data.result)
                    }
                })
        } catch (err) {
            console.log("Message Error Fetching Data :::", err)
        }
    }

    async function deleteBook(id) {
        try {
            await axios.get("http://localhost:8080/checkSession", { withCredentials: true })
                .then(res => {
                    if (res.data.message === "return") {
                        Swal.fire({
                            icon: "error",
                            title: "Please, Login!",
                            timer: "2000",
                        })
                        router.push("/login")
                    } else {
                        Swal.fire({
                            title: "Are you sure?",
                            text: "You won't be able to revert this!",
                            icon: "warning",
                            showCancelButton: true,
                            confirmButtonColor: "#3085d6",
                            cancelButtonColor: "#d33",
                            confirmButtonText: "Yes, delete it!"
                        }).then((result) => {
                            if (result.isConfirmed) {
                                axios.delete(`http://localhost:8080/deleteBook/${id}`, { withCredentials: true })
                                    .then(res => {
                                        if (res.data.message === "return") {
                                            Swal.fire({
                                                icon: "error",
                                                title: "Delete failed!",
                                                text: "PLease, Login!"
                                            })
                                            router.push("/login")
                                        } else {
                                            Swal.fire({
                                                icon: "success",
                                                title: "Delete Successful!",
                                                footer: "Thanks!",
                                                timer: "1000"
                                            })
                                            getBooks()
                                        }
                                    })
                            }
                        });
                    }
                })
        } catch (err) {
            console.log(err)
        }
    }

    async function checkSession() {
        await axios.get("http://localhost:8080/checkSession", { withCredentials: true })
            .then(res => {
                if (res.data.message === "return") {
                    Swal.fire({
                        icon: "error",
                        title: "Please, Login!",
                        timer: "2000",
                    })
                    router.push("/login")
                } else {
                    getBooks()
                }
            })
    }
    useEffect(() => {
        checkSession()
        getBooks()
    }, [])
    console.log(books)


    return (
        <div className={`${style.container}`}>
            <div className="overflow-x-auto">
                <table className="table">
                    {/* head */}
                    <thead>
                        <tr>
                            <th>Title</th>
                            <th>Description</th>
                            <th>Operation</th>
                        </tr>
                    </thead>
                    <tbody>
                        {
                            books.map((val, key) => (
                                <tr key={key}>
                                    <td>
                                        <div className="flex items-center gap-3">
                                            <div className="avatar">
                                                <div className="mask mask-squircle w-12 h-12">
                                                    <img src={val.image} alt="Avatar Tailwind CSS Component" />
                                                </div>
                                            </div>
                                            <div>
                                                <div className="font-bold">{val.name}</div>
                                                <div className="text-sm opacity-50">{val.author}</div>
                                            </div>
                                        </div>
                                    </td>
                                    <td className="max-w-sm">
                                        <span className="badge badge-ghost badge-sm">{val.description}</span>
                                    </td>
                                    <th>
                                        <div className="flex gap-4">
                                            <button
                                                onClick={() => deleteBook(val.ID)}
                                                className={style.delete}
                                            >
                                                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path d="M135.2 17.7L128 32H32C14.3 32 0 46.3 0 64S14.3 96 32 96H416c17.7 0 32-14.3 32-32s-14.3-32-32-32H320l-7.2-14.3C307.4 6.8 296.3 0 284.2 0H163.8c-12.1 0-23.2 6.8-28.6 17.7zM416 128H32L53.2 467c1.6 25.3 22.6 45 47.9 45H346.9c25.3 0 46.3-19.7 47.9-45L416 128z" /></svg>
                                            </button>
                                            <Link
                                                className={style.update}
                                                href={`/books/update/${val.ID}`}
                                            >
                                                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M471.6 21.7c-21.9-21.9-57.3-21.9-79.2 0L362.3 51.7l97.9 97.9 30.1-30.1c21.9-21.9 21.9-57.3 0-79.2L471.6 21.7zm-299.2 220c-6.1 6.1-10.8 13.6-13.5 21.9l-29.6 88.8c-2.9 8.6-.6 18.1 5.8 24.6s15.9 8.7 24.6 5.8l88.8-29.6c8.2-2.7 15.7-7.4 21.9-13.5L437.7 172.3 339.7 74.3 172.4 241.7zM96 64C43 64 0 107 0 160V416c0 53 43 96 96 96H352c53 0 96-43 96-96V320c0-17.7-14.3-32-32-32s-32 14.3-32 32v96c0 17.7-14.3 32-32 32H96c-17.7 0-32-14.3-32-32V160c0-17.7 14.3-32 32-32h96c17.7 0 32-14.3 32-32s-14.3-32-32-32H96z" /></svg>
                                            </Link>
                                        </div>
                                    </th>
                                </tr>
                            ))
                        }
                    </tbody>
                </table>
            </div>
        </div>
    )
}